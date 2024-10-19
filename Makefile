gen:
	cd protobuf && \
	protoc \
	    --go_out=protocol_demo \
		--go_opt=paths=source_relative \
		--go-grpc_out=protocol_demo \
		--go-grpc_opt=paths=source_relative \
		protocol.proto
