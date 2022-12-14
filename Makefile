EXE=./aoc2022
TESTFLAGS= -cover
TESTFLAGS+=-race

build:
	@go build

test:
	@go test $(TESTFLAGS) ./...

test-short:
	@go test $(TESTFLAGS) -short ./...

run:
	@$(EXE) -day $(day) -puzzle $(puzzle)

lint:
	@golangci-lint run

clean:
	@-rm $(EXE)
	@go clean -testcache

all: clean test build run
