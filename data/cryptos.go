package data

import models "github.com/chutified/crypto-currencies/models"

type Service struct {
	Currencies map[string]*models.Currency
}

func New() *Service {
	return &Service{
		Currencies: make(map[string]*models.Currency),
	}
}
