VERSION=$(shell git describe --tags --always)

all: build test

init:
	curl -L https://github.com/qiniu/goc/releases/download/v1.4.4/goc-v1.4.4-linux-amd64.tar.gz | tar -zx && chmod +x goc && mv goc /usr/local/bin

build:
	mkdir -p bin/ && cd ../cmd/librarian && goc build --buildflags="-ldflags '-X main.Version=$(VERSION)'" -o ../../tests/bin/ .
	#cd ./fakes3 && go build -o ../bin/ .
	cd ./client && go build -o ../bin/ .

test:
	cd bin && goc server &
	#cd bin && ./fakes3 &
	sleep 1s
	cd bin && CREATE_ADMIN=true ./librarian -conf ../../configs/example-for-testing.yaml &
	cd bin && ./client
	goc profile -o ../coverage-goc.out
