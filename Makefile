.PHONY: build check

build:
	go build -o bin/wallet ./cmd/cli

check: build
	@test -z "$$(gofmt -l .)" || (echo "non formatés :"; gofmt -l .; exit 1)
	go vet ./...
	./bin/wallet currencies
