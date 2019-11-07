package app

import "time"

// TeamStats domain entity.
type TeamStats struct {
	FixtureID     uint64 `json:"fixture_id"`
	TeamID        uint64 `json:"team_id"`
	TeamShots     `json:"shots"`
	TeamPasses    `json:"passes"`
	TeamAttacks   `json:"attacks"`
	Fouls         *int      `json:"fouls"`
	Corners       *int      `json:"corners"`
	Offsides      *int      `json:"offsides"`
	Possession    *int      `json:"possession"`
	YellowCards   *int      `json:"yellow_cards"`
	RedCards      *int      `json:"red_cards"`
	Saves         *int      `json:"saves"`
	Substitutions *int      `json:"substitutions"`
	GoalKicks     *int      `json:"goal_kicks"`
	GoalAttempts  *int      `json:"goal_attempts"`
	FreeKicks     *int      `json:"free_kicks"`
	ThrowIns      *int      `json:"throw_ins"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// TeamShots domain sub entity.
type TeamShots struct {
	Total      *int `json:"total"`
	OnGoal     *int `json:"on_goal"`
	OffGoal    *int `json:"off_goal"`
	Blocked    *int `json:"blocked"`
	InsideBox  *int `json:"inside_box"`
	OutsideBox *int `json:"outside_box"`
}

// TeamPasses domain sub entity.
type TeamPasses struct {
	Total      *int `json:"total"`
	Accuracy   *int `json:"accuracy"`
	Percentage *int `json:"percentage"`
}

// TeamAttacks domain sub entity.
type TeamAttacks struct {
	Total     *int `json:"total"`
	Dangerous *int `json:"dangerous"`
}

// TeamStatsRepository provides an interface to persist TeamStats domain struct objects to a storage engine.
type TeamStatsRepository interface {
	InsertTeamStats(m *TeamStats) error
	UpdateTeamStats(m *TeamStats) error
	ByFixtureAndTeam(fixtureID, teamID uint64) (*TeamStats, error)
}
