package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	playing       string
	pointsO       int
	pointsX       int
	win           string
	alter         int
	difficulty    int
	countdownTime time.Time
	countdown     int
}

const (
	StateMenu = iota
	StateLoading
	StatePlaying
	StatePause
	StateGameOver
)

var gameState = StateMenu

// Handle game logic
func (g *Game) Update() error {
	switch gameState {
	case StateMenu:
		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			gameState = StateLoading
			g.countdownTime = time.Now()
			g.countdown = 3
		}
		if ebiten.IsKeyPressed(ebiten.Key1) {
			g.difficulty = 1
		}
		if ebiten.IsKeyPressed(ebiten.Key2) {
			g.difficulty = 2
		}
		if ebiten.IsKeyPressed(ebiten.Key3) {
			g.difficulty = 3
		}
	case StateLoading:
		if g.countdown > 0 {
			elapsed := time.Since(g.countdownTime)
			if elapsed >= time.Second {
				g.countdownTime = time.Now()
				g.countdown--
			}
		} else {
			gameState = StatePlaying
		}
	}
	return nil
}

func (g *Game) Layout(int, int) (int, int) {
	return sWidth, sHeight
}
