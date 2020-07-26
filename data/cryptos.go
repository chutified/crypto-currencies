package data

import (
	"reflect"
	"time"

	models "github.com/chutified/crypto-currencies/models"
	"github.com/pkg/errors"
)

type Service struct {
	Currencies map[string]*models.Currency
}

func New() *Service {
	return &Service{
		Currencies: make(map[string]*models.Currency),
	}
}

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
