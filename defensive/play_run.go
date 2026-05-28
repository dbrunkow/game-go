package defensive

import "github.com/brunkow/football/game"

// PlayRun is a placeholder defensive play against the run.
type PlayRun struct {
	BaseDefensivePlay
}

// NewPlayRun creates a new defensive PlayRun.
func NewPlayRun(g *game.Game) *PlayRun {
	return &PlayRun{
		BaseDefensivePlay: NewBaseDefensivePlay(g),
	}
}

// Run executes the defensive play (placeholder, returns 0).
func (p *PlayRun) Run() float64 {
	return 0.0
}
