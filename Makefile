.PHONY: run lint fmt test

run:
	go run cmd/main.go

lint:
	golangci-lint run ./..

lint-fix:
	golangci-lint run --fix ./...

fmt:
	gofmt -w .
	goimports -w .

test:
	go test ./.. -v