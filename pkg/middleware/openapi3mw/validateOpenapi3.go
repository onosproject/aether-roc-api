// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
//

package openapi3mw

import (
	"github.com/getkin/kin-openapi/openapi3"
	legacyrouter "github.com/getkin/kin-openapi/routers/legacy"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger("middleware", "openapi3mw")

type (
	// ValidateOpenapi3Config defines the config for ValidateOpenapi3 middleware.
	ValidateOpenapi3Config struct {
		// Skipper defines a function to skip middleware.
		Skipper middleware.Skipper

		// OpenAPI3 is a router specific to the section
		OpenAPI3 *openapi3.T
	}
)

var (
	// DefaultValidateOpenapi3Config is the default ValidateOpenapi3Config middleware config.
	DefaultValidateOpenapi3Config = ValidateOpenapi3Config{
		Skipper: middleware.DefaultSkipper,
	}
)

// ValidateOpenapi3 returns a ValidateOpenapi3 middleware.
func ValidateOpenapi3(openapi3 *openapi3.T) echo.MiddlewareFunc {
	c := ValidateOpenapi3Config{
		OpenAPI3: openapi3,
	}
	return ValidateOpenapi3WithConfig(c)
}

// ValidateOpenapi3WithConfig returns a ValidateOpenapi3 middleware with config.
// See: `ValidateOpenapi3()`.
func ValidateOpenapi3WithConfig(config ValidateOpenapi3Config) echo.MiddlewareFunc {
	// Defaults
	if config.Skipper == nil {
		config.Skipper = DefaultValidateOpenapi3Config.Skipper
	}

	openapi3Router, err := legacyrouter.NewRouter(config.OpenAPI3)
	if err != nil {
		log.Errorf("error loading swagger %s\n", err.Error())
		return nil
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			if config.Skipper(ctx) {
				return next(ctx)
			}
			// Request
			//var resBody *bytes.Buffer
			//if config.ValidateResponse {
			//	ctx.Response().Writer, resBody = ResponseWriter(ctx)
			//}

			log.Infof("Validating %s %s request\n", ctx.Path(), ctx.Request().Method)
			if _, err = ValidateRequest(ctx, openapi3Router); err != nil {
				return err
			}
			//Response
			//if config.ValidateResponse && ctx.Response().Size > 0 {
			//	return ValidateResponse(ctx, rvi, resBody)
			//}

			return next(ctx)
		}
	}
}
