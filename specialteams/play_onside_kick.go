package specialteams

import (
	"fmt"

	"github.com/brunkow/football/game"
)

// PlayOnsideKick represents an onside kick play.
type PlayOnsideKick struct {
	BaseSpecialTeamsPlay
}

// NewPlayOnsideKick creates a new PlayOnsideKick.
func NewPlayOnsideKick(g *game.Game) *PlayOnsideKick {
	return &PlayOnsideKick{
		BaseSpecialTeamsPlay: NewBaseSpecialTeamsPlay(g),
	}
}

// Run executes the onside kick.
func (p *PlayOnsideKick) Run() {
	fmt.Println("Onside kick")
	recover := p.Rand.Intn(1000)
	if recover > 500 {
		// Recovered by kicking team
		p.Game.ChangePossession()
		fmt.Printf("Recovered by kicking team %s\n", p.GetOffenseTeam().Name)
	} else {
		// Recovered by receiving team
		fmt.Printf("Recovered by receiving team %s\n", p.GetOffenseTeam().Name)
	}
	if !p.Game.IsDirection() {
		p.Game.YardLine = 60.0
	} else {
		p.Game.YardLine = 40.0
	}
	p.Game.AddClock(6)
}
