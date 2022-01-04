export CGO_ENABLED=0
export GO111MODULE=on

.PHONY: build

AETHER_ROC_API_VERSION := latest
ONOS_BUILD_VERSION := v0.6.9
OAPI_CODEGEN_VERSION := v1.7.0

build: # @HELP build the Go binaries and run all validations (default)
build:
	CGO_ENABLED=1 go build -o build/_output/aether-roc-api ./cmd/aether-roc-api
	CGO_ENABLED=1 go build -o build/_output/aether-roc-websocket ./cmd/aether-roc-websocket

test: # @HELP run the unit tests and source code validation
test: build deps linters license_check openapi-linters
	CGO_ENABLED=1 go test -race github.com/onosproject/aether-roc-api/pkg/...
	CGO_ENABLED=1 go test -race github.com/onosproject/aether-roc-api/cmd/...

jenkins-test:  # @HELP run the unit tests and source code validation producing a junit style report for Jenkins
jenkins-test: build-tools deps license_check linters # openapi-linters
	CGO_ENABLED=1 TEST_PACKAGES=github.com/onosproject/aether-roc-api/... ./../build-tools/build/jenkins/make-unit

deps: # @HELP ensure that the required dependencies are in place
	go build -v ./...
	bash -c "diff -u <(echo -n) <(git diff go.mod)"
	bash -c "diff -u <(echo -n) <(git diff go.sum)"

linters: golang-ci # @HELP examines Go source code and reports coding problems
	golangci-lint run --timeout 5m

build-tools: # @HELP install the ONOS build tools if needed
	@if [ ! -d "../build-tools" ]; then cd .. && git clone https://github.com/onosproject/build-tools.git; fi

jenkins-tools: # @HELP installs tooling needed for Jenkins
	cd .. && go get -u github.com/jstemmer/go-junit-report && go get github.com/t-yuki/gocover-cobertura

golang-ci: # @HELP install golang-ci if not present
	golangci-lint --version || curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b `go env GOPATH`/bin v1.42.0

oapi-codegen:
	oapi-codegen || ( cd .. && go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen@${OAPI_CODEGEN_VERSION})

openapi-spec-validator: # @HELP install openapi-spec-validator
	openapi-spec-validator -h || python -m pip install openapi-spec-validator==0.3.1

license_check: # @HELP examine and ensure license headers exist
license_check: build-tools
	./../build-tools/licensing/boilerplate.py -v --rootdir=${CURDIR} --boilerplate LicenseRef-ONF-Member-1.0

openapi-linters: # @HELP lints the Open API specifications
openapi-linters: openapi-spec-validator
	openapi-spec-validator api/aether-top-level-openapi3.yaml
	openapi-spec-validator api/aether-4.0.0-openapi3.yaml
	openapi-spec-validator api/aether-5.0.0-openapi3.yaml

gofmt: # @HELP run the Go format validation
	bash -c "diff -u <(echo -n) <(gofmt -d pkg/ cmd/ tests/)"

