package data

// symbolsConv defines a map of symbols and crypto currency names.
type symbolsConv map[string]string

// fromCurrencies updates the symbolsConv with the map of name and currency model.
func (s *Service) updateSymConv() {

	// construct
	for name, curr := range s.Currencies {
		s.symconv[curr.Symbol] = name
	}
}
