package data

import models "github.com/chutified/crypto-currencies/models"

// symbolsConv defines a map of symbols and crypto currency names.
type symbolsConv map[string]string

// fromCurrencies updates the symbolsConv with the map of name and currency model.
func (cs symbolsConv) fromCurrencies(ccs map[string]*models.Currency) {

	// construct
	for name, curr := range ccs {
		cs[curr.Symbol] = name
	}
}
