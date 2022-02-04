// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
//

package southbound

import (
	"context"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/openconfig/gnmi/proto/gnmi"
	"google.golang.org/grpc"
)

var log = logging.GetLogger("southbound")

// GNMIProvisioner handles provisioning of device configuration via gNMI interface.
type GNMIProvisioner struct {
	gnmi gnmi.GNMIClient
}

// Init initializes the gNMI provisioner
func (p *GNMIProvisioner) Init(gnmiConn *grpc.ClientConn) error {
	log.Infof("Initializing new GnmiProvisioner to %s", gnmiConn.Target())
	p.gnmi = gnmi.NewGNMIClient(gnmiConn)
	return nil
}

// Get passes a gNMI GetRequest to the server which synchronously replies with a GetResponse
func (p *GNMIProvisioner) Get(ctx context.Context, request *gnmi.GetRequest) (*gnmi.GetResponse, error) {
	return p.gnmi.Get(ctx, request)
}

// Set passes a gNMI SetRequest to the server which synchronously replies with a SetResponse
func (p *GNMIProvisioner) Set(ctx context.Context, request *gnmi.SetRequest) (*gnmi.SetResponse, error) {
	return p.gnmi.Set(ctx, request)
}
