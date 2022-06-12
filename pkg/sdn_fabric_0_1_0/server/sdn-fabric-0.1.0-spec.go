// Code generated by oapi-codegen. DO NOT EDIT.
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
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"0bq7nvha197dzIPS0u3/BQAA//9ssvVjz5cAAA==",
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
