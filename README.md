# Crypto-currencies

The Crypto-currencies is a microservice, which is using <a href="https://grpc.io/" target="_blank">gRPC technology</a>.
The service provides the latest data about all cryptocurrencies on the market.
It supports both unary and bidirectional streaming calls, which allows data update every 15 seconds.
When an error occurs, it handles it in a non-fatal way with the error message.

The whole service is containerized using a Docker engine and everything can be easily run and deployed with the pre-prepared `make` commands in the Makefile.

The Cryptocurrencies obtains all necessary data from the <a href="https://coinmarketcap.com/all/views/all" target="_blank">CoinMarketCap</a> website. The algorithm does not infringe any copyrights nor the website robots exclusion protocol.

## Installation

### Requirements
- <a href="https://git-scm.com/downloads" target="_blank">Git</a>
- <a href="https://docs.docker.com/get-docker/" target="_blank">Docker Engine</a>

### Linux/Mac
This is the exact way to download and run the service. On a Windows machine, the installation process would be slightly different.
```bash
$ git clone https://github.com/chutified/crypto-currencies.git     # download repository
$ cd crypto-currencies        # move to repository dir
$ make build                  # build docker image
$ make run                    # initialize service
```

## Supported crypto currencies
The service supports a large number of cryptocurrencies, so they are listed <a href="https://github.com/chutified/crypto-currencies/blob/master/docs/currencies.md" taget="_blan">here</a>.

**Note:**
*The Crypto request holds the key "Name" and its value is **not** case sensitive.*
*So the Crypto names must not be completely lowercase nor uppercase to be found.*

## Client
All clients can be built with the help of the <a href="https://grpc.io/docs/protoc-installation/" target="_blank">Protocol Buffer Compiler</a> with the <a href="https://grpc.io/" target="_blank">gRPC</a> plugin.

*The protobuffer of the services:* <a href="https://github.com/chutified/crypto-currencies/blob/master/protos/crypto.proto">commodity.proto</a> TODO CHECK URL

## Directory tree
```bash
_
├── config
│   ├── tests
│   │   ├── config_0.yaml
│   │   └── config_1.yaml
│   ├── config.go
│   └── config_test.go
├── data
│   ├── cryptos.go
│   ├── cryptos_test.go
│   ├── fetch.go
│   ├── fetch_test.go
│   └── symbols.go
├── models
│   └── currency.go
├── protos
│   ├── crypto
│   │   └── crypto.pb.go
│   ├── google
│   │   └── rpc
│   │       └── status.proto
│   └── crypto.proto
├── server
│   ├── crypto.go
│   ├── crypto_test.go
│   └── handlers.go
├── config.yaml
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
├── Makefile
└── README.md
```

source: https://coinmarketcap.com/all/views/all/
