package game

import "fmt"

// Team index constants
const (
	TeamA = 0
	TeamB = 1
)

// Game holds the complete state of a football game.
type Game struct {
	Teams       [2]Team
	OffenseTeam int
	DefenseTeam int
	Score       [2]int
	YardLine    float64
	Direction   int // 0 or 1
	Series      float64
	Td          bool
	FieldGoal   bool
	Turnover    bool
	Fumble      bool
	Interception bool
	Clock       int
	Down        int
	Quarter     int
	Stats       *Stats
}

// NewGame creates a new Game with two teams.
func NewGame(teamA, teamB *Team) *Game {
	g := &Game{
		OffenseTeam: TeamA,
		DefenseTeam: TeamB,
		YardLine:    20.0,
		Direction:   1,
		Down:        1,
		Series:      0.0,
		Turnover:    false,
	}
	g.Teams[TeamA] = *teamA
	g.Teams[TeamB] = *teamB
	g.Score[TeamA] = 0
	g.Score[TeamB] = 0
	g.Stats = NewStats(g)
	return g
}

// IsTouchdown checks if the current yard line represents a touchdown.
func (g *Game) IsTouchdown() bool {
	if g.Direction == 1 {
		return g.YardLine >= 100 || g.Td
	}
	return g.YardLine <= 0 || g.Td
}

// IsActiveDrive checks if the current drive is still active.
func (g *Game) IsActiveDrive() bool {
	return g.Down < 5 && !g.Td && !g.Turnover
}

// IsFirstDown checks if the offense has earned a first down.
func (g *Game) IsFirstDown() bool {
	return g.Series >= 10.0
}

// IsFourthDown checks if it's fourth down without a first down.
func (g *Game) IsFourthDown() bool {
	return g.Series < 10 && g.Down == 4
}

// AddYards adds yards to the current position based on direction.
func (g *Game) AddYards(yards float64) {
	if g.Direction == 1 {
		g.YardLine += yards
	} else {
		g.YardLine -= yards
	}
	g.Series += yards
}

// AddClock adds seconds to the game clock.
func (g *Game) AddClock(clock int) {
	g.Clock += clock
}

// ChangePossession switches offense and defense.
func (g *Game) ChangePossession() {
	g.Turnover = false
	g.Series = 0.0
	g.Down = 1
	g.Direction = 1 - g.Direction
	g.SwitchOffenseDefense()
}

// SwitchOffenseDefense swaps the offense and defense team indices.
func (g *Game) SwitchOffenseDefense() {
	g.OffenseTeam, g.DefenseTeam = g.DefenseTeam, g.OffenseTeam
}

// GetYardsToTD returns the yards remaining to the end zone.
func (g *Game) GetYardsToTD() float64 {
	if g.Direction == 1 {
		return 100.0 - g.YardLine
	}
	return g.YardLine
}

// IsDirection returns true if direction is 1.
func (g *Game) IsDirection() bool {
	return g.Direction == 1
}

// SetDirection sets the direction from a boolean.
func (g *Game) SetDirection(dir bool) {
	if dir {
		g.Direction = 1
	} else {
		g.Direction = 0
	}
}

// AddDown increments the down counter.
func (g *Game) AddDown() {
	g.Down++
}

// SetTurnover marks the play as a turnover.
func (g *Game) SetTurnover() {
	g.Turnover = true
}

// IsTied returns true if the game is tied.
func (g *Game) IsTied() bool {
	return g.Score[TeamA] == g.Score[TeamB]
}

// LogScore prints the current score.
func LogScore(g *Game) {
	fmt.Printf("%s %d %s %d\n",
		g.Teams[TeamA].Name, g.Score[TeamA],
		g.Teams[TeamB].Name, g.Score[TeamB])
}

// LogOffenseDefense prints the current offense/defense.
func LogOffenseDefense(g *Game) {
	fmt.Printf("Offense %s    Defense %s\n",
		g.Teams[g.OffenseTeam].Name,
		g.Teams[g.DefenseTeam].Name)
}

// RunOffensivePlay executes an offensive play (defensive play is a placeholder).
func (g *Game) RunOffensivePlay(oplay OffensivePlayInterface, dplay DefensivePlayInterface) {
	oplay.Run()
}

// RunSpecialTeamsPlay executes a special teams play.
func (g *Game) RunSpecialTeamsPlay(ostPlay SpecialTeamsPlayInterface, dstPlay SpecialTeamsPlayInterface) {
	ostPlay.Run()
}

// SetScore adds points to the appropriate team based on direction.
func (g *Game) SetScore(score int) {
	if g.Direction == 1 {
		g.Score[TeamA] += score
	} else {
		g.Score[TeamB] += score
	}
}
