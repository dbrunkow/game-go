package offensive

import (
	"fmt"

	"github.com/brunkow/football/game"
)

// OffensiveCoordinator makes offensive play-calling decisions.
type OffensiveCoordinator struct {
	game.Coordinator
}

// NewOffensiveCoordinator creates a new OffensiveCoordinator for the given team.
func NewOffensiveCoordinator(team int) *OffensiveCoordinator {
	return &OffensiveCoordinator{
		Coordinator: game.NewCoordinator(team),
	}
}

// GetWinningBy returns how many points this team is winning by.
func (oc *OffensiveCoordinator) GetWinningBy(g *game.Game) int {
	return g.Score[oc.Team] - g.Score[1-oc.Team]
}

// GetOffensiveLineup returns the offensive lineup from the roster.
func (oc *OffensiveCoordinator) GetOffensiveLineup(g *game.Game) map[string]*game.Player {
	roster := oc.GetRoster(g)
	lineup := make(map[string]*game.Player)
	if oc.Rand.Intn(30) < 29 {
		lineup["QB"] = roster["QB1"]
	} else {
		lineup["QB"] = roster["QB2"]
	}
	if oc.Rand.Intn(3) < 2 {
		lineup["HB"] = roster["RB1"]
	} else {
		lineup["HB"] = roster["RB2"]
	}
	lineup["WR1"] = roster["WR1"]
	lineup["WR2"] = roster["WR2"]
	lineup["WR3"] = roster["WR3"]
	lineup["TE"] = roster["TE1"]
	lineup["C"] = roster["C1"]
	lineup["RG"] = roster["OG1"]
	lineup["LG"] = roster["OG2"]
	lineup["RT"] = roster["OT1"]
	lineup["LT"] = roster["OT2"]
	return lineup
}

// CreatePlay decides which offensive play to call based on game situation.
func (oc *OffensiveCoordinator) CreatePlay(g *game.Game) game.OffensivePlayInterface {
	// End of half field goal attempt
	if g.Clock >= 870 && g.GetYardsToTD() <= 30 && g.Quarter == 2 {
		return NewPlayFieldGoal(g)
	}

	// End of game field goal attempt when close and losing by 3 or less
	if g.Clock >= 870 && g.GetYardsToTD() <= 35 &&
		g.Quarter >= 4 &&
		oc.GetWinningBy(g) >= -3 && oc.GetWinningBy(g) <= 0 {
		return NewPlayFieldGoal(g)
	}

	// Fourth down decisions
	if g.IsFourthDown() {
		if g.GetYardsToTD() < 10 {
			fmt.Println("Going for it!")
			return NewPlayRun(g)
		} else if g.Clock >= 600 && g.Quarter >= 4 && oc.GetWinningBy(g) <= 0 {
			if oc.Rand.Intn(3) == 0 {
				return NewPlayRun(g)
			}
			return NewPlayPass(g)
		} else if g.Clock >= 400 && g.GetYardsToTD() <= 50 &&
			g.Quarter >= 4 && oc.GetWinningBy(g) <= 0 {
			if oc.Rand.Intn(3) == 0 {
				return NewPlayRun(g)
			}
			return NewPlayPass(g)
		} else if g.GetYardsToTD() < 35 {
			return NewPlayFieldGoal(g)
		}
		return NewPlayPunt(g)
	}

	// Normal play calling
	if g.Quarter >= 4 && oc.GetWinningBy(g) <= 0 {
		// More passing when losing in Q4
		if oc.Rand.Intn(6) == 0 {
			return NewPlayRun(g)
		}
		return NewPlayPass(g)
	}

	// 50/50 run/pass
	if oc.Rand.Intn(2) == 0 {
		return NewPlayRun(g)
	}
	return NewPlayPass(g)
}
