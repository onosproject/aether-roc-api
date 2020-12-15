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

func Test_gnmiGetAetherV200targetSubscriber(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apFromGnmi, err := ioutil.ReadFile("../testdata/AccessProfileFromGnmi.json")
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
	apResource, err := serverImpl.gnmiGetAccessProfile(
		context.Background(), "/aether/v2.0.0/internal/access-profile", "internal")
	assert.NilError(t, err, "unexpected error on GetRequest")
	assert.Assert(t, apResource != nil)

	assert.Equal(t, 2, len(*apResource.AccessProfile))
	for _, ap := range *apResource.AccessProfile {
		switch *ap.Id {
		case "ap1":
			assert.Equal(t, "Sample ap type", *ap.Type)
		case "ap2":
			assert.Equal(t, "2nd Sample ap type", *ap.Type)
		default:
			t.Errorf("unexpected AP ID %v", ap.Id)
		}
	}

	apnResource, err := serverImpl.gnmiGetApnProfile(
		context.Background(), "/aether/v2.0.0/internal/apn-profile", "internal")
	assert.NilError(t, err, "unexpected error on GetRequest")
	assert.Assert(t, apnResource != nil)

	assert.Equal(t, 1, len(*apnResource.ApnProfile))
	for _, apn := range *apnResource.ApnProfile {
		switch *apn.Id {
		case "apn1":
			assert.Equal(t, "ap1 display name", *apn.DisplayName)
		default:
			t.Errorf("unexpected AP ID %v", *apn.Id)
		}
	}
}
