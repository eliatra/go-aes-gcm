all: build
build:
	goreleaser build --single-target --snapshot --clean