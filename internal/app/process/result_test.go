package process_test

import (
	"errors"
	"github.com/jonboulle/clockwork"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/statistico/statistico-football-data/internal/app"
	"github.com/statistico/statistico-football-data/internal/app/mock"
	"github.com/statistico/statistico-football-data/internal/app/process"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResultProcessor_Process(t *testing.T) {
	t.Run("inserts new result into repository when processing result by season id command", func(t *testing.T) {
		t.Helper()

		resultRepo := new(mock.ResultRepository)
		seasonRepo := new(mock.SeasonRepository)
		requester := new(mock.ResultRequester)
		clock := clockwork.NewFakeClock()
		logger, hook := test.NewNullLogger()

		processor := process.NewResultProcessor(resultRepo, seasonRepo, requester, clock, logger)

		done := make(chan bool)

		res := newResult(34)

		results := make([]app.Result, 1)
		results[0] = res

		ch := resultChannel(results)

		requester.On("ResultsBySeasonIDs", []uint64{34}).Return(ch)
		resultRepo.On("ByFixtureID", uint64(34)).Return(&app.Result{}, errors.New("not found"))
		resultRepo.On("Insert", &res).Return(nil)
		processor.Process("results:by-season-id", "34", done)

		<-done

		requester.AssertExpectations(t)
		resultRepo.AssertExpectations(t)
		assert.Nil(t, hook.LastEntry())
	})

	t.Run("updates existing result into repository when processing result by season id command", func(t *testing.T) {
		t.Helper()

		resultRepo := new(mock.ResultRepository)
		seasonRepo := new(mock.SeasonRepository)
		requester := new(mock.ResultRequester)
		clock := clockwork.NewFakeClock()
		logger, hook := test.NewNullLogger()

		processor := process.NewResultProcessor(resultRepo, seasonRepo, requester, clock, logger)

		done := make(chan bool)

		res := newResult(34)

		results := make([]app.Result, 1)
		results[0] = res

		ch := resultChannel(results)

		requester.On("ResultsBySeasonIDs", []uint64{34}).Return(ch)
		resultRepo.On("ByFixtureID", uint64(34)).Return(&app.Result{}, nil)
		resultRepo.On("Update", &res).Return(nil)
		processor.Process("results:by-season-id", "34", done)

		<-done

		requester.AssertExpectations(t)
		resultRepo.AssertExpectations(t)
		assert.Nil(t, hook.LastEntry())
	})

	t.Run("logs error when unable to insert result into repository when processing result by season id command", func(t *testing.T) {
		t.Helper()

		resultRepo := new(mock.ResultRepository)
		seasonRepo := new(mock.SeasonRepository)
		requester := new(mock.ResultRequester)
		clock := clockwork.NewFakeClock()
		logger, hook := test.NewNullLogger()

		processor := process.NewResultProcessor(resultRepo, seasonRepo, requester, clock, logger)

		done := make(chan bool)

		res := newResult(34)

		results := make([]app.Result, 1)
		results[0] = res

		ch := resultChannel(results)

		requester.On("ResultsBySeasonIDs", []uint64{34}).Return(ch)
		resultRepo.On("ByFixtureID", uint64(34)).Return(&app.Result{}, errors.New("not found"))
		resultRepo.On("Insert", &res).Return(errors.New("error occurred"))
		processor.Process("results:by-season-id", "34", done)

		<-done

		requester.AssertExpectations(t)
		resultRepo.AssertExpectations(t)
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	})

	t.Run("logs error when unable to update result into repository when processing result by season id command", func(t *testing.T) {
		t.Helper()

		resultRepo := new(mock.ResultRepository)
		seasonRepo := new(mock.SeasonRepository)
		requester := new(mock.ResultRequester)
		clock := clockwork.NewFakeClock()
		logger, hook := test.NewNullLogger()

		processor := process.NewResultProcessor(resultRepo, seasonRepo, requester, clock, logger)

		done := make(chan bool)

		res := newResult(34)

		results := make([]app.Result, 1)
		results[0] = res

		ch := resultChannel(results)

		requester.On("ResultsBySeasonIDs", []uint64{34}).Return(ch)
		resultRepo.On("ByFixtureID", uint64(34)).Return(&app.Result{}, nil)
		resultRepo.On("Update", &res).Return(errors.New("error occurred"))
		processor.Process("results:by-season-id", "34", done)

		<-done

		requester.AssertExpectations(t)
		resultRepo.AssertExpectations(t)
		seasonRepo.AssertExpectations(t)
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	})

	t.Run("inserts new result into repository when processing result by competition id command", func(t *testing.T) {
		t.Helper()

		resultRepo := new(mock.ResultRepository)
		seasonRepo := new(mock.SeasonRepository)
		requester := new(mock.ResultRequester)
		clock := clockwork.NewFakeClock()
		logger, hook := test.NewNullLogger()

		processor := process.NewResultProcessor(resultRepo, seasonRepo, requester, clock, logger)

		done := make(chan bool)

		res := newResult(34)

		results := make([]app.Result, 1)
		results[0] = res

		ch := resultChannel(results)

		seasonRepo.On("ByCompetitionId", uint64(5), "name_asc").Return(
			[]app.Season{
				*newSeason(1, false),
				*newSeason(2, true),
			},
			nil,
		)

		requester.On("ResultsBySeasonIDs", []uint64{1, 2}).Return(ch)
		resultRepo.On("ByFixtureID", uint64(34)).Return(&app.Result{}, errors.New("not found"))
		resultRepo.On("Insert", &res).Return(nil)
		processor.Process("results:by-competition-id", "5", done)

		<-done

		requester.AssertExpectations(t)
		resultRepo.AssertExpectations(t)
		assert.Nil(t, hook.LastEntry())
	})

	t.Run("updates existing result into repository when processing result by competition id command", func(t *testing.T) {
		t.Helper()

		resultRepo := new(mock.ResultRepository)
		seasonRepo := new(mock.SeasonRepository)
		requester := new(mock.ResultRequester)
		clock := clockwork.NewFakeClock()
		logger, hook := test.NewNullLogger()

		processor := process.NewResultProcessor(resultRepo, seasonRepo, requester, clock, logger)

		done := make(chan bool)

		res := newResult(34)

		results := make([]app.Result, 1)
		results[0] = res

		ch := resultChannel(results)

		seasonRepo.On("ByCompetitionId", uint64(5), "name_asc").Return(
			[]app.Season{
				*newSeason(1, false),
				*newSeason(2, true),
			},
			nil,
		)

		requester.On("ResultsBySeasonIDs", []uint64{1, 2}).Return(ch)
		resultRepo.On("ByFixtureID", uint64(34)).Return(&app.Result{}, nil)
		resultRepo.On("Update", &res).Return(nil)
		processor.Process("results:by-competition-id", "5", done)

		<-done

		requester.AssertExpectations(t)
		resultRepo.AssertExpectations(t)
		seasonRepo.AssertExpectations(t)
		assert.Nil(t, hook.LastEntry())
	})

	t.Run("logs error when unable to insert result into repository when processing result by competition id command", func(t *testing.T) {
		t.Helper()

		resultRepo := new(mock.ResultRepository)
		seasonRepo := new(mock.SeasonRepository)
		requester := new(mock.ResultRequester)
		clock := clockwork.NewFakeClock()
		logger, hook := test.NewNullLogger()

		processor := process.NewResultProcessor(resultRepo, seasonRepo, requester, clock, logger)

		done := make(chan bool)

		res := newResult(34)

		results := make([]app.Result, 1)
		results[0] = res

		ch := resultChannel(results)

		seasonRepo.On("ByCompetitionId", uint64(5), "name_asc").Return(
			[]app.Season{
				*newSeason(1, false),
				*newSeason(2, true),
			},
			nil,
		)

		requester.On("ResultsBySeasonIDs", []uint64{1, 2}).Return(ch)
		resultRepo.On("ByFixtureID", uint64(34)).Return(&app.Result{}, errors.New("not found"))
		resultRepo.On("Insert", &res).Return(errors.New("error occurred"))
		processor.Process("results:by-competition-id", "5", done)

		<-done

		requester.AssertExpectations(t)
		resultRepo.AssertExpectations(t)
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	})

	t.Run("logs error when unable to update result into repository when processing result by competition id command", func(t *testing.T) {
		t.Helper()

		resultRepo := new(mock.ResultRepository)
		seasonRepo := new(mock.SeasonRepository)
		requester := new(mock.ResultRequester)
		clock := clockwork.NewFakeClock()
		logger, hook := test.NewNullLogger()

		processor := process.NewResultProcessor(resultRepo, seasonRepo, requester, clock, logger)

		done := make(chan bool)

		res := newResult(34)

		results := make([]app.Result, 1)
		results[0] = res

		ch := resultChannel(results)

		seasonRepo.On("ByCompetitionId", uint64(5), "name_asc").Return(
			[]app.Season{
				*newSeason(1, false),
				*newSeason(2, true),
			},
			nil,
		)

		requester.On("ResultsBySeasonIDs", []uint64{1, 2}).Return(ch)
		resultRepo.On("ByFixtureID", uint64(34)).Return(&app.Result{}, nil)
		resultRepo.On("Update", &res).Return(errors.New("error occurred"))
		processor.Process("results:by-competition-id", "5", done)

		<-done

		requester.AssertExpectations(t)
		resultRepo.AssertExpectations(t)
		seasonRepo.AssertExpectations(t)
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	})

	t.Run("inserts new result into repository when processing result current season command", func(t *testing.T) {
		t.Helper()

		resultRepo := new(mock.ResultRepository)
		seasonRepo := new(mock.SeasonRepository)
		requester := new(mock.ResultRequester)
		clock := clockwork.NewFakeClock()
		logger, hook := test.NewNullLogger()

		processor := process.NewResultProcessor(resultRepo, seasonRepo, requester, clock, logger)

		done := make(chan bool)

		res := newResult(34)

		results := make([]app.Result, 1)
		results[0] = res

		ch := resultChannel(results)

		seasonRepo.On("CurrentSeasonIDs").Return([]uint64{1, 2}, nil)

		requester.On("ResultsBySeasonIDs", []uint64{1, 2}).Return(ch)
		resultRepo.On("ByFixtureID", uint64(34)).Return(&app.Result{}, errors.New("not found"))
		resultRepo.On("Insert", &res).Return(nil)
		processor.Process("results:current-season", "5", done)

		<-done

		requester.AssertExpectations(t)
		resultRepo.AssertExpectations(t)
		assert.Nil(t, hook.LastEntry())
	})

	t.Run("updates existing result into repository when processing result current season command", func(t *testing.T) {
		t.Helper()

		resultRepo := new(mock.ResultRepository)
		seasonRepo := new(mock.SeasonRepository)
		requester := new(mock.ResultRequester)
		clock := clockwork.NewFakeClock()
		logger, hook := test.NewNullLogger()

		processor := process.NewResultProcessor(resultRepo, seasonRepo, requester, clock, logger)

		done := make(chan bool)

		res := newResult(34)

		results := make([]app.Result, 1)
		results[0] = res

		ch := resultChannel(results)

		seasonRepo.On("CurrentSeasonIDs").Return([]uint64{1, 2}, nil)

		requester.On("ResultsBySeasonIDs", []uint64{1, 2}).Return(ch)
		resultRepo.On("ByFixtureID", uint64(34)).Return(&app.Result{}, nil)
		resultRepo.On("Update", &res).Return(nil)
		processor.Process("results:current-season", "5", done)

		<-done

		requester.AssertExpectations(t)
		resultRepo.AssertExpectations(t)
		seasonRepo.AssertExpectations(t)
		assert.Nil(t, hook.LastEntry())
	})

	t.Run("logs error when unable to insert result into repository when processing result current season command", func(t *testing.T) {
		t.Helper()

		resultRepo := new(mock.ResultRepository)
		seasonRepo := new(mock.SeasonRepository)
		requester := new(mock.ResultRequester)
		clock := clockwork.NewFakeClock()
		logger, hook := test.NewNullLogger()

		processor := process.NewResultProcessor(resultRepo, seasonRepo, requester, clock, logger)

		done := make(chan bool)

		res := newResult(34)

		results := make([]app.Result, 1)
		results[0] = res

		ch := resultChannel(results)

		seasonRepo.On("CurrentSeasonIDs").Return([]uint64{1, 2}, nil)

		requester.On("ResultsBySeasonIDs", []uint64{1, 2}).Return(ch)
		resultRepo.On("ByFixtureID", uint64(34)).Return(&app.Result{}, errors.New("not found"))
		resultRepo.On("Insert", &res).Return(errors.New("error occurred"))
		processor.Process("results:current-season", "5", done)

		<-done

		requester.AssertExpectations(t)
		resultRepo.AssertExpectations(t)
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	})

	t.Run("logs error when unable to update result into repository when processing result current season command", func(t *testing.T) {
		t.Helper()

		resultRepo := new(mock.ResultRepository)
		seasonRepo := new(mock.SeasonRepository)
		requester := new(mock.ResultRequester)
		clock := clockwork.NewFakeClock()
		logger, hook := test.NewNullLogger()

		processor := process.NewResultProcessor(resultRepo, seasonRepo, requester, clock, logger)

		done := make(chan bool)

		res := newResult(34)

		results := make([]app.Result, 1)
		results[0] = res

		ch := resultChannel(results)

		seasonRepo.On("CurrentSeasonIDs").Return([]uint64{1, 2}, nil)

		requester.On("ResultsBySeasonIDs", []uint64{1, 2}).Return(ch)
		resultRepo.On("ByFixtureID", uint64(34)).Return(&app.Result{}, nil)
		resultRepo.On("Update", &res).Return(errors.New("error occurred"))
		processor.Process("results:current-season", "5", done)

		<-done

		requester.AssertExpectations(t)
		resultRepo.AssertExpectations(t)
		seasonRepo.AssertExpectations(t)
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	})
}

func newResult(f uint64) app.Result {
	return app.Result{FixtureID: f}
}

func resultChannel(results []app.Result) chan app.Result {
	ch := make(chan app.Result, len(results))

	for _, c := range results {
		ch <- c
	}

	close(ch)

	return ch
}
