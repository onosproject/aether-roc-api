# Aether ROC API

Open API 3 YAML files are generated from config models.

## Retrieve model at runtime
The model files can be downloaded dynamically at runtime
```bash
curl -H "Accept: application/yaml" http://localhost:8181/spec
curl -H "Accept: application/yaml" http://localhost:8181/spec/aether-2.0.0-openapi3.yaml
curl -H "Accept: application/yaml" http://localhost:8181/spec/aether-4.0.0-openapi3.yaml
```

The first URL downloads the `aether-top-level-openapi3.yaml`,
and points at the other 2 models internally.

## Build
For example
```bash
cd ~/go/src/github.com/onosproject/config-models/models/aether-2.0.x
go run ./cmd/openapi-gen -o ~/go/src/github.com/onosproject/aether-roc-api/api/aether-2.0.0-openapi3.yaml
``` 

To generate Go code from this OpenAPI defintion, install and run **openapi-gen**
```bash
go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen
make oapi-codegen-aether-4.0.0 oapi-codegen-aether-2.0.0
```
