export CGO_ENABLED=0
export GO111MODULE=on

.PHONY: build

AETHER_ROC_API_VERSION := latest
ONOS_BUILD_VERSION := v0.6.3

build: # @HELP build the Go binaries and run all validations (default)
build:
	CGO_ENABLED=1 go build -o build/_output/aether-roc-api ./cmd/aether-roc-api

test: # @HELP run the unit tests and source code validation
test: build deps linters license_check
	CGO_ENABLED=1 go test -race github.com/onosproject/aether-roc-api/pkg/...
	CGO_ENABLED=1 go test -race github.com/onosproject/aether-roc-api/cmd/...
#	CGO_ENABLED=1 go test -race github.com/onosproject/aether-roc-api/api/...

coverage: # @HELP generate unit test coverage data
coverage: build deps
	./../build-tools/build/coveralls/coveralls-coverage aether-roc-api 7KVoIxwT8yII0oTiSxQK7lDsnmzN4gabB

deps: # @HELP ensure that the required dependencies are in place
	go build -v ./...
	bash -c "diff -u <(echo -n) <(git diff go.mod)"
	bash -c "diff -u <(echo -n) <(git diff go.sum)"

linters: # @HELP examines Go source code and reports coding problems
	golangci-lint run --timeout 30m

license_check: # @HELP examine and ensure license headers exist
	@if [ ! -d "../build-tools" ]; then cd .. && git clone https://github.com/onosproject/build-tools.git; fi
	./../build-tools/licensing/boilerplate.py -v --rootdir=${CURDIR} --boilerplate LicenseRef-ONF-Member-1.0

gofmt: # @HELP run the Go format validation
	bash -c "diff -u <(echo -n) <(gofmt -d pkg/ cmd/ tests/)"

oapi-codegen-aether-1.0.0: # @HELP generate openapi types from aether-1.0.0-openapi3.yaml
	oapi-codegen -generate types -package types -o pkg/aether_1_0_0/types/aether-1.0.0-types.go api/aether-1.0.0-openapi3.yaml
	oapi-codegen -generate types -package server -templates pkg/codegen/templates -o pkg/aether_1_0_0/server/aether-1.0.0-gnmi.go api/aether-1.0.0-openapi3.yaml
	oapi-codegen -generate server -package server -templates pkg/codegen/modified -o pkg/aether_1_0_0/server/aether-1.0.0-server.go api/aether-1.0.0-openapi3.yaml
	sed -i "s/Target/types.Target/g" pkg/aether_1_0_0/server/aether-1.0.0-server.go
	oapi-codegen -generate server -package server -templates pkg/codegen/templates -o pkg/aether_1_0_0/server/aether-1.0.0-impl.go api/aether-1.0.0-openapi3.yaml
	sed -i "s/Target/types.Target/g" pkg/aether_1_0_0/server/aether-1.0.0-impl.go
	oapi-codegen -generate spec -package server -o pkg/aether_1_0_0/server/aether-1.0.0-spec.go api/aether-1.0.0-openapi3.yaml
	for f in pkg/aether_1_0_0/**/aether-1.0.0*.go; do \
  		sed -i '1i// Code generated by oapi-codegen. DO NOT EDIT.' $$f; \
		sed -i 's/rbac/aether/g' $$f;done

oapi-codegen-rbac: # @HELP generate openapi types from rbac-1.0.0-openapi3.yaml
	oapi-codegen -generate types -package types -o pkg/rbac_1_0_0/types/rbac-1.0.0-types.go api/rbac-1.0.0-openapi3.yaml
	oapi-codegen -generate types -package server -templates pkg/codegen/templates -o pkg/rbac_1_0_0/server/rbac-1.0.0-gnmi.go api/rbac-1.0.0-openapi3.yaml
	oapi-codegen -generate server -package server -templates pkg/codegen/modified -o pkg/rbac_1_0_0/server/rbac-1.0.0-server.go api/rbac-1.0.0-openapi3.yaml
	oapi-codegen -generate types -package server -templates pkg/codegen/convert-oapi-gnmi -o pkg/rbac_1_0_0/server/convert-oapi-gnmi.go api/rbac-1.0.0-openapi3.yaml
	oapi-codegen -generate types -package server -templates pkg/codegen/convert-gnmi-oapi -o pkg/rbac_1_0_0/server/convert-gnmi-oapi.go api/rbac-1.0.0-openapi3.yaml
	sed -i "s/Target/types.Target/g" pkg/rbac_1_0_0/server/rbac-1.0.0-server.go
	oapi-codegen -generate server -package server -templates pkg/codegen/templates -o pkg/rbac_1_0_0/server/rbac-1.0.0-impl.go api/rbac-1.0.0-openapi3.yaml
	sed -i "s/Target/types.Target/g" pkg/rbac_1_0_0/server/rbac-1.0.0-impl.go
	oapi-codegen -generate spec -package server -o pkg/rbac_1_0_0/server/rbac-1.0.0-spec.go api/rbac-1.0.0-openapi3.yaml
	sed -i "s/model_0_0_0/rbac_1_0_0/g" pkg/rbac_1_0_0/*/*.go
	sed -i "s/model-0.0.0/rbac-1.0.0/g" pkg/rbac_1_0_0/*/*.go
	for f in pkg/rbac_1_0_0/**/rbac-1.0.0*.go; do sed -i '1i// Code generated by oapi-codegen. DO NOT EDIT.' $$f; done

