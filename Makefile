
default:
	go build


update:
	go get -u
	go mod tidy

gosec:
	gosec ./...

release:
	echo TODO
