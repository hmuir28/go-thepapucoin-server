build:
	@go build -o ./bin/go-thepapucoin-server

run: build
	@./bin/go-thepapucoin-server

test:
	go test -v ./...
