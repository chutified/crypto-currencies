package server

import (
	"bytes"
	"log"
	"testing"

	"github.com/chutified/crypto-currencies/data"
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
	}{}
}
