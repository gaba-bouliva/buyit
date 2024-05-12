build:
	@go build -o bin/buyit	cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/buyit