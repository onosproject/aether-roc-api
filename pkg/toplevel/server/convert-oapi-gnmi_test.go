// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package server

import (
	"encoding/json"
	types2 "github.com/onosproject/aether-roc-api/pkg/aether_2_0_0/types"
	"github.com/onosproject/aether-roc-api/pkg/toplevel/types"
	"gotest.tools/assert"
	"io/ioutil"
	"strings"
	"testing"
)

func Test_encodeToGnmiPatchBody(t *testing.T) {
	patchBodyExampleJSON, err := ioutil.ReadFile("../testdata/PatchBody_Example.json")
	assert.NilError(t, err, "error loading testdata file")
	jsonObj := new(types.PatchBody)
	err = json.Unmarshal(patchBodyExampleJSON, jsonObj)
	assert.NilError(t, err)

	testMap := make(map[string]interface{})
	err = json.Unmarshal(patchBodyExampleJSON, &testMap)
	assert.NilError(t, err)

	updates, deletes, ext101Version, ext102Type, defaultTarget, err :=
		encodeToGnmiPatchBody(jsonObj)
	assert.NilError(t, err)
	assert.Assert(t, ext101Version != nil)
	assert.Assert(t, ext102Type != nil)
	if ext101Version != nil {
		assert.Equal(t, "2.0.0", *ext101Version)
	}
	if ext102Type != nil {
		assert.Equal(t, "Aether", *ext102Type)
	}
	assert.Equal(t, "connectivity-service-v2", defaultTarget)
	assert.Equal(t, 18, len(updates))
	for _, upd := range updates {
		switch path := strings.ReplaceAll(upd.Path.String(), "  ", " "); path {
		case `elem:{name:"access-profile"} elem:{name:"access-profile" key:{key:"id" value:"ap1"}} elem:{name:"id"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"ap1"`, upd.Val.String())
		case `elem:{name:"access-profile"} elem:{name:"access-profile" key:{key:"id" value:"ap1"}} elem:{name:"description"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"Sample access profile"`, upd.Val.String())
		case `elem:{name:"access-profile"} elem:{name:"access-profile" key:{key:"id" value:"ap1"}} elem:{name:"display-name"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"Sample display name"`, upd.Val.String())
		case `elem:{name:"access-profile"} elem:{name:"access-profile" key:{key:"id" value:"ap1"}} elem:{name:"type"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"Sample ap type"`, upd.Val.String())
		case `elem:{name:"access-profile"} elem:{name:"access-profile" key:{key:"id" value:"ap1"}} elem:{name:"filter"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"Sample ap filter"`, upd.Val.String())
		case `elem:{name:"access-profile"} elem:{name:"access-profile" key:{key:"id" value:"ap2"}} elem:{name:"id"} target:"test1"`:
			assert.Equal(t, `string_val:"ap2"`, upd.Val.String())
		case `elem:{name:"access-profile"} elem:{name:"access-profile" key:{key:"id" value:"ap2"}} elem:{name:"description"} target:"test1"`:
			assert.Equal(t, `string_val:"2nd Sample access profile"`, upd.Val.String())
		case `elem:{name:"access-profile"} elem:{name:"access-profile" key:{key:"id" value:"ap2"}} elem:{name:"display-name"} target:"test1"`:
			assert.Equal(t, `string_val:"2nd Sample display name"`, upd.Val.String())
		case `elem:{name:"access-profile"} elem:{name:"access-profile" key:{key:"id" value:"ap2"}} elem:{name:"type"} target:"test1"`:
			assert.Equal(t, `string_val:"2nd Sample ap type"`, upd.Val.String())
		case `elem:{name:"access-profile"} elem:{name:"access-profile" key:{key:"id" value:"ap2"}} elem:{name:"filter"} target:"test1"`:
			assert.Equal(t, `string_val:"2nd Sample ap filter"`, upd.Val.String())
		case `elem:{name:"apn-profile"} elem:{name:"apn-profile" key:{key:"id" value:"apn1"}} elem:{name:"id"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"apn1"`, upd.Val.String())
		case `elem:{name:"apn-profile"} elem:{name:"apn-profile" key:{key:"id" value:"apn1"}} elem:{name:"description"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"apn1 Description"`, upd.Val.String())
		case `elem:{name:"apn-profile"} elem:{name:"apn-profile" key:{key:"id" value:"apn1"}} elem:{name:"display-name"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"apn1 display name"`, upd.Val.String())
		case `elem:{name:"apn-profile"} elem:{name:"apn-profile" key:{key:"id" value:"apn1"}} elem:{name:"apn-name"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"apn1 name"`, upd.Val.String())
		case `elem:{name:"apn-profile"} elem:{name:"apn-profile" key:{key:"id" value:"apn1"}} elem:{name:"dns-primary"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"10.1.2.3"`, upd.Val.String())
		case `elem:{name:"apn-profile"} elem:{name:"apn-profile" key:{key:"id" value:"apn1"}} elem:{name:"dns-secondary"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"10.1.2.4"`, upd.Val.String())
		case `elem:{name:"apn-profile"} elem:{name:"apn-profile" key:{key:"id" value:"apn1"}} elem:{name:"gx-enabled"} target:"connectivity-service-v2"`:
			assert.Equal(t, `bool_val:false`, upd.Val.String())
		case `elem:{name:"apn-profile"} elem:{name:"apn-profile" key:{key:"id" value:"apn1"}} elem:{name:"mtu"} target:"connectivity-service-v2"`:
			assert.Equal(t, `uint_val:9600`, upd.Val.String())
		default:
			t.Fatalf("unexpected path %s", path)
		}
	}

	assert.Equal(t, 1, len(deletes))
	for _, del := range deletes {
		switch path := strings.ReplaceAll(del.String(), "  ", " "); path {
		case `elem:{name:"access-profile"} elem:{name:"access-profile" key:{key:"id" value:"ap3d"}} elem:{name:"id"} target:"connectivity-service-v2"`:
		default:
			t.Fatalf("unexpected path %s", path)
		}
	}

}

func Test_addProps(t *testing.T) {
	desc1 := "desc1"
	disp1 := "display 1"
	filter1 := "filter 1"
	id1 := "id1"
	type1 := "type1"
	target1 := "target1"
	addProps := make(map[string]types2.AdditionalPropertyTarget)

	addProps["additional-properties"] = types2.AdditionalPropertyTarget{Target: &target1}

	ap1 := types2.AccessProfileAccessProfile{
		Description:          &desc1,
		DisplayName:          &disp1,
		Filter:               &filter1,
		Id:                   &id1,
		Type:                 &type1,
		AdditionalProperties: addProps,
	}

	bytes, err := json.Marshal(ap1)
	assert.NilError(t, err)
	assert.Equal(t,
		`{"additional-properties":{"target":"target1"},"description":"desc1","display-name":"display 1","filter":"filter 1","id":"id1","type":"type1"}`,
		string(bytes))
}
