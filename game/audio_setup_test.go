// Copyright (c) 2025 Elian Waeber & Valentin Roch
// SPDX-License-Identifier: Apache-2.0

package game

import (
	a "GoRythm/internal/audio"
	"testing"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

var audioContext *audio.Context

// setup initializes the audio context with the specified sample rate
func setup() {
	audioContext = audio.NewContext(a.SampleRate)
}

// TestMain sets up the audio context before running tests
func TestMain(m *testing.M) {
	setup()
	m.Run()
}

// TestAudioContext tests if the audio context is initialized correctly
func TestAudioContext(t *testing.T) {
	if audioContext == nil {
		t.Error("Audio context failed to initialize")
	}
}
