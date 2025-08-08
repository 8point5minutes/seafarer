package main

type Actor interface {
	noAction(*World) bool
	openMenu(*World) bool
	closeMenu(*World) bool
	sail(*World) bool
	anchor(*World) bool
}

func (p *Player) noAction(world *World) bool {
	return false
}
func NoAction(i Actor) func(*World) bool { return i.noAction }

func (p *Player) openMenu(world *World) bool {
	p.CurrentMenu = 1
	return false
}
func OpenMenu(i Actor) func(*World) bool { return i.openMenu }

func (p *Player) closeMenu(world *World) bool {
	p.CurrentMenu = 0
	return false
}
func CloseMenu(i Actor) func(*World) bool { return i.closeMenu }

func (s *Ship) sail(world *World) bool {
	s.CurrentMovePoints = s.Type.MaxMovePoints
	originalPosition := s.CurrentDirection
	//determining turns in certain directions
	if s.RudderPosition == TurningPort {
		s.PortTurn()
		if IsTurnWindward(originalPosition, s.CurrentDirection, world.CurrentWind.WindDirection) {
			//turning only one if we turn into the wind
			s.CurrentMovePoints--
		} else {
			//we turn as much as the ship can otherwise
			for i := 0; i < s.Type.MaxTurnSpeed-1; i++ {
				s.PortTurn()
			}
		}
	} else if s.RudderPosition == TurningStarboard {
		s.StarboardTurn()
		if IsTurnWindward(originalPosition, s.CurrentDirection, world.CurrentWind.WindDirection) {
			//turning only one if we turn into the wind
			s.CurrentMovePoints--
		} else {
			//we turn as much as the ship can otherwise
			for i := 0; i < s.Type.MaxTurnSpeed-1; i++ {
				s.StarboardTurn()
			}
		}
	}
	if s.Anchored {
		s.CurrentMovePoints = 0
	} else {
		wind := world.CurrentWind
		point_of_sail := PointsOfSail(s.CurrentDirection, wind.WindDirection)
		if point_of_sail == IntoWind {
			s.CurrentMovePoints = 0
		} else if point_of_sail == CloseHaul {
			if wind.windSpeed == SteadyGusts {
				s.CurrentMovePoints = 2
			} else {
				s.CurrentMovePoints = 1
			}
		} else if point_of_sail == BeamReach {
			if wind.windSpeed == SteadyGusts {
				s.CurrentMovePoints = 2
			} else {
				s.CurrentMovePoints = 1
			}
		} else if point_of_sail == BroadReach {
			if wind.windSpeed == SteadyGusts {
				s.CurrentMovePoints = 1
			} else if wind.windSpeed == LightWinds {
				s.CurrentMovePoints = 2
			} else {
				s.CurrentMovePoints = 3
			}
		} else if point_of_sail == RunDownwind {
			s.CurrentMovePoints = 1
		}
	}
	for i := 0; i < s.CurrentMovePoints; i++ {
		s.Move()
	}
	return true
}

func Sail(i Actor) func(*World) bool { return i.sail }

func (s *Ship) anchor(world *World) bool {
	if s.Anchored {
		s.Anchored = false
	} else {
		s.Anchored = true
	}
	return true
}

func Anchor(i Actor) func(*World) bool { return i.anchor }
