// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package southbound

import (
	"context"
	"github.com/onosproject/onos-lib-go/pkg/grpc/retry"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/openconfig/gnmi/proto/gnmi"
	"google.golang.org/grpc"
	"time"
)

var log = logging.GetLogger("southbound")

// GNMIProvisioner handles provisioning of device configuration via gNMI interface.
type GNMIProvisioner struct {
	gnmi gnmi.GNMIClient
}

// Init initializes the gNMI provisioner
func (p *GNMIProvisioner) Init(gnmiEndpoint string, opts ...grpc.DialOption) error {
	optsWithRetry := []grpc.DialOption{
		grpc.WithStreamInterceptor(retry.RetryingStreamClientInterceptor(retry.WithInterval(100 * time.Millisecond))),
	}
	optsWithRetry = append(opts, optsWithRetry...)
	gnmiConn, err := grpc.Dial(gnmiEndpoint, optsWithRetry...)
	if err != nil {
		log.Error("Unable to connect to onos-config", err)
		return err
	}
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
