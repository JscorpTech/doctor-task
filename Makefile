
air:
	@air --build.cmd "go build -o bin/api ./cmd/main.go" --build.bin "./bin/api runserver"

run:
	@go run ./cmd/main.go runserver

build:
	@go build -o ./bin/app ./cmd/main.go

migrate:
	@go run ./cmd/main.go migrate