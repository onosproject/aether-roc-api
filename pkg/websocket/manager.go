// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package websocket

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/websocket"
)

var mgr Manager

var openWebSockets = make(map[*websocket.Conn]chan string)
var hbInterval uint = 30

// Manager single point of entry for the ROC system.
type Manager struct {
	echoRouter *echo.Echo
}

// NewManager - create a new instance of Manager
func NewManager(allowCorsOrigins []string, heartbeat uint) (*Manager, error) {
	hbInterval = heartbeat
	mgr = Manager{}
	var err error
	mgr.echoRouter = echo.New()
	if len(allowCorsOrigins) > 0 {
		mgr.echoRouter.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: allowCorsOrigins,
			AllowHeaders: []string{echo.HeaderAccessControlAllowOrigin, echo.HeaderContentType, echo.HeaderAuthorization},
		}))
	}

	mgr.echoRouter.Use(middleware.Logger())
	mgr.echoRouter.Use(middleware.Recover())
	mgr.echoRouter.GET("/ws", serveWS)
	mgr.echoRouter.POST("/webhook", webhook)
	return &mgr, err
}

// Run - start the manager
func (m *Manager) Run(port uint) {
	address := fmt.Sprintf(":%d", port)
	mgr.echoRouter.Logger.Fatal(m.echoRouter.Start(address))

}
