BINDIR      := $(CURDIR)/bin

all: build

deps:
	go mod download
	go mod tidy

format:
	gofmt -w -e .

build: deps format
	CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-w -s" -o '$(BINDIR)'/go-aes-gcm ./