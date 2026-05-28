package offensive

import (
	"fmt"

	"github.com/brunkow/football/game"
)

// PlayPass represents a passing play.
type PlayPass struct {
	BaseOffensivePlay
}

// NewPlayPass creates a new PlayPass.
func NewPlayPass(g *game.Game) *PlayPass {
	return &PlayPass{
		BaseOffensivePlay: NewBaseOffensivePlay(g),
	}
}

// Run executes the passing play.
func (p *PlayPass) Run() float64 {
	yards := p.passOrInterception()
	if p.Game.IsTouchdown() {
		p.Game.Td = true
		p.SetScore(6)
		fmt.Println("Touchdown!")
		kicking := p.Rand.Intn(100)
		if kicking < 95 {
			fmt.Println("Kick is good")
			p.SetScore(1)
		} else {
			fmt.Println("Missed extra point")
		}
		p.Game.Down = 1
		game.LogScore(p.Game)
		p.Game.Turnover = false
		p.Game.Series = 0
		if p.Game.IsDirection() {
			p.Game.YardLine = 80
		} else {
			p.Game.YardLine = 20
		}
		p.Game.ChangePossession()
	} else if p.Game.Fumble {
		p.Game.ChangePossession()
		fmt.Printf("Fumble! %s have the ball\n", p.GetOffenseTeam().Name)
		p.Game.Fumble = false
	} else if p.Game.Interception {
		p.Game.ChangePossession()
		fmt.Printf("Interception! %s have the ball\n", p.GetOffenseTeam().Name)
		p.Game.Interception = false
	} else {
		if p.Game.IsFourthDown() {
			fmt.Println("Turnover on downs")
			p.Game.ChangePossession()
		} else if p.Game.IsFirstDown() {
			p.Game.Series = 0.0
			p.Game.Down = 1
		} else {
			p.Game.AddDown()
		}
	}
	return yards
}

func (p *PlayPass) passOrInterception() float64 {
	interception := p.Rand.Intn(1000)
	yards := p.pass()
	if interception > 950 && !p.Game.IsTouchdown() {
		p.Game.Interception = true
		p.Game.AddClock(6)
	} else {
		if yards > 0.0 {
			p.Game.AddClock(30)
		} else {
			p.Game.AddClock(6)
		}
	}
	return yards
}

func (p *PlayPass) pass() float64 {
	var yards float64
	diff := p.GetTeamDiff() * 15.0
	group := float64(p.Rand.Intn(1000)) + diff

	if group <= 450 {
		// Incomplete
		yards = 0.0
	} else if group < 470.0 {
		// Touchdown pass
		yards = p.Game.GetYardsToTD()
	} else if group < 850.0 {
		// 10-20 yards
		yards = float64(p.Rand.Intn(100))/10.0 + 10.0
		if p.Game.GetYardsToTD() < yards {
			yards = p.Game.GetYardsToTD()
		}
	} else if group < 950.0 {
		// 20-30 yards
		yards = float64(p.Rand.Intn(100))/10.0 + 20.0
		if p.Game.GetYardsToTD() < yards {
			yards = p.Game.GetYardsToTD()
		}
	} else {
		// 0-10 yards
		yards = float64(p.Rand.Intn(100)) / 10.0
		if p.Game.GetYardsToTD() < yards {
			yards = p.Game.GetYardsToTD()
		}
	}
	p.Game.AddYards(yards)

	completion := 0
	if yards > 0.0 {
		completion = 1
	}
	p.Game.Stats.AddPassStats(p.Game.OffenseTeam, "QB1", 1, completion, yards)

	fmt.Printf("%d P %.1f %.1f\n", p.Game.Down, yards, p.Game.YardLine)
	return yards
}
