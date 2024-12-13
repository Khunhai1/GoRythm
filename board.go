package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

var gameImage = ebiten.NewImage(sWidth, sWidth)
var XImage, OImage *ebiten.Image

func (g *Game) Init() {
	re := newRandom().Intn(2)
	if re == 0 {
		g.playing = "O"
		g.alter = 0
	} else {
		g.playing = "X"
		g.alter = 1
	}
}
