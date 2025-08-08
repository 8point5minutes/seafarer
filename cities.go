package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type City struct {
	Name   string
	X      int
	Y      int
	Sprite *ebiten.Image
}

func NewCity(cityName string, sprite *ebiten.Image) *City {
	c := &City{Name: cityName, X: 0, Y: 0, Sprite: sprite}
	return c
}

func (c *City) SetCityPos(x int, y int) {
	c.X = x
	c.Y = y
}

func (c *City) CityRender(screen *ebiten.Image) {
	gd := NewGameData()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.X*gd.TileWidth), float64(c.Y*gd.TileHeight))
	screen.DrawImage(c.Sprite, op)
}
