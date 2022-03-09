// Code generated by GENERATOR. DO NOT EDIT.
// Package server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.7.0 DO NOT EDIT.
package server

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	externalRef0 "github.com/onosproject/aether-roc-api/pkg/aether_2_0_0/server"
	externalRef1 "github.com/onosproject/aether-roc-api/pkg/aether_4_0_0/server"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"AA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	pathPrefix := path.Dir(pathToFile)

	for rawPath, rawFunc := range externalRef0.PathToRawSpec(path.Join(pathPrefix, "./aether-2.0.0-openapi3.yaml")) {
		if _, ok := res[rawPath]; ok {
			// it is not possible to compare functions in golang, so always overwrite the old value
		}
		res[rawPath] = rawFunc
	}
	for rawPath, rawFunc := range externalRef1.PathToRawSpec(path.Join(pathPrefix, "./aether-4.0.0-openapi3.yaml")) {
		if _, ok := res[rawPath]; ok {
			// it is not possible to compare functions in golang, so always overwrite the old value
		}
		res[rawPath] = rawFunc
	}
	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
