package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/chutified/crypto-currencies/config"
	"github.com/chutified/crypto-currencies/data"
	"github.com/chutified/crypto-currencies/protos/crypto"
	"github.com/chutified/crypto-currencies/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	l := log.New(os.Stdout, "[CRYPTOCURRENCY SERVICE] ", log.LstdFlags)

	// configuration
	cfg, err := config.GetConfig("config.yaml")
	if err != nil {
		l.Fatal(err)
	}
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	// data service
	ds := data.New()
	err = ds.Update()
	if err != nil {
		l.Fatal(err)
	}

	// servers
	cs := server.New(l, ds)
	gs := grpc.NewServer()

	// registrations
	crypto.RegisterCryptoServer(gs, cs)
	reflection.Register(gs)

	// listen
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		l.Fatal(err)
	}

	// initialize the server
	l.Printf("[start] launch server on %s\n", addr)
	gs.Serve(listen)
}
