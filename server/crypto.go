package server

import (
	"context"
	"io"
	"log"
	"strings"
	"time"

	data "github.com/chutified/crypto-currencies/data"
	crypto "github.com/chutified/crypto-currencies/protos/crypto"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	go func() {
		c.handleUpdatesCrypto(15*time.Second, "https://coinmarketcap.com/all/views/all/")
	}()

	return c
}

// GetCrypto handles the GetCrypto gRPC calls.
func (c *Crypto) GetCrypto(ctx context.Context, req *crypto.GetCryptoRequest) (*crypto.GetCryptoResponse, error) {

	// handle request
	resp, err := c.handleGetCryptoRequest(req)
	if err != nil {
		c.log.Printf("[error] handle GetCryptoRequest: %v", err)

		// define error status
		gErr := status.Newf(
			codes.NotFound,
			"cryptocurrency '%s' not found", req.GetName(),
		)

		return nil, gErr.Err()
	}

	// success
	c.log.Printf("[success] handled request of '%s' currency", resp.GetName())
	return resp, nil
}

// SubscribeCrypto handles the SubscribeCrypto gRPC calls.
func (c *Crypto) SubscribeCrypto(srv crypto.Crypto_SubscribeCryptoServer) error {

	id := uuid.New().String()
	c.log.Printf("[success] new client (%s)", id)
	// handle requests
	for {

		// receive request
		req, err := srv.Recv()
		if err == io.EOF {
			c.log.Printf("[cancel] client canceled connection (%s)", id)

			// cancel all subscriptions
			delete(c.subs, srv)
			c.log.Printf("[server] delete client's subscriptions (%s)", id)

			return nil
		}
		if err != nil {
			c.log.Printf("[error] receive error (%s)", id)

			// cancel all subscriptions
			delete(c.subs, srv)
			c.log.Printf("[server] delete client's subscriptions (%s)", id)

			return errors.Wrap(err, "receiving client's request")
		}
		req.Name = strings.ToUpper(req.GetName())

		// validate request
		m, err := c.ds.GetCurrency(req.Name)
		if err != nil {
			c.log.Printf("[invalid] invalid request, currency: %s (%s)", req.Name, id)

			// define error status
			gErr := status.Newf(
				codes.NotFound,
				"cryptocurrency '%s' not found", req.GetName(),
			)

			// send error message
			err = srv.Send(&crypto.SubscribeCryptoResponse{
				Message: &crypto.SubscribeCryptoResponse_Error{
					Error: gErr.Proto(),
				},
			})

			continue
		}
		req.Name = m.Name

		// create server key if it does not exist
		if _, ok := c.subs[srv]; !ok {
			c.subs[srv] = []*crypto.GetCryptoRequest{}
		}

		// check if client has already subscribed
		var duplicit *status.Status
		for _, r := range c.subs[srv] {

			// compare names
			if r.Name == req.Name {

				// define error status
				duplicit = status.Newf(
					codes.AlreadyExists,
					"client has already subscribed for the currency '%s'", req.Name,
				)

				break
			}
		}
		// check duplicit
		if duplicit != nil {
			c.log.Printf("[invalid] invalid request, currency: '%s' already subscribed (%s)", req.Name, id)

			// send error message
			srv.Send(&crypto.SubscribeCryptoResponse{
				Message: &crypto.SubscribeCryptoResponse_Error{
					Error: duplicit.Proto(),
				},
			})

			continue
		}

		// append
		c.log.Printf("[success] currency: '%s' subscribed (%s)", req.Name, id)
		c.subs[srv] = append(c.subs[srv], req)
	}
}

// handleUpdatesCrypto will inform if the data service receives new data values.
func (c *Crypto) handleUpdatesCrypto(interval time.Duration, url string) {

	// prepare channels
	updates, errs := c.ds.MonitorData(interval, url)

	// handle erorrs
	go func() {

		// receive errors
		for err := range errs {
			c.log.Printf("[error] monitoring service data: %v", err)
		}
	}()

	// handle updates
	for range updates {
		c.log.Printf("[update] cryptocurrencies data updated")

		// range over clients
		for client, subs := range c.subs {

			// range over subsciptions
			for _, req := range subs {

				// handle subscription
				resp, err := c.handleGetCryptoRequest(req)
				if err != nil {
					c.log.Printf("[error] handle GetCryptoRequest: %v", err)

					// define error status
					gErr := status.Newf(
						codes.NotFound,
						"cryptocurrency '%s' no longer exists in service database", req.GetName(),
					)

					// send error message
					err = client.Send(&crypto.SubscribeCryptoResponse{
						Message: &crypto.SubscribeCryptoResponse_Error{
							Error: gErr.Proto(),
						},
					})

					continue
				}

				err = client.Send(&crypto.SubscribeCryptoResponse{
					Message: &crypto.SubscribeCryptoResponse_GetCryptoResponse{
						GetCryptoResponse: resp,
					},
				})
				if err != nil {
					c.log.Printf("[error] send response: %v", err)
					continue
				}
			}
		}
	}
}
