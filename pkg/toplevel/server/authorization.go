// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/onosproject/onos-lib-go/pkg/auth"
	"net/http"
	"strings"
)

func checkAuthorization(httpContext echo.Context, allowedGroups ...string) error {
	authHeader := httpContext.Request().Header.Get(authorization)
	if authHeader == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "no Authorization token")
	}

	jwtAuth := new(auth.JwtAuthenticator)
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorization header is not Bearer token")
	}
	authClaims, err := jwtAuth.ParseAndValidate(authHeader[7:])
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, fmt.Sprintf("Bad request. Bearer token. %s", err.Error()))
	}
	if err = authClaims.Valid(); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized,
			fmt.Sprintf("Bad request. Auth header not valid. %s", err.Error()))
	}

	username := ""
	if name, ok := authClaims["name"]; ok {
		username = name.(string)
	}

	groupsIf, ok := authClaims["groups"].([]interface{})
	if ok {
		for _, group := range groupsIf {
			for _, allowed := range allowedGroups {
				if group.(string) == allowed {
					log.Infof("%s called endpoint %s", username, httpContext.Request().URL)
					return nil
				}
			}
		}
	}

	return echo.NewHTTPError(http.StatusUnauthorized,
		fmt.Sprintf("User %s is not in %v", username, allowedGroups))
}
