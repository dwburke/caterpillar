
GORELEASER_VERSION=v0.159.0

default:
	$(MAKE) build

build:
	CGO_ENABLED=0 go build -ldflags="-extldflags=-static"


update:
	go get -u
	go mod tidy

gosec:
	gosec ./...

setup:
	rm -rf goreleaser
	mkdir goreleaser
	cd goreleaser && wget https://github.com/goreleaser/goreleaser/releases/download/${GORELEASER_VERSION}/goreleaser_Linux_x86_64.tar.gz
	cd goreleaser && tar xvzf goreleaser_Linux_x86_64.tar.gz

release:
	git push --tags
	goreleaser/goreleaser --rm-dist
