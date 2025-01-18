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

//go:embed assets/audio/1.mp3
var mp3Data []byte

type AudioPlayer struct {
	context *audio.Context
	player  *audio.Player
}

func NewAudioPlayer(ctx *audio.Context) (*AudioPlayer, error) {
	// Decode MP3 file
	stream, err := mp3.DecodeWithSampleRate(sampleRate, bytes.NewReader(mp3Data))
	if err != nil {
		return nil, fmt.Errorf("failed to decode mp3 file: %w", err)
	}

	// Create audio player
	player, err := audioContext.NewPlayer(stream)
	if err != nil {
		return nil, fmt.Errorf("failed to create audio player: %w", err)
	}

	// Set volume
	player.SetVolume(0.05)

	// Create AudioPlayer instance
	ap := &AudioPlayer{
		context: audioContext,
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
