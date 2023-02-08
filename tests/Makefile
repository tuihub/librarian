VERSION=$(shell git describe --tags --always)

all: build test

init:
	curl -s https://api.github.com/repos/qiniu/goc/releases/latest | grep "browser_download_url.*-linux-amd64.tar.gz" | cut -d : -f 2,3 | tr -d \" | xargs -n 1 curl -L | tar -zx && chmod +x goc && mv goc /usr/local/bin

build:
	mkdir -p bin/ && cd ../cmd/librarian && goc build --buildflags="-ldflags '-X main.Version=$(VERSION)'" -o ../../tests/bin/ .
	cd ./fakes3 && go build -o ../bin/ .
	cd ./client && go build -o ../bin/ .

test:
	cd bin && goc server &
	cd bin && ./fakes3 &
	sleep 1s
	cd bin && UNLOCK_CREATE=true ./librarian &
	cd bin && ./client
	goc profile -o ../coverage-goc.out
