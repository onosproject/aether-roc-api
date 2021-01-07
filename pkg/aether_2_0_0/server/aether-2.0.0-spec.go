// Code generated by oapi-codegen. DO NOT EDIT.
// Package server provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package server

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"strings"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"a1YCm3kAAA==",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
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

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
