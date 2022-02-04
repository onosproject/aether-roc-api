// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
//

package manager

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	aether_2_0_0 "github.com/onosproject/aether-roc-api/pkg/aether_2_0_0/server"
	aether_4_0_0 "github.com/onosproject/aether-roc-api/pkg/aether_4_0_0/server"
	"github.com/onosproject/aether-roc-api/pkg/southbound"
	toplevel "github.com/onosproject/aether-roc-api/pkg/toplevel/server"
	"github.com/onosproject/onos-api/go/onos/config/diags"
	"github.com/onosproject/onos-lib-go/pkg/grpc/retry"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"google.golang.org/grpc"
	"time"
)

var mgr Manager

var log = logging.GetLogger("manager")

// Manager single point of entry for the ROC system.
type Manager struct {
	echoRouter    *echo.Echo
	openapis      map[string]interface{}
	authorization bool
}

// NewManager -
func NewManager(gnmiEndpoint string, allowCorsOrigins []string,
	validateResponses bool, authorization bool, opts ...grpc.DialOption) (*Manager, error) {
	mgr = Manager{authorization: authorization}
	optsWithRetry := []grpc.DialOption{
		grpc.WithStreamInterceptor(retry.RetryingStreamClientInterceptor(retry.WithInterval(100 * time.Millisecond))),
	}
	optsWithRetry = append(opts, optsWithRetry...)
	gnmiConn, err := grpc.Dial(gnmiEndpoint, optsWithRetry...)
	if err != nil {
		log.Error("Unable to connect to onos-config", err)
		return nil, err
	}

	gnmiClient := new(southbound.GNMIProvisioner)
	err = gnmiClient.Init(gnmiConn)
	if err != nil {
		log.Error("Unable to setup GNMI provisioner", err)
		return nil, err
	}

	transactionServiceClient := diags.NewChangeServiceClient(gnmiConn)

	mgr.openapis = make(map[string]interface{})
	aether20APIImpl := &aether_2_0_0.ServerImpl{
		GnmiClient: gnmiClient,
	}
	mgr.openapis["Aether-2.0.0"] = aether20APIImpl
	aether40APIImpl := &aether_4_0_0.ServerImpl{
		GnmiClient: gnmiClient,
	}
	mgr.openapis["Aether-4.0.0"] = aether40APIImpl
	topLevelAPIImpl := &toplevel.ServerImpl{
		GnmiClient:    gnmiClient,
		ConfigClient:  transactionServiceClient,
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
	if err := aether_2_0_0.RegisterHandlers(mgr.echoRouter, aether20APIImpl, validateResponses); err != nil {
		return nil, fmt.Errorf("aether_2_0_0.RegisterHandlers()  %s", err)
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
