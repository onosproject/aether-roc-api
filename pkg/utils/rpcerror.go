// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package utils

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// ConvertGrpcError - capture gRPC error messages properly
func ConvertGrpcError(err error) *echo.HTTPError {
	if strings.HasPrefix(err.Error(), "rpc error: code = Internal desc = rpc error: code = InvalidArgument") {
		return echo.NewHTTPError(http.StatusNoContent, err.Error())
	} else if strings.HasPrefix(err.Error(), "rpc error: code = InvalidArgument desc =") {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	} else if strings.HasPrefix(err.Error(), "rpc error: code = Unauthenticated desc =") {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	} else {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
}