oapi-codegen-aether-4.0.0: # @HELP generate openapi types from aether-4.0.0-openapi3.yaml
oapi-codegen-aether-4.0.0: oapi-codegen
	mkdir -p pkg/aether_4_0_0/types pkg/aether_4_0_0/server
	oapi-codegen -generate types -package types -o pkg/aether_4_0_0/types/aether-4.0.0-types.go api/aether-4.0.0-openapi3.yaml
	oapi-codegen -generate spec -package server -o pkg/aether_4_0_0/server/aether-4.0.0-spec.go api/aether-4.0.0-openapi3.yaml
	oapi-codegen \
		-generate types,server \
		-import-mapping externalRef0:"github.com/onosproject/aether-roc-api/pkg/aether_4_0_0/types",externalRef1:"github.com/onosproject/config-models/modelplugin/aether-4.0.0/aether_4_0_0" \
		-package server \
		-templates pkg/codegen/templates \
		-o pkg/aether_4_0_0/server/aether-4.0.0-impl.go \
		api/aether-4.0.0-openapi3.yaml
	sed -i "s/target Target/target externalRef0.Target/g" pkg/aether_4_0_0/server/aether-4.0.0-impl.go
	oapi-codegen \
		-generate server \
		-import-mapping externalRef0:"github.com/onosproject/aether-roc-api/pkg/aether_4_0_0/types" \
		-package server \
		-templates pkg/codegen/modified \
		-o pkg/aether_4_0_0/server/aether-4.0.0-server.go \
		api/aether-4.0.0-openapi3.yaml
	sed -i "s/target Target/target externalRef0.Target/g" pkg/aether_4_0_0/server/aether-4.0.0-server.go
	oapi-codegen \
		-generate types \
		-import-mapping externalRef0:"github.com/onosproject/config-models/modelplugin/aether-4.0.0/aether_4_0_0" \
		-package server \
		-templates pkg/codegen/convert-oapi-gnmi \
		-o pkg/aether_4_0_0/server/aether-4.0.0-convert-oapi-gnmi.go \
		api/aether-4.0.0-openapi3.yaml
	oapi-codegen \
		-generate types \
		-import-mapping externalRef0:"github.com/onosproject/config-models/modelplugin/aether-4.0.0/aether_4_0_0" \
		-package server \
		-templates pkg/codegen/convert-gnmi-oapi \
		-o pkg/aether_4_0_0/server/aether-4.0.0-convert-gnmi-oapi.go \
		api/aether-4.0.0-openapi3.yaml
	for f in pkg/aether_4_0_0/**/aether-4.0.0*.go; do sed -i '1i// Code generated by oapi-codegen. DO NOT EDIT.' $$f; done

oapi-codegen-aether-5.0.0: # @HELP generate openapi types from aether-5.0.0-openapi3.yaml
oapi-codegen-aether-5.0.0: oapi-codegen
	mkdir -p pkg/aether_5_0_0/types pkg/aether_5_0_0/server
	oapi-codegen -generate types -package types -o pkg/aether_5_0_0/types/aether-5.0.0-types.go api/aether-5.0.0-openapi3.yaml
	oapi-codegen -generate spec -package server -o pkg/aether_5_0_0/server/aether-5.0.0-spec.go api/aether-5.0.0-openapi3.yaml
	oapi-codegen \
		-generate types,server \
		-import-mapping externalRef0:"github.com/onosproject/aether-roc-api/pkg/aether_5_0_0/types",externalRef1:"github.com/onosproject/config-models/modelplugin/aether-5.0.0/aether_5_0_0" \
		-package server \
		-templates pkg/codegen/templates \
		-o pkg/aether_5_0_0/server/aether-5.0.0-impl.go \
		api/aether-5.0.0-openapi3.yaml
	sed -i "s/target Target/target externalRef0.Target/g" pkg/aether_5_0_0/server/aether-5.0.0-impl.go
	oapi-codegen \
		-generate server \
		-import-mapping externalRef0:"github.com/onosproject/aether-roc-api/pkg/aether_5_0_0/types" \
		-package server \
		-templates pkg/codegen/modified \
		-o pkg/aether_5_0_0/server/aether-5.0.0-server.go \
		api/aether-4.0.0-openapi3.yaml
	sed -i "s/target Target/target externalRef0.Target/g" pkg/aether_5_0_0/server/aether-5.0.0-server.go
	oapi-codegen \
		-generate types \
		-import-mapping externalRef0:"github.com/onosproject/config-models/modelplugin/aether-5.0.0/aether_5_0_0" \
		-package server \
		-templates pkg/codegen/convert-oapi-gnmi \
		-o pkg/aether_5_0_0/server/aether-5.0.0-convert-oapi-gnmi.go \
		api/aether-5.0.0-openapi3.yaml
	oapi-codegen \
		-generate types \
		-import-mapping externalRef0:"github.com/onosproject/config-models/modelplugin/aether-5.0.0/aether_5_0_0" \
		-package server \
		-templates pkg/codegen/convert-gnmi-oapi \
		-o pkg/aether_5_0_0/server/aether-5.0.0-convert-gnmi-oapi.go \
		api/aether-5.0.0-openapi3.yaml
	for f in pkg/aether_5_0_0/**/aether-5.0.0*.go; do sed -i '1i// Code generated by oapi-codegen. DO NOT EDIT.' $$f; done

