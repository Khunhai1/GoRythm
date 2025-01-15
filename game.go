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

var numpadToBoard = map[ebiten.Key][2]int{
	ebiten.KeyKP1: {0, 2},
	ebiten.KeyKP2: {1, 2},
	ebiten.KeyKP3: {2, 2},
	ebiten.KeyKP4: {0, 1},
	ebiten.KeyKP5: {1, 1},
	ebiten.KeyKP6: {2, 1},
	ebiten.KeyKP7: {0, 0},
	ebiten.KeyKP8: {1, 0},
	ebiten.KeyKP9: {2, 0},
}

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
				g.countdown--
				g.countdownTime = time.Now()
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
			var x, y int
			if g.difficulty == 1 {
				x, y = g.EasyCpu()
			} else if g.difficulty == 3 {
				x, y = g.HardCpu()
			}
			g.placeSymbol(x, y)
		} else {
			for key, pos := range numpadToBoard {
				if inpututil.IsKeyJustPressed(key) {
					x, y := pos[0], pos[1]
					if g.board[x][y] == "" {
						g.placeSymbol(x, y)
					}
				}
			}
		}
		g.win, _ = g.CheckWin()
		if g.win != "" {
			gameState = StateGameOver
			if g.win == "O" {
				g.pointsO++
			} else {
				g.pointsX++
			}
		}
		if g.IsBoardFull() {
			gameState = StateGameOver
		}
	case StateGameOver:
		g.audioPlayer.Stop() // Arrêter la musique
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			gameImage.Clear()
			boardImage.Clear()
			// Réinitialiser l'état du jeu ici
			err := g.Init()
			if err != nil {
				return err
			}
			gameState = StateMenu // Retour au menu
		}
	}

	return nil
}

func (g *Game) Layout(int, int) (int, int) {
	return sWidth, sHeight
}
