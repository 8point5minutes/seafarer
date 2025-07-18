package main

type World struct {
	CurrentWind  Wind
	CurrentLevel Level
}

func NewWorld() *World {
	wind := Wind{windSpeed: 2, windDirection: 0}
	w := &World{CurrentWind: wind, CurrentLevel: NewLevel()}
	w.CurrentWind.DisplayWind()
	return w
}
