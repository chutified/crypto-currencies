package server

import (
	"strings"

	crypto "github.com/chutified/crypto-currencies/protos/crypto"
	"github.com/pkg/errors"
)

func (c *Crypto) handleGetCryptoRequest(req *crypto.GetCryptoRequest) (*crypto.GetCryptoResponse, error) {

	// get name
	name := req.GetName()
	name = strings.ToUpper(name)
	// get currency
	crc, err := c.ds.GetCurrency(name)
	if err != nil {
		return nil, errors.Wrap(err, "call data service GetCurrency")
	}

	// construct response
	resp := &crypto.GetCryptoResponse{
		Name:              crc.Name,
		Symbol:            crc.Symbol,
		MarketCapUSD:      crc.MarketCapUSD,
		Price:             crc.Price,
		CirculatingSupply: crc.CirculatingSupply,
		Mineable:          crc.Mineable,
		Volume:            crc.Volume,
		ChangeHour:        crc.ChangeHour,
		ChangeDay:         crc.ChangeDay,
		ChangeWeek:        crc.ChangeWeek,
	}
	return resp, nil
}
