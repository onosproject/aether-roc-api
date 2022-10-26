// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
//

package server

import (
	"encoding/json"
	"fmt"
	types2 "github.com/onosproject/aether-roc-api/pkg/aether_2_1_0/types"
	"github.com/onosproject/aether-roc-api/pkg/toplevel/types"
	"gotest.tools/assert"
	"os"
	"strings"
	"testing"
)

func Test_encodeToGnmiPatchBody(t *testing.T) {
	patchBodyExampleJSON, err := os.ReadFile("../testdata/PatchBody_Example.json")
	assert.NilError(t, err, "error loading testdata file")
	jsonObj := new(types.PatchBody)
	err = json.Unmarshal(patchBodyExampleJSON, jsonObj)
	assert.NilError(t, err)

	testMap := make(map[string]interface{})
	err = json.Unmarshal(patchBodyExampleJSON, &testMap)
	assert.NilError(t, err)

	pb, err :=
		encodeToGnmiPatchBody(jsonObj)
	assert.NilError(t, err)
	assert.Assert(t, pb.Ext101Version != nil)
	assert.Assert(t, pb.Ext102Type != nil)
	if pb.Ext100Name != nil {
		assert.Equal(t, "test-name", *pb.Ext100Name)
	}
	if pb.Ext101Version != nil {
		assert.Equal(t, "2.1.x", *pb.Ext101Version)
	}
	if pb.Ext102Type != nil {
		assert.Equal(t, "Aether", *pb.Ext102Type)
	}
	assert.Equal(t, "defaultent", pb.DefaultTarget)
	assert.Equal(t, 427, len(pb.Updates))
	for _, upd := range pb.Updates {
		switch tgt := upd.Path.Target; tgt {
		case `defaultent`:
		case `starbucks`:
		case "acme":
			switch path0 := strings.ReplaceAll(upd.Path.Elem[0].String(), "  ", " "); path0 {
			case `name:"application" key:{key:"application-id" value:"acme-dataacquisition"}`:
				switch path1 := strings.ReplaceAll(upd.Path.Elem[1].String(), "  ", " "); path1 {
				case `name:"address"`:
					assert.Equal(t, `string_val:"da.acme.com"`, upd.Val.String())
				case `name:"application-id"`:
					assert.Equal(t, `string_val:"acme-dataacquisition"`, upd.Val.String())
				case `name:"description"`:
					assert.Equal(t, `string_val:"Data Acquisition"`, upd.Val.String())
				case `name:"display-name"`:
					assert.Equal(t, `string_val:"DA"`, upd.Val.String())
				case `name:"endpoint" key:{key:"endpoint-id" value:"da"}`:
					switch path2 := strings.ReplaceAll(upd.Path.Elem[2].String(), "  ", " "); path2 {
					case `name:"endpoint-id"`:
						assert.Equal(t, `string_val:"da"`, upd.Val.String())
					case `name:"display-name"`:
						assert.Equal(t, `string_val:"data acquisition endpoint"`, upd.Val.String())
					case `name:"protocol"`:
						assert.Equal(t, `string_val:"TCP"`, upd.Val.String())
					case `name:"traffic-class"`:
						assert.Equal(t, `string_val:"class-2"`, upd.Val.String())
					case `name:"port-start"`:
						assert.Equal(t, `uint_val:7585`, upd.Val.String())
					case `name:"port-end"`:
						assert.Equal(t, `uint_val:7588`, upd.Val.String())
					case `name:"mbr"`:
						switch path3 := strings.ReplaceAll(upd.Path.Elem[3].String(), "  ", " "); path3 {
						case `name:"uplink"`:
							assert.Equal(t, `uint_val:2000000`, upd.Val.String())
						case `name:"downlink"`:
							assert.Equal(t, `uint_val:1000000`, upd.Val.String())
						default:
							t.Logf("unhandled mbr update %s", path2)
						}
					default:
						t.Logf("unhandled da update %s", path2)
					}
				default:
					t.Logf("unhandled acme-dataacquisition update %s", path1)
				}
			case `name:"site" key:{key:"site-id" value:"acme-chicago"}`:
			case `name:"template" key:{key:"template-id" value:"template-1"}`:
			case `name:"template" key:{key:"template-id" value:"template-2"}`:
			case `name:"traffic-class" key:{key:"traffic-class-id" value:"class-1"}`:
			case `name:"traffic-class" key:{key:"traffic-class-id" value:"class-2"}`:
			case `name:"traffic-class" key:{key:"traffic-class-id" value:"class-3"}`:
			default:
				t.Logf("unhandled acme update %s", path0)
			}
		case "connectivity-service-v2":
			switch path0 := upd.Path.Elem[0].Name; path0 {
			case "connectivity-services":
				switch path1 := strings.ReplaceAll(upd.Path.Elem[1].String(), "  ", " "); path1 {
				case `name:"connectivity-service" key:{key:"connectivity-service-id" value:"cs5gtest"}`:
					switch path2 := upd.Path.Elem[2].String(); path2 {
					case `name:"acc-prometheus-url"`:
						assert.Equal(t, `string_val:"./prometheus-acc"`, upd.Val.String())
					case `name:"core-5g-endpoint"`:
						assert.Equal(t, `string_val:"http://aether-roc-umbrella-sdcore-test-dummy/v1/config/5g"`, upd.Val.String())
					case `name:"description"`:
						assert.Equal(t, `string_val:"5G Test"`, upd.Val.String())
					case `name:"display-name"`:
						assert.Equal(t, `string_val:"ROC 5G Test Connectivity Service"`, upd.Val.String())
					case `name:"connectivity-service-id"`:
						assert.Equal(t, `string_val:"cs5gtest"`, upd.Val.String())
					default:
						t.Logf("unhandled v2 update %s %s %s", path0, path1, path2)
					}

				case `name:"connectivity-service" key:{key:"connectivity-service-id" value:"cs4gtest"}`:
					switch path2 := upd.Path.Elem[2].String(); path2 {
					case `name:"acc-prometheus-url"`:
						assert.Equal(t, `string_val:"./prometheus-acc"`, upd.Val.String())
					case `name:"core-5g-endpoint"`:
						assert.Equal(t, `string_val:"http://aether-roc-umbrella-sdcore-test-dummy/v1/config/5g"`, upd.Val.String())
					case `name:"description"`:
						assert.Equal(t, `string_val:"ROC 4G Test Connectivity Service"`, upd.Val.String())
					case `name:"display-name"`:
						assert.Equal(t, `string_val:"4G Test"`, upd.Val.String())
					case `name:"connectivity-service-id"`:
						assert.Equal(t, `string_val:"cs4gtest"`, upd.Val.String())
					default:
						t.Logf("unhandled v2 update %s %s %s", path0, path1, path2)
					}

				default:
					t.Logf("unhandled v2 update %s %s", path0, path1)
				}
			}
		default:
			t.Logf("unhandled target %s", tgt)
		}
	}

	assert.Equal(t, 4, len(pb.Deletes))
	for _, del := range pb.Deletes {
		switch path := strings.ReplaceAll(del.String(), "  ", " "); path {
		case `elem:{name:"application" key:{key:"application-id" value:"app-10"}} elem:{name:"endpoint" key:{key:"endpoint-id" value:"ep-10"}} elem:{name:"endpoint-id"} target:"acme"`:
		case `elem:{name:"application" key:{key:"application-id" value:"app-10"}} elem:{name:"endpoint" key:{key:"endpoint-id" value:"ep-11"}} elem:{name:"endpoint-id"} target:"acme"`:
		case `elem:{name:"traffic-class" key:{key:"traffic-class-id" value:"class-4"}} elem:{name:"traffic-class-id"} target:"acme"`:
		case `elem:{name:"traffic-class" key:{key:"traffic-class-id" value:"class-5"}} elem:{name:"traffic-class-id"} target:"acme"`:
		default:
			t.Fatalf("unexpected path %s", path)
		}
	}
}

