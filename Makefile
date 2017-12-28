
RELEASE_VERSION="v1.0.0"

#all: test build
all: test lint

setup:
	go get -u github.com/golang/dep/cmd/dep
	go get -u gopkg.in/alecthomas/gometalinter.v2
	dep ensure

dep:
	dep ensure	

test:
	go test -race $(go list ./... | grep -v /vendor/) -v

coverage:
	go test -race $(go list ./... | grep -v /vendor/) -v -coverprofile .testCoverage.txt

build: build-dawrin build-linux

build-dawrin: 
	GOOS=darwin GOARCH=386 go build -v -o artifacts/aws-helper_darwin-arch386
	ls -laF artifacts/

build-linux:
	GOOS=linux GOARCH=amd64 go build -v -o artifacts/aws-helper_linux-amd64
	ls -laF artifacts/

lint:
	gometalinter -j 1 --vendor --deadline 70s --disable gotype --enable gofmt --enable goimports --enable misspell --enable unused --disable gas
	
