package game

import "math/rand"

// Coordinator is the base struct for all coordinators.
type Coordinator struct {
	Team int
	Rand *rand.Rand
}

// NewCoordinator creates a new Coordinator for the given team.
func NewCoordinator(team int) Coordinator {
	return Coordinator{
		Team: team,
		Rand: rand.New(rand.NewSource(rand.Int63())),
	}
}

// GetTeam returns the Team struct for this coordinator's team.
func (c *Coordinator) GetTeam(g *Game) *Team {
	return &g.Teams[c.Team]
}

// GetRoster returns the roster for this coordinator's team.
func (c *Coordinator) GetRoster(g *Game) map[string]*Player {
	return c.GetTeam(g).Roster
}
