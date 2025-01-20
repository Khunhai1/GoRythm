// Copyright (c) 2025 Elian Waeber & Valentin Roch
// SPDX-License-Identifier: Apache-2.0

package generation

import (
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

// TestGenerateBoard tests the GenerateBoard function.
// Checks if the board is generated correctly.
// Checks if the board dimensions are correct.
func TestGenerateBoard(t *testing.T) {
	size := 480
	screen := ebiten.NewImage(size, size)
	board := GenerateBoard(screen, size)

	if board == nil {
		t.Error("Expected board to be generated, got nil")
	}

	if board.Bounds().Dx() != size || board.Bounds().Dy() != size {
		t.Errorf("Expected board dimensions to be %dx%d, got %dx%d", size, size, board.Bounds().Dx(), board.Bounds().Dy())
	}
}

// TestGenerateSymbols tests the GenerateSymbols function.
// Checks if all symbols are generated correctly.
// Checks if the symbols dimensions are correct.
func TestGenerateSymbols(t *testing.T) {
	size := 480
	screen := ebiten.NewImage(size, size)
	imageX, imageO, imageXHighlighted, imageOHighlighted, imageEmpty := GenerateSymbols(screen)

	if imageX == nil || imageO == nil || imageXHighlighted == nil || imageOHighlighted == nil || imageEmpty == nil {
		t.Error("Expected all symbols to be generated, got nil for one or more symbols")
	}

	expectedSize := EffectiveCellSize
	if imageX.Bounds().Dx() != expectedSize || imageX.Bounds().Dy() != expectedSize {
		t.Errorf("Expected imageX dimensions to be %dx%d, got %dx%d", expectedSize, expectedSize, imageX.Bounds().Dx(), imageX.Bounds().Dy())
	}

	if imageO.Bounds().Dx() != expectedSize || imageO.Bounds().Dy() != expectedSize {
		t.Errorf("Expected imageO dimensions to be %dx%d, got %dx%d", expectedSize, expectedSize, imageO.Bounds().Dx(), imageO.Bounds().Dy())
	}

	if imageXHighlighted.Bounds().Dx() != expectedSize || imageXHighlighted.Bounds().Dy() != expectedSize {
		t.Errorf("Expected imageXHighlighted dimensions to be %dx%d, got %dx%d", expectedSize, expectedSize, imageXHighlighted.Bounds().Dx(), imageXHighlighted.Bounds().Dy())
	}

	if imageOHighlighted.Bounds().Dx() != expectedSize || imageOHighlighted.Bounds().Dy() != expectedSize {
		t.Errorf("Expected imageOHighlighted dimensions to be %dx%d, got %dx%d", expectedSize, expectedSize, imageOHighlighted.Bounds().Dx(), imageOHighlighted.Bounds().Dy())
	}

	expectedEmptySize := EmptyImageSize
	if imageEmpty.Bounds().Dx() != expectedEmptySize || imageEmpty.Bounds().Dy() != expectedEmptySize {
		t.Errorf("Expected imageEmpty dimensions to be %dx%d, got %dx%d", expectedEmptySize, expectedEmptySize, imageEmpty.Bounds().Dx(), imageEmpty.Bounds().Dy())
	}
}
