BINARY:=xnotitle
# Check if monova installed
MONOVA:=$(shell monova -version dot 2> /dev/null)
PWD:=$(shell pwd)
VERSION=0.0.0

version:
ifdef MONOVA
override VERSION="$(shell monova)"
else
	$(info "Install monova to automatically generate package version")
endif

test:
	go test

build: test version
	go build -ldflags="-X main.version=${VERSION}"

run: build
	./${BINARY}
