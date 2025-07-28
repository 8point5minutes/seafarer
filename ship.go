package main

import "github.com/hajimehoshi/ebiten/v2"

type ShipType struct {
	Name  string
	Model *ebiten.Image
}

type Ship struct {
	Type ShipType
	Name string
	X    int
	Y    int
}

func NewShipType(name string, modelPath string) ShipType {
	s := ShipType{Name: name, Model: ImageFromPath(modelPath)}
	return s
}

func NewShip(name string, shipType ShipType) *Ship {
	s := &Ship{Name: name, Type: shipType, X: 5, Y: 6}
	return s
}