func Test_addProps(t *testing.T) {
	desc1 := "desc1"
	disp1 := "display 1"
	mnc := "456"
	var id1 types2.ListKey = "id1"
	enterprise := int32(789)
	unchangedSite := "enterprise"
	unchangedImsi := "mcc"
	addPropsSite := make(map[string]types2.AdditionalPropertyUnchanged)
	addPropsSite["additional-properties"] = types2.AdditionalPropertyUnchanged{Unchanged: &unchangedSite}
	addPropsImsi := make(map[string]types2.AdditionalPropertyUnchanged)
	addPropsImsi["additional-properties"] = types2.AdditionalPropertyUnchanged{Unchanged: &unchangedImsi}

	ap1 := types2.Site{
		Description: &desc1,
		DisplayName: &disp1,
		SiteId:      id1,
		ImsiDefinition: &types2.SiteImsiDefinition{
			Enterprise:           enterprise,
			Format:               "CCCNNNEEESSSSSS",
			Mnc:                  mnc,
			AdditionalProperties: addPropsImsi,
		},
	}

	bytes, err := json.Marshal(ap1)
	assert.NilError(t, err)
	assert.Equal(t,
		`{"description":"desc1","display-name":"display 1","imsi-definition":{"additional-properties":{"unchanged":"mcc"},"enterprise":789,"format":"CCCNNNEEESSSSSS","mcc":"","mnc":"456"},"site-id":"id1"}`,
		string(bytes))
}

func TestEncodeToGnmiElements(t *testing.T) {

	tests := []struct {
		name  string
		args  *types.Elements
		error error
	}{
		{"valid-elements-with-sites", &types.Elements{
			Site210: &types2.SiteList{
				{SiteId: "site-1"},
				{SiteId: "site-2"},
			},
		}, nil},
		{"invalid-site", &types.Elements{
			Site210: &types2.SiteList{
				{SiteId: "site-1"},
				{SiteId: "site-2"},
				{SiteId: undefined},
			},
		}, fmt.Errorf("code=422, message=site-id-cannot-be-undefined")},
		{"invalid-slice", &types.Elements{
			Site210: &types2.SiteList{
				{
					SiteId: "site-1",
					Slice: &types2.SiteSliceList{
						{
							DefaultBehavior: "ALLOW-ALL",
							SliceId:         "slice-1",
						},
						{
							SliceId: undefined,
						},
					},
				},
			},
		}, fmt.Errorf("code=422, message=slice-id-cannot-be-undefined")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := encodeToGnmiElements(tt.args, "target", false)
			if tt.error != nil {
				assert.Error(t, err, tt.error.Error())
			} else {
				assert.NilError(t, err)
			}
		})
	}
}
