package defensive

import "github.com/brunkow/football/game"

// PlayBlockKick is a placeholder defensive play for blocking kicks.
type PlayBlockKick struct {
	BaseDefensivePlay
}

// NewPlayBlockKick creates a new defensive PlayBlockKick.
func NewPlayBlockKick(g *game.Game) *PlayBlockKick {
	return &PlayBlockKick{
		BaseDefensivePlay: NewBaseDefensivePlay(g),
	}
}

// Run executes the block kick play (placeholder, returns 0).
func (p *PlayBlockKick) Run() float64 {
	return 0.0
}
