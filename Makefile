all: build
build:
	goreleaser build --single-target --snapshot --clean
test: build
	cd gcm; go test
	./test.sh
