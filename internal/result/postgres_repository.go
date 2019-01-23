package result

import (
	"errors"
	"database/sql"
	"github.com/joesweeny/statshub/internal/model"
	_ "github.com/lib/pq"
	"time"
)

var ErrNotFound = errors.New("not found")

type PostgresResultRepository struct {
	Connection *sql.DB
}

func (p *PostgresResultRepository) Insert(r *model.Result) error {
	query := `
	INSERT INTO sportmonks_result (fixture_id, pitch_condition, home_formation, away_formation, home_score, away_score,
	home_pen_score, away_pen_score, half_time_score, full_time_score, extra_time_score, home_league_position,
	away_league_position, minutes, seconds, added_time, extra_time, injury_time, created_at, updated_at) VALUES ($1, $2, 
	$3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20)`

	_, err := p.Connection.Exec(
		query,
		r.FixtureID,
		r.PitchCondition,
		r.HomeFormation,
		r.AwayFormation,
		r.HomeScore,
		r.AwayScore,
		r.HomePenScore,
		r.AwayPenScore,
		r.HalfTimeScore,
		r.FullTimeScore,
		r.ExtraTimeScore,
		r.HomeLeaguePosition,
		r.AwayLeaguePosition,
		r.Minutes,
		r.Seconds,
		r.AddedTime,
		r.ExtraTime,
		r.InjuryTime,
		r.CreatedAt.Unix(),
		r.UpdatedAt.Unix(),
	)

	return err
}

func (p *PostgresResultRepository) Update(r *model.Result) error {
	_, err := p.GetByFixtureId(r.FixtureID)

	if err != nil {
		return err
	}

	query := `
	UPDATE sportmonks_result SET pitch_condition = $2, home_formation = $3, away_formation = $4, home_score = $5, 
    away_score = $6, home_pen_score = $7, away_pen_score = $8, half_time_score = $9, full_time_score = $10, 
	extra_time_score = $11, home_league_position = $12, away_league_position = $13, minutes = $14, seconds = $15, 
	added_time = $16, extra_time = $17, injury_time = $18, updated_at = $19 WHERE fixture_id = $1`

	_, err = p.Connection.Exec(
		query,
		r.FixtureID,
		r.PitchCondition,
		r.HomeFormation,
		r.AwayFormation,
		r.HomeScore,
		r.AwayScore,
		r.HomePenScore,
		r.AwayPenScore,
		r.HalfTimeScore,
		r.FullTimeScore,
		r.ExtraTimeScore,
		r.HomeLeaguePosition,
		r.AwayLeaguePosition,
		r.Minutes,
		r.Seconds,
		r.AddedTime,
		r.ExtraTime,
		r.InjuryTime,
		r.UpdatedAt.Unix(),
	)

	return err
}

func (p *PostgresResultRepository) GetByFixtureId(id int) (*model.Result, error) {
	query := `SELECT * FROM sportmonks_result where fixture_id = $1`
	row := p.Connection.QueryRow(query, id)

	return rowToResult(row)
}

func rowToResult(r *sql.Row) (*model.Result, error) {
	var fixtureId         int
	var pitch             *string
	var homeFormation     *string
	var awayFormation     *string
	var homeScore         *int
	var awayScore         *int
	var homePenScore      *int
	var awayPenScore      *int
	var halfTimeScore     *string
	var fullTimeScore     *string
	var extraTimeScore    *string
	var homePosition      *int
	var awayPosition      *int
	var mins              *int
	var secs              *int
	var added             *int
	var extra             *int
	var injury            *int
	var created           int64
	var updated           int64

	m := model.Result{}

	err := r.Scan(&fixtureId, &pitch, &homeFormation, &awayFormation, &homeScore, &awayScore, &homePenScore, &awayPenScore,
		&halfTimeScore, &fullTimeScore, &extraTimeScore, &homePosition, &awayPosition, &mins, &secs, &added, &extra,
		&injury, &created, &updated)

	if err != nil {
		return &m, ErrNotFound
	}

	m.FixtureID = fixtureId
	m.PitchCondition = pitch
	m.HomeFormation = homeFormation
	m.AwayFormation = awayFormation
	m.HomeScore = homeScore
	m.AwayScore = awayScore
	m.HomePenScore = homePenScore
	m.AwayPenScore = awayPenScore
	m.HalfTimeScore = halfTimeScore
	m.FullTimeScore = fullTimeScore
	m.ExtraTimeScore = extraTimeScore
	m.HomeLeaguePosition = homePosition
	m.AwayLeaguePosition = awayPosition
	m.Minutes = mins
	m.Seconds = secs
	m.AddedTime = added
	m.ExtraTime = extra
	m.InjuryTime = injury
	m.CreatedAt = time.Unix(created, 0)
	m.UpdatedAt = time.Unix(updated, 0)

	return &m, nil
}