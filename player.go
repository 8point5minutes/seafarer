package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Player struct {
	Name string
	*Ship
	CurrentAction func(Actor) func(*World) bool
	CurrentMenu   int
}

func NewPlayer(name string) *Player {
	st := NewShipType("Carrack", "assets/carrack.png", 3, 2)
	p := &Player{Name: name, Ship: NewShip("Terrible", st), CurrentAction: NoAction, CurrentMenu: 0}
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
	if player.CurrentMenu == 1 {
		vector.DrawFilledRect(screen, 20, 120, 280, 60, color.NRGBA{0x60, 0x30, 0x00, 0xff}, false)
	}
}
