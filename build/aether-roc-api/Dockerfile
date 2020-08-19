ARG AETHER_ROC_API_BASE_VERSION=latest

FROM onosproject/aether-roc-api-base:$AETHER_ROC_API_BASE_VERSION as base

FROM alpine:3.12
RUN apk add libc6-compat

USER nobody

COPY --from=base /go/src/github.com/onosproject/aether-roc-api/build/_output/aether-roc-api /usr/local/bin/aether-roc-api

ENTRYPOINT ["aether-roc-api"]
