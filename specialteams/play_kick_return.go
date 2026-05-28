package specialteams

import (
	"fmt"

	"github.com/brunkow/football/game"
)

// PlayKickReturn represents a kick return play.
type PlayKickReturn struct {
	BaseSpecialTeamsPlay
}

// NewPlayKickReturn creates a new PlayKickReturn.
func NewPlayKickReturn(g *game.Game) *PlayKickReturn {
	return &PlayKickReturn{
		BaseSpecialTeamsPlay: NewBaseSpecialTeamsPlay(g),
	}
}

// Run executes the kick return (placeholder).
func (p *PlayKickReturn) Run() {
	fmt.Println("Kickoff return")
}
