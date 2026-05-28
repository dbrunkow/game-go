package offensive

import (
	"math/rand"

	"github.com/brunkow/football/game"
)

// BaseOffensivePlay provides common functionality for offensive plays.
type BaseOffensivePlay struct {
	Game *game.Game
	Rand *rand.Rand
}

// NewBaseOffensivePlay creates a new BaseOffensivePlay.
func NewBaseOffensivePlay(g *game.Game) BaseOffensivePlay {
	return BaseOffensivePlay{
		Game: g,
		Rand: rand.New(rand.NewSource(rand.Int63())),
	}
}

// GetOffenseTeam returns the current offensive team.
func (b *BaseOffensivePlay) GetOffenseTeam() *game.Team {
	return &b.Game.Teams[b.Game.OffenseTeam]
}

// GetDefenseTeam returns the current defensive team.
func (b *BaseOffensivePlay) GetDefenseTeam() *game.Team {
	return &b.Game.Teams[b.Game.DefenseTeam]
}

// GetTeamDiff returns the difference between offense and defense ratings.
func (b *BaseOffensivePlay) GetTeamDiff() float64 {
	return b.GetOffenseTeam().Offense - b.GetDefenseTeam().Defense
}

// SetScore adds points to the appropriate team based on direction.
func (b *BaseOffensivePlay) SetScore(score int) {
	b.Game.SetScore(score)
}
