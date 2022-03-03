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
	"dGl3htpX5rXdxP4TAAD//zqfC+nKJgAA",
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
