protoc:
	protoc --gofast_out=plugins=grpc:. proto/book/book.proto
	protoc --gofast_out=plugins=grpc:. proto/hello/hello.proto
