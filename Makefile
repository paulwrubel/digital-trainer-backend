GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
GOVET=$(GOCMD) vet
GOLINT=golint
BINARY_NAME=digital_trainer_backend

GO_SOURCES=$(shell find . -type f -name "*.go") go.mod go.sum
all: build

build: $(GO_SOURCES)
	GOOS=linux \
	GOARCH=amd64 \
	$(GOBUILD) -o $(BINARY_NAME) .
vet:
	# go vet
	$(GOVET) cmd

lint:
	# golint
	$(GOLINT) ./...

perfect: lint vet