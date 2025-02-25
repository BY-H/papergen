.PHONY: build
build:
	go build -o _output/papergen cmd/papergen/main.go

.PHONY: clean
clean:
	rm -rf _output