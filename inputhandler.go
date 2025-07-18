package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (actor *Actor) HandleInput() {
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		fmt.Println("Northwest")
		actor.SetAction(SailNW)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		fmt.Println("North")
		actor.SetAction(SailN)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyE) {
		fmt.Println("Northeast")
		actor.SetAction(SailNE)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		fmt.Println("West")
		actor.SetAction(SailW)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		fmt.Println("East")
		actor.SetAction(SailE)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyZ) {
		fmt.Println("Southwest")
		actor.SetAction(SailSW)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyX) {
		fmt.Println("South")
		actor.SetAction(SailS)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		fmt.Println("Southeast")
		actor.SetAction(SailSE)
	} else {
		actor.SetAction(NoAction)
	}
}
