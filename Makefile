test-all:
	go test ./test/...

build:
	go build -o bin/ ./cmd/...

release-ip-lang:
	goreleaser release