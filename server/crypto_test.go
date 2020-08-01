package server

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"reflect"
	"testing"
	"time"

	data "github.com/chutified/crypto-currencies/data"
	crypto "github.com/chutified/crypto-currencies/protos/crypto"
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

// TODO SubscribeCrypto

func TestHandleUpdatesCrypto(t *testing.T) {

	log := log.New(bytes.NewBufferString(""), "", log.LstdFlags)
	ds := data.New()
	err := ds.Update("https://coinmarketcap.com/all/views/all/")
	if err != nil {
		t.Fatalf("unexpected data service update error: %v", err)
	}
	c := New(log, ds)

	tests := []struct {
		name string
		url  string
		err  string
	}{
		{
			name: "ok",
			url:  "https://coinmarketcap.com/all/views/all",
		},
		{
			name: "invalid url",
			url:  "invalid",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t1 *testing.T) {

			cache := c.ds.Currencies

			go func() {
				c.handleUpdatesCrypto(400*time.Millisecond, test.url)
			}()

			if test.url == "invalid" {

				time.Sleep(1 * time.Second)
				equals := reflect.DeepEqual(c.ds.Currencies, cache)
				assert.Equal(t1, equals, true)

			} else {

				time.Sleep(2 * time.Minute)
				equals := reflect.DeepEqual(c.ds.Currencies, cache)
				assert.Equal(t1, equals, false)
			}
		})
	}

	// TODO add subscribtions
}
