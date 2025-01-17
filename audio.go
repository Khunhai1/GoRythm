package main

import (
	"bytes"
	_ "embed"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
)

const (
	sampleRate = 44100
)

//go:embed audio/1.mp3
var mp3Data []byte

type AudioPlayer struct {
	context *audio.Context
	player  *audio.Player
}

func NewAudioPlayer() (*AudioPlayer, error) {
	// Create audio context
	context := audio.NewContext(sampleRate)

	// Decode MP3 file
	stream, err := mp3.DecodeWithSampleRate(sampleRate, bytes.NewReader(mp3Data))
	if err != nil {
		return nil, fmt.Errorf("failed to decode mp3 file: %w", err)
	}

	// Create audio player
	player, err := context.NewPlayer(stream)
	if err != nil {
		return nil, fmt.Errorf("failed to create audio player: %w", err)
	}

	// Create AudioPlayer instance
	ap := &AudioPlayer{
		context: context,
		player:  player,
	}

	return ap, nil
}

func (ap *AudioPlayer) Play() {
	ap.player.Play()
}

func (ap *AudioPlayer) Stop() {
	ap.player.Close()
}
