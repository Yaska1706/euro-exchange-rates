build:
	go build -o bin/rakuten-interview main.go

run:
	go run main.go

test:
	go test -v ./test/...
