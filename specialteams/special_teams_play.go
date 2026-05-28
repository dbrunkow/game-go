package specialteams

import (
	"math/rand"

	"github.com/brunkow/football/game"
)

// BaseSpecialTeamsPlay provides common functionality for special teams plays.
type BaseSpecialTeamsPlay struct {
	Game *game.Game
	Rand *rand.Rand
}

// NewBaseSpecialTeamsPlay creates a new BaseSpecialTeamsPlay.
func NewBaseSpecialTeamsPlay(g *game.Game) BaseSpecialTeamsPlay {
	return BaseSpecialTeamsPlay{
		Game: g,
		Rand: rand.New(rand.NewSource(rand.Int63())),
	}
}

// GetOffenseTeam returns the current offensive team.
func (b *BaseSpecialTeamsPlay) GetOffenseTeam() *game.Team {
	return &b.Game.Teams[b.Game.OffenseTeam]
}

// GetDefenseTeam returns the current defensive team.
func (b *BaseSpecialTeamsPlay) GetDefenseTeam() *game.Team {
	return &b.Game.Teams[b.Game.DefenseTeam]
}

// GetTeamDiff returns the difference between offense and defense ratings.
func (b *BaseSpecialTeamsPlay) GetTeamDiff() float64 {
	return b.GetOffenseTeam().Offense - b.GetDefenseTeam().Defense
}
