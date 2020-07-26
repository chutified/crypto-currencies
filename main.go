package main

import (
	"log"
	"net"
	"os"

	"github.com/chutified/crypto-currencies/data"
	"github.com/chutified/crypto-currencies/protos/crypto"
	"github.com/chutified/crypto-currencies/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	l := log.New(os.Stdout, "[CRYPTOCURRENCY SERVICE] ", log.LstdFlags)

	ds := data.New()
	err := ds.Update()
	if err != nil {
		l.Fatal(err)
	}

	cs := server.New(l, ds)
	gs := grpc.NewServer()

	crypto.RegisterCryptoServer(gs, cs)
	reflection.Register(gs)

	listen, err := net.Listen("tcp", "localhost:10503")
	if err != nil {
		l.Fatal(err)
	}
	l.Printf("START")
	gs.Serve(listen)
}
