package offensive

import (
	"fmt"

	"github.com/brunkow/football/game"
)

// PlayRun represents a rushing play.
type PlayRun struct {
	BaseOffensivePlay
}

// NewPlayRun creates a new PlayRun.
func NewPlayRun(g *game.Game) *PlayRun {
	return &PlayRun{
		BaseOffensivePlay: NewBaseOffensivePlay(g),
	}
}

// Run executes the rushing play.
func (p *PlayRun) Run() float64 {
	yards := p.runOrFumble()
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

func (p *PlayRun) runOrFumble() float64 {
	yards := p.runBall()
	fumble := p.Rand.Intn(1000)
	if fumble > 980 && !p.Game.IsTouchdown() {
		p.Game.Fumble = true
		p.Game.AddClock(6)
	} else {
		p.Game.AddClock(30)
	}
	return yards
}

func (p *PlayRun) runBall() float64 {
	var yards float64
	diff := p.GetTeamDiff() * 15.0
	group := float64(p.Rand.Intn(1000)) + diff

	if group <= 120.0 {
		// Negative yards
		yards = -1.0 * float64(p.Rand.Intn(100)) / 10.0
	} else if group < 210 {
		// 10-20 yards
		yards = float64(p.Rand.Intn(100))/10.0 + 10.0
		if p.Game.GetYardsToTD() < yards {
			yards = p.Game.GetYardsToTD()
		}
	} else if group < 230 {
		// 20+ breakaway
		yards = 100.0
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
	p.Game.Stats.AddRunStats(p.Game.OffenseTeam, "RB1", 1, yards)

	fmt.Printf("%d R %.1f %.1f\n", p.Game.Down, yards, p.Game.YardLine)
	return yards
}
