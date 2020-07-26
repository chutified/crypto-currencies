package data

import (
	"fmt"
	"reflect"
	"time"

	models "github.com/chutified/crypto-currencies/models"
	"github.com/pkg/errors"
)

// Service is a data service of the whole app which handles data operations,
// including fetching, handling data request or updating.
type Service struct {
	Currencies map[string]*models.Currency
}

// New is a constructor for the Service.
func New() *Service {
	return &Service{
		Currencies: make(map[string]*models.Currency),
	}
}

// GetCurrency finds the currency by its name.
func (s *Service) GetCurrency(c string) (*models.Currency, error) {

	// search
	cc, ok := s.Currencies[c]
	if !ok {
		return nil, fmt.Errorf("currency %s not found", c)
	}

	// success
	return cc, nil
}

// Update updates the database with the latest data.
func (s *Service) Update() error {

	// fetch
	rs, err := fetchRecords("https://coinmarketcap.com/all/views/all/")
	if err != nil {
		return errors.Wrap(err, "fetching records")
	}

	// parse records
	ccs, err := parseRecords(rs)
	if err != nil {
		return errors.Wrap(err, "parsing records")
	}

	// update
	s.Currencies = ccs
	return nil
}

// MonitorData handles the new updates.
func (s *Service) MonitorData(interval time.Duration) (chan struct{}, chan error) {

	// prepare channels
	upd := make(chan struct{})
	errs := make(chan error)

	go func() {
		ticker := time.Tick(interval)
		for range ticker {

			// clone
			cache := s.Currencies
			err := s.Update()
			if err != nil {
				errs <- err
				continue
			}

			// compare
			if reflect.DeepEqual(s.Currencies, cache) {
				// update, new values
				upd <- struct{}{}
			}
		}
	}()

	return upd, errs
}
