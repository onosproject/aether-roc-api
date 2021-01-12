// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package server

import (
	"encoding/json"
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
	updates, deletes, err := encodeToGnmiPatchBody(jsonObj)
	assert.NilError(t, err)
	assert.Equal(t, 18, len(updates))
	for _, upd := range updates {
		switch path := strings.ReplaceAll(upd.Path.String(), "  ", " "); path {
		case `elem:{name:"access-profile"} elem:{name:"access-profile" key:{key:"id" value:"ap1"}} elem:{name:"id"}`:
			assert.Equal(t, `string_val:"ap1"`, upd.Val.String())
		case `elem:{name:"access-profile"} elem:{name:"access-profile" key:{key:"id" value:"ap1"}} elem:{name:"description"}`:
			assert.Equal(t, `string_val:"Sample access profile"`, upd.Val.String())
		case `elem:{name:"access-profile"} elem:{name:"access-profile" key:{key:"id" value:"ap1"}} elem:{name:"display-name"}`:
			assert.Equal(t, `string_val:"Sample display name"`, upd.Val.String())
		case `elem:{name:"access-profile"} elem:{name:"access-profile" key:{key:"id" value:"ap1"}} elem:{name:"type"}`:
			assert.Equal(t, `string_val:"Sample ap type"`, upd.Val.String())
		case `elem:{name:"access-profile"} elem:{name:"access-profile" key:{key:"id" value:"ap1"}} elem:{name:"filter"}`:
			assert.Equal(t, `string_val:"Sample ap filter"`, upd.Val.String())
		case `elem:{name:"access-profile"} elem:{name:"access-profile" key:{key:"id" value:"ap2"}} elem:{name:"id"}`:
			assert.Equal(t, `string_val:"ap2"`, upd.Val.String())
		case `elem:{name:"access-profile"} elem:{name:"access-profile" key:{key:"id" value:"ap2"}} elem:{name:"description"}`:
			assert.Equal(t, `string_val:"2nd Sample access profile"`, upd.Val.String())
		case `elem:{name:"access-profile"} elem:{name:"access-profile" key:{key:"id" value:"ap2"}} elem:{name:"display-name"}`:
			assert.Equal(t, `string_val:"2nd Sample display name"`, upd.Val.String())
		case `elem:{name:"access-profile"} elem:{name:"access-profile" key:{key:"id" value:"ap2"}} elem:{name:"type"}`:
			assert.Equal(t, `string_val:"2nd Sample ap type"`, upd.Val.String())
		case `elem:{name:"access-profile"} elem:{name:"access-profile" key:{key:"id" value:"ap2"}} elem:{name:"filter"}`:
			assert.Equal(t, `string_val:"2nd Sample ap filter"`, upd.Val.String())
		case `elem:{name:"apn-profile"} elem:{name:"apn-profile" key:{key:"id" value:"apn1"}} elem:{name:"id"}`:
			assert.Equal(t, `string_val:"apn1"`, upd.Val.String())
		case `elem:{name:"apn-profile"} elem:{name:"apn-profile" key:{key:"id" value:"apn1"}} elem:{name:"description"}`:
			assert.Equal(t, `string_val:"apn1 Description"`, upd.Val.String())
		case `elem:{name:"apn-profile"} elem:{name:"apn-profile" key:{key:"id" value:"apn1"}} elem:{name:"display-name"}`:
			assert.Equal(t, `string_val:"apn1 display name"`, upd.Val.String())
		case `elem:{name:"apn-profile"} elem:{name:"apn-profile" key:{key:"id" value:"apn1"}} elem:{name:"apn-name"}`:
			assert.Equal(t, `string_val:"apn1 name"`, upd.Val.String())
		case `elem:{name:"apn-profile"} elem:{name:"apn-profile" key:{key:"id" value:"apn1"}} elem:{name:"dns-primary"}`:
			assert.Equal(t, `string_val:"10.1.2.3"`, upd.Val.String())
		case `elem:{name:"apn-profile"} elem:{name:"apn-profile" key:{key:"id" value:"apn1"}} elem:{name:"dns-secondary"}`:
			assert.Equal(t, `string_val:"10.1.2.4"`, upd.Val.String())
		case `elem:{name:"apn-profile"} elem:{name:"apn-profile" key:{key:"id" value:"apn1"}} elem:{name:"gx-enabled"}`:
			assert.Equal(t, `bool_val:false`, upd.Val.String())
		case `elem:{name:"apn-profile"} elem:{name:"apn-profile" key:{key:"id" value:"apn1"}} elem:{name:"mtu"}`:
			assert.Equal(t, `uint_val:9600`, upd.Val.String())
		default:
			t.Fatalf("unexpected path %s", path)
		}
	}

	assert.Equal(t, 1, len(deletes))
	for _, del := range deletes {
		switch path := strings.ReplaceAll(del.String(), "  ", " "); path {
		case `elem:{name:"access-profile"} elem:{name:"access-profile" key:{key:"id" value:"ap3d"}} elem:{name:"id"}`:
		default:
			t.Fatalf("unexpected path %s", path)
		}
	}

}
