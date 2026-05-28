package defensive

import "github.com/brunkow/football/game"

// PlayPass is a placeholder defensive play against the pass.
type PlayPass struct {
	BaseDefensivePlay
}

// NewPlayPass creates a new defensive PlayPass.
func NewPlayPass(g *game.Game) *PlayPass {
	return &PlayPass{
		BaseDefensivePlay: NewBaseDefensivePlay(g),
	}
}

// Run executes the defensive play (placeholder, returns 0).
func (p *PlayPass) Run() float64 {
	return 0.0
}
