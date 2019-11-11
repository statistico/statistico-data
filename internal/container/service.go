package container

import (
	"github.com/statistico/statistico-data/internal/app/grpc"
	"github.com/statistico/statistico-data/internal/fixture"
	"github.com/statistico/statistico-data/internal/result"
)

func (c Container) FixtureService() *grpc.FixtureService {
	return &grpc.FixtureService{
		FixtureRepo: c.FixtureRepository(),
		Handler: fixture.Handler{
			CompetitionRepo: c.CompetitionRepository(),
			RoundRepo:       c.RoundRepository(),
			SeasonRepo:      c.SeasonRepository(),
			TeamRepo:        c.TeamRepository(),
			VenueRepo:       c.VenueRepository(),
			Logger:          c.Logger,
		},
		Logger: c.Logger,
	}
}

func (c Container) ResultService() *grpc.ResultService {
	return &grpc.ResultService{
		FixtureRepo: c.FixtureRepository(),
		ResultRepo:  c.ResultRepository(),
		Handler: result.Handler{
			CompetitionRepo: c.CompetitionRepository(),
			RoundRepo:       c.RoundRepository(),
			SeasonRepo:      c.SeasonRepository(),
			TeamRepo:        c.TeamRepository(),
			VenueRepo:       c.VenueRepository(),
			Logger:          c.Logger,
		},
		Logger: c.Logger,
	}
}

func (c Container) PlayerStatsService() *grpc.PlayerStatsService {
	return &grpc.PlayerStatsService{
		PlayerRepository: c.PlayerStatsRepository(),
		FixtureRepo:      c.FixtureRepository(),
		Logger:           c.Logger,
	}
}

func (c Container) TeamStatsService() *grpc.TeamStatsService {
	return &grpc.TeamStatsService{
		TeamRepository:    c.TeamStatsRepository(),
		FixtureRepository: c.FixtureRepository(),
		Logger:            c.Logger,
	}
}
