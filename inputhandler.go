package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (player *Player) HandleInput() {
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		player.SetAction(SailNW)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		player.SetAction(SailN)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyE) {
		player.SetAction(SailNE)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		player.SetAction(SailW)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		player.SetAction(SailE)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyZ) {
		player.SetAction(SailSW)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyX) {
		player.SetAction(SailS)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		player.SetAction(SailSE)
	} else {
		player.SetAction(NoAction)
	}
}
