BINARY_NAME=monkey

all: build test run

build:
	go build -o bin/${BINARY_NAME} main.go

test:
	go test -v ./lexer

run:
	./bin/${BINARY_NAME}

clean:
	go clean
	rm bin/${BINARY_NAME}
