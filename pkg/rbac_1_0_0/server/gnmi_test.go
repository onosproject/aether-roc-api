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

func Test_gnmiGetRbacV100targetRbac(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	rbacFromGnmi, err := ioutil.ReadFile("../testdata/RbacFromGnmi.json")
	assert.NilError(t, err, "error loadign testdata file")
	mockClient := southbound.NewMockGnmiClient(ctrl)
	mockClient.EXPECT().Get(gomock.Any(), gomock.Any()).DoAndReturn(
		func(ctx context.Context, request *gnmi.GetRequest) (*gnmi.GetResponse, error) {
			gr := gnmi.GetResponse{
				Notification: []*gnmi.Notification{
					{
						Update: []*gnmi.Update{
							{
								Val: &gnmi.TypedValue{
									Value: &gnmi.TypedValue_JsonVal{JsonVal: rbacFromGnmi},
								},
							},
						},
					},
				},
			}
			return &gr, nil
		},
	)

	serverImpl := ServerImpl{GnmiClient: mockClient}
	rbacResource, err := serverImpl.gnmiGetRbac(
		context.Background(), "/rbac/v1.0.0/internal/rbac", "internal")

	assert.NilError(t, err, "unexpected error on GetRequest")
	assert.Assert(t, rbacResource != nil)
	assert.Equal(t, 2, len(*rbacResource.Role), "expected 2 roles")
	for _, r := range *rbacResource.Role {
		switch rID := *r.Roleid; rID {
		case "aether-ops":
			assert.Equal(t, "Aether Operations", *r.Description)
			assert.Assert(t, r.Permission != nil)
			assert.Assert(t, r.Permission.LeafListNoun != nil)
			assert.DeepEqual(t, []string{"/internal/*", "/internal/rbac/*"}, *r.Permission.LeafListNoun)
			assert.Assert(t, r.Permission.Operation != nil)
			assert.Equal(t, "READ", *r.Permission.Operation)
			assert.Assert(t, r.Permission.Type != nil)
			assert.Equal(t, "CONFIG", *r.Permission.Type)
		case "aether-admin":
			assert.Equal(t, "Aether Admin", *r.Description)
			assert.Assert(t, r.Permission != nil)
			assert.Assert(t, r.Permission.LeafListNoun != nil)
			assert.DeepEqual(t, []string{"/internal/aether/*", "/internal/rbac/*"}, *r.Permission.LeafListNoun)
			assert.Assert(t, r.Permission.Operation != nil)
			assert.Equal(t, "ALL", *r.Permission.Operation)
			assert.Assert(t, r.Permission.Type != nil)
			assert.Equal(t, "CONFIG", *r.Permission.Type)
		default:
			t.Errorf("Unhandled %s", rID)
		}
	}

	assert.Equal(t, 1, len(*rbacResource.Group), "expected 1 group")
	group1 := (*rbacResource.Group)[0]
	assert.Assert(t, group1.Groupid != nil)
	assert.Equal(t, "menlo-admins", *group1.Groupid)
	assert.Equal(t, 2, len(*group1.Role), "expecting 2")
	for _, g := range *group1.Role {
		switch gID := *g.Roleid; gID {
		case "aether-admin":
			assert.Equal(t, "As role Aether Admins", *g.Description)
		case "aether-ops":
			assert.Equal(t, "As role Aether Ops", *g.Description)
		default:
			t.Errorf("Unhandled %s", gID)
		}
	}
}
