package defensive

import (
	"math/rand"

	"github.com/brunkow/football/game"
)

// BaseDefensivePlay provides common functionality for defensive plays.
type BaseDefensivePlay struct {
	Game *game.Game
	Rand *rand.Rand
}

// NewBaseDefensivePlay creates a new BaseDefensivePlay.
func NewBaseDefensivePlay(g *game.Game) BaseDefensivePlay {
	return BaseDefensivePlay{
		Game: g,
		Rand: rand.New(rand.NewSource(rand.Int63())),
	}
}

// GetOffenseTeam returns the current offensive team.
func (b *BaseDefensivePlay) GetOffenseTeam() *game.Team {
	return &b.Game.Teams[b.Game.OffenseTeam]
}

// GetDefenseTeam returns the current defensive team.
func (b *BaseDefensivePlay) GetDefenseTeam() *game.Team {
	return &b.Game.Teams[b.Game.DefenseTeam]
}

// GetTeamDiff returns the difference between offense and defense ratings.
func (b *BaseDefensivePlay) GetTeamDiff() float64 {
	return b.GetOffenseTeam().Offense - b.GetDefenseTeam().Defense
}

// SetScore adds points to the appropriate team based on direction.
func (b *BaseDefensivePlay) SetScore(score int) {
	b.Game.SetScore(score)
}
