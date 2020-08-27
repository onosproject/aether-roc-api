# Aether ROC API

Open API 3 YAML files are generated from config models.

For example
```bash
cd ~/go/src/github.com/onosproject/config-models/modelplugin/aether-1.0.0
go run ./cmd/openapi-gen -o ~/go/src/github.com/onosproject/aether-roc-api/api/aether-1.0.0-openapi3.yaml
``` 

There is a similar tool for **rbac-1.0.0** in its **modelplugin** folder.

To generate Go code from this OpenAPI defintion, install and run **openapi-gen**
```bash
go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen
make oapi-codegen-aether oapi-codegen-rbac
```
