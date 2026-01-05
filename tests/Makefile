VERSION := $(shell git describe --tags --always)
GOHOSTOS := $(shell go env GOHOSTOS)
GOHOSTARCH := $(shell go env GOHOSTARCH)

GOC_OS :=
GOC_ARCH :=

ifeq ($(GOHOSTOS),darwin)
	GOC_OS := darwin
	ifeq ($(GOHOSTARCH),arm64)
		GOC_ARCH := arm64
	else
		GOC_ARCH := amd64
	endif
endif

ifeq ($(GOHOSTOS),linux)
	GOC_OS := linux
	ifeq ($(GOHOSTARCH),amd64)
		GOC_ARCH := amd64
	else ifeq ($(GOHOSTARCH),386)
		GOC_ARCH := 386
	endif
endif

ifeq ($(GOC_OS),)
$(error Unsupported Operating System: $(GOHOSTOS))
endif

ifeq ($(GOC_ARCH),)
$(error Unsupported Architecture: $(GOHOSTARCH) on $(GOHOSTOS))
endif

all: build test

init:
	mkdir -p bin
	@echo "Detected $(GOHOSTOS)/$(GOHOSTARCH), downloading goc for $(GOC_OS)-$(GOC_ARCH)..."
	curl -s https://api.github.com/repos/qiniu/goc/releases/latest | grep "browser_download_url.*-$(GOC_OS)-$(GOC_ARCH).tar.gz" | cut -d : -f 2,3 | tr -d \" | xargs -n 1 curl -L | tar -zx -C bin
	chmod +x bin/goc

build:
	mkdir -p bin/
	cd .. && tests/bin/goc build --buildflags="-ldflags '-X main.Version=$(VERSION)'" -o tests/bin/librarian .
	cd ./client && go build -o ../bin/ .

test:
	@# Check if processes are running using lsof/netstat on ports if possible, or just trap exit in script
	# Using a wrapper script logic within Makefile to manage PIDs
	mkdir -p bin/data
	( \
		cd bin && ./goc server > /dev/null 2>&1 & GOC_PID=$$! ; \
		echo "Started goc server (PID: $$GOC_PID)" ; \
		sleep 2 ; \
		cd bin && CREATE_ADMIN_USER=admin CREATE_ADMIN_PASS=admin ./librarian serve -c ../../configs/example-for-testing.toml -d data > /dev/null 2>&1 & LIB_PID=$$! ; \
		echo "Started librarian server (PID: $$LIB_PID)" ; \
		sleep 20 ; \
		cd bin && ./client ; \
		EXIT_CODE=$$? ; \
		./goc profile -o ./coverage-goc.out ; \
		echo "Stopping librarian server (PID: $$LIB_PID)..." ; \
		kill $$LIB_PID || true ; \
		echo "Stopping goc server (PID: $$GOC_PID)..." ; \
		kill $$GOC_PID || true ; \
		exit $$EXIT_CODE \
	)
