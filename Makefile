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
	ag -l | entr -r -s "HOST=localhost:8080 CERT_FILE=keys/fullchain.pem KEY_FILE=keys/privkey.pem make serve"
