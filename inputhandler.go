package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (player *Player) HandleInput() {
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		player.RudderPosition = TurningPort
		player.SetAction(Sail)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		player.RudderPosition = HeadingStraight
		player.SetAction(Sail)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyE) {
		player.RudderPosition = TurningStarboard
		player.SetAction(Sail)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		player.SetAction(Anchor)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyO) {
		if player.CurrentMenu == 0 {
			player.SetAction(OpenMenu)
		} else {
			player.SetAction(CloseMenu)
		}
	} else {
		player.SetAction(NoAction)
	}
}
