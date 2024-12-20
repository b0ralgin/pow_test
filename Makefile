test:
	go test ./... -v


build-client:
	go build -o ./bin/client ./cmd/client/
build-server:
	go build -o ./bin/server ./cmd/server/

build: build-client build-server




docker_build:
	docker build --target=client -t client -f Dockerfile .
	docker build --target=server -t server -f Dockerfile .

run_client:
	docker run client

run_server:
	docker run -p 8080:8080 server