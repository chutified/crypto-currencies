# Crypto-currencies

The Crypto-currencies is a microservice, which is using <a href="https://grpc.io/" target="_blank">gRPC technology</a>.
The service provides the latest data about all cryptocurrencies on the market.
It supports both unary and bidirectional streaming calls, which allows data update every 15 seconds.
When an error occurs, it handles it in a non-fatal way with the error message.

The whole service is containerized using a Docker engine and everything can be easily run and deployed with the pre-prepared `make` commands in the Makefile.

The Cryptocurrencies obtains all necessary data from the <a href="https://coinmarketcap.com/all/views/all" target="_blank">CoinMarketCap</a> website. The algorithm does not infringe any copyrights nor the website robots exclusion protocol.


#### Directory tree
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
