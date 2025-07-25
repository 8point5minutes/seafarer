package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	Name string
	Ship *Ship
	*Actor
}

func NewPlayer(name string) *Player {
	st := NewShipType("Carrack", "assets/carrack.png")
	p := &Player{Name: name, Ship: NewShip("Terrible", st), Actor: NewActor()}
	return p
}
func (player *Player) PlayerRender(screen *ebiten.Image) {
	gd := NewGameData()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(player.X*gd.TileWidth), float64(player.Y*gd.TileHeight))
	screen.DrawImage(player.Ship.Type.Model, op)
}
