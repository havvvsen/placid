all: run

build: cmd/app/main.go
	go build -o ./bin/app ./cmd/app/main.go

run: build
	./bin/app
