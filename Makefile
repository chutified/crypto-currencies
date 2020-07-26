.PHONY: protogen

protogen:
	protoc -I protos/ --go_out=plugins=grpc:protos/crypto protos/crypto.proto
