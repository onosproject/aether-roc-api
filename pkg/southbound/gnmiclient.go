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
type GnmiClient interface {
	Init(opts ...grpc.DialOption) error
	Get(ctx context.Context, request *gnmi.GetRequest) (*gnmi.GetResponse, error)
}
