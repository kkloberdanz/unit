.PHONY: test
test:
	go test ./...

.PHONY: gc-analysis
gc-analysis:
	go build -gcflags -m

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: lint
lint:
	golangci-lint run
