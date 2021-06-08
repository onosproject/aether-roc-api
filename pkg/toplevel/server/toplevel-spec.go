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
	externalRef0 "github.com/onosproject/aether-roc-api/pkg/aether_2_1_0/server"
	externalRef1 "github.com/onosproject/aether-roc-api/pkg/aether_3_0_0/server"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/6xWW2/cNhP9KwTzAUnwiZJ2AwSInuqmdhs0sXNxnmwjoMnRLh2JZEhqXdXQfy+GkvYe",
	"GK31tp7hnDPXYz1QYWprNOjgafFAvVhCzePP0wrq0WydseCCgvgXFwK8Z9aZUlXA5ukszdH+PwclLWia",
	"cQhLcL2DGQuaW/UqbXldPcs2fNlAlp3s4NEuodzqCeE3YIgtjNYgglqp0DIPbqXEBCRvj6D+lO1Vmh9n",
	"i44nsIEO4KxTfoKKTtdYe8hPzn4X+YeZcJU+mZ098iAahz2ajODLHmLP0o914Uxjp6DYgtvGt6ZSop2O",
	"oMfbZnDNND3aoEX05tYLp25HiCdhr7EQubHTTfar3cy06xL6kQex/NXI9lAAf4MKQv9zYDu+56OCdgk9",
	"/SuA9sroGCUBi7BBGU0L+sFIqEhoLRCuJVmBw4fElOR54G4B4TkxmpTK+UCEA45h5KpS+vvNi2UI1hdZ",
	"Jo3wqdHGW2fuQITUuEWGfzNhdKkW8UG20LX6ButUsmeNB2ZKtjaxWT5jEuL4hjyY0sxDcPCjAR9e0mSv",
	"GWLJ9QKY5jWwWR5HEFSogBYHroRikbSgPjilF9iXGmtnaGazfL4dvOf5aeyY5iyfHYZvOw8QumT9GhtD",
	"No3ZvDW32E1k+2ol/5dDl1Dypgqsn+Lh4C+jnbzoG06wTy9JMKTxQG5bMoQTVRJtAvEWhCoVSNwGpaVa",
	"yYZXpOnzymS/lIR7wjXhUipk4RUZxtXi5HgI4JD6KmdvOPv7+ppdX6ffbv5/tDs4c+VA0uJqv5abLhnS",
	"P+c1HJ6IHqz7oOswj3HxqQpQP9rWLa5unSp3jrc9qNKlidtodOAi9hpqripaRNcveP8awr1x35Ve4HXQ",
	"ZEiSXljQ5HztJGem0TJeGU1o4xBjPLMjMF1CKyVAe9jUTd/3ls9QsovzM/YBalQ/VKjHEbN7i0cbQIes",
	"sZXh0mfzfJ5n+etsC+xCVy37Yspwzx2wgZCtZmmeWlnGLoGr/UU5yPF/4Jy9yfJ55OzhlV4wriV79/Ez",
	"Ozv5xDArlr+OfJtbOom6Sy6NJe9hBRVN6HCGtKD9x0OX0EGQ16a4nMu4B6NyOyNYfPKAPrHEH7hjcTTv",
	"JC16jb40duQZRGpU7aGi+M1qbaVEjMzuPKYyfuw+tnmb/wNd1x+Ft0b7fnfnveDtnnVMFmQ8Id/UNXct",
	"5npy+fYPwgMJSyDBWFJh0qjye+ViWNafWeQYpGO38uGIdmo/zGuSBuzca2zBbrm/n16Siz8JUu5WjI6x",
	"EHJCKuUDlnvX+EB6c5S8CNr9EwAA//9SqoHbigwAAA==",
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

	for rawPath, rawFunc := range externalRef0.PathToRawSpec(path.Join(pathPrefix, "./aether-2.1.0-openapi3.yaml")) {
		if _, ok := res[rawPath]; ok {
			// it is not possible to compare functions in golang, so always overwrite the old value
		}
		res[rawPath] = rawFunc
	}
	for rawPath, rawFunc := range externalRef1.PathToRawSpec(path.Join(pathPrefix, "./aether-3.0.0-openapi3.yaml")) {
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
