package main

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	playing       string
	board         [3][3]string
	pointsO       int
	pointsX       int
	win           string
	alter         int
	difficulty    int
	countdownTime time.Time
	countdown     int
	rounds        int
	player        string
	audioPlayer   *AudioPlayer
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
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			gameState = StateLoading
			g.countdownTime = time.Now()
			g.countdown = 3
		}
		if inpututil.IsKeyJustPressed(ebiten.Key1) {
			g.difficulty = 1
		}
		if inpututil.IsKeyJustPressed(ebiten.Key2) {
			g.difficulty = 2
		}
		if inpututil.IsKeyJustPressed(ebiten.Key3) {
			g.difficulty = 3
		}
	case StateLoading:
		g.player = "human"
		if g.countdown > 0 {
			elapsed := time.Since(g.countdownTime)
			if elapsed >= time.Second {
				g.countdownTime = time.Now()
				g.countdown--
			}
		} else {
			gameState = StatePlaying
			if g.audioPlayer != nil {
				g.audioPlayer.Play()
			} else {
				return fmt.Errorf("audio player is nil")
			}
		}
	case StatePlaying:
		if g.player == "ai" {
			if g.difficulty == 1 {
				g.EasyCpu()
				g.player = "human"
			}
		} else {
			if inpututil.IsKeyJustPressed(ebiten.KeyKP1) {
				g.placeSymbol(0, 320)
			}
			if inpututil.IsKeyJustPressed(ebiten.KeyKP2) {
				g.placeSymbol(160, 320)
			}
			if inpututil.IsKeyJustPressed(ebiten.KeyKP3) {
				g.placeSymbol(320, 320)
			}
			if inpututil.IsKeyJustPressed(ebiten.KeyKP4) {
				g.placeSymbol(0, 160)
			}
			if inpututil.IsKeyJustPressed(ebiten.KeyKP5) {
				g.placeSymbol(160, 160)
			}
			if inpututil.IsKeyJustPressed(ebiten.KeyKP6) {
				g.placeSymbol(320, 160)
			}
			if inpututil.IsKeyJustPressed(ebiten.KeyKP7) {
				g.placeSymbol(0, 0)
			}
			if inpututil.IsKeyJustPressed(ebiten.KeyKP8) {
				g.placeSymbol(160, 0)
			}
			if inpututil.IsKeyJustPressed(ebiten.KeyKP9) {
				g.placeSymbol(320, 0)
			}
		}
		g.CheckWin()
		if g.win != "" {
			gameState = StateGameOver
			if g.win == "O" {
				g.pointsO++
			} else {
				g.pointsX++
			}
		}
		if g.rounds == 9 && g.win == "" {
			gameState = StateGameOver
		}
	case StateGameOver:
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			err := g.Init()
			if err != nil {
				return err
			}
			g.board = [3][3]string{}
			g.rounds = 0
			g.win = ""
			gameState = StateMenu
		}
	}
	return nil
}

func (g *Game) Layout(int, int) (int, int) {
	return sWidth, sHeight
}
