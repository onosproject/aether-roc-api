// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package server

import (
	"context"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/onosproject/aether-roc-api/pkg/aether_1_0_0/types"
	"github.com/onosproject/aether-roc-api/pkg/southbound"
	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/gnmi/proto/gnmi_ext"
	"gotest.tools/assert"
	"io/ioutil"
	"testing"
)

func TestServerImpl_PostAetherV100targetAccessProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	aetherFromGnmi, err := ioutil.ReadFile("../testdata/sample-create-access-profile.json")
	assert.NilError(t, err, "error loading testdata file")
	mockClient := southbound.NewMockGnmiClient(ctrl)
	mockClient.EXPECT().Set(gomock.Any(), gomock.Any()).DoAndReturn(
		func(ctx context.Context, request *gnmi.SetRequest) (*gnmi.SetResponse, error) {
			sr := gnmi.SetResponse{
				Response: []*gnmi.UpdateResult{
					{
						Path: &gnmi.Path{
							Elem: []*gnmi.PathElem{
								{
									Name: "access-profiles",
								},
							},
							Target: "spgw-1",
						},
					},
				},
				Extension: []*gnmi_ext.Extension{
					{
						Ext: &gnmi_ext.Extension_RegisteredExt{
							RegisteredExt: &gnmi_ext.RegisteredExtension{
								Id:  100,
								Msg: []byte("test-commit"),
							},
						},
					},
				},
			}
			return &sr, nil
		},
	)
	serverImpl := ServerImpl{GnmiClient: mockClient}
	gnmiChangeName, err := serverImpl.gnmiPostAetherV100targetAccessProfile(
		context.Background(), aetherFromGnmi, "/rbac/v1.0.0/internal/aether-profiles",
		"internal")
	assert.NilError(t, err, "error creating access profile")
	assert.Equal(t, "test-commit", *gnmiChangeName, "unexpected value for gNMI Set response")

}

func Test_jsonMarshallGnmiObj(t *testing.T) {
	aetherFromGnmi, err := ioutil.ReadFile("../testdata/sample-create-access-profile.json")
	assert.NilError(t, err, "error loading testdata file")

	const expectedJSON = `{"AccessProfile":{"profile1":{"Description":"test profile 1","Filter":null,"Id":"profile1","Type":null},"profile2":{"Description":"test profile 2","Filter":null,"Id":"profile2","Type":null}}}`

	jsonObj := new(types.AetherV100targetAccessProfile)
	err = json.Unmarshal(aetherFromGnmi, jsonObj)
	assert.NilError(t, err, "error converting json to Obj")

	gnmiObj, err := encodeToGnmiAetherV100targetAccessProfile(jsonObj)
	assert.NilError(t, err, "error converting jsonObj to gnmiObj")

	gnmiJSONPayload, err := json.Marshal(gnmiObj)
	assert.NilError(t, err, "error converting gnmi to Json")
	assert.Equal(t, expectedJSON, string(gnmiJSONPayload))
}
