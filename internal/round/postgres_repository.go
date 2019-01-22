package round

import (
	"database/sql"
	"github.com/joesweeny/statshub/internal/model"
	_ "github.com/lib/pq"
	"time"
	"github.com/pkg/errors"
)

var ErrNotFound = errors.New("not found")

type PostgresRoundRepository struct {
	Connection *sql.DB
}

func (p *PostgresRoundRepository) Insert(r *model.Round) error {
	query := `
	INSERT INTO sportmonks_round (id, name, season_id, start_date, end_date, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := p.Connection.Exec(
		query,
		r.ID,
		r.Name,
		r.SeasonID,
		r.StartDate.Unix(),
		r.EndDate.Unix(),
		r.CreatedAt.Unix(),
		r.UpdatedAt.Unix(),
	)

	return err
}

func (p *PostgresRoundRepository) GetById(id int) (*model.Round, error) {
	query := `SELECT * FROM sportmonks_round where id = $1`
	row := p.Connection.QueryRow(query, id)

	return rowToRound(row)
}

func rowToRound(r *sql.Row) (*model.Round, error) {
	var id int
	var name string
	var season int
	var start int64
	var end int64
	var created int64
	var updated int64

	m := model.Round{}

	err := r.Scan(&id, &name, &season, &start, &end, &created, &updated)

	if err != nil {
		return &m, ErrNotFound
	}

	m.ID = id
	m.Name = name
	m.SeasonID = season
	m.StartDate = time.Unix(start, 0)
	m.EndDate = time.Unix(end, 0)
	m.CreatedAt = time.Unix(created, 0)
	m.UpdatedAt = time.Unix(updated, 0)

	return &m, nil
}
