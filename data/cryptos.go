package data

import (
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
