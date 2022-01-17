FROM onosproject/golang-build:v1.0.0 as build

ENV GO111MODULE=on
ARG ONOS_MAKE_TARGET=build

COPY Makefile go.mod go.sum /go/src/github.com/onosproject/aether-roc-api/
COPY cmd/ /go/src/github.com/onosproject/aether-roc-api/cmd/
COPY pkg/ /go/src/github.com/onosproject/aether-roc-api/pkg/
COPY vendor/ /go/src/github.com/onosproject/aether-roc-api/vendor/

RUN cd /go/src/github.com/onosproject/aether-roc-api && GOFLAGS=-mod=vendor make ${ONOS_MAKE_TARGET}

FROM alpine:3.12
RUN apk add libc6-compat

USER nobody

COPY --from=build /go/src/github.com/onosproject/aether-roc-api/build/_output/aether-roc-api /usr/local/bin/aether-roc-api
COPY assets/ /assets
COPY api/ /

ENTRYPOINT ["aether-roc-api"]
