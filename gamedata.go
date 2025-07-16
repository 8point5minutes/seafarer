package main

type GameData struct {
	ScreenWidth  int
	ScreenHeight int
	TileWidth    int
	TileHeight   int
}

func NewGameData() GameData {
	g := GameData{
		ScreenWidth:  40,
		ScreenHeight: 30,
		TileWidth:    16,
		TileHeight:   16,
	}
	return g
}
