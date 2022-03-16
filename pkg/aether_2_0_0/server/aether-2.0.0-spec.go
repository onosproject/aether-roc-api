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
