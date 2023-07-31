build:
	@go build -o bin/bookxd

run: build
	@./bin/bookxd

test:
	@go test -v ./...