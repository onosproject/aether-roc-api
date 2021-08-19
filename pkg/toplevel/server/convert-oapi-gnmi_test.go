// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package server

import (
	"encoding/json"
	types2 "github.com/onosproject/aether-roc-api/pkg/aether_3_0_0/types"
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

	updates, deletes, ext100Name, ext101Version, ext102Type, defaultTarget, err :=
		encodeToGnmiPatchBody(jsonObj)
	assert.NilError(t, err)
	assert.Assert(t, ext101Version != nil)
	assert.Assert(t, ext102Type != nil)
	if ext100Name != nil {
		assert.Equal(t, "test-name", *ext100Name)
	}
	if ext101Version != nil {
		assert.Equal(t, "2.1.0", *ext101Version)
	}
	if ext102Type != nil {
		assert.Equal(t, "Aether", *ext102Type)
	}
	assert.Equal(t, "connectivity-service-v2", defaultTarget)
	assert.Equal(t, 20, len(updates))
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
		case `elem:{name:"site"} elem:{name:"site" key:{key:"id" value:"starbucks-newyork"}} elem:{name:"id"} target:"connectivity-service-v3"`:
			assert.Equal(t, `string_val:"starbucks-newyork"`, upd.Val.String())
		case `elem:{name:"site"} elem:{name:"site" key:{key:"id" value:"starbucks-newyork"}} elem:{name:"imsi-definition"} elem:{name:"enterprise"} target:"connectivity-service-v3"`:
			assert.Equal(t, `uint_val:223`, upd.Val.String())
		default:
			t.Fatalf("unexpected path %s", path)
		}
	}

	assert.Equal(t, 6, len(deletes))
	for _, del := range deletes {
		switch path := strings.ReplaceAll(del.String(), "  ", " "); path {
		case `elem:{name:"access-profile"} elem:{name:"access-profile" key:{key:"id" value:"ap3d"}} elem:{name:"id"} target:"connectivity-service-v2"`:
		case `elem:{name:"vcs"} elem:{name:"vcs" key:{key:"id" value:"vcs-to-delete-from"}} elem:{name:"application" key:{key:"application" value:"application-to-delete-the-allow-from"}} elem:{name:"allow"} target:"connectivity-service-v3"`:
		case `elem:{name:"vcs"} elem:{name:"vcs" key:{key:"id" value:"vcs-to-delete-from-2"}} elem:{name:"application" key:{key:"application" value:"application-to-delete"}} elem:{name:"application"} target:"connectivity-service-v3"`:
		case `elem:{name:"vcs"} elem:{name:"vcs" key:{key:"id" value:"vcs-to-delete-from-3"}} elem:{name:"device-group" key:{key:"device-group" value:"device-group-to-delete-the-enable-from"}} elem:{name:"enable"} target:"connectivity-service-v3"`:
		case `elem:{name:"vcs"} elem:{name:"vcs" key:{key:"id" value:"vcs-to-delete-from-4"}} elem:{name:"device-group" key:{key:"device-group" value:"device-group-to-delete"}} elem:{name:"device-group"} target:"connectivity-service-v3"`:
		case `elem:{name:"vcs"} elem:{name:"vcs" key:{key:"id" value:"vcs-to-delete-from-5"}} elem:{name:"description"} target:"connectivity-service-v3"`:
		default:
			t.Fatalf("unexpected path %s", path)
		}
	}
}

func Test_addProps(t *testing.T) {
	desc1 := "desc1"
	disp1 := "display 1"
	mnc := int32(456)
	id1 := "id1"
	enterprise := int32(789)
	unchangedSite := "enterprise"
	unchangedImsi := "mcc"
	addPropsSite := make(map[string]types2.AdditionalPropertyUnchanged)
	addPropsSite["additional-properties"] = types2.AdditionalPropertyUnchanged{Unchanged: &unchangedSite}
	addPropsImsi := make(map[string]types2.AdditionalPropertyUnchanged)
	addPropsImsi["additional-properties"] = types2.AdditionalPropertyUnchanged{Unchanged: &unchangedImsi}

	ap1 := types2.SiteSite{
		Description: &desc1,
		DisplayName: &disp1,
		Id:          id1,
		ImsiDefinition: &types2.SiteSiteImsiDefinition{
			Enterprise:           enterprise,
			Format:               "CCCNNNEEESSSSSS",
			Mnc:                  mnc,
			AdditionalProperties: addPropsImsi,
		},
		AdditionalProperties: addPropsSite,
	}

	bytes, err := json.Marshal(ap1)
	assert.NilError(t, err)
	assert.Equal(t,
		`{"additional-properties":{"unchanged":"enterprise"},"description":"desc1","display-name":"display 1","enterprise":"","id":"id1","imsi-definition":{"additional-properties":{"unchanged":"mcc"},"enterprise":789,"format":"CCCNNNEEESSSSSS","mcc":0,"mnc":456}}`,
		string(bytes))
}
