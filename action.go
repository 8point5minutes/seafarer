package main

type Actor interface {
	noAction(*World) bool
	sailN(*World) bool
	sailNW(*World) bool
	sailNE(*World) bool
	sailE(*World) bool
	sailW(*World) bool
	sailSE(*World) bool
	sailSW(*World) bool
	sailS(*World) bool
}

func (p *Player) noAction(world *World) bool {
	return false
}
func NoAction(i Actor) func(*World) bool { return i.noAction }

func (s *Ship) sailN(world *World) bool {
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
	s.Y = MaxValue(s.Y-speed, 0)
	return true
}

func SailN(i Actor) func(*World) bool { return i.sailN }

func (s *Ship) sailNE(world *World) bool {
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
	s.Y = MaxValue(s.Y-speed, 0)
	s.X = MinValue(s.X+speed, gd.ScreenWidth-1)
	return true
}

func SailNE(i Actor) func(*World) bool { return i.sailNE }

func (s *Ship) sailE(world *World) bool {
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
	s.X = MinValue(s.X+speed, gd.ScreenWidth-1)
	return true
}

func SailE(i Actor) func(*World) bool { return i.sailE }

func (p *Player) sailS(world *World) bool {
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
	p.Y = MinValue(p.Y+speed, gd.ScreenHeight-1)
	return true
}
func SailS(i Actor) func(*World) bool { return i.sailS }

func (s *Ship) sailSE(world *World) bool {
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
	s.X = MinValue(s.X+speed, gd.ScreenWidth-1)
	s.Y = MinValue(s.Y+speed, gd.ScreenHeight-1)
	return true
}
func SailSE(i Actor) func(*World) bool { return i.sailSE }

func (s *Ship) sailSW(world *World) bool {
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
	s.X = MaxValue(s.X-speed, 0)
	s.Y = MinValue(s.Y+speed, gd.ScreenHeight-1)
	return true
}
func SailSW(i Actor) func(*World) bool { return i.sailSW }

func (s *Ship) sailW(world *World) bool {
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
	s.X = MaxValue(s.X-speed, 0)
	return true
}
func SailW(i Actor) func(*World) bool { return i.sailW }

func (s *Ship) sailNW(world *World) bool {
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
	s.X = MaxValue(s.X-speed, 0)
	s.Y = MaxValue(s.Y-speed, 0)
	return true
}
func SailNW(i Actor) func(*World) bool { return i.sailNW }
