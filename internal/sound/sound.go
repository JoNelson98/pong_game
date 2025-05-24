package sound

import (
	"bytes"
	_ "embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

//go:embed assets/electric_beep.wav
var beepWav []byte

var (
	soundVolume float64= 0.2
	audioContext *audio.Context
	beepPlayer *audio.Player
)

func Init() {
	const sampleRate = 44100
	audioContext = audio.NewContext(sampleRate)

	stream, err := wav.DecodeWithSampleRate(sampleRate,bytes.NewReader(beepWav))
	if err != nil {
		log.Printf("Failed to decode audio: %v", err)
		return
	}

	beepPlayer, err = audio.NewPlayer(audioContext,stream)
	if err != nil {
		log.Printf("Failed to create audio player: %v",err)
		return	
	}
}

func PlayBeep() {
	if beepPlayer == nil {
		return
	}
	beepPlayer.SetVolume((soundVolume))
	beepPlayer.Rewind()
	beepPlayer.Play()
}