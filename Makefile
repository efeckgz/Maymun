BINARY_NAME=maymun

all: test build

build:
	go build -o bin/${BINARY_NAME} main.go

test:
	# go test -v ./lexer
	go test -v ./parser

run:
	./bin/${BINARY_NAME}

clean:
	go clean
	rm bin/${BINARY_NAME}
