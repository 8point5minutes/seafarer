package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Level struct {
	Tiles []*MapTile
}

func NewLevel() Level {
	l := Level{}
	l.Tiles = l.CreateTiles()
	return l
}

// GetIndexFromXY gets the index of the map array from a given X,Y TILE coordinate.
// This coordinate is logical tiles, not pixels.
func (level *Level) GetIndexFromXY(x int, y int) int {
	gd := NewGameData()
	return (y * gd.ScreenWidth) + x
}

func (level *Level) GetTileFromIndex(x int, y int) *MapTile {
	t := level.Tiles[level.GetIndexFromXY(x, y)]
	return t
}

func (level *Level) CreateTiles() []*MapTile {
	gd := NewGameData()
	tiles := make([]*MapTile, 0)
	//grass := NewTileType(false, "grass.png")
	sea := NewTileType(true)
	sea.NewImage("sea1.png")
	sea.NewImage("sea2.png")
	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			tile := NewTile(x*gd.TileHeight, y*gd.TileWidth, sea)
			tile.RandomizeTypeIndex()
			tiles = append(tiles, tile)
		}
	}

	return tiles
}

func (level *Level) DrawLevel(screen *ebiten.Image) {
	gd := NewGameData()
	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			tile := level.Tiles[level.GetIndexFromXY(x, y)]
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
			screen.DrawImage(tile.Image(), op)
		}
	}
}
