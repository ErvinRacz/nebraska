// Package codegen provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.2 DO NOT EDIT.
package codegen

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