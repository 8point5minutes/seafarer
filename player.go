package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Player struct {
	Name string
	X    int
	Y    int
	Ship *Ship
}

func NewPlayer(name string) *Player {
	p := &Player{Name: name, X: 20, Y: 6, Ship: NewShip("Terrible", "carrack.png")}
	return p
}

func (player *Player) PlayerRender(screen *ebiten.Image) {
	gd := NewGameData()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(player.X*gd.TileHeight), float64(player.Y*gd.TileWidth))
	screen.DrawImage(player.Ship.Model, op)

}

type Ship struct {
	Name  string
	Model *ebiten.Image
}

func NewShip(name string, modelPath string) *Ship {
	s := &Ship{Name: name, Model: ImageFromPath(modelPath)}
	return s
}

func (player *Player) Move(level Level) {
	gd := NewGameData()
	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		if player.Y > 0 {
			nextTile := level.GetTileFromIndex(player.X, player.Y-1)
			if nextTile.Type.Navigable {
				player.Y--

			} else {
				fmt.Println("Your ship cannot enter land !")
			}

		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		if player.X > 0 {
			nextTile := level.GetTileFromIndex(player.X-1, player.Y)
			if nextTile.Type.Navigable {
				player.X--

			} else {
				fmt.Println("Your ship cannot enter land !")
			}

		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		if player.Y < gd.ScreenHeight {
			nextTile := level.GetTileFromIndex(player.X, player.Y+1)
			if nextTile.Type.Navigable {
				player.Y++

			} else {
				fmt.Println("Your ship cannot enter land !")
			}

		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		if player.X < gd.ScreenWidth {
			nextTile := level.GetTileFromIndex(player.X+1, player.Y)
			if nextTile.Type.Navigable {
				player.X++

			} else {
				fmt.Println("Your ship cannot enter land !")
			}

		}
	}
}
