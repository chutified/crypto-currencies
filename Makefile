.PHONY: protogen, build, run

protogen:
	protoc -I protos/ --go_out=plugins=grpc:protos/crypto protos/crypto.proto

build:
	docker build -t cryptocurrencies .

run:
	docker run -it --network="host" --name cryptosrv --rm cryptocurrencies
