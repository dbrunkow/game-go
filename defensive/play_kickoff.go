package defensive

import "github.com/brunkow/football/game"

// PlayKickoff is a placeholder defensive play for kickoffs.
type PlayKickoff struct {
	BaseDefensivePlay
}

// NewPlayKickoff creates a new defensive PlayKickoff.
func NewPlayKickoff(g *game.Game) *PlayKickoff {
	return &PlayKickoff{
		BaseDefensivePlay: NewBaseDefensivePlay(g),
	}
}

// Run executes the defensive kickoff play (placeholder, returns 0).
func (p *PlayKickoff) Run() float64 {
	return 0.0
}
