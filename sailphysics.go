package main

import (
	"fmt"
	"strconv"
)

const NDir = 0
const NEDir = 1
const EDir = 2
const SEDir = 3
const SDir = 4
const SWDir = 5
const WDir = 6
const NWDir = 7

const IntoWind = 0
const CloseHaul = 1
const BeamReach = 2
const BroadReach = 3
const RunDownwind = 4

const TurningPort = 0
const HeadingStraight = 1
const TurningStarboard = 2

const FeebleBreeze = 1
const LightWinds = 2
const SteadyGusts = 3
const HeavySqualls = 4

type Wind struct {
	windSpeed     int
	WindDirection int
}

func (w Wind) DisplayWind() {
	fmt.Println("Wind Strength: " + strconv.Itoa(w.windSpeed) + " Wind Direction: " + w.Direction())
}

func (w Wind) Direction() string {
	//wind angle uses a clock style system
	//0 = N, 1 = NE, 2 = E, 3 = SE, 4 = S, 5 = SW, 6 = W, 7 = NW
	switch w.WindDirection {
	case NDir:
		return "N"
	case NEDir:
		return "NE"
	case EDir:
		return "E"
	case SEDir:
		return "SE"
	case SDir:
		return "S"
	case SWDir:
		return "SW"
	case WDir:
		return "W"
	case NWDir:
		return "NW"
	}
	return ""
}

func IsTurnWindward(previousSailDir int, presentSailDir int, windDir int) bool {
	if (presentSailDir < presentSailDir) || (presentSailDir == NWDir && previousSailDir == NDir) {
		//checking to see if the turn was to Portside
		if presentSailDir < windDir || (presentSailDir == NWDir && windDir == NDir) {
			return true
		}
		return false
	} else {
		if presentSailDir > windDir || (presentSailDir == NDir && windDir == NWDir) {
			return true
		}
		return false
	}
}

func PointsOfSail(sailDir int, windDir int) int {
	switch sailDir {
	case NDir:
		if windDir == NDir {
			return IntoWind
		} else if windDir == NEDir || windDir == NWDir {
			return CloseHaul
		} else if windDir == EDir || windDir == WDir {
			return BeamReach
		} else if windDir == SEDir || windDir == SWDir {
			return BroadReach
		} else {
			return RunDownwind
		}
	case NEDir:
		if windDir == NEDir {
			return IntoWind
		} else if windDir == NDir || windDir == EDir {
			return CloseHaul
		} else if windDir == SEDir || windDir == NWDir {
			return BeamReach
		} else if windDir == SDir || windDir == WDir {
			return BroadReach
		} else {
			return RunDownwind
		}
	case EDir:
		if windDir == EDir {
			return IntoWind
		} else if windDir == NEDir || windDir == SEDir {
			return CloseHaul
		} else if windDir == NDir || windDir == SDir {
			return BeamReach
		} else if windDir == SWDir || windDir == NWDir {
			return BroadReach
		} else {
			return RunDownwind
		}
	case SEDir:
		if windDir == SEDir {
			return IntoWind
		} else if windDir == EDir || windDir == SDir {
			return CloseHaul
		} else if windDir == NEDir || windDir == SWDir {
			return BeamReach
		} else if windDir == NDir || windDir == WDir {
			return BroadReach
		} else {
			return RunDownwind
		}
	case SDir:
		if windDir == SDir {
			return IntoWind
		} else if windDir == SEDir || windDir == SWDir {
			return CloseHaul
		} else if windDir == EDir || windDir == WDir {
			return BeamReach
		} else if windDir == NEDir || windDir == NWDir {
			return BroadReach
		} else {
			return RunDownwind
		}
	case SWDir:
		if windDir == SWDir {
			return IntoWind
		} else if windDir == SDir || windDir == WDir {
			return CloseHaul
		} else if windDir == SEDir || windDir == NWDir {
			return BeamReach
		} else if windDir == EDir || windDir == NDir {
			return BroadReach
		} else {
			return RunDownwind
		}
	case WDir:
		if windDir == WDir {
			return IntoWind
		} else if windDir == SWDir || windDir == NWDir {
			return CloseHaul
		} else if windDir == SDir || windDir == NDir {
			return BeamReach
		} else if windDir == SEDir || windDir == NEDir {
			return BroadReach
		} else {
			return RunDownwind
		}
	case NWDir:
		if windDir == NWDir {
			return IntoWind
		} else if windDir == WDir || windDir == NDir {
			return CloseHaul
		} else if windDir == SWDir || windDir == NEDir {
			return BeamReach
		} else if windDir == EDir || windDir == SDir {
			return BroadReach
		} else {
			return RunDownwind
		}
	}
	return 99
}
