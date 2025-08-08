package main

import (
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func ImageFromPath(imagePath string) *ebiten.Image {
	image, _, err := ebitenutil.NewImageFromFile(imagePath)
	if err != nil {
		log.Fatal(err)
	}
	return image
}

type TileType struct {
	Name      string
	Navigable bool
	Images    []*ebiten.Image
}

func NewTileType(name string, navigable bool) TileType {
	tileType := TileType{Name: name, Navigable: navigable, Images: make([]*ebiten.Image, 0)}
	return tileType
}

func (tt *TileType) NewImage(imagePath string) {
	tt.Images = append(tt.Images, ImageFromPath(imagePath))
}

func (tt *TileType) LoadVariations(imagePath string, num_variations int) {
	for x := 0; x < num_variations; x++ {
		FilePath := imagePath + strconv.Itoa(x) + ".png"
		tt.Images = append(tt.Images, ImageFromPath(FilePath))
	}
}

type MapTile struct {
	PixelX         int
	PixelY         int
	TypeImageIndex int
	Type           TileType
	City           *City
}

func NewTile(x int, y int, tileType TileType) *MapTile {
	t := &MapTile{PixelX: x, PixelY: y, TypeImageIndex: 0, Type: tileType, City: nil}
	return t
}

func (t *MapTile) RandomizeImageIndex() {
	t.TypeImageIndex = GetRandomInt(len(t.Type.Images))
}

func (t *MapTile) Image() *ebiten.Image {
	return t.Type.Images[t.TypeImageIndex]
}

func (t *MapTile) SetTileType(newType TileType) {
	t.Type = newType
	t.TypeImageIndex = 0
}

func (t *MapTile) HasCity() bool {
	return t.City != nil
}

func (t *MapTile) SetCity(city *City) {
	t.City = city
}
