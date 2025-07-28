package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	Name string
	*Ship
	CurrentAction func(Actor) func(*World) bool
}

func NewPlayer(name string) *Player {
	st := NewShipType("Carrack", "assets/carrack.png")
	p := &Player{Name: name, Ship: NewShip("Terrible", st), CurrentAction: NoAction}
	return p
}

func (player *Player) SetAction(newAction func(Actor) func(*World) bool) {
	player.CurrentAction = newAction
}

func (player *Player) GetShip() *Ship {
	return player.Ship
}

func (player *Player) PlayerRender(screen *ebiten.Image) {
	gd := NewGameData()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(player.Ship.X*gd.TileWidth), float64(player.Ship.Y*gd.TileHeight))
	screen.DrawImage(player.Ship.Type.Model, op)
}
