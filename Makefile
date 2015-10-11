build: go-get install test

go-get:
	go get golang.org/x/tools/cmd/vet
	go get github.com/golang/lint/golint
	go get golang.org/x/tools/cmd/goimports
	go get github.com/opennota/check/cmd/defercheck
	go get github.com/opennota/check/cmd/structcheck
	go get github.com/opennota/check/cmd/varcheck
	go get gopkg.in/alecthomas/kingpin.v2
	go get github.com/stretchr/testify
	go get github.com/mtojek/localserver

install:
	go get -t -v ./...

test:
	go test -v ./...
	go test -race  -i ./...
	go list ./... | sed -e 's;github.com/mtojek/go-telnet;.;' | xargs defercheck
	go list ./... | sed -e 's;github.com/mtojek/go-telnet;.;' | xargs varcheck
	go list ./... | sed -e 's;github.com/mtojek/go-telnet;.;' | xargs structcheck
	golint ./...
	go tool vet -v=true .
	test -z "`gofmt -d .`"
	test -z "`goimports -l .`"

cc: #cleancode
	gofmt -s -w .
	goimports -w .

dev: install
	cat resources/input-data/wp.pl_1.bin | go-telnet www.wp.pl 80
