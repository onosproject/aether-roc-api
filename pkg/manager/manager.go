// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package manager

import (
	"github.com/labstack/echo/v4"
	"github.com/onosproject/aether-roc-api/pkg/aether_1_0_0"
	"github.com/onosproject/aether-roc-api/pkg/rbac_1_0_0"
	"github.com/onosproject/aether-roc-api/pkg/southbound"
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
func NewManager(opts ...grpc.DialOption) (*Manager, error) {
	mgr = Manager{}

	var err error
	mgr.gnmiClient = new(southbound.GNMIProvisioner)
	err = mgr.gnmiClient.Init(opts...)
	if err != nil {
		log.Error("Unable to setup GNMI provisioner", err)
		return nil, err
	}

	mgr.openapis = make(map[string]interface{})
	rbacAPIImpl := &rbac_1_0_0.ServerImpl{
		GnmiClient: mgr.gnmiClient,
	}
	mgr.openapis["Rbac-1.0.0"] = rbacAPIImpl
	aetherAPIImpl := &aether_1_0_0.ServerImpl{
		GnmiClient: mgr.gnmiClient,
	}
	mgr.openapis["Aether-1.0.0"] = aetherAPIImpl

	mgr.echoRouter = echo.New()
	rbac_1_0_0.RegisterHandlers(mgr.echoRouter, rbacAPIImpl)
	aether_1_0_0.RegisterHandlers(mgr.echoRouter, aetherAPIImpl)

	return &mgr, nil
}

// Run starts the northbound services.
func (m *Manager) Run() {
	log.Warn("Starting Manager")

	m.echoRouter.Logger.Fatal(m.echoRouter.Start(":8181"))

	log.Warn("Manager Stopping")
}
