package offensive

import (
	"fmt"

	"github.com/brunkow/football/game"
)

// PlayPunt represents a punt play.
type PlayPunt struct {
	BaseOffensivePlay
}

// NewPlayPunt creates a new PlayPunt.
func NewPlayPunt(g *game.Game) *PlayPunt {
	return &PlayPunt{
		BaseOffensivePlay: NewBaseOffensivePlay(g),
	}
}

// Run executes the punt.
func (p *PlayPunt) Run() float64 {
	fmt.Printf("Punting to %s from %.1f", p.GetDefenseTeam().Name, p.Game.YardLine)
	distance := 40.0
	p.Game.SetTurnover()
	p.Game.AddYards(distance)
	fmt.Printf(" to %.1f\n", p.Game.YardLine)

	if p.Game.YardLine <= 0 && !p.Game.IsDirection() {
		p.Game.YardLine = 20.0
		fmt.Println("Touchback")
	} else if p.Game.YardLine >= 100 && p.Game.IsDirection() {
		p.Game.YardLine = 80.0
		fmt.Println("Touchback")
	}
	p.Game.ChangePossession()
	p.Game.AddClock(6)
	return 0.0
}