oapi-codegen-aether-2.0.0: # @HELP generate openapi types from aether-2.0.0-openapi3.yaml
	oapi-codegen -generate types -package types -o pkg/aether_2_0_0/types/aether-2.0.0-types.go api/aether-2.0.0-openapi3.yaml
	oapi-codegen -generate types -package server -templates pkg/codegen/templates -o pkg/aether_2_0_0/server/aether-2.0.0-gnmi.go api/aether-2.0.0-openapi3.yaml
	oapi-codegen -generate server -package server -templates pkg/codegen/modified -o pkg/aether_2_0_0/server/aether-2.0.0-server.go api/aether-2.0.0-openapi3.yaml
	oapi-codegen -generate types -package server -templates pkg/codegen/convert-oapi-gnmi -o pkg/aether_2_0_0/server/aether-2.0.0-convert-oapi-gnmi.go api/aether-2.0.0-openapi3.yaml
	oapi-codegen -generate types -package server -templates pkg/codegen/convert-gnmi-oapi -o pkg/aether_2_0_0/server/aether-2.0.0-convert-gnmi-oapi.go api/aether-2.0.0-openapi3.yaml
	sed -i "s/Target/types.Target/g" pkg/aether_2_0_0/server/aether-2.0.0-server.go
	oapi-codegen -generate server -package server -templates pkg/codegen/templates -o pkg/aether_2_0_0/server/aether-2.0.0-impl.go api/aether-2.0.0-openapi3.yaml
	sed -i "s/Target/types.Target/g" pkg/aether_2_0_0/server/aether-2.0.0-impl.go
	oapi-codegen -generate spec -package server -o pkg/aether_2_0_0/server/aether-2.0.0-spec.go api/aether-2.0.0-openapi3.yaml
	sed -i "s/model_0_0_0/aether_2_0_0/g" pkg/aether_2_0_0/*/*.go
	sed -i "s/model-0.0.0/aether-2.0.0/g" pkg/aether_2_0_0/*/*.go
	for f in pkg/aether_2_0_0/**/aether-2.0.0*.go; do sed -i '1i// Code generated by oapi-codegen. DO NOT EDIT.' $$f; done

aether-roc-api-base-docker: # @HELP build aether-roc-api base Docker image
	@go mod vendor
	docker build . -f build/base/Dockerfile \
		--build-arg ONOS_BUILD_VERSION=${ONOS_BUILD_VERSION} \
		--build-arg ONOS_MAKE_TARGET=build \
		-t onosproject/aether-roc-api-base:${AETHER_ROC_API_VERSION}
	@rm -rf vendor

aether-roc-api-docker: aether-roc-api-base-docker # @HELP build aether-roc-api Docker image
	docker build . -f build/aether-roc-api/Dockerfile \
		--build-arg AETHER_ROC_API_BASE_VERSION=${AETHER_ROC_API_VERSION} \
		-t onosproject/aether-roc-api:${AETHER_ROC_API_VERSION}

images: # @HELP build all Docker images
images: build aether-roc-api-docker

kind: # @HELP build Docker images and add them to the currently configured kind cluster
kind: images
	@if [ "`kind get clusters`" = '' ]; then echo "no kind cluster found" && exit 1; fi
	kind load docker-image onosproject/aether-roc-api:${AETHER_ROC_API_VERSION}

all: build images

publish: # @HELP publish version on github and dockerhub
	./../build-tools/publish-version ${VERSION} onosproject/aether-roc-api

bumponosdeps: # @HELP update "onosproject" go dependencies and push patch to git.
	./../build-tools/bump-onos-deps ${VERSION}

clean: # @HELP remove all the build artifacts
	rm -rf ./build/_output ./vendor ./cmd/aether-roc-api/aether-roc-api ./cmd/onos/onos
	go clean -testcache github.com/onosproject/aether-roc-api/...

help:
	@grep -E '^.*: *# *@HELP' $(MAKEFILE_LIST) \
    | sort \
    | awk ' \
        BEGIN {FS = ": *# *@HELP"}; \
        {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}; \
    '
