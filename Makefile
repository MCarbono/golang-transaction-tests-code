default: run

run:
	go run cmd/main.go

tests:
	go test ./... -v