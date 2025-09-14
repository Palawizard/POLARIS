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
	inited    bool
	baseSR    = beep.SampleRate(44100)
	bgCtrl    *beep.Ctrl
	bgCloser  io.Closer
	mu        sync.RWMutex
	sfxCache  = map[string]*beep.Buffer{}
	sfxFmt    = beep.Format{}
	masterVol = -1.0
	sfxVol    = 0.0
)

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

func resampleToBase(streamer beep.Streamer, format beep.Format) beep.Streamer {
	if format.SampleRate == baseSR {
		return streamer
	}
	return beep.Resample(4, format.SampleRate, baseSR, streamer)
}

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

func PlaySFX(path string) error {
	if err := PreloadSFX(path, path); err != nil {
		return err
	}
	return PlaySFXCached(path)
}

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

func StopMusic() {
	if bgCtrl != nil {
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
