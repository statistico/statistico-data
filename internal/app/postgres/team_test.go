package postgres_test

import (
	"github.com/statistico/statistico-data/internal/app"
	"github.com/statistico/statistico-data/internal/app/postgres"
	"github.com/statistico/statistico-data/internal/app/test"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTeamRepository_Insert(t *testing.T) {
	conn, cleanUp := test.GetConnection(t, "sportmonks_team")
	repo := postgres.NewTeamRepository(conn, test.Clock)

	t.Run("increases table count", func(t *testing.T) {
		t.Helper()
		defer cleanUp()

		for i := 1; i < 4; i++ {
			c := newTeam(int64(i))

			if err := repo.Insert(c); err != nil {
				t.Errorf("Error when inserting record into the database: %s", err.Error())
			}

			row := conn.QueryRow("select count(*) from sportmonks_team")

			var count int

			if err := row.Scan(&count); err != nil {
				t.Errorf("Error when scanning rows returned by the database: %s", err.Error())
			}

			assert.Equal(t, i, count)
		}
	})

	t.Run("returns error when ID primary key violates unique constraint", func(t *testing.T) {
		t.Helper()
		defer cleanUp()
		c := newTeam(50)

		if err := repo.Insert(c); err != nil {
			t.Errorf("Test failed, expected nil, got %s", err)
		}

		if e := repo.Insert(c); e == nil {
			t.Fatalf("Test failed, expected %s, got nil", e)
		}
	})
}

func TestTeamRepository_ByID(t *testing.T) {
	conn, cleanUp := test.GetConnection(t, "sportmonks_team")
	repo := postgres.NewTeamRepository(conn, test.Clock)

	t.Run("team can be retrieved by ID", func(t *testing.T) {
		t.Helper()
		defer cleanUp()

		m := newTeam(43)

		if err := repo.Insert(m); err != nil {
			t.Errorf("Error when inserting record into the database: %s", err.Error())
		}

		r, err := repo.ByID(43)

		if err != nil {
			t.Errorf("Error when retrieving a record from the database: %s", err.Error())
		}

		a := assert.New(t)

		a.Equal(int64(43), r.ID)
		a.Equal("West Ham United", r.Name)
		a.Equal(int64(560), r.VenueID)
		a.Equal(false, r.NationalTeam)
		a.Nil(r.ShortCode)
		a.Nil(r.CountryID)
		a.Nil(r.Founded)
		a.Nil(r.Logo)
		a.Equal("2019-01-14 11:25:00 +0000 UTC", r.CreatedAt.String())
		a.Equal("2019-01-14 11:25:00 +0000 UTC", r.UpdatedAt.String())
	})

	t.Run("returns error if team does not exist", func(t *testing.T) {
		t.Helper()
		defer cleanUp()

		if _, err := repo.ByID(4); err == nil {
			t.Fatalf("Test failed, expected %v, got nil", err)
		}
	})
}

func TestTeamRepository_Update(t *testing.T) {
	conn, cleanUp := test.GetConnection(t, "sportmonks_team")
	repo := postgres.NewTeamRepository(conn, test.Clock)

	t.Run("modifies existing team", func(t *testing.T) {
		t.Helper()
		defer cleanUp()

		m := newTeam(43)

		if err := repo.Insert(m); err != nil {
			t.Errorf("Error when inserting record into the database: %s", err.Error())
		}

		var shortCode = "WHU"
		var founded = 1898
		var logo = "http://path.to/logo"
		var d = time.Date(2019, 01, 14, 11, 25, 00, 00, time.UTC)

		m.Name = "West Ham London Boooo"
		m.ShortCode = &shortCode
		m.CountryID = int64(5)
		m.Founded = &founded
		m.Logo = &logo
		m.UpdatedAt = d

		if err := repo.Update(m); err != nil {
			t.Errorf("Error when updating a record in the database: %s", err.Error())
		}

		r, err := repo.ByID(43)

		if err != nil {
			t.Errorf("Error when updating a record in the database: %s", err.Error())
		}

		a := assert.New(t)

		a.Equal(int64(43), r.ID)
		a.Equal("West Ham London Boooo", r.Name)
		a.Equal(int64(560), r.VenueID)
		a.Equal(false, r.NationalTeam)
		a.Equal("WHU", *r.ShortCode)
		a.Equal(int64(5), r.CountryID)
		a.Equal(1898, *r.Founded)
		a.Equal("http://path.to/logo", *r.Logo)
		a.Equal("2019-01-14 11:25:00 +0000 UTC", r.CreatedAt.String())
		a.Equal("2019-01-14 11:25:00 +0000 UTC", r.UpdatedAt.String())
	})

	t.Run("returns an error if team does not exist", func(t *testing.T) {
		t.Helper()
		defer cleanUp()
		c := newTeam(146)

		err := repo.Update(c)

		if err == nil {
			t.Fatalf("Test failed, expected %v, got nil", err)
		}
	})
}

func newTeam(id int64) *app.Team {
	return &app.Team{
		ID:           id,
		Name:         "West Ham United",
		VenueID:      560,
		NationalTeam: false,
		CreatedAt:    time.Unix(1546965200, 0),
		UpdatedAt:    time.Unix(1546965200, 0),
	}
}