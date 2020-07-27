package server

import (
	"context"
	"io"
	"log"
	"strings"

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
		// TODO
	}

	// success
	return resp, nil
}

// SubscribeCrypto handles the SubscribeCrypto gRPC calls.
func (c *Crypto) SubscribeCrypto(srv crypto.Crypto_SubscribeCryptoServer) error {

	// handle requests
	for {

		// receive request
		req, err := srv.Recv()
		if err == io.EOF {

			// cancel all subscriptions
			delete(c.subs, srv)

			return nil
		}
		if err != nil {

			// cancel all subscriptions
			delete(c.subs, srv)

			return err
			// TODO
		}
		name := strings.ToUpper(req.GetName())

		// validate request
		_, err = c.ds.GetCurrency(name)
		if err != nil {
			// TODO

			continue
		}

		// create server key if it does not exist
		if _, ok := c.subs[srv]; !ok {
			c.subs[srv] = []*crypto.GetCryptoRequest{}
		}

		// check if client has already subscribed
		var duplicit error
		for _, r := range c.subs[srv] {

			// compare names
			if r.Name == name {

				duplicit = errors.Errorf("client has already subscribed to %s", name)
				// TODO

				break
			}
		}
		// check duplicit
		if duplicit != nil {

			// cancel all subscriptions
			delete(c.subs, srv)
			// TODO

			continue
		}

		// append
		c.subs[srv] = append(c.subs[srv], req)
	}
}
