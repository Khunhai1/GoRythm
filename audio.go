package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
)

const (
	sampleRate = 44100
)

type AudioPlayer struct {
	context *audio.Context
	player  *audio.Player
}

func NewAudioPlayer(music int) (*AudioPlayer, error) {
	// Create audio context
	context := audio.NewContext(sampleRate)

	// Load MP3 file from "music" parameter
	mp3Data, err := os.ReadFile(fmt.Sprintf("audio/%d.mp3", music))
	if err != nil {
		return nil, fmt.Errorf("failed to read audio file: %w", err)
	}

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
