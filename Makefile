EXE=./aoc2022
TESTFLAGS= -cover
TESTFLAGS+=-race

build:
	go build

test:
	go test $(TESTFLAGS) ./...

run:
	$(EXE) -day $(day) -puzzle $(puzzle)

lint:
	golangci-lint run

clean:
	-rm $(EXE)

all: test build run