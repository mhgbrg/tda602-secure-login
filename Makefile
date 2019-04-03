.PHONY: all
all: build

.PHONY: build
build:
	go build server.go

.PHONY: serve
serve: build
	./server

.PHONY: watch
watch:
	ag -l -u | entr -r make serve
