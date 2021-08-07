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
