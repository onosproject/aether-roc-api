// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
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

func Test_gnmiGetAetherV300targetConnSvc(t *testing.T) {
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
	apResource, err := serverImpl.gnmiGetConnectivityService(
		context.Background(), "/aether/v3.0.0/internal/connectivity-service", "internal")
	assert.NilError(t, err, "unexpected error on GetRequest")
	assert.Assert(t, apResource != nil)

	siteResource, err := serverImpl.gnmiGetSite(
		context.Background(), "/aether/v3.0.0/internal/site", "internal")
	assert.NilError(t, err, "unexpected error on GetRequest")
	assert.Assert(t, siteResource != nil)
	siteContainer := *siteResource.Site
	assert.Equal(t, 1, len(siteContainer))
	assert.Assert(t, siteContainer[0].Description != nil)
	assert.Equal(t, "Global Default Site", *siteContainer[0].Description)
	assert.Equal(t, "001", siteContainer[0].ImsiDefinition.Mnc)

	appResource, err := serverImpl.gnmiGetApplication(
		context.Background(), "/aether/v3.0.0/internal/application", "internal")
	assert.NilError(t, err, "unexpected error on GetRequest")
	assert.Assert(t, appResource != nil)
	appContainer := *appResource.Application
	assert.Equal(t, 1, len(appContainer))
	assert.Assert(t, appContainer[0].Description != nil)
	assert.Equal(t, "Network Video Recorder", *appContainer[0].Description)

	assert.Assert(t, appContainer[0].Endpoint != nil)
	endPoint := *appContainer[0].Endpoint
	assert.Equal(t, 1, len(endPoint))
	assert.Assert(t, endPoint[0].PortEnd != nil)
	assert.Equal(t, 3330, *endPoint[0].PortEnd)  // Mandatory
	assert.Equal(t, 3316, endPoint[0].PortStart) // Optional

}
