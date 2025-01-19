package game

import (
	"testing"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

var audioContext *audio.Context

// setup initializes the audio context with the specified sample rate
func setup() {
	audioContext = audio.NewContext(sampleRate)
}

// TestMain sets up the audio context before running tests
func TestMain(m *testing.M) {
	setup()
	m.Run()
}
