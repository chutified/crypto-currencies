package server

import (
	"context"
	"log"
	"strings"

	"github.com/chutified/crypto-currencies/data"
	"github.com/chutified/crypto-currencies/protos/crypto"
	"github.com/pkg/errors"
)

// Crypto is a server for handling crypto calls.
type Crypto struct {
	log *log.Logger
	ds  *data.Service
}

// New defines a constructor for the Crypto server.
func New(log *log.Logger, ds *data.Service) *Crypto {
	c := &Crypto{
		log: log,
		ds:  ds,
	}

	return c
}

// GetCrypto handles the GetCrypto gRPC calls.
func (c *Crypto) GetCrypto(ctx context.Context, req *crypto.GetCryptoRequest) (*crypto.GetCryptoResponse, error) {

	// handle request
	resp, err := c.handleGetCryptoRequest(req)
	if err != nil {
		return nil, errors.Wrap(err, "handling GetCryptoRequest")
	}

	// success
	return resp, nil
}

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
