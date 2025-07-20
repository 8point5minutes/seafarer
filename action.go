package main

type Action func(world *World, actor *Actor) (turnOver bool)

type Actor struct {
	CurrentAction Action
	X             int
	Y             int
}

func NewActor() *Actor {
	a := &Actor{CurrentAction: NoAction, X: 5, Y: 5}
	return a
}

func (actor *Actor) SetAction(newAction Action) {
	actor.CurrentAction = newAction
}

func NoAction(world *World, actor *Actor) bool {
	return false
}

func SailN(world *World, actor *Actor) bool {
	speed := 0
	if world.CurrentWind.Direction() == "SW" || world.CurrentWind.Direction() == "SE" {
		//can't sail into upwind so we don't even consider it
		//broad reach (45 degree to the back) is at +1 versus wind speed
		speed = (world.CurrentWind.windSpeed + 1)
	} else if world.CurrentWind.Direction() == "E" || world.CurrentWind.Direction() == "W" {
		//beam reach (wind perpendicular) are at wind speed
		speed = world.CurrentWind.windSpeed
	} else if world.CurrentWind.Direction() == "S" || world.CurrentWind.Direction() == "NW" || world.CurrentWind.Direction() == "NE" {
		//running (wind behind) or close haul (forward and to the side) is at -1 versus wind speed
		speed = (world.CurrentWind.windSpeed - 1)
	}
	actor.Y = MaxValue(actor.Y-speed, 0)
	return true
}

func SailNE(world *World, actor *Actor) bool {
	gd := NewGameData()
	speed := 0
	if world.CurrentWind.Direction() == "W" || world.CurrentWind.Direction() == "S" {
		//can't sail into upwind so we don't even consider it
		//broad reach (45 degree to the back) is at +1 versus wind speed
		speed = (world.CurrentWind.windSpeed/2 + 1)
	} else if world.CurrentWind.Direction() == "SE" || world.CurrentWind.Direction() == "NW" {
		//beam reach (wind perpendicular) are at wind speed
		speed = (world.CurrentWind.windSpeed) / 2
	} else if world.CurrentWind.Direction() == "N" || world.CurrentWind.Direction() == "E" || world.CurrentWind.Direction() == "SW" {
		//running (wind behind) or close haul (forward and to the side) is at -1 versus wind speed
		speed = (world.CurrentWind.windSpeed - 1) / 2

	}
	actor.Y = MaxValue(actor.Y-speed, 0)
	actor.X = MinValue(actor.X+speed, gd.ScreenWidth-1)
	return true
}

func SailE(world *World, actor *Actor) bool {
	gd := NewGameData()
	speed := 0
	if world.CurrentWind.Direction() == "NW" || world.CurrentWind.Direction() == "SW" {
		//can't sail into upwind so we don't even consider it
		//broad reach (45 degree to the back) is at +1 versus wind speed
		speed = (world.CurrentWind.windSpeed + 1)
	} else if world.CurrentWind.Direction() == "N" || world.CurrentWind.Direction() == "S" {
		//beam reach (wind perpendicular) are at wind speed
		speed = world.CurrentWind.windSpeed
	} else if world.CurrentWind.Direction() == "W" || world.CurrentWind.Direction() == "NE" || world.CurrentWind.Direction() == "SE" {
		//running (wind behind) or close haul (forward and to the side) is at -1 versus wind speed
		speed = (world.CurrentWind.windSpeed - 1)
	}
	actor.X = MinValue(actor.X+speed, gd.ScreenWidth-1)
	return true
}
func SailS(world *World, actor *Actor) bool {
	gd := NewGameData()
	speed := 0
	if world.CurrentWind.Direction() == "NW" || world.CurrentWind.Direction() == "NE" {
		//can't sail into upwind so we don't even consider it
		//broad reach (45 degree to the back) is at +1 versus wind speed
		speed = (world.CurrentWind.windSpeed + 1)
	} else if world.CurrentWind.Direction() == "E" || world.CurrentWind.Direction() == "W" {
		//beam reach (wind perpendicular) are at wind speed
		speed = world.CurrentWind.windSpeed
	} else if world.CurrentWind.Direction() == "N" || world.CurrentWind.Direction() == "SW" || world.CurrentWind.Direction() == "SE" {
		//running (wind behind) or close haul (forward and to the side) is at -1 versus wind speed
		speed = (world.CurrentWind.windSpeed - 1)
	}
	actor.Y = MinValue(actor.Y+speed, gd.ScreenHeight-1)
	return true

}
func SailSE(world *World, actor *Actor) bool {
	gd := NewGameData()
	speed := 0
	if world.CurrentWind.Direction() == "N" || world.CurrentWind.Direction() == "W" {
		//can't sail into upwind so we don't even consider it
		//broad reach (45 degree to the back) is at +1 versus wind speed
		speed = (world.CurrentWind.windSpeed)/2 + 1
	} else if world.CurrentWind.Direction() == "NE" || world.CurrentWind.Direction() == "SW" {
		//beam reach (wind perpendicular) are at wind speed
		speed = world.CurrentWind.windSpeed / 2
	} else if world.CurrentWind.Direction() == "S" || world.CurrentWind.Direction() == "NW" || world.CurrentWind.Direction() == "E" {
		//running (wind behind) or close haul (forward and to the side) is at -1 versus wind speed
		speed = (world.CurrentWind.windSpeed - 1) / 2
	}
	actor.X = MinValue(actor.X+speed, gd.ScreenWidth-1)
	actor.Y = MinValue(actor.Y+speed, gd.ScreenHeight-1)
	return true
}

