// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package manager

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	aether_3_0_0 "github.com/onosproject/aether-roc-api/pkg/aether_3_0_0/server"
	aether_4_0_0 "github.com/onosproject/aether-roc-api/pkg/aether_4_0_0/server"
	"github.com/onosproject/aether-roc-api/pkg/southbound"
	toplevel "github.com/onosproject/aether-roc-api/pkg/toplevel/server"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"google.golang.org/grpc"
)

var mgr Manager

var log = logging.GetLogger("manager")

// Manager single point of entry for the ROC system.
type Manager struct {
	gnmiClient    southbound.GnmiClient
	echoRouter    *echo.Echo
	openapis      map[string]interface{}
	authorization bool
}

// NewManager -
func NewManager(gnmiEndpoint string, allowCorsOrigins []string,
	validateResponses bool, authorization bool, opts ...grpc.DialOption) (*Manager, error) {
	mgr = Manager{authorization: authorization}

	var err error
	mgr.gnmiClient = new(southbound.GNMIProvisioner)
	err = mgr.gnmiClient.Init(gnmiEndpoint, opts...)
	if err != nil {
		log.Error("Unable to setup GNMI provisioner", err)
		return nil, err
	}

	mgr.openapis = make(map[string]interface{})
	aether30APIImpl := &aether_3_0_0.ServerImpl{
		GnmiClient: mgr.gnmiClient,
	}
	mgr.openapis["Aether-3.0.0"] = aether30APIImpl
	aether40APIImpl := &aether_4_0_0.ServerImpl{
		GnmiClient: mgr.gnmiClient,
	}
	mgr.openapis["Aether-4.0.0"] = aether40APIImpl
	topLevelAPIImpl := &toplevel.ServerImpl{
		GnmiClient:    mgr.gnmiClient,
		Authorization: authorization,
	}
	mgr.openapis["TopLevel"] = topLevelAPIImpl

	mgr.echoRouter = echo.New()
	if len(allowCorsOrigins) > 0 {
		mgr.echoRouter.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: allowCorsOrigins,
			AllowHeaders: []string{echo.HeaderAccessControlAllowOrigin, echo.HeaderContentType, echo.HeaderAuthorization},
		}))
	}
	mgr.echoRouter.File("/", "assets/index.html")
	mgr.echoRouter.Static("/", "assets")
	if err := aether_3_0_0.RegisterHandlers(mgr.echoRouter, aether30APIImpl, validateResponses); err != nil {
		return nil, fmt.Errorf("aether_3_0_0.RegisterHandlers()  %s", err)
	}
	if err := aether_4_0_0.RegisterHandlers(mgr.echoRouter, aether40APIImpl, validateResponses); err != nil {
		return nil, fmt.Errorf("aether_4_0_0.RegisterHandlers()  %s", err)
	}
	if err := toplevel.RegisterHandlers(mgr.echoRouter, topLevelAPIImpl); err != nil {
		return nil, fmt.Errorf("toplevel.RegisterHandlers()  %s", err)
	}

	return &mgr, nil
}

// Run starts the northbound services.
func (m *Manager) Run(port uint) {
	log.Infof("Starting Manager on port %d", port)

	m.echoRouter.Logger.Fatal(m.echoRouter.Start(fmt.Sprintf(":%d", port)))

	log.Warn("Manager Stopping")
}
