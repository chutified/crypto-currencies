package data

import models "github.com/chutified/crypto-currencies/models"

type symbolsConv map[string]string

func (cs symbolsConv) fromCurrencies(ccs map[string]*models.Currency) {

	// construct
	for name, curr := range ccs {
		cs[curr.Symbol] = name
	}
}
