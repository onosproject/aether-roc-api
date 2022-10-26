// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
//

package server

import (
	legacyrouter "github.com/getkin/kin-openapi/routers/legacy"
	"github.com/labstack/echo/v4"
	"github.com/onosproject/aether-roc-api/pkg/middleware/openapi3mw"
	"gotest.tools/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Passing a field like description with no length should give an error
// in validation because the minLength=1
func Test_ValidateRequestGivesError(t *testing.T) {
	e := echo.New()
	testDelete := `{
    "default-target": "defaultent",
    "Deletes": {
    },
    "Updates": {
        "traffic-class-2.1.0": [
			{
				"traffic-class-id": "sed",
				"description": ""
			}
		]
    },
    "Extensions": {
        "model-version-101": "2.1.0",
        "model-type-102": "Aether"
    }
}`
	req := httptest.NewRequest(http.MethodPatch, "/aether-roc-api", strings.NewReader(testDelete))
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	assert.Assert(t, c != nil)
	h := func(c echo.Context) error {
		body, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return err
		}
		return c.String(http.StatusOK, string(body))
	}
	assert.Assert(t, h != nil)
	openAPIDefinition, err := GetSwagger()
	assert.NilError(t, err)

	openapi3Router, err := legacyrouter.NewRouter(openAPIDefinition)
	assert.NilError(t, err)

	_, err = openapi3mw.ValidateRequest(c, openapi3Router)
	expectError := `code=400, message=Error at "/Updates/traffic-class-2.1.0/0/description": minimum string length is 1
Schema:
  {
    "description": "long description field",
    "maxLength": 1024,
    "minLength": 1,
    "title": "description",
    "type": "string"
  }

Value:
  ""
`
	assert.Error(t, err, expectError)
}

// Passing a description with a length of 0 in a delete is OK though
// because in the API, this is our way of telling the gNMI to delete it
func Test_ValidateRequestDeleteDesc(t *testing.T) {
	e := echo.New()
	testDelete := `{
    "default-target": "defaultent",
    "Updates": {
    },
    "Deletes": {
        "application-2.1.0": [
			{
				"application-id": "sed",
				"endpoint": [
					{
						"endpoint-id": "ep-1",
						"display-name": ""
					}
				],
				"additionalProperties": {
				  	"enterprise-id": "acme",
					"unchanged": "address"
				}
			}
		]
    },
    "Extensions": {
        "model-version-101": "2.1.0",
        "model-type-102": "Aether"
    }
}`
	req := httptest.NewRequest(http.MethodPatch, "/aether-roc-api", strings.NewReader(testDelete))
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	assert.Assert(t, c != nil)
	h := func(c echo.Context) error {
		body, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return err
		}
		return c.String(http.StatusOK, string(body))
	}
	assert.Assert(t, h != nil)
	openAPIDefinition, err := GetSwagger()
	assert.NilError(t, err)

	openapi3Router, err := legacyrouter.NewRouter(openAPIDefinition)
	assert.NilError(t, err)

	_, err = openapi3mw.ValidateRequest(c, openapi3Router)
	assert.NilError(t, err)
	assert.Equal(t, rec.Code, 200)
}
