.PHONY: build
build:
	go build -o build/service .

.PHONY: vendor
vendor:
	go mod vendor -v

# Full build inside a docker container for a clean release build
docker-build: vendor
	docker build -t devicechain-io/devicemanagement . -f docker/Dockerfile
