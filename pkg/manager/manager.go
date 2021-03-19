// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package manager

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	aether_1_0_0 "github.com/onosproject/aether-roc-api/pkg/aether_1_0_0/server"
	aether_2_0_0 "github.com/onosproject/aether-roc-api/pkg/aether_2_0_0/server"
	aether_2_1_0 "github.com/onosproject/aether-roc-api/pkg/aether_2_1_0/server"
	"github.com/onosproject/aether-roc-api/pkg/southbound"
	toplevel "github.com/onosproject/aether-roc-api/pkg/toplevel/server"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"google.golang.org/grpc"
)

var mgr Manager

var log = logging.GetLogger("manager")

// Manager single point of entry for the ROC system.
type Manager struct {
	gnmiClient southbound.GnmiClient
	echoRouter *echo.Echo
	openapis   map[string]interface{}
}

// NewManager -
func NewManager(gnmiEndpoint string, allowCorsOrigins []string,
	validateResponses bool, opts ...grpc.DialOption) (*Manager, error) {
	mgr = Manager{}

	var err error
	mgr.gnmiClient = new(southbound.GNMIProvisioner)
	err = mgr.gnmiClient.Init(gnmiEndpoint, opts...)
	if err != nil {
		log.Error("Unable to setup GNMI provisioner", err)
		return nil, err
	}

	mgr.openapis = make(map[string]interface{})
	aetherAPIImpl := &aether_1_0_0.ServerImpl{
		GnmiClient: mgr.gnmiClient,
	}
	mgr.openapis["Aether-1.0.0"] = aetherAPIImpl
	aether2APIImpl := &aether_2_0_0.ServerImpl{
		GnmiClient: mgr.gnmiClient,
	}
	mgr.openapis["Aether-2.0.0"] = aether2APIImpl
	aether21APIImpl := &aether_2_1_0.ServerImpl{
		GnmiClient: mgr.gnmiClient,
	}
	mgr.openapis["Aether-2.1.0"] = aether21APIImpl
	topLevelAPIImpl := &toplevel.ServerImpl{
		GnmiClient: mgr.gnmiClient,
	}
	mgr.openapis["TopLevel"] = topLevelAPIImpl

	mgr.echoRouter = echo.New()
	if len(allowCorsOrigins) > 0 {
		mgr.echoRouter.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: allowCorsOrigins,
			AllowHeaders: []string{echo.HeaderAccessControlAllowOrigin, echo.HeaderContentType, echo.HeaderAuthorization},
		}))
	}
	aether_1_0_0.RegisterHandlers(mgr.echoRouter, aetherAPIImpl)
	if err := aether_2_0_0.RegisterHandlers(mgr.echoRouter, aether2APIImpl, validateResponses); err != nil {
		return nil, fmt.Errorf("aether_2_0_0.RegisterHandlers()  %s", err)
	}
	if err := aether_2_1_0.RegisterHandlers(mgr.echoRouter, aether21APIImpl, validateResponses); err != nil {
		return nil, fmt.Errorf("aether_2_1_0.RegisterHandlers()  %s", err)
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
