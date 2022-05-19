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

func Test_GnmiGetAetherV210targetSite(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apFromGnmi, err := ioutil.ReadFile("../testdata/ConfigFromGnmi.json")
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

	sitesList, err := serverImpl.GnmiGetSiteList(
		context.Background(), "/aether/v2.1.0/internal/site", "internal")
	assert.NilError(t, err)
	assert.Equal(t, 1, len(*sitesList))
	site0 := (*sitesList)[0]
	assert.Assert(t, site0.Description != nil)
	assert.Equal(t, "Global Default Site", *site0.Description)
	assert.Equal(t, "001", site0.ImsiDefinition.Mnc)

	appsList, err := serverImpl.GnmiGetApplicationList(context.Background(),
		"/aether/v2.1.0/internal/application", "internal")
	assert.NilError(t, err)
	assert.Equal(t, 1, len(*appsList))
	app0 := (*appsList)[0]
	assert.Assert(t, app0.Description != nil)
	assert.Equal(t, "Network Video Recorder", *app0.Description)

	assert.Assert(t, app0.Endpoint != nil)
	endPoint := app0.Endpoint
	assert.Equal(t, 1, len(*endPoint))
	ep0 := (*endPoint)[0]
	assert.Assert(t, ep0.PortEnd != nil)
	assert.Equal(t, 3330, *ep0.PortEnd)  // Optional
	assert.Equal(t, 3316, ep0.PortStart) // Optional

	tcList, err := serverImpl.GnmiGetTrafficClassList(context.Background(),
		"/aether/v2.1.0/internal/traffic-class", "internal")
	assert.NilError(t, err)
	assert.Equal(t, 1, len(*tcList))
	tc0 := (*tcList)[0]
	assert.Assert(t, tc0.Description != nil)
	assert.Equal(t, "High Priority TC", *tc0.Description)
	assert.Equal(t, "Class 1", *tc0.DisplayName)
	assert.Equal(t, 10, *tc0.Qci)
	assert.Equal(t, 11, *tc0.Arp)

}

func Test_gnmiGetAetherV210targetSmallCellSingle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apFromGnmi, err := ioutil.ReadFile("../testdata/ConfigFromGnmi.json")
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

	site1Sc1Resource, err := serverImpl.GnmiGetSiteSmallCell(
		context.Background(), "/aether/v2.0.0/internal/site/small-cell",
		"internal", "defaultent-defaultsite", "sc1")
	assert.NilError(t, err, "unexpected error on GetRequest")
	assert.Assert(t, site1Sc1Resource != nil)
	assert.Equal(t, "sc1", string(site1Sc1Resource.SmallCellId))
}

func Test_gnmiGetAetherV210targetSmallCellMultiple(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apFromGnmi, err := ioutil.ReadFile("../testdata/ConfigFromGnmi.json")
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

	site1ScellsResource, err := serverImpl.GnmiGetSiteSmallCellList(
		context.Background(), "/aether/v2.0.0/internal/enterprises/enterprise/small-cell",
		"internal", "defaultent-defaultsite")
	assert.NilError(t, err, "unexpected error on GetRequest")
	assert.Assert(t, site1ScellsResource != nil)
	assert.Equal(t, 2, len(*site1ScellsResource))
	for _, sc := range *site1ScellsResource {
		assert.Assert(t, sc.Description == nil)
		switch sc.SmallCellId {
		case "sc1":
			assert.Equal(t, "sc1.onf.org", *sc.Address)
			assert.Equal(t, "tac-1", sc.Tac)
			assert.Assert(t, *sc.Enable)
		case "sc2":
			assert.Equal(t, "sc2.onf.org", *sc.Address)
			assert.Equal(t, "tac-2", sc.Tac)
			assert.Assert(t, *sc.Enable)
		default:
			t.Errorf("error - unexpected SC id %s", sc.SmallCellId)
		}
	}
}
