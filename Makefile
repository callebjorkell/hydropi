GIT_HASH:=$(shell git rev-parse --short HEAD)
DIRTY:=$(shell test -z "`git status --porcelain`" || echo "-dirty")
VERSION:=$(GIT_HASH)$(DIRTY)
TIME:=$(shell date -u -Iseconds)

BIN:=hydropi
PACKAGE:=./cmd/hydropi

.PHONY: dev cross-pi pi deps update-deps vet test-server fmt test

dev: test vet deps fmt
	go build -ldflags "-X main.buildVersion=$(VERSION) -X main.buildTime=$(TIME)" -o $(BIN) $(PACKAGE)

test: deps vet
	go test ./...

fmt:
	@test -z $(shell gofmt -l .)

vet:
	go vet ./...

update-deps:
	go get -u ./...
	go mod tidy

deps:
	go mod download

pi: deps
	# GOOS=linux GOARCH=arm GOARM=6
	go build -o $(BIN) -tags=pi -ldflags "-X main.buildVersion=$(VERSION) -X main.buildTime=$(TIME)" $(PACKAGE)

run:
	go run $(PACKAGE) start
