package server

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/chutified/crypto-currencies/data"
	"github.com/chutified/crypto-currencies/protos/crypto"
	"gopkg.in/go-playground/assert.v1"
)

func TestCrypto(t *testing.T) {

	log := log.New(bytes.NewBufferString(""), "", log.LstdFlags)
	ds := data.New()
	err := ds.Update("https://coinmarketcap.com/all/views/all/")
	if err != nil {
		t.Fatalf("unexpected data service update error: %v", err)
	}

	// >>> New()
	c := New(log, ds)

	// >>> GetCrypto()
	tests := []struct {
		name     string
		currName string
		err      string
	}{
		{
			name:     "ok",
			currName: "BITCOIN",
			err:      "",
		},
		{
			name:     "not found",
			currName: "invalid",
			err:      "cryptocurrency '.*' not found",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t1 *testing.T) {

			resp, err := c.GetCrypto(context.Background(), &crypto.GetCryptoRequest{
				Name: test.currName,
			})
			if err != nil {

				exp := fmt.Sprintf("%s.*", test.err)
				assert.MatchRegex(t1, err.Error(), exp)

			} else {

				assert.Equal(t1, "", test.err)
				assert.Equal(t1, resp.Name, test.currName)
			}
		})
	}
}
