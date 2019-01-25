package season

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/joesweeny/sportmonks-go-client"
	"github.com/joesweeny/statshub/internal/model"
	"github.com/jonboulle/clockwork"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestProcess(t *testing.T) {
	repo := new(mockRepository)

	server := newTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, req.URL.String(), "http://example.com/api/v2.0/seasons?api_token=my-key&page=1")
		b, _ := json.Marshal(seasonResponse())
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBuffer(b)),
		}
	})

	client := sportmonks.Client{
		Client:  server,
		BaseURL: "http://example.com",
		ApiKey:  "my-key",
	}

	service := Service{
		Repository: repo,
		Factory:    Factory{clockwork.NewFakeClock()},
		Client:     &client,
		Logger:     log.New(ioutil.Discard, "", 0),
	}

	t.Run("inserts new season", func(t *testing.T) {
		repo.On("GetById", 100).Return(&model.Season{}, errors.New("not Found"))
		repo.On("Insert", mock.Anything).Return(nil)
		repo.AssertNotCalled(t, "Update", mock.Anything)
		service.Process()
	})

	t.Run("updates existing competition", func(t *testing.T) {
		c := newSeason(1)
		repo.On("GetById", 100).Return(c, nil)
		repo.On("Update", &c).Return(nil)
		repo.MethodCalled("Update", &c)
		repo.AssertNotCalled(t, "Insert", mock.Anything)
		service.Process()
	})
}

type mockRepository struct {
	mock.Mock
}

func (m mockRepository) Insert(c *model.Season) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m mockRepository) Update(c *model.Season) error {
	args := m.Called(&c)
	return args.Error(0)
}

func (m mockRepository) GetById(id int) (*model.Season, error) {
	args := m.Called(id)
	c := args.Get(0).(*model.Season)
	return c, args.Error(1)
}

func (m mockRepository) GetIds() ([]int, error) {
	args := m.Called()
	return args.Get(0).([]int), args.Error(1)
}

func (m mockRepository) GetCurrentSeasons() ([]model.Season, error) {
	args := m.Called()
	return args.Get(0).([]model.Season), args.Error(1)
}

type roundTripFunc func(req *http.Request) *http.Response

func (r roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return r(req), nil
}

func newTestClient(fn roundTripFunc) *http.Client {
	return &http.Client{
		Transport: roundTripFunc(fn),
	}
}

func seasonResponse() sportmonks.SeasonsResponse {
	s := newClientSeason()

	m := sportmonks.Meta{}
	m.Pagination.Total = 1
	m.Pagination.Count = 1
	m.Pagination.PerPage = 1
	m.Pagination.CurrentPage = 1
	m.Pagination.TotalPages = 1

	res := sportmonks.SeasonsResponse{}
	res.Data = append(res.Data, *s)
	res.Meta = m

	return res
}
