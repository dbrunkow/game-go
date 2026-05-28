package specialteams

import (
	"fmt"

	"github.com/brunkow/football/game"
)

// PlayKickoff represents a kickoff play.
type PlayKickoff struct {
	BaseSpecialTeamsPlay
}

// NewPlayKickoff creates a new PlayKickoff.
func NewPlayKickoff(g *game.Game) *PlayKickoff {
	return &PlayKickoff{
		BaseSpecialTeamsPlay: NewBaseSpecialTeamsPlay(g),
	}
}

// Run executes the kickoff.
func (p *PlayKickoff) Run() {
	fmt.Printf("Kicking off to %s\n", p.GetOffenseTeam().Name)
	p.Game.AddClock(6)
}
