build:
	go build -o gmoji cmd/gmoji/main.go

test:
	go test -v ./...
