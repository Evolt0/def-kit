protoc:
	protoc --gofast_out=plugins=grpc:. proto/${package}/${proto}.proto