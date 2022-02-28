// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
//

package utils

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

const (
	notEqual   = "not equal to any target nodes"
	hasLrPath  = "has leafref path"
	schemaPath = "schema path"
)

const (
	respInternalInvalid   = `rpc error: code = Internal desc = rpc error: code = InvalidArgument`
	respInvalidValidation = `rpc error: code = InvalidArgument desc = rpc error: code = InvalidArgument desc = validation error field name`
	respInvalidBase       = `rpc error: code = InvalidArgument desc =`
	respUnauthorized      = `rpc error: code = Unauthenticated desc =`
)

// ConvertGrpcError - capture gRPC error messages properly
func ConvertGrpcError(err error) *echo.HTTPError {

	// if the error is already the right type, just return it
	switch e := err.(type) {
	case *echo.HTTPError:
		return e
	}

	if strings.HasPrefix(err.Error(), respInternalInvalid) {
		return echo.NewHTTPError(http.StatusNoContent, err.Error())
	} else if strings.HasPrefix(err.Error(), respInvalidValidation) {
		var msg string
		remainingErr := err.Error()[110:]
		if strings.HasSuffix(remainingErr, notEqual) {
			firstMsg := remainingErr[:strings.Index(remainingErr, notEqual)] // Message is repeated - strip out first instance
			if strings.Contains(firstMsg, " (") && strings.Contains(firstMsg, hasLrPath) && strings.Contains(firstMsg, schemaPath) {
				idxFieldEnd := strings.Index(firstMsg, " (")
				offendingInst := firstMsg[:idxFieldEnd]
				endSpIdx := strings.Index(firstMsg, schemaPath) + len(schemaPath)
				hasLrPathIdx := strings.Index(firstMsg, hasLrPath)
				orphanPath := firstMsg[endSpIdx : hasLrPathIdx-1]
				offendingPath := firstMsg[hasLrPathIdx+len(hasLrPath)+1 : len(firstMsg)-1]
				msg = fmt.Sprintf("Change gives LeafRef error on %s. %s not present. From path: %s", offendingPath, offendingInst, orphanPath)
			} else {
				msg = firstMsg
			}
		} else {
			msg = remainingErr
		}
		return echo.NewHTTPError(http.StatusBadRequest, msg)
	} else if strings.HasPrefix(err.Error(), respInvalidBase) {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	} else if strings.HasPrefix(err.Error(), respUnauthorized) {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	} else {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
}
