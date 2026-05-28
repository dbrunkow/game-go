package game

import "fmt"

// GameEngine drives the game simulation.
type GameEngine struct {
	Game *Game
}

// NewGameEngine creates a new GameEngine for the given game.
func NewGameEngine(g *Game) *GameEngine {
	return &GameEngine{Game: g}
}

// Halftime resets state for the second half.
func (e *GameEngine) Halftime() {
	e.Game.SetDirection(false)
	e.Game.OffenseTeam = TeamB
	e.Game.DefenseTeam = TeamA
	e.Game.YardLine = 80.0
	e.Game.Down = 1
	e.Game.Series = 0.0
	e.Game.Turnover = false
}

// GetOffensiveCoordinator returns the offensive coordinator for the current offense.
func (e *GameEngine) GetOffensiveCoordinator() OffensiveCoordinatorInterface {
	return e.Game.Teams[e.Game.OffenseTeam].OffensiveCoordinator
}

// GetDefensiveCoordinator returns the defensive coordinator for the current defense.
func (e *GameEngine) GetDefensiveCoordinator() DefensiveCoordinatorInterface {
	return e.Game.Teams[e.Game.DefenseTeam].DefensiveCoordinatorRef
}

// GetOSpecialTeamsCoordinator returns the special teams coordinator for the offense.
func (e *GameEngine) GetOSpecialTeamsCoordinator() SpecialTeamsCoordinatorInterface {
	return e.Game.Teams[e.Game.OffenseTeam].SpecialTeamsCoordinator
}

// GetDSpecialTeamsCoordinator returns the special teams coordinator for the defense.
func (e *GameEngine) GetDSpecialTeamsCoordinator() SpecialTeamsCoordinatorInterface {
	return e.Game.Teams[e.Game.DefenseTeam].SpecialTeamsCoordinator
}

// RunQuarter simulates one quarter of play.
func (e *GameEngine) RunQuarter(quarter int) {
	e.Game.Clock = 0
	e.Game.Quarter = quarter
	for e.Game.Clock < 900 {
		oPlay := e.GetOffensiveCoordinator().CreatePlay(e.Game)
		dPlay := e.GetDefensiveCoordinator().CreatePlay(e.Game)
		e.Game.RunOffensivePlay(oPlay, dPlay)

		if e.Game.Td || e.Game.FieldGoal {
			ostPlay := e.GetOSpecialTeamsCoordinator().CreatePlay(e.Game)
			dstPlay := e.GetDSpecialTeamsCoordinator().CreatePlay(e.Game)
			e.Game.RunSpecialTeamsPlay(ostPlay, dstPlay)
			e.Game.Td = false
			e.Game.FieldGoal = false
		} else if e.Game.GetYardsToTD() < 0 || e.Game.GetYardsToTD() > 100 {
			fmt.Println("Safety")
			e.AddDefensiveScore(2)
			LogScore(e.Game)
			e.Game.Td = false
			e.Game.Turnover = false
			e.Game.Series = 0
			if e.Game.IsDirection() {
				e.Game.SetDirection(false)
				e.Game.YardLine = 80
			} else {
				e.Game.YardLine = 20
				e.Game.SetDirection(true)
			}
			e.Game.Down = 1
			e.Game.SwitchOffenseDefense()
		}
	}
	fmt.Printf("Total yards: %.1f\n", e.Game.YardLine)
	LogScore(e.Game)
	fmt.Printf("End of quarter %d\n", quarter)
	fmt.Println("===========================================================")
}

// AddScore adds points to the offensive team.
func (e *GameEngine) AddScore(score int) {
	if e.Game.IsDirection() {
		e.Game.Score[TeamA] += score
	} else {
		e.Game.Score[TeamB] += score
	}
}

// AddDefensiveScore adds points to the defensive team.
func (e *GameEngine) AddDefensiveScore(score int) {
	if !e.Game.IsDirection() {
		e.Game.Score[TeamA] += score
	} else {
		e.Game.Score[TeamB] += score
	}
}
