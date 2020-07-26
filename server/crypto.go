package server

import (
	"log"

	"github.com/chutified/crypto-currencies/data"
)

type Crypto struct {
	log *log.Logger
	ds  *data.Service
}

func New(log *log.Logger, ds *data.Service) *Crypto {
	c := &Crypto{
		log: log,
		ds:  ds,
	}

	return c
}
