package process

import (
	"github.com/statistico/statistico-data/internal/app"
	"log"
)

const country = "country"

// CountryProcessor is used to retrieve data from an external data source and to this applications
// chosen data store
type CountryProcessor struct {
	repository app.CountryRepository
	requester  app.CountryRequester
	logger     *log.Logger
}

func (p CountryProcessor) Process(command string, option string, done chan bool) {
	if command != country {
		p.logger.Fatalf("Command %s is not supported", command)
		return
	}

	countries := make(chan *app.Country, 50)

	go func() {
		err := p.requester.Countries(countries)

		if err != nil {
			p.logger.Fatalf("Error when processing countries in CountryProcessor '%s'", err.Error())
		}
	}()

	go p.persistCountries(countries, done)
}

// Loop through provided channel and persist Country struct(s) to database, once the channel
// is empty the channel passed as the second argument is notified
func (p CountryProcessor) persistCountries(ch <-chan *app.Country, done chan bool) {
	for country := range ch {
		p.persist(country)
	}

	done <- true
}

// Persist Country struct to the database, update if record exists, create new if not
func (p CountryProcessor) persist(c *app.Country) {
	country, err := p.repository.GetById(c.ID)

	if err != nil {
		if err := p.repository.Insert(country); err != nil {
			log.Printf("Error '%s' occurred when inserting Country struct: %+v\n,", err.Error(), country)
		}

		return
	}

	if err := p.repository.Update(country); err != nil {
		log.Printf("Error '%s' occurred when updating Competition struct: %+v\n,", err.Error(), country)
	}

	return
}

func NewCountryProcessor(r app.CountryRepository, s app.CountryRequester, l *log.Logger) *CountryProcessor {
	return &CountryProcessor{repository: r, requester: s, logger: l}
}