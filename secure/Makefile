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
	ag -l | entr -r -s "HOSTNAME=localhost HTTP_PORT=8080 HTTPS_PORT=8081 CERT_FILE=keys/fullchain.pem KEY_FILE=keys/privkey.pem make serve"
