package offensive

import (
	"fmt"

	"github.com/brunkow/football/game"
)

// PlayFieldGoal represents a field goal attempt.
type PlayFieldGoal struct {
	BaseOffensivePlay
}

// NewPlayFieldGoal creates a new PlayFieldGoal.
func NewPlayFieldGoal(g *game.Game) *PlayFieldGoal {
	return &PlayFieldGoal{
		BaseOffensivePlay: NewBaseOffensivePlay(g),
	}
}

// Run executes the field goal attempt.
func (p *PlayFieldGoal) Run() float64 {
	fmt.Printf("Attempting field goal! %.1f\n", p.Game.GetYardsToTD())
	good := p.kick()
	p.Game.SetTurnover()
	if good {
		fmt.Println("Kick is good")
		p.SetScore(3)
		p.Game.FieldGoal = true
		game.LogScore(p.Game)
		if !p.Game.IsDirection() {
			p.Game.YardLine = 20.0
		} else {
			p.Game.YardLine = 80.0
		}
	} else {
		fmt.Println("Missed field goal")
	}
	p.Game.ChangePossession()
	p.Game.Down = 1
	p.Game.AddClock(6)
	return 0.0
}

func (p *PlayFieldGoal) kick() bool {
	kicking := p.Rand.Intn(100)
	yardsToTD := p.Game.GetYardsToTD()
	if yardsToTD <= 25 {
		return kicking <= 95
	} else if yardsToTD < 35 {
		return kicking <= 80
	} else if yardsToTD < 45 {
		return kicking <= 70
	} else if yardsToTD < 55 {
		return kicking <= 50
	}
	return kicking <= 10
}
