package fixture

import (
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/statistico/statistico-data/internal/competition"
	"github.com/statistico/statistico-data/internal/model"
	"github.com/statistico/statistico-data/internal/round"
	"github.com/statistico/statistico-data/internal/season"
	"github.com/statistico/statistico-data/internal/team"
	"github.com/statistico/statistico-data/internal/venue"
	pbCompetition "github.com/statistico/statistico-data/proto/competition"
	pbFixture "github.com/statistico/statistico-data/proto/fixture"
	pbRound "github.com/statistico/statistico-data/proto/round"
	pbSeason "github.com/statistico/statistico-data/proto/season"
	pbTeam "github.com/statistico/statistico-data/proto/team"
	pbVenue "github.com/statistico/statistico-data/proto/venue"
	"time"
)

type Handler struct {
	CompetitionRepo competition.Repository
	RoundRepo       round.Repository
	SeasonRepo      season.Repository
	TeamRepo        team.Repository
	VenueRepo       venue.Repository
}

func (h Handler) HandleFixture(f *model.Fixture) (*pbFixture.Fixture, error) {
	s, err := h.SeasonRepo.Id(f.SeasonID)

	if err != nil {
		return nil, err
	}

	c, err := h.CompetitionRepo.GetById(s.LeagueID)

	if err != nil {
		return nil, err
	}

	home, err := h.TeamRepo.GetById(f.HomeTeamID)

	if err != nil {
		return nil, err
	}

	away, err := h.TeamRepo.GetById(f.AwayTeamID)

	if err != nil {
		return nil, err
	}

	proto := pbFixture.Fixture{
		Id:          int64(f.ID),
		Competition: competitionToProto(c),
		Season:      seasonToProto(s),
		HomeTeam:    teamToProto(home),
		AwayTeam:    teamToProto(away),
		DateTime:    f.Date.Unix(),
	}

	if f.RoundID != nil {
		r, err := h.RoundRepo.GetById(*f.RoundID)

		if err != nil {
			return nil, err
		}

		proto.Round = roundToProto(r)
	}

	if f.VenueID != nil {
		v, err := h.VenueRepo.GetById(*f.VenueID)

		if err != nil {
			return nil, err
		}

		proto.Venue = venueToProto(v)
	}

	if f.RefereeID != nil {
		ref := wrappers.Int64Value{}
		ref.Value = int64(*f.RefereeID)
		proto.RefereeId = &ref
	}

	return &proto, nil
}

func teamToProto(t *model.Team) *pbTeam.Team {
	var x pbTeam.Team
	x.Id = int64(t.ID)
	x.Name = t.Name

	return &x
}

func competitionToProto(c *model.Competition) *pbCompetition.Competition {
	var x pbCompetition.Competition
	x.Id = int64(c.ID)
	x.Name = c.Name
	x.IsCup = &wrappers.BoolValue{
		Value: c.IsCup,
	}

	return &x
}

func roundToProto(r *model.Round) *pbRound.Round {
	return &pbRound.Round{
		Id:        int64(r.ID),
		Name:      r.Name,
		SeasonId:  int64(r.SeasonID),
		StartDate: r.StartDate.Format(time.RFC3339),
		EndDate:   r.EndDate.Format(time.RFC3339),
	}
}

func seasonToProto(s *model.Season) *pbSeason.Season {
	var x pbSeason.Season
	x.Id = int64(s.ID)
	x.Name = s.Name
	x.IsCurrent = &wrappers.BoolValue{
		Value: s.IsCurrent,
	}

	return &x
}

func venueToProto(v *model.Venue) *pbVenue.Venue {
	id := wrappers.Int64Value{}
	id.Value = int64(v.ID)
	name := wrappers.StringValue{}
	name.Value = v.Name

	ven := pbVenue.Venue{}
	ven.Id = &id
	ven.Name = &name

	return &ven
}
