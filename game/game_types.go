// Copyright (c) 2025 Elian Waeber & Valentin Roch
// SPDX-License-Identifier: Apache-2.0

package game

// A GameState type represent the different states of a Game.
type GameState int

const (
	StateMenu GameState = iota
	StateLoading
	StatePlaying
	StatePause
	StateGameOver
)

// A GamePlayer type represent the different type of players of a Game.
type PlayerType string

const (
	NO_PLAYER  PlayerType = ""
	HUMAN_TYPE PlayerType = "human"
	AI_TYPE    PlayerType = "ai"
)

// A GamePlaying type represent which side is playing.
type SymbolPlaying string

const (
	NONE_PLAYING SymbolPlaying = ""
	X_PLAYING    SymbolPlaying = "X"
	O_PLAYING    SymbolPlaying = "O"
)

// A GameMode type represent the different game modes of a Game.
type GameMode int

const (
	NO_MODE GameMode = iota
	EASY_AI_MODE
	HARD_AI_MODE
	GORYTHM_MODE
)
