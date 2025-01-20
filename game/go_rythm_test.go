// Copyright (c) 2025 Elian Waeber & Valentin Roch
// SPDX-License-Identifier: Apache-2.0

package game

import (
	"GoRythm/internal/audio"
	"testing"
	"time"
)

// TestNewGoRythm tests the NewGoRythm function.
// Checks if the created GoRythm instance is not nil and have the expected default attributes.
func TestNewGoRythm(t *testing.T) {
	gr := NewGoRythm()
	if gr == nil {
		t.Fatal("Expected GoRythm instance, got nil")
	}
	if len(gr.movesO) != 0 || len(gr.movesX) != 0 {
		t.Fatal("Expected empty move channels")
	}
	if gr.toBeRemovedO != noMove || gr.toBeRemovedX != noMove {
		t.Fatal("Expected noMove for toBeRemovedO and toBeRemovedX")
	}
	if !gr.startTime.IsZero() || !gr.circleColorChangeTime.IsZero() {
		t.Fatal("Expected zero startTime and circleColorChangeTime")
	}
}

// TestStart tests the Start function.
// Checks if the start time is set correctly.
func TestStart(t *testing.T) {
	gr := NewGoRythm()
	startTime := time.Now()
	gr.Start(startTime)
	if gr.startTime != startTime {
		t.Fatalf("Expected startTime %v, got %v", startTime, gr.startTime)
	}
}

// TestUpdate tests the Update function.
// Checks if the function returns the expected values for remove, highlight, toRemove, and toHighlight.
func TestUpdate(t *testing.T) {
	gr := NewGoRythm()
	gr.Start(time.Now())

	// First move
	remove, highlight, toBeRemoved, toBeHighlighted := gr.Update(X_PLAYING, 1, 1)
	if remove || highlight && toBeRemoved != noMove && toBeHighlighted != noMove {
		t.Fatal("Expected no remove or highlight on first move")
	}

	// Second move
	remove, highlight, toBeRemoved, toBeHighlighted = gr.Update(X_PLAYING, 2, 2)
	if remove || highlight && toBeRemoved != noMove && toBeHighlighted != noMove {
		t.Fatal("Expected no remove or highlight on second move")
	}

	// Third move
	remove, highlight, toBeRemoved, toBeHighlighted = gr.Update(X_PLAYING, 3, 3)
	if remove || !highlight && toBeRemoved != noMove && toBeHighlighted == noMove {
		t.Fatal("Expected not remove and highlight on third move")
	}
	if toBeHighlighted != [2]int{1, 1} {
		t.Fatalf("Expected toBeHighlighted [1, 1], got %v", toBeHighlighted)
	}

	// Fourth move
	remove, highlight, toRemove, toHighlight := gr.Update(X_PLAYING, 1, 2)
	if !remove || !highlight {
		t.Fatal("Expected remove and highlight on fourth move")
	}
	if toRemove != [2]int{1, 1} || toHighlight != [2]int{2, 2} {
		t.Fatalf("Expected toRemove [1, 1] and toHighlight [2, 2], got %v and %v", toRemove, toHighlight)
	}
}

// TestCalculateScore tests the CalculateScore function.
// Checks if the score is perfect, good, ok, or missed based on the time the player makes a move.
func TestCalculateScore(t *testing.T) {
	gr := NewGoRythm()
	gr.Start(time.Now())

	const beatInterval float64 = 1.0

	// Mock beat map
	gr.beatMap = []audio.Beat{
		{Time: beatInterval, BeatNum: 1},
		{Time: beatInterval + 1, BeatNum: 2},
		{Time: beatInterval + 2, BeatNum: 3},
	}

	// Perfect score
	timeToSleep := beatInterval
	time.Sleep(time.Duration(timeToSleep) * time.Second)
	score := gr.CalculateScore()
	if score != perfectScore {
		t.Fatalf("Expected score %d, got %d", perfectScore, score)
	}

	// Good score
	timeToSleep = beatInterval + perfectPrec - goodPrec
	time.Sleep(time.Duration(timeToSleep) * time.Second)
	score = gr.CalculateScore()
	if score != perfectScore {
		t.Fatalf("Expected score %d, got %d", goodScore, score)
	}

	// Ok score
	timeToSleep = beatInterval + goodPrec - okPrec
	time.Sleep(time.Duration(timeToSleep) * time.Second)
	score = gr.CalculateScore()
	if score != perfectScore {
		t.Fatalf("Expected score %d, got %d", okScore, score)
	}
}

// TestCalculateScoreMissed tests the CalculateScore function.
// Checks if the score is 0 when the player misses a beat.
func TestCalculateScoreMissed(t *testing.T) {
	gr := NewGoRythm()
	gr.Start(time.Now())

	const beatInterval float64 = 1.0

	// Mock beat map
	gr.beatMap = []audio.Beat{
		{Time: beatInterval, BeatNum: 1},
		{Time: beatInterval + 1, BeatNum: 2},
		{Time: beatInterval + 2, BeatNum: 3},
	}

	// Missed score
	score := gr.CalculateScore()
	if score != 0 {
		t.Fatalf("Expected score 0, got %d", score)
	}
}
