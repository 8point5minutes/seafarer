package main

type World struct {
	CurrentWind  Wind
	CurrentLevel Level
}

func NewWorld() *World {
	wind := Wind{windSpeed: GetDiceRoll(4), WindDirection: GetRandomInt(8)}
	w := &World{CurrentWind: wind, CurrentLevel: NewLevel()}
	w.CurrentWind.DisplayWind()
	return w
}
