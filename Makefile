VERSION ?= 0.0.1
BUILDDIR ?= $(CURDIR)/build
FUNCTIONAL_AREA ?= device-management

.PHONY: build
build:
	go build -o $(BUILDDIR)/service .

.PHONY: build-stripped
build-stripped:
	go build -ldflags '-s -w' -o $(BUILDDIR)/service .

.PHONY: vendor
vendor:
	go mod vendor

clean:
	rm -rf $(BUILDDIR)/*

# Build a docker image based on the build target
.PHONY: docker-build
docker-build: vendor build-stripped
	docker build -t devicechain-io/${FUNCTIONAL_AREA}:${VERSION} . -f docker/Dockerfile

# Run the docker image
.PHONY: docker-run
docker-run: docker-build
	docker run -it --env DC_INSTANCE_ID=dc1 --env DC_TENANT_MICROSERVICE_ID=tms-tenant1-devicemanagement devicechain-io/${FUNCTIONAL_AREA}:${VERSION}
