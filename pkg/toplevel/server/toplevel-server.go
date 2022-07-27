// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
//

package server

import (
	"github.com/labstack/echo/v4"
	"github.com/onosproject/aether-roc-api/pkg/middleware/openapi3mw"
)

// TopLevelServerInterface represents all server handlers.
type TopLevelServerInterface interface {
	// PATCH at the top level of aether-roc-api
	// (PATCH /aether-roc-api)
	PatchAetherRocAPI(ctx echo.Context) error
	// GET /targets A list of just target names
	// (GET /targets)
	GetTargets(ctx echo.Context) error
	// (GET /transactions)
	GetTransactions(ctx echo.Context) error
	// (POST /sdcore/synchronize/{address})
	PostSdcoreSynchronize(ctx echo.Context) error
	// GET /spec The OpenAPI specification for this service
	GetSpec(ctx echo.Context) error
	// GET /spec/aether-2.0.0-openapi3.yaml The OpenAPI specification for Aether 2.0.0
	GetAether200Spec(ctx echo.Context) error
	// GET /spec/aether-2.1.0-openapi3.yaml The OpenAPI specification for Aether 2.1.0
	GetAether210Spec(ctx echo.Context) error
	// GET /spec/aether-app-gtwy-openapi3.yaml The OpenAPI specification for Aether App Gateway
	GetAetherAppGtwySpec(ctx echo.Context) error
	// GET /spec/sdn-fabric-0.1.0-openapi3.yaml The OpenAPI specification for SDN Fabric
	GetSdnFabric010Spec(ctx echo.Context) error
}

// TopLevelInterfaceWrapper converts echo contexts to parameters.
type TopLevelInterfaceWrapper struct {
	Handler TopLevelServerInterface
}

// PatchAetherRocAPI converts echo context to params.
func (w *TopLevelInterfaceWrapper) PatchAetherRocAPI(ctx echo.Context) error {

	// Invoke the callback with all the unmarshalled arguments
	return w.Handler.PatchAetherRocAPI(ctx)
}

// GetTargets - get the full list of targets (devices)
func (w *TopLevelInterfaceWrapper) GetTargets(ctx echo.Context) error {

	// Invoke the callback with all the unmarshalled arguments
	return w.Handler.GetTargets(ctx)
}

// GetTransactions - get the full list of transactions (network-changes)
func (w *TopLevelInterfaceWrapper) GetTransactions(ctx echo.Context) error {

	// Invoke the callback with all the unmarshalled arguments
	return w.Handler.GetTransactions(ctx)
}

// PostSdcoreSynchronize - call synchronize on the sdcore adapter
func (w *TopLevelInterfaceWrapper) PostSdcoreSynchronize(ctx echo.Context) error {

	// Invoke the callback with all the unmarshalled arguments
	return w.Handler.PostSdcoreSynchronize(ctx)
}

// GetSpec - Get the OpenAPI3 specification in YAML format
func (w *TopLevelInterfaceWrapper) GetSpec(ctx echo.Context) error {

	// Invoke the callback with all the unmarshalled arguments
	return w.Handler.GetSpec(ctx)
}

// GetAether200Spec - Get the Aether 2.0.0 part of the OpenAPI3 specification
func (w *TopLevelInterfaceWrapper) GetAether200Spec(ctx echo.Context) error {

	// Invoke the callback with all the unmarshalled arguments
	return w.Handler.GetAether200Spec(ctx)
}

// GetAether210Spec - Get the Aether 2.1.0 part of the OpenAPI3 specification
func (w *TopLevelInterfaceWrapper) GetAether210Spec(ctx echo.Context) error {

	// Invoke the callback with all the unmarshalled arguments
	return w.Handler.GetAether210Spec(ctx)
}

// GetAetherAppGtwySpec - Get the Aether app gateway part of the OpenAPI3 specification
func (w *TopLevelInterfaceWrapper) GetAetherAppGtwySpec(ctx echo.Context) error {

	// Invoke the callback with all the unmarshalled arguments
	return w.Handler.GetAetherAppGtwySpec(ctx)
}

// GetSdnFabric010Spec - Get the SDN Fabric part of the OpenAPI3 specification
func (w *TopLevelInterfaceWrapper) GetSdnFabric010Spec(ctx echo.Context) error {

	// Invoke the callback with all the unmarshalled arguments
	return w.Handler.GetSdnFabric010Spec(ctx)
}

// EchoRouter is a simple interface which specifies echo.Route addition functions which
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
func RegisterHandlers(router EchoRouter, si TopLevelServerInterface) error {
	openAPIDefinition, err := GetSwagger()
	if err != nil {
		return err
	}

	wrapper := TopLevelInterfaceWrapper{
		Handler: si,
	}

	router.PATCH("/aether-roc-api", wrapper.PatchAetherRocAPI, openapi3mw.ValidateOpenapi3(openAPIDefinition))
	router.GET("/targets", wrapper.GetTargets)
	router.GET("/transactions", wrapper.GetTransactions)
	router.GET("/aether-top-level-openapi3.yaml", wrapper.GetSpec)
	router.GET("/aether-2.0.0-openapi3.yaml", wrapper.GetAether200Spec)
	router.GET("/aether-2.1.0-openapi3.yaml", wrapper.GetAether210Spec)
	router.GET("/aether-app-gtwy-openapi3.yaml", wrapper.GetAetherAppGtwySpec)
	router.GET("/sdn-fabric-0.1.0-openapi3.yaml", wrapper.GetSdnFabric010Spec)
	router.POST("/sdcore/synchronize/:service", wrapper.PostSdcoreSynchronize)

	return nil
}
