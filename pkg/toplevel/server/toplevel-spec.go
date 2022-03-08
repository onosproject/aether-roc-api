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
