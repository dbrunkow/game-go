package game

// Player represents a football player with a power rating.
type Player struct {
	Power     int
	FirstName string
	LastName  string
}

// NewPlayer creates a new Player with the given power rating.
func NewPlayer(power int) *Player {
	return &Player{Power: power}
}
