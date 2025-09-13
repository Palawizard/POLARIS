package audiosystem

import (
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

var (
	inited   bool
	baseSR   = beep.SampleRate(44100)
	bgCtrl   *beep.Ctrl
	bgCloser io.Closer
)

const masterVol = -1.7

func Init() error {
	if inited {
		return nil
	}
	if err := speaker.Init(baseSR, baseSR.N(time.Second/10)); err != nil {
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
	ext := filepath.Ext(path)
	switch ext {
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
		return nil, beep.Format{}, nil, err
	}
}

func resampleToBase(streamer beep.Streamer, format beep.Format) beep.Streamer {
	if format.SampleRate == baseSR {
		return streamer
	}
	return beep.Resample(4, format.SampleRate, baseSR, streamer)
}

func PlaySFX(path string) error {
	if err := Init(); err != nil {
		return err
	}
	stream, format, closer, err := decode(path)
	if err != nil {
		return err
	}
	rs := resampleToBase(stream, format)
	vol := &effects.Volume{Streamer: rs, Base: 2, Volume: masterVol}
	speaker.Play(beep.Seq(vol, beep.Callback(func() { _ = closer.Close() })))
	return nil
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
