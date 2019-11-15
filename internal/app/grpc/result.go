package grpc

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/statistico/statistico-data/internal/app"
	"github.com/statistico/statistico-data/internal/app/grpc/factory"
	"github.com/statistico/statistico-data/internal/app/grpc/proto"
	"time"
)

const maxLimit = 10000

var ErrTimeParse = errors.New("unable to parse date provided in Request")

type ResultService struct {
	fixtureRepo app.FixtureRepository
	factory *factory.ResultFactory
	logger *logrus.Logger
}

func (s ResultService) GetHistoricalResultsForFixture(r *proto.HistoricalResultRequest, stream proto.ResultService_GetHistoricalResultsForFixtureServer) error {
	date, err := time.Parse(time.RFC3339, r.DateBefore)

	if err != nil {
		return ErrTimeParse
	}

	limit := uint64(r.Limit)

	query := app.FixtureRepositoryQuery{
		HomeTeamID: &r.HomeTeamId,
		AwayTeamID: &r.AwayTeamId,
		DateTo: &date,
		Limit:      &limit,
	}

	fixtures, err := s.fixtureRepo.Get(query)

	if err != nil {
		s.logger.Printf("Error retrieving Fixture(s) in Result Service. Error: %s", err.Error())
		return fmt.Errorf("server error: Unable to fulfil Request")
	}

	return s.sendResults(fixtures, stream)
}

func (s ResultService) GetResultsForTeam(r *proto.TeamRequest, stream proto.ResultService_GetResultsForTeamServer) error {
	date, err := time.Parse(time.RFC3339, r.DateBefore)

	if err != nil {
		return ErrTimeParse
	}

	limit := r.Limit.GetValue()

	if limit == 0 {
		limit = maxLimit
	}

	fixtures, err := s.fixtureRepo.ByTeamID(uint64(r.TeamId), limit, date)

	if err != nil {
		s.logger.Printf("Error retrieving Fixture(s) in Result Service. Error: %s", err.Error())
		return fmt.Errorf("server error: Unable to fulfil Request")
	}

	return s.sendResults(fixtures, stream)
}

func (s ResultService) GetResultsForSeason(r *proto.SeasonRequest, stream proto.ResultService_GetResultsForSeasonServer) error {
	date, err := time.Parse(time.RFC3339, r.DateBefore)

	if err != nil {
		return ErrTimeParse
	}

	id := uint64(r.SeasonId)

	query := app.FixtureRepositoryQuery{
		SeasonID:   &id,
		DateTo: &date,
	}

	fixtures, err := s.fixtureRepo.Get(query)

	if err != nil {
		s.logger.Printf("Error retrieving Fixture(s) in Result Service. Error: %s", err.Error())
		return fmt.Errorf("server error: Unable to fulfil Request")
	}

	return s.sendResults(fixtures, stream)
}

func (s ResultService) sendResults(f []app.Fixture, stream proto.ResultService_GetResultsForTeamServer) error {
	for _, fix := range f {
		x, err := s.factory.BuildResult(&fix)

		if err != nil {
			s.logger.Warnf("Error hydrating Result. Error: %s", err.Error())
			return err
		}

		if err := stream.Send(x); err != nil {
			s.logger.Warnf("Error streaming Result back to client. Error: %s", err.Error())
			continue
		}
	}

	return nil
}

func NewResultService(r app.FixtureRepository, f *factory.ResultFactory, log *logrus.Logger) *ResultService {
	return &ResultService{fixtureRepo: r, factory: f, logger: log}
}
