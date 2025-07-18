package main

import (
	"fmt"
	"strconv"
)

type Wind struct {
	windSpeed     int
	windDirection int
}

func (w Wind) DisplayWind() {
	fmt.Println("Wind Strength: " + strconv.Itoa(w.windSpeed) + " Wind Direction: " + w.Direction())
}

func (w Wind) Direction() string {
	//wind angle uses a clock style system
	//0 = N, 1 = NE, 2 = E, 3 = SE, 4 = S, 5 = SW, 6 = W, 7 = NW
	switch w.windDirection {
	case 0:
		return "N"
	case 1:
		return "NE"
	case 2:
		return "E"
	case 3:
		return "SE"
	case 4:
		return "S"
	case 5:
		return "SW"
	case 6:
		return "W"
	case 7:
		return "NW"
	}
	return ""
}
