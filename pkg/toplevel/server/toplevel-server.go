// Code generated by oapi-codegen. DO NOT EDIT.
// Package server provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package server

import (
	"github.com/labstack/echo/v4"
	"github.com/onosproject/aether-roc-api/pkg/middleware/openapi3mw"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// PATCH at the top level of aether-roc-api
	// (PATCH /aether-roc-api)
	PatchAetherRocApi(ctx echo.Context) error
	// GET /targets A list of just target names
	// (GET /targets)
	GetTargets(ctx echo.Context) error
	// (POST /sdcore/synchronize/{address})
	PostSdcoreSynchronize(ctx echo.Context) error
	// GET /spec The OpenAPI specification for this service
	GetSpec(ctx echo.Context) error
	// GET /spec/aether-2.1.0-openapi3.yaml The OpenAPI specification for Aether 2.1.0
	GetAether210Spec(ctx echo.Context) error
	// GET /spec/aether-3.0.0-openapi3.yaml The OpenAPI specification for Aether 3.0.0
	GetAether300Spec(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// PatchAetherRocApi converts echo context to params.
func (w *ServerInterfaceWrapper) PatchAetherRocApi(ctx echo.Context) error {

	// Invoke the callback with all the unmarshalled arguments
	return w.Handler.PatchAetherRocApi(ctx)
}

// GetTargets - get the full list of targets (devices)
func (w *ServerInterfaceWrapper) GetTargets(ctx echo.Context) error {

	// Invoke the callback with all the unmarshalled arguments
	return w.Handler.GetTargets(ctx)
}

// PostSdcoreSynchronize - call synchronize on the sdcore adapter
func (w *ServerInterfaceWrapper) PostSdcoreSynchronize(ctx echo.Context) error {

	// Invoke the callback with all the unmarshalled arguments
	return w.Handler.PostSdcoreSynchronize(ctx)
}

// GetSpec - Get the OpenAPI3 specification in YAML format
func (w *ServerInterfaceWrapper) GetSpec(ctx echo.Context) error {

	// Invoke the callback with all the unmarshalled arguments
	return w.Handler.GetSpec(ctx)
}

// GetAether210Spec - Get the Aether 2.1.0 part of the OpenAPI3 specification
func (w *ServerInterfaceWrapper) GetAether210Spec(ctx echo.Context) error {

	// Invoke the callback with all the unmarshalled arguments
	return w.Handler.GetAether210Spec(ctx)
}

// GetAether300Spec - Get the Aether 3.0.0 part of the OpenAPI3 specification
func (w *ServerInterfaceWrapper) GetAether300Spec(ctx echo.Context) error {

	// Invoke the callback with all the unmarshalled arguments
	return w.Handler.GetAether300Spec(ctx)
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) error {
	openApiDefinition, err := GetSwagger()
	if err != nil {
		return err
	}

	wrapper := ServerInterfaceWrapper{
		Handler:        si,
	}

	router.PATCH("/aether-roc-api", wrapper.PatchAetherRocApi, openapi3mw.ValidateOpenapi3(openApiDefinition))
	router.GET("/targets", wrapper.GetTargets)
	router.GET("/aether-top-level-openapi3.yaml", wrapper.GetSpec)
	router.GET("/aether-2.1.0-openapi3.yaml", wrapper.GetAether210Spec)
	router.GET("/aether-3.0.0-openapi3.yaml", wrapper.GetAether300Spec)
	router.POST("/sdcore/synchronize/:service", wrapper.PostSdcoreSynchronize)

	return nil
}
