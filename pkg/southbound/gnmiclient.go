// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package southbound

import (
	"context"
	"github.com/openconfig/gnmi/proto/gnmi"
	"google.golang.org/grpc"
)

// GnmiClient - a way of making gNMI calls
// If the interface is changed, generate new mocks with:
// mockgen -package southbound -source pkg/southbound/gnmiclient.go -mock_names GnmiClient=MockGnmiClient > /tmp/gnmiclient_mock.go
// mv /tmp/gnmiclient_mock.go pkg/southbound
type GnmiClient interface {
	Init(gnmiConn *grpc.ClientConn) error
	Get(ctx context.Context, request *gnmi.GetRequest) (*gnmi.GetResponse, error)
	Set(ctx context.Context, request *gnmi.SetRequest) (*gnmi.SetResponse, error)
}
