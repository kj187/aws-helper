#all: test build
all: test lint

setup:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

setup-coverage:
	go get golang.org/x/tools/cmd/cover
	go get github.com/axw/gocov/gocov
	go get github.com/mattn/goveralls

dep:
	dep ensure	

test:
	go test -race $(go list ./... | grep -v /vendor/) -v

build: build-dawrin build-linux

build-dawrin: 
	GOOS=darwin GOARCH=386 go build -v -o artifacts/aws-helper_darwin-arch386
	ls -laF artifacts/

build-linux:
	GOOS=linux GOARCH=amd64 go build -v -o artifacts/aws-helper_linux-amd64
	ls -laF artifacts/
