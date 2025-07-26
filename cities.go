package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type City struct {
	Name   string
	X      int
	Y      int
	Images []*ebiten.Image
}
