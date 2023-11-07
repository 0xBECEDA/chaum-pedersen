
proto:
	protoc api/v1/*.proto \
    		--go_out=. \
    		--go-grpc_out=. \
    		--go_opt=paths=source_relative \
    		--go-grpc_opt=paths=source_relative \
    		--proto_path=.


	protoc api/v2/*.proto \
    		--go_out=. \
    		--go-grpc_out=. \
    		--go_opt=paths=source_relative \
    		--go-grpc_opt=paths=source_relative \
    		--proto_path=.

build:
	docker build -f ./docker/Dockerfile.server -t server .
	docker build -f ./docker/Dockerfile.client -t client .

run: build
	docker compose -f docker/docker-compose.yml up

stop:
	docker compose -f docker/docker-compose.yml down
