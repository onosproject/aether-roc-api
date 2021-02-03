// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package utils

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/metadata"
	"io"
)

const (
	authorization = "Authorization"
	host          = "Host"
	userAgent     = "User-Agent"
	remoteAddr    = "remoteaddr"
)

//ReadRequestBody - read the bytes from the Request Body
func ReadRequestBody(bodyReader io.ReadCloser) ([]byte, error) {
	body := make([]byte, 0)
	buf := make([]byte, 100)
	for {
		count, err := bodyReader.Read(buf)
		body = append(body, buf[:count]...)
		if err == io.EOF {
			bodyReader.Close()
			break
		}
		if err != nil {
			return nil, fmt.Errorf("unable to read POST body %v", err)
		}
	}
	return body, nil
}

// NewGnmiContext - convert the HTTP context in to a gRPC Context
func NewGnmiContext(httpContext echo.Context) context.Context {

	return metadata.AppendToOutgoingContext(context.Background(),
		authorization, httpContext.Request().Header.Get(authorization),
		host, httpContext.Request().Host,
		"ua", httpContext.Request().Header.Get(userAgent), // `User-Agent` would be over written by gRPC
		remoteAddr, httpContext.Request().RemoteAddr)
}
