package game

import (
	"fmt"
	"math"
	"strings"
)

// Stat type constants
const (
	StatRun              = "RUN"
	StatRunAttempts      = "RUN_ATTEMPTS"
	StatPass             = "PASS"
	StatPassAttempts     = "PASS_ATTEMPTS"
	StatPassCompletions  = "PASS_COMPLETIONS"
	StatPassTDs          = "PASS_TDS"
	StatPassInterceptions = "PASS_INTERCEPTIONS"
	StatFumble           = "FUMBLE"
	StatTeamRun          = "TEAM_RUN"
	StatTeamPass         = "TEAM_PASS"
)

// Stats tracks all game statistics.
type Stats struct {
	PlayerStats [2]map[string]map[string]float64 // [team][statType][playerKey] = value
	TeamsStats  [2]map[string]float64            // [team][statType] = value
	game        *Game
}

// NewStats creates a new Stats instance for the given game.
func NewStats(g *Game) *Stats {
	s := &Stats{game: g}
	for i := 0; i < 2; i++ {
		s.PlayerStats[i] = createGains()
		s.TeamsStats[i] = make(map[string]float64)
	}
	return s
}

func createGains() map[string]map[string]float64 {
	gains := make(map[string]map[string]float64)
	gains[StatRun] = make(map[string]float64)
	gains[StatRunAttempts] = make(map[string]float64)
	gains[StatFumble] = make(map[string]float64)
	gains[StatPassAttempts] = make(map[string]float64)
	gains[StatPassCompletions] = make(map[string]float64)
	gains[StatPassInterceptions] = make(map[string]float64)
	gains[StatPass] = make(map[string]float64)
	return gains
}

// AddPlayerStats adds a stat value for a player on a team.
func (s *Stats) AddPlayerStats(team int, key string, player string, value float64) {
	s.PlayerStats[team][key][player] += value
}

// AddTeamStats adds a stat value for a team.
func (s *Stats) AddTeamStats(team int, key string, value float64) {
	s.TeamsStats[team][key] += value
}

// AddPassStats records passing statistics.
func (s *Stats) AddPassStats(team int, player string, attempt int, completion int, yards float64) {
	s.AddPlayerStats(team, StatPassAttempts, player, float64(attempt))
	s.AddPlayerStats(team, StatPassCompletions, player, float64(completion))
	s.AddPlayerStats(team, StatPass, player, yards)
	s.AddTeamStats(team, StatTeamPass, yards)
}

// AddRunStats records rushing statistics.
func (s *Stats) AddRunStats(team int, player string, attempts int, yards float64) {
	s.AddPlayerStats(team, StatRun, player, yards)
	s.AddPlayerStats(team, StatRunAttempts, player, float64(attempts))
	s.AddTeamStats(team, StatTeamRun, yards)
}

// AddFumble records a fumble for a player.
func (s *Stats) AddFumble(team int, player string) {
	s.AddPlayerStats(team, StatFumble, player, 1)
}

// PrintStats prints all game statistics.
func (s *Stats) PrintStats() {
	fmt.Println("\n=== GAME STATISTICS ===")
	s.printPlayerStats(StatRun, TeamA)
	s.printPlayerStats(StatRunAttempts, TeamA)
	s.printPlayerStats(StatRun, TeamB)
	s.printPlayerStats(StatRunAttempts, TeamB)
	s.printPlayerStats(StatPass, TeamA)
	s.printPlayerStats(StatPassCompletions, TeamA)
	s.printPlayerStats(StatPassAttempts, TeamA)
	s.printPlayerStats(StatPassInterceptions, TeamA)
	s.printPlayerStats(StatPass, TeamB)
	s.printPlayerStats(StatPassCompletions, TeamB)
	s.printPlayerStats(StatPassAttempts, TeamB)
	s.printPlayerStats(StatPassInterceptions, TeamB)
	s.printPlayerStats(StatFumble, TeamA)
	s.printPlayerStats(StatFumble, TeamB)
	s.printTeamYards(TeamA, StatTeamRun)
	s.printTeamYards(TeamB, StatTeamRun)
	s.printTeamYards(TeamA, StatTeamPass)
	s.printTeamYards(TeamB, StatTeamPass)
}

func (s *Stats) printTeamYards(team int, key string) {
	val := s.TeamsStats[team][key]
	teamName := rightPad(s.game.Teams[team].Name, 6)
	keyPad := rightPad(key, 12)
	fmt.Printf("%s %s %10.1f\n", teamName, keyPad, roundTo(val, 1))
}

func (s *Stats) printPlayerStats(posKey string, team int) {
	for player, val := range s.PlayerStats[team][posKey] {
		teamName := rightPad(s.game.Teams[team].Name, 6)
		keyPad := rightPad(posKey, 16)
		fmt.Printf("%s %s %6.1f  %s\n", teamName, keyPad, roundTo(val, 1), player)
	}
}

func rightPad(s string, length int) string {
	if len(s) >= length {
		return s
	}
	return s + strings.Repeat(" ", length-len(s))
}

func roundTo(val float64, places int) float64 {
	pow := math.Pow(10, float64(places))
	return math.Round(val*pow) / pow
}
