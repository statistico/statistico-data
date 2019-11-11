package player_stats

import (
	"context"
	"errors"
	"fmt"
	"github.com/statistico/statistico-data/internal/app"
	"github.com/statistico/statistico-data/internal/app/proto"
	"log"
)

type Service struct {
	PlayerRepository app.PlayerStatsRepository
	FixtureRepo      app.FixtureRepository
	Logger           *log.Logger
}

func (s Service) GetPlayerStatsForFixture(c context.Context, r *proto.FixtureRequest) (*proto.PlayerStatsResponse, error) {
	fix, err := s.FixtureRepo.ByID(r.FixtureId)

	if err != nil {
		m := fmt.Sprintf("Fixture with ID %d does not exist", r.FixtureId)
		return nil, errors.New(m)
	}

	res := proto.PlayerStatsResponse{}

	home, err := s.PlayerRepository.ByFixtureAndTeam(fix.ID, fix.HomeTeamID)

	if err != nil {
		e := fmt.Errorf("error when retrieving player stats: FixtureID %d, Home Team ID %d", fix.ID, fix.HomeTeamID)
		s.Logger.Println(e)
		return nil, e
	}

	res.HomeTeam = HandlePlayerStats(home)

	away, err := s.PlayerRepository.ByFixtureAndTeam(uint64(fix.ID), uint64(fix.AwayTeamID))

	if err != nil {
		e := fmt.Errorf("error when retrieving player stats: FixtureID %d, Away Team ID %d", fix.ID, fix.HomeTeamID)
		s.Logger.Println(e)
		return nil, e
	}

	res.AwayTeam = HandlePlayerStats(away)

	return &res, nil
}

func (s Service) GetLineUpForFixture(c context.Context, r *proto.FixtureRequest) (*proto.LineupResponse, error) {
	fix, err := s.FixtureRepo.ByID(r.FixtureId)

	if err != nil {
		m := fmt.Sprintf("Fixture with ID %d does not exist", r.FixtureId)
		return nil, errors.New(m)
	}

	res := proto.LineupResponse{}

	home, err := s.PlayerRepository.ByFixtureAndTeam(uint64(fix.ID), uint64(fix.HomeTeamID))

	if err != nil {
		e := fmt.Errorf("error when retrieving player stats: FixtureID %d, Home Team ID %d", fix.ID, fix.HomeTeamID)
		s.Logger.Println(e)
		return nil, e
	}

	homeLineup := proto.Lineup{
		Start: HandleStartingLineupPlayers(home),
		Bench: HandleSubstituteLineupPlayers(home),
	}

	res.HomeTeam = &homeLineup

	away, err := s.PlayerRepository.ByFixtureAndTeam(uint64(fix.ID), uint64(fix.AwayTeamID))

	if err != nil {
		e := fmt.Errorf("error when retrieving player stats: FixtureID %d, Away Team ID %d", fix.ID, fix.AwayTeamID)
		s.Logger.Println(e)
		return nil, e
	}

	awayLineup := proto.Lineup{
		Start: HandleStartingLineupPlayers(away),
		Bench: HandleSubstituteLineupPlayers(away),
	}

	res.AwayTeam = &awayLineup

	return &res, nil
}
