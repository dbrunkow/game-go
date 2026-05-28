package main

import (
	"fmt"

	"github.com/brunkow/football/defensive"
	"github.com/brunkow/football/game"
	"github.com/brunkow/football/offensive"
	"github.com/brunkow/football/specialteams"
)

func createBears() *game.Team {
	team := game.NewTeam("Bears", game.TeamA, 9, 9)
	team.OffensiveCoordinator = offensive.NewOffensiveCoordinator(game.TeamA)
	team.DefensiveCoordinatorRef = defensive.NewDefensiveCoordinator(game.TeamA)
	team.SpecialTeamsCoordinator = specialteams.NewSpecialTeamsCoordinator(game.TeamA)
	return team
}

func createJets() *game.Team {
	team := game.NewTeam("Jets", game.TeamB, 9, 26)
	team.OffensiveCoordinator = offensive.NewOffensiveCoordinator(game.TeamB)
	team.DefensiveCoordinatorRef = defensive.NewDefensiveCoordinator(game.TeamB)
	team.SpecialTeamsCoordinator = specialteams.NewSpecialTeamsCoordinator(game.TeamB)
	return team
}

func main() {
	teamA := createBears()
	teamB := createJets()

	g := game.NewGame(teamA, teamB)
	engine := game.NewGameEngine(g)

	fmt.Printf("Starting %s\n", g.Teams[g.OffenseTeam].Name)
	engine.RunQuarter(1)
	engine.RunQuarter(2)
	engine.Halftime()
	fmt.Printf("Starting %s\n", g.Teams[g.OffenseTeam].Name)
	engine.RunQuarter(3)
	engine.RunQuarter(4)

	if g.IsTied() {
		engine.Halftime()
		engine.RunQuarter(5)
	}

	g.Stats.PrintStats()
}
