package defensive

import (
	"github.com/brunkow/football/game"
)

// DefensiveCoordinator makes defensive play-calling decisions.
type DefensiveCoordinator struct {
	game.Coordinator
}

// NewDefensiveCoordinator creates a new DefensiveCoordinator for the given team.
func NewDefensiveCoordinator(team int) *DefensiveCoordinator {
	return &DefensiveCoordinator{
		Coordinator: game.NewCoordinator(team),
	}
}

// GetDefensiveLineup returns the defensive lineup from the roster.
func (dc *DefensiveCoordinator) GetDefensiveLineup(roster map[string]*game.Player) map[string]*game.Player {
	lineup := make(map[string]*game.Player)
	lineup["LE"] = roster["DE1"]
	lineup["RE"] = roster["DE2"]
	lineup["LT"] = roster["LT1"]
	lineup["RT"] = roster["LT2"]
	lineup["LOLB"] = roster["OLB1"]
	lineup["MLB"] = roster["ILB1"]
	lineup["RILB"] = roster["OLB2"]
	lineup["LCB"] = roster["CB1"]
	lineup["RCB"] = roster["CB2"]
	lineup["SS"] = roster["SS1"]
	lineup["FS"] = roster["FS1"]
	return lineup
}

// CreatePlay returns a defensive play (currently always DefensivePlayRun as placeholder).
func (dc *DefensiveCoordinator) CreatePlay(g *game.Game) game.DefensivePlayInterface {
	return NewPlayRun(g)
}
