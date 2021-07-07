// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package openapi3mw

import (
	"bytes"
	"context"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
	"strings"
)

// CustomParamDecoder a helper for openapi3filter for decoding text parameters
func CustomParamDecoder(param *openapi3.Parameter, values []string) (interface{}, *openapi3.Schema, error) {
	if len(values) == 0 {
		return nil, nil, fmt.Errorf("no value for param %s in CustomParamDecoder()", param.Name)
	}
	schema := param.Content.Get("text/plain; charset=utf-8").Schema.Value
	return values[0], schema, nil
}

// ValidateRequest - validate that a HTTP request matches the OpenAPI3 schema
func ValidateRequest(ctx echo.Context, openAPI3Router routers.Router) (*openapi3filter.RequestValidationInput, error) {
	url := *ctx.Request().URL
	url.Host = ctx.Request().Host
	url.Scheme = ctx.Scheme()
	route, pathParams, err := openAPI3Router.FindRoute(ctx.Request())
	if err != nil {
		switch typedErr := err.(type) {
		default:
			return nil, ctx.JSON(http.StatusBadRequest, typedErr.Error())
		}
	}
	// Validate request
	requestValidationInput := &openapi3filter.RequestValidationInput{
		Request:      ctx.Request(),
		PathParams:   pathParams,
		Route:        route,
		ParamDecoder: CustomParamDecoder,
	}
	if err = openapi3filter.ValidateRequest(context.TODO(), requestValidationInput); err != nil {
		if err != nil {
			switch typedErr := err.(type) {
			case *openapi3filter.RequestError:
				if typedErr.Reason == "doesn't match the schema" {
					switch reasonErr := typedErr.Err.(type) {
					case *openapi3.SchemaError:
						errString := reasonErr.Error()
						if strings.HasPrefix(errString, "Error at \"/Deletes") &&
							(reasonErr.SchemaField == "minLength" || reasonErr.SchemaField == "required") {
							return requestValidationInput, nil
						}
						return nil, ctx.JSON(http.StatusBadRequest, reasonErr.Error())
					}
					return nil, ctx.JSON(http.StatusBadRequest, typedErr.Error())
				}
				return nil, ctx.JSON(http.StatusBadRequest, typedErr.Error())
			default:
				return nil, ctx.JSON(http.StatusInternalServerError, typedErr.Error())
			}
		}
	}
	return requestValidationInput, nil
}

// ValidateResponse - validate the response matches the schema before sending it out
func ValidateResponse(ctx echo.Context, rvi *openapi3filter.RequestValidationInput, resBody *bytes.Buffer) error {
	responseValidationInput := &openapi3filter.ResponseValidationInput{
		RequestValidationInput: rvi,
		Status:                 ctx.Response().Status,
		Header: http.Header{
			"Content-Type": []string{ctx.Response().Header().Get("Content-Type")},
		},
		Body: ioutil.NopCloser(bytes.NewReader(resBody.Bytes())),
	}

	// Validate response.
	if err := openapi3filter.ValidateResponse(context.TODO(), responseValidationInput); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return nil
}
