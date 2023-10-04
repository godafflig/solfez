package utils

import (
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

type audioPanel struct {
	sampleRate beep.SampleRate
	streamer   beep.StreamSeeker
	ctrl       *beep.Ctrl
	resampler  *beep.Resampler
	volume     *effects.Volume
}

func newAudioPanel(sampleRate beep.SampleRate, streamer beep.StreamSeeker) *audioPanel {
	ctrl := &beep.Ctrl{Streamer: beep.Loop(-1, streamer)}
	resampler := beep.ResampleRatio(4, 1, ctrl)
	volume := &effects.Volume{Streamer: resampler, Base: 2}
	return &audioPanel{sampleRate, streamer, ctrl, resampler, volume}
}

func (ap *audioPanel) play() {
	speaker.Play(ap.volume)
}

func Game() {

	soundFile, err := os.Open("./notes/C6.mp3")
	if err != nil {
		log.Fatal(err)
	}
	defer soundFile.Close()

	// Decode the sound file.
	streamer, format, err := mp3.Decode(soundFile)
	if err != nil {
		log.Fatal(err)
	}

	err1 := speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/30))
	if err1 != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	ap := newAudioPanel(format.SampleRate, streamer)

	// Play the sound.
	done := make(chan bool)
	go func() {

		ap.play()
		log.Print("PLAY sound")
		if err != nil {
			log.Fatal(err)
		}
		done <- true
	}()

	// Wait for the sound to finish playing.
	<-done
}