func SailSW(world *World, actor *Actor) bool {
	gd := NewGameData()
	speed := 0
	if world.CurrentWind.Direction() == "N" || world.CurrentWind.Direction() == "E" {
		//can't sail into upwind so we don't even consider it
		//broad reach (45 degree to the back) is at +1 versus wind speed
		speed = world.CurrentWind.windSpeed/2 + 1
	} else if world.CurrentWind.Direction() == "SE" || world.CurrentWind.Direction() == "NW" {
		//beam reach (wind perpendicular) are at wind speed
		speed = world.CurrentWind.windSpeed / 2
	} else if world.CurrentWind.Direction() == "S" || world.CurrentWind.Direction() == "NE" || world.CurrentWind.Direction() == "W" {
		//running (wind behind) or close haul (forward and to the side) is at -1 versus wind speed
		speed = (world.CurrentWind.windSpeed - 1) / 2
	}
	actor.X = MaxValue(actor.X-speed, 0)
	actor.Y = MinValue(actor.Y+speed, gd.ScreenHeight-1)
	return true
}

func SailW(world *World, actor *Actor) bool {
	speed := 0
	if world.CurrentWind.Direction() == "NE" || world.CurrentWind.Direction() == "SE" {
		//can't sail into upwind so we don't even consider it
		//broad reach (45 degree to the back) is at +1 versus wind speed
		speed = (world.CurrentWind.windSpeed + 1)
	} else if world.CurrentWind.Direction() == "N" || world.CurrentWind.Direction() == "S" {
		//beam reach (wind perpendicular) are at wind speed
		speed = world.CurrentWind.windSpeed
	} else if world.CurrentWind.Direction() == "SW" || world.CurrentWind.Direction() == "NW" || world.CurrentWind.Direction() == "E" {
		//running (wind behind) or close haul (forward and to the side) is at -1 versus wind speed
		speed = (world.CurrentWind.windSpeed - 1)
	}
	actor.X = MaxValue(actor.X-speed, 0)
	return true
}

func SailNW(world *World, actor *Actor) bool {
	speed := 0
	if world.CurrentWind.Direction() == "E" || world.CurrentWind.Direction() == "S" {
		//can't sail into upwind so we don't even consider it
		//broad reach (45 degree to the back) is at +1 versus wind speed
		speed = (world.CurrentWind.windSpeed)/2 + 1
	} else if world.CurrentWind.Direction() == "NE" || world.CurrentWind.Direction() == "SW" {
		//beam reach (wind perpendicular) are at wind speed
		speed = world.CurrentWind.windSpeed / 2
	} else if world.CurrentWind.Direction() == "N" || world.CurrentWind.Direction() == "SE" || world.CurrentWind.Direction() == "W" {
		//running (wind behind) or close haul (forward and to the side) is at -1 versus wind speed
		speed = (world.CurrentWind.windSpeed - 1) / 2
	}
	actor.X = MaxValue(actor.X-speed, 0)
	actor.Y = MaxValue(actor.Y-speed, 0)
	return true
}
