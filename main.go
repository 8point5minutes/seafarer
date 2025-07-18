package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	world  *World
	player *Player
}

func NewGame() *Game {
	g := &Game{world: NewWorld(), player: NewPlayer("Dread Pirate Roberts")}
	return g
}

// Update is called each tic.
func (g *Game) Update() error {
	p := g.player
	p.HandleInput()
	g.player.CurrentAction(g.world, p.Actor)

	return nil
}

// Draw is called each draw cycle and is where we will blit.
func (g *Game) Draw(screen *ebiten.Image) {
	w := g.world
	p := g.player
	w.CurrentLevel.DrawLevel(screen)
	p.PlayerRender(screen)
}

// Layout will return the screen dimensions.
func (g *Game) Layout(w, h int) (int, int) { return 640, 480 }

func main() {
	g := NewGame()
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Seafarer")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
