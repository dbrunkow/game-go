package specialteams

import (
	"github.com/brunkow/football/game"
)

// SpecialTeamsCoordinator makes special teams play-calling decisions.
type SpecialTeamsCoordinator struct {
	game.Coordinator
}

// NewSpecialTeamsCoordinator creates a new SpecialTeamsCoordinator for the given team.
func NewSpecialTeamsCoordinator(team int) *SpecialTeamsCoordinator {
	return &SpecialTeamsCoordinator{
		Coordinator: game.NewCoordinator(team),
	}
}

// GetWinningBy returns how many points this team is winning by.
func (stc *SpecialTeamsCoordinator) GetWinningBy(g *game.Game) int {
	return g.Score[stc.Team] - g.Score[1-stc.Team]
}

// CreatePlay decides which special teams play to call.
func (stc *SpecialTeamsCoordinator) CreatePlay(g *game.Game) game.SpecialTeamsPlayInterface {
	if g.Td || g.FieldGoal {
		if g.Quarter >= 4 && g.Clock > 400 && stc.GetWinningBy(g) > 0 {
			return NewPlayOnsideKick(g)
		}
		return NewPlayKickoff(g)
	}
	return NewPlayKickoff(g)
}
