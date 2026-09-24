.PHONY: build check

build:
	go build -o bin/wallet ./cmd/cli

check:
	gofmt -l . && go vet ./... && $(MAKE) build && ./bin/wallet currencies
