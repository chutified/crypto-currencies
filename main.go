package main

import (
	"fmt"
	"log"
	"net"
	"os"

	config "github.com/chutommy/crypto-currencies/config"
	data "github.com/chutommy/crypto-currencies/data"
	crypto "github.com/chutommy/crypto-currencies/protos/crypto"
	server "github.com/chutommy/crypto-currencies/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	l := log.New(os.Stdout, "[CRYPTOCURRENCY SERVICE] ", log.LstdFlags)

	// configuration
	cfg, err := config.GetConfig("config.yaml")
	if err != nil {
		l.Fatalf("[error] get config %v", err)
	}
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	// data service
	ds := data.New()
	err = ds.Update("https://coinmarketcap.com/all/views/all/")
	if err != nil {
		l.Fatalf("[error] update %v", err)
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
		l.Fatalf("[error] tcp listen %v", err)
	}

	// initialize the server
	l.Printf("[start] launch server on %s\n", addr)
	gs.Serve(listen)
}
