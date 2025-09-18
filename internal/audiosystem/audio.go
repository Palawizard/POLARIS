package audiosystem

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

var (
	// lazily-initialized speaker state and shared controls
	inited   bool
	baseSR   = beep.SampleRate(44100)
	bgCtrl   *beep.Ctrl
	bgCloser io.Closer

	// small lock-protected SFX buffer cache and format info
	mu       sync.RWMutex
	sfxCache = map[string]*beep.Buffer{}
	sfxFmt   = beep.Format{}

	// volumes in effects.Volume "Volume" units (log scale, base 2)
	masterVol = -1.0 // background music volume
	sfxVol    = 0.0  // one-shot SFX volume
)

// Init ensures the audio speaker is ready. Safe to call multiple times.
func Init() error {
	if inited {
		return nil
	}
	if err := speaker.Init(baseSR, baseSR.N(50*time.Millisecond)); err != nil {
		return err
	}
	inited = true
	return nil
}

// decode opens and decodes a supported audio file, returning a seekable
// streamer, its format, and a closer you must Close when finished.
func decode(path string) (beep.StreamSeekCloser, beep.Format, io.Closer, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, beep.Format{}, nil, err
	}
	switch ext := filepath.Ext(path); ext {
	case ".wav", ".WAV":
		s, format, err := wav.Decode(f)
		if err != nil {
			_ = f.Close()
			return nil, beep.Format{}, nil, err
		}
		return s, format, s, nil
	case ".mp3", ".MP3":
		s, format, err := mp3.Decode(f)
		if err != nil {
			_ = f.Close()
			return nil, beep.Format{}, nil, err
		}
		return s, format, s, nil
	default:
		_ = f.Close()
		return nil, beep.Format{}, nil, fmt.Errorf("unsupported audio type: %s", ext)
	}
}

// resampleToBase converts a stream to the engine's base sample rate if needed.
func resampleToBase(streamer beep.Streamer, format beep.Format) beep.Streamer {
	if format.SampleRate == baseSR {
		return streamer
	}
	// quality 4 is a good trade-off for game audio
	return beep.Resample(4, format.SampleRate, baseSR, streamer)
}

// PreloadSFX decodes an effect and stores it in memory under the given id.
// If already cached, it returns immediately.
func PreloadSFX(id, path string) error {
	if err := Init(); err != nil {
		return err
	}
	mu.RLock()
	if _, ok := sfxCache[id]; ok {
		mu.RUnlock()
		return nil
	}
	mu.RUnlock()

	stream, format, closer, err := decode(path)
	if err != nil {
		return err
	}
	defer closer.Close()

	rs := resampleToBase(stream, format)

	// normalize buffer format to our base sample rate
	bufFmt := beep.Format{
		SampleRate:  baseSR,
		NumChannels: format.NumChannels,
		Precision:   format.Precision,
	}
	buf := beep.NewBuffer(bufFmt)
	buf.Append(rs)

	mu.Lock()
	sfxCache[id] = buf
	if sfxFmt.SampleRate == 0 {
		sfxFmt = bufFmt
	}
	mu.Unlock()
	return nil
}

// PlaySFXCached plays a preloaded SFX by id. It does nothing if the id was not cached.
func PlaySFXCached(id string) error {
	if err := Init(); err != nil {
		return err
	}
	mu.RLock()
	buf, ok := sfxCache[id]
	mu.RUnlock()
	if !ok {
		return fmt.Errorf("sfx not preloaded: %s", id)
	}
	stream := buf.Streamer(0, buf.Len())
	vol := &effects.Volume{Streamer: stream, Base: 2, Volume: sfxVol}
	speaker.Play(vol)
	return nil
}

// PlaySFX plays an SFX from disk, preloading it into the cache on first use.
func PlaySFX(path string) error {
	if err := PreloadSFX(path, path); err != nil {
		return err
	}
	return PlaySFXCached(path)
}

// PlayMusicLoop starts background music from the given file and loops it.
// Any currently playing music is stopped first.
func PlayMusicLoop(path string) error {
	if err := Init(); err != nil {
		return err
	}
	StopMusic()

	stream, format, closer, err := decode(path)
	if err != nil {
		return err
	}
	bgCloser = closer

	loop := beep.Loop(-1, stream)
	rs := resampleToBase(loop, format)

	vol := &effects.Volume{Streamer: rs, Base: 2, Volume: masterVol}
	bgCtrl = &beep.Ctrl{Streamer: vol}
	speaker.Play(bgCtrl)
	return nil
}

// StopMusic halts the current background track and releases its resources.
func StopMusic() {
	if bgCtrl != nil {
		// Lock to avoid races with the audio thread.
		speaker.Lock()
		bgCtrl.Paused = true
		bgCtrl = nil
		speaker.Unlock()
	}
	if bgCloser != nil {
		_ = bgCloser.Close()
		bgCloser = nil
	}
}
