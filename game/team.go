package game

import "math/rand"

// Team represents a football team.
type Team struct {
	Name                    string
	Index                   int
	Offense                 float64
	Defense                 float64
	Roster                  map[string]*Player
	OffensiveCoordinator    OffensiveCoordinatorInterface
	DefensiveCoordinatorRef DefensiveCoordinatorInterface
	SpecialTeamsCoordinator SpecialTeamsCoordinatorInterface
}

// OffensiveCoordinatorInterface defines the interface for offensive coordinators.
type OffensiveCoordinatorInterface interface {
	CreatePlay(g *Game) OffensivePlayInterface
}

// DefensiveCoordinatorInterface defines the interface for defensive coordinators.
type DefensiveCoordinatorInterface interface {
	CreatePlay(g *Game) DefensivePlayInterface
}

// SpecialTeamsCoordinatorInterface defines the interface for special teams coordinators.
type SpecialTeamsCoordinatorInterface interface {
	CreatePlay(g *Game) SpecialTeamsPlayInterface
}

// OffensivePlayInterface defines the interface for offensive plays.
type OffensivePlayInterface interface {
	Run() float64
}

// DefensivePlayInterface defines the interface for defensive plays.
type DefensivePlayInterface interface {
	Run() float64
}

// SpecialTeamsPlayInterface defines the interface for special teams plays.
type SpecialTeamsPlayInterface interface {
	Run()
}

// NewTeam creates a new team with the given parameters.
func NewTeam(name string, index int, offense, defense float64) *Team {
	return &Team{
		Name:    name,
		Index:   index,
		Offense: offense,
		Defense: defense,
		Roster:  createRoster(),
	}
}

func createRoster() map[string]*Player {
	r := rand.New(rand.NewSource(rand.Int63()))
	roster := make(map[string]*Player)
	positions := []string{
		"RB1", "RB2", "RB3",
		"QB1", "QB2", "QB3",
		"FB1",
		"WR1", "WR2", "WR3", "WR4", "WR5",
		"TE1", "TE2",
		"C1", "C2",
		"OG1", "OG2", "OG3",
		"OT1", "OT2", "OT3", "OT4",
		"DT1", "DT2", "DT3", "DT4",
		"DE1", "DE2", "DE3", "DE4",
		"ILB1", "ILB2",
		"OLB1", "OLB2", "OLB3",
		"CB1", "CB2", "CB3", "CB4",
		"SS1", "SS2",
		"FS1", "FS2",
		"P1", "K1",
	}
	for _, pos := range positions {
		roster[pos] = NewPlayer(r.Intn(20))
	}
	return roster
}
