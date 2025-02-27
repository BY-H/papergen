.PHONY: build
build:
	@go build -o _output/papergen cmd/papergen/main.go

.PHONY: format
format:
	@gofmt -s -w ./

.PHONY: clean
clean:
	@rm -vrf _output
