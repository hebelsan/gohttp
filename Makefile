GO=go

.PHONY: help build test lint clean

help:
	@echo "Usage:"
	@echo ""
	@echo "	make <commands>"
	@echo ""
	@echo "The commands are:"
	@echo ""
	@echo "	build               Build the package"
	@echo "	clean               Run go clean"
	@echo "	help                Print this help text"
	@echo "	lint                Run golangci-lint"
	@echo "	test                Run go test"

build:
	$(GO) build -v ./...

install:
	$(GO) install

test:
	$(GO) test -v ./...

lint:
	golangci-lint run

clean:
	$(GO) clean