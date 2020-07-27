package server

import (
	"context"
	"log"

	"github.com/chutified/crypto-currencies/data"
	"github.com/chutified/crypto-currencies/protos/crypto"
	"github.com/pkg/errors"
)

// Crypto is a server for handling crypto calls.
type Crypto struct {
	log  *log.Logger
	ds   *data.Service
	subs map[crypto.Crypto_SubscribeCryptoServer][]*crypto.GetCryptoRequest
}

// New defines a constructor for the Crypto server.
func New(log *log.Logger, ds *data.Service) *Crypto {
	c := &Crypto{
		log:  log,
		ds:   ds,
		subs: make(map[crypto.Crypto_SubscribeCryptoServer][]*crypto.GetCryptoRequest),
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
