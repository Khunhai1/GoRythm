package main

import (
	"math/rand"
	"time"
)

func newRandom() *rand.Rand {
	s1 := rand.NewSource(time.Now().UnixNano())
	return rand.New(s1)
}

func (g *Game) ResetPoints() {
	g.pointsO = 0
	g.pointsX = 0
}
