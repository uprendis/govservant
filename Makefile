.PHONY: all
all: govservant

.PHONY: govservant
govservant:
	go build -o build/govservant ./cmd/govservant

.PHONY: test
test:
	go test ./...

.PHONY: clean
clean:
	rm -fr ./build/*
