ARG ONOS_BUILD_VERSION=undefined

FROM onosproject/golang-build:$ONOS_BUILD_VERSION

ENV GO111MODULE=on
ARG ONOS_MAKE_TARGET=build

COPY Makefile go.mod go.sum /go/src/github.com/onosproject/aether-roc-api/
COPY cmd/ /go/src/github.com/onosproject/aether-roc-api/cmd/
COPY pkg/ /go/src/github.com/onosproject/aether-roc-api/pkg/
COPY vendor/ /go/src/github.com/onosproject/aether-roc-api/vendor/

RUN cd /go/src/github.com/onosproject/aether-roc-api && GOFLAGS=-mod=vendor make ${ONOS_MAKE_TARGET}
