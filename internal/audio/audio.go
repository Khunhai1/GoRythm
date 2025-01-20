// Copyright (c) 2025 Elian Waeber & Valentin Roch
// SPDX-License-Identifier: Apache-2.0

// Package audio provides audio player feature for playing music in the GoRythm
// game as well as the beatmap struct, which contains the time and beat number
// of each beat.
package audio

import (
	"bytes"
	_ "embed"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
)

const (
	SampleRate = 44100
	Volume     = 0.05
)

//go:embed assets/audio/1.mp3
var mp3Data []byte

// AudioPlayer is a struct that contains the audio context and player for the game.
type AudioPlayer struct {
	context *audio.Context
	player  *audio.Player
}

// NewAudioPlayer creates a new AudioPlayer instance with the given audio context.
// It decodes the MP3 file and create a player with it.
func NewAudioPlayer(ctx *audio.Context) (*AudioPlayer, error) {
	// Decode MP3 file
	stream, err := mp3.DecodeWithSampleRate(SampleRate, bytes.NewReader(mp3Data))
	if err != nil {
		return nil, fmt.Errorf("failed to decode mp3 file: %w", err)
	}

	// Create audio player
	player, err := ctx.NewPlayer(stream)
	if err != nil {
		return nil, fmt.Errorf("failed to create audio player: %w", err)
	}

	// Create AudioPlayer instance
	ap := &AudioPlayer{
		context: ctx,
		player:  player,
	}

	return ap, nil
}

// Play plays the audio player with a fixed volume.
func (ap *AudioPlayer) Play() {
	ap.player.SetVolume(Volume)
	ap.player.Play()
}

// Restart stops the audio player, rewinds it to the beginning and pauses it.
func (ap *AudioPlayer) Restart() error {
	if err := ap.player.Rewind(); err != nil {
		return err
	}
	ap.player.Pause()
	return nil
}

// Close closes the audio player.
func (ap *AudioPlayer) Close() error {
	return ap.player.Close()
}

// IsPlaying returns whether the audio player is playing or not.
func (ap *AudioPlayer) IsPlaying() bool {
	return ap.player.IsPlaying()
}
