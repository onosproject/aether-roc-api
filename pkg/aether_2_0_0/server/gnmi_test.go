// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
//

package server

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/onosproject/aether-roc-api/pkg/southbound"
	"github.com/openconfig/gnmi/proto/gnmi"
	"gotest.tools/assert"
	"io/ioutil"
	"testing"
)

func Test_gnmiGetAetherV200targetConnSvc(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apFromGnmi, err := ioutil.ReadFile("../testdata/ConnectivityServiceFromGnmi.json")
	assert.NilError(t, err, "error loading testdata file")
	mockClient := southbound.NewMockGnmiClient(ctrl)
	mockClient.EXPECT().Get(gomock.Any(), gomock.Any()).DoAndReturn(
		func(ctx context.Context, request *gnmi.GetRequest) (*gnmi.GetResponse, error) {
			gr := gnmi.GetResponse{
				Notification: []*gnmi.Notification{
					{
						Update: []*gnmi.Update{
							{
								Val: &gnmi.TypedValue{
									Value: &gnmi.TypedValue_JsonVal{JsonVal: apFromGnmi},
								},
							},
						},
					},
				},
			}
			return &gr, nil
		},
	).AnyTimes()
	serverImpl := ServerImpl{GnmiClient: mockClient}
	apResource, err := serverImpl.gnmiGetConnectivityServices(
		context.Background(), "/aether/v2.0.0/internal/connectivity-services", "internal")
	assert.NilError(t, err, "unexpected error on GetRequest")
	assert.Assert(t, apResource != nil)

	entResource, err := serverImpl.gnmiGetEnterprises(
		context.Background(), "/aether/v2.0.0/internal/enterprises", "internal")
	assert.NilError(t, err, "unexpected error on GetRequest")
	assert.Assert(t, entResource != nil)
	entContainer := *entResource.Enterprise
	assert.Equal(t, 1, len(entContainer))
	assert.Assert(t, entContainer[0].Description != nil)
	assert.Equal(t, "Test Enterprise", *entContainer[0].Description)
	siteContainer := *entContainer[0].Site
	assert.Equal(t, 1, len(siteContainer))
	assert.Assert(t, siteContainer[0].Description != nil)
	assert.Equal(t, "Global Default Site", *siteContainer[0].Description)

	assert.Equal(t, "001", siteContainer[0].ImsiDefinition.Mnc)

	appContainer := *entContainer[0].Application
	assert.Equal(t, 1, len(appContainer))
	assert.Assert(t, appContainer[0].Description != nil)
	assert.Equal(t, "Network Video Recorder", *appContainer[0].Description)

	assert.Assert(t, appContainer[0].Endpoint != nil)
	endPoint := *appContainer[0].Endpoint
	assert.Equal(t, 1, len(endPoint))
	assert.Assert(t, endPoint[0].PortEnd != nil)
	assert.Equal(t, 3330, *endPoint[0].PortEnd)   // Optional
	assert.Equal(t, 3316, *endPoint[0].PortStart) // Optional

	tcContainer := *entContainer[0].TrafficClass
	assert.Equal(t, 1, len(tcContainer))
	assert.Assert(t, tcContainer[0].Description != nil)
	assert.Equal(t, "High Priority TC", *tcContainer[0].Description)
	assert.Equal(t, "Class 1", *tcContainer[0].DisplayName)
	assert.Equal(t, 10, *tcContainer[0].Qci)
	assert.Equal(t, 11, *tcContainer[0].Arp)

}
