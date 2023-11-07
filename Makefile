
proto:
	protoc internal/api/v1/*.proto \
    		--go_out=. \
    		--go-grpc_out=. \
    		--go_opt=paths=source_relative \
    		--go-grpc_opt=paths=source_relative \
    		--proto_path=.

	protoc internal/api/v2/*.proto \
    		--go_out=. \
    		--go-grpc_out=. \
    		--go_opt=paths=source_relative \
    		--go-grpc_opt=paths=source_relative \
    		--proto_path=.
