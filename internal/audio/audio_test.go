// Copyright (c) 2025 Elian Waeber & Valentin Roch
// SPDX-License-Identifier: Apache-2.0

package audio

import (
	"testing"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

var ctx *audio.Context

// TestAudioContext tests if the audio context is created without any errors.
func TestNewAudioContext(t *testing.T) {
	ctx = audio.NewContext(SampleRate)
	if ctx == nil {
		t.Error("Audio context failed to initialize")
	}
}

// TestNewAudioPlayer tests the NewAudioPlayer function.
// Checks if the AudioPlayer is created and not nil.
func TestNewAudioPlayer(t *testing.T) {
	ap, err := NewAudioPlayer(ctx)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if ap == nil {
		t.Fatalf("Expected non-nil AudioPlayer, got nil")
	}
}

// TestAudioPlayer_Play tests the Play method of the AudioPlayer.
// Checks if the player is playing after calling the Play method.
func TestAudioPlayer_Play(t *testing.T) {
	ap, err := NewAudioPlayer(ctx)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	ap.Play()
	if !ap.player.IsPlaying() {
		t.Fatalf("Expected player to be playing, but it is not")
	}
}

// TestAudioPlayer_Restart tests the Restart method of the AudioPlayer.
// Checks if the player is paused after calling the Restart method.
func TestAudioPlayer_Restart(t *testing.T) {
	ap, err := NewAudioPlayer(ctx)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	ap.Play()
	err = ap.Restart()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if ap.player.IsPlaying() {
		t.Fatalf("Expected player to be paused, but it is playing")
	}
}

// TestAudioPlayer_Close tests the Close method of the AudioPlayer.
// Checks if the player is closed without any errors.
func TestAudioPlayer_Close(t *testing.T) {
	ap, err := NewAudioPlayer(ctx)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	err = ap.Close()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}