aether-top-level: # @HELP generate openapi types from aether-top-level-openapi3.yaml
aether-top-level: oapi-codegen
	oapi-codegen -generate types -package types \
	-import-mapping \
	./aether-5.0.0-openapi3.yaml:github.com/onosproject/aether-roc-api/pkg/aether_5_0_0/types,\	
	./aether-4.0.0-openapi3.yaml:github.com/onosproject/aether-roc-api/pkg/aether_4_0_0/types \
	-o pkg/toplevel/types/toplevel-types.go api/aether-top-level-openapi3.yaml

	oapi-codegen -generate spec -package server \
	-import-mapping \
	./aether-5.0.0-openapi3.yaml:github.com/onosproject/aether-roc-api/pkg/aether_5_0_0/types,\	
	./aether-4.0.0-openapi3.yaml:github.com/onosproject/aether-roc-api/pkg/aether_4_0_0/types \
	-o pkg/toplevel/server/toplevel-spec.go api/aether-top-level-openapi3.yaml

aether-roc-api-docker: # @HELP build aether-roc-api Docker image
	@go mod vendor
	docker build . -f build/aether-roc-api/Dockerfile \
		-t onosproject/aether-roc-api:${AETHER_ROC_API_VERSION}
	@rm -rf vendor

aether-roc-websocket-docker: # @HELP build aether-roc-websocket Docker image
	@go mod vendor
	docker build . -f build/aether-roc-websocket/Dockerfile \
		-t onosproject/aether-roc-websocket:${AETHER_ROC_API_VERSION}
	@rm -rf vendor

images: # @HELP build all Docker images
images: build aether-roc-api-docker aether-roc-websocket-docker

kind: # @HELP build Docker images and add them to the currently configured kind cluster
kind: images
	@if [ "`kind get clusters`" = '' ]; then echo "no kind cluster found" && exit 1; fi
	kind load docker-image onosproject/aether-roc-api:${AETHER_ROC_API_VERSION}
	kind load docker-image onosproject/aether-roc-websocket:${AETHER_ROC_API_VERSION}

all: build images

publish: # @HELP publish version on github and dockerhub
	./../build-tools/publish-version ${VERSION} onosproject/aether-roc-api onosproject/aether-roc-websocket

jenkins-publish: build-tools jenkins-tools # @HELP Jenkins calls this to publish artifacts
	./build/bin/push-images
	../build-tools/release-merge-commit

bumponosdeps: # @HELP update "onosproject" go dependencies and push patch to git.
	./../build-tools/bump-onos-deps ${VERSION}

generated: # @HELP create generated artifacts
generated: oapi-codegen-aether-4.0.0 oapi-codegen-aether-5.0.0

clean: # @HELP remove all the build artifacts
	rm -rf ./build/_output ./vendor ./cmd/aether-roc-api/aether-roc-api ./cmd/aether-roc-websocket/aether-roc-websocket
	go clean -testcache github.com/onosproject/aether-roc-api/...

clean-generated: # @HELP remove generated artifacts
	rm -f pkg/aether_4_0_0/**/aether-4.0.0*.go
	rm -f pkg/aether_5_0_0/**/aether-5.0.0*.go

.PHONY: local-models
local-models:
	rm -rf local-models
	cp -a ../config-models/modelplugin/aether-5.0.0 local-models
redo: clean-generated local-models generated
	m-ara && p-ara	

help:
	@grep -E '^.*: *# *@HELP' $(MAKEFILE_LIST) \
    | sort \
    | awk ' \
        BEGIN {FS = ": *# *@HELP"}; \
        {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}; \
    '
