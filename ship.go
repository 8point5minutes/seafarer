package main

import "github.com/hajimehoshi/ebiten/v2"

type ShipType struct {
	Name          string
	Model         *ebiten.Image
	MaxMovePoints int
	MaxTurnSpeed  int
}

type Ship struct {
	Type              ShipType
	Name              string
	X                 int
	Y                 int
	RudderPosition    int
	CurrentDirection  int
	CurrentMovePoints int
	Anchored          bool
}

func NewShipType(name string, modelPath string, movePoints int, turnSpeed int) ShipType {
	s := ShipType{Name: name, Model: ImageFromPath(modelPath), MaxMovePoints: movePoints, MaxTurnSpeed: turnSpeed}
	return s
}

func NewShip(name string, shipType ShipType) *Ship {
	s := &Ship{Name: name, Type: shipType, X: 5, Y: 6, CurrentDirection: NDir, RudderPosition: HeadingStraight, Anchored: false}
	return s
}

func (s *Ship) PortTurn() {
	if s.CurrentDirection == NDir {
		s.CurrentDirection = NWDir
	} else {
		s.CurrentDirection--
	}
}

func (s *Ship) StarboardTurn() {
	if s.CurrentDirection == NWDir {
		s.CurrentDirection = NDir
	} else {
		s.CurrentDirection++
	}
}

func (s *Ship) Move() {
	gd := NewGameData()
	switch s.CurrentDirection {
	case NDir:
		if s.Y > 0 {
			s.Y--
		}
	case NEDir:
		if s.Y > 0 {
			s.Y--
		}
		if s.X < gd.ScreenWidth-1 {
			s.X++
		}
	case EDir:
		if s.X < gd.ScreenWidth-1 {
			s.X++
		}
	case SEDir:
		if s.Y < gd.ScreenHeight-1 {
			s.Y++
		}
		if s.X < gd.ScreenWidth-1 {
			s.X++
		}
	case SDir:
		if s.Y < gd.ScreenHeight-1 {
			s.Y++
		}
	case SWDir:
		if s.Y < gd.ScreenHeight-1 {
			s.Y++
		}
		if s.X > 0 {
			s.X--
		}
	case WDir:
		if s.X > 0 {
			s.X--
		}
	case NWDir:
		if s.X > 0 {
			s.X--
		}
		if s.Y > 0 {
			s.Y--
		}
	}
}
