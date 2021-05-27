// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package server

import (
	"encoding/json"
	"github.com/onosproject/aether-roc-api/pkg/aether_2_0_0/types"
	"gotest.tools/assert"
	"io/ioutil"
	"strings"
	"testing"
)

func Test_encodeToGnmiAccessProfile(t *testing.T) {
	apn1ID := "apn1"
	apn1Name := "APN1 Name"
	apn1Desc := "APN1 Desc"
	apn1Disp := "APN1 Display Name"
	apn1Mtu := int32(9601)
	apn1GxEnabled := true
	apn1DnsPri := "1.1.1.1"
	apn1DnsSec := "1.1.1.0"

	apn2ID := "apn2"
	apn2Name := "APN2 Name"
	apn2Desc := "APN2 Desc"
	apn2Mtu := int32(9602)
	apn2GxEnabled := false
	apn2DnsPri := "2.2.2.2"
	apn2DnsSec := "2.2.2.0"

	apList := []types.ApnProfileApnProfile{
		{
			ApnName:      &apn1Name,
			Description:  &apn1Desc,
			DisplayName:  &apn1Disp,
			DnsPrimary:   &apn1DnsPri,
			DnsSecondary: &apn1DnsSec,
			GxEnabled:    &apn1GxEnabled,
			Id:           &apn1ID,
			Mtu:          &apn1Mtu,
		},
		{
			ApnName:      &apn2Name,
			Description:  &apn2Desc,
			DnsPrimary:   &apn2DnsPri,
			DnsSecondary: &apn2DnsSec,
			GxEnabled:    &apn2GxEnabled,
			Id:           &apn2ID,
			Mtu:          &apn2Mtu,
		},
	}

	jsonObj := types.ApnProfile{
		ApnProfile: &apList,
	}

	gnmiUpdates, err := EncodeToGnmiApnProfile(&jsonObj, false, false, "target1", "/apn-profile")
	assert.NilError(t, err)
	assert.Equal(t, 15, len(gnmiUpdates))
	for _, gnmiUpdate := range gnmiUpdates {
		switch path := strings.ReplaceAll(gnmiUpdate.String(), "  ", " "); path {
		case
			`path:{elem:{name:"apn-profile"} elem:{name:"apn-profile" key:{key:"id" value:"apn1"}} elem:{name:"apn-name"}} val:{string_val:"APN1 Name"}`,
			`path:{elem:{name:"apn-profile"} elem:{name:"apn-profile" key:{key:"id" value:"apn1"}} elem:{name:"description"}} val:{string_val:"APN1 Desc"}`,
			`path:{elem:{name:"apn-profile"} elem:{name:"apn-profile" key:{key:"id" value:"apn1"}} elem:{name:"display-name"}} val:{string_val:"APN1 Display Name"}`,
			`path:{elem:{name:"apn-profile"} elem:{name:"apn-profile" key:{key:"id" value:"apn1"}} elem:{name:"dns-primary"}} val:{string_val:"1.1.1.1"}`,
			`path:{elem:{name:"apn-profile"} elem:{name:"apn-profile" key:{key:"id" value:"apn1"}} elem:{name:"dns-secondary"}} val:{string_val:"1.1.1.0"}`,
			`path:{elem:{name:"apn-profile"} elem:{name:"apn-profile" key:{key:"id" value:"apn1"}} elem:{name:"gx-enabled"}} val:{bool_val:true}`,
			`path:{elem:{name:"apn-profile"} elem:{name:"apn-profile" key:{key:"id" value:"apn1"}} elem:{name:"id"}} val:{string_val:"apn1"}`,
			`path:{elem:{name:"apn-profile"} elem:{name:"apn-profile" key:{key:"id" value:"apn1"}} elem:{name:"mtu"}} val:{uint_val:9601}`,
			// And for second instance
			`path:{elem:{name:"apn-profile"} elem:{name:"apn-profile" key:{key:"id" value:"apn2"}} elem:{name:"apn-name"}} val:{string_val:"APN2 Name"}`,
			`path:{elem:{name:"apn-profile"} elem:{name:"apn-profile" key:{key:"id" value:"apn2"}} elem:{name:"description"}} val:{string_val:"APN2 Desc"}`,
			`path:{elem:{name:"apn-profile"} elem:{name:"apn-profile" key:{key:"id" value:"apn2"}} elem:{name:"dns-primary"}} val:{string_val:"2.2.2.2"}`,
			`path:{elem:{name:"apn-profile"} elem:{name:"apn-profile" key:{key:"id" value:"apn2"}} elem:{name:"dns-secondary"}} val:{string_val:"2.2.2.0"}`,
			`path:{elem:{name:"apn-profile"} elem:{name:"apn-profile" key:{key:"id" value:"apn2"}} elem:{name:"gx-enabled"}} val:{bool_val:false}`,
			`path:{elem:{name:"apn-profile"} elem:{name:"apn-profile" key:{key:"id" value:"apn2"}} elem:{name:"id"}} val:{string_val:"apn2"}`,
			`path:{elem:{name:"apn-profile"} elem:{name:"apn-profile" key:{key:"id" value:"apn2"}} elem:{name:"mtu"}} val:{uint_val:9602}`:

		default:
			t.Logf("unexpected: %s", path)
		}
	}

}

func Test_encodeToGnmiSubscriberUe(t *testing.T) {
	ue1Id := "Ue1"
	ue1DispName := "UE1 Displ Name"
	ue1Enabled := true
	ue1Priority := int32(10)
	ue1RequestedApn := "UE1ReqApn"

	ue1RangeFrom := int64(1<<63 - 10)
	ue1RangeTo := int64(1<<63 - 1)
	ap1 := "ap1"
	ap1allowed := true
	ap2 := "ap2"
	ap2allowed := true

	apn1 := "apn1"
	qos1 := "qos1"
	sec1 := "sec1"
	up1 := "up1"

	mcc := int32(123)
	mnc := int32(456)
	tac := int32(789)

	ent1 := "ent1"

	accessProfiles := []types.SubscriberUeProfilesAccessProfile{
		{
			AccessProfile: &ap1,
			Allowed:       &ap1allowed,
		},
		{
			AccessProfile: &ap2,
			Allowed:       &ap2allowed,
		},
	}

	subscriberUeList := []types.SubscriberUe{
		{
			Profiles: &types.SubscriberUeProfiles{
				AccessProfile:   &accessProfiles,
				ApnProfile:      &apn1,
				QosProfile:      &qos1,
				SecurityProfile: &sec1,
				UpProfile:       &up1,
			},
			ServingPlmn: &types.SubscriberUeServingPlmn{
				Mcc: &mcc,
				Mnc: &mnc,
				Tac: &tac,
			},
			ImsiRangeFrom: &ue1RangeFrom,
			ImsiRangeTo:   &ue1RangeTo,
			DisplayName:   &ue1DispName,
			Enabled:       &ue1Enabled,
			Enterprise:    &ent1,
			Id:            &ue1Id,
			Priority:      &ue1Priority,
			RequestedApn:  &ue1RequestedApn,
		},
	}

	jsonObj := types.Subscriber{
		Ue: &subscriberUeList,
	}
	gnmiUpdates, err := EncodeToGnmiSubscriber(&jsonObj, false, false, "target1", "/subscriber")
	assert.NilError(t, err)
	assert.Equal(t, 19, len(gnmiUpdates))
	for _, upd := range gnmiUpdates {
		switch path := strings.ReplaceAll(upd.Path.String(), "  ", " "); path {
		case `elem:{name:"subscriber"} elem:{name:"ue" key:{key:"id" value:"Ue1"}} elem:{name:"profiles"} elem:{name:"apn-profile"} target:"target1"`:
			assert.Equal(t, `string_val:"apn1"`, upd.Val.String())
		case `elem:{name:"subscriber"} elem:{name:"ue" key:{key:"id" value:"Ue1"}} elem:{name:"profiles"} elem:{name:"qos-profile"} target:"target1"`:
			assert.Equal(t, `string_val:"qos1"`, upd.Val.String())
		case `elem:{name:"subscriber"} elem:{name:"ue" key:{key:"id" value:"Ue1"}} elem:{name:"profiles"} elem:{name:"security-profile"} target:"target1"`:
			assert.Equal(t, `string_val:"sec1"`, upd.Val.String())
		case `elem:{name:"subscriber"} elem:{name:"ue" key:{key:"id" value:"Ue1"}} elem:{name:"profiles"} elem:{name:"up-profile"} target:"target1"`:
			assert.Equal(t, `string_val:"up1"`, upd.Val.String())
		case `elem:{name:"subscriber"} elem:{name:"ue" key:{key:"id" value:"Ue1"}} elem:{name:"profiles"} elem:{name:"access-profile" key:{key:"access-profile" value:"ap1"}} elem:{name:"access-profile"} target:"target1"`:
			assert.Equal(t, `string_val:"ap1"`, upd.Val.String())
		case `elem:{name:"subscriber"} elem:{name:"ue" key:{key:"id" value:"Ue1"}} elem:{name:"profiles"} elem:{name:"access-profile" key:{key:"access-profile" value:"ap1"}} elem:{name:"allowed"} target:"target1"`:
			assert.Equal(t, `bool_val:true`, upd.Val.String())
		case `elem:{name:"subscriber"} elem:{name:"ue" key:{key:"id" value:"Ue1"}} elem:{name:"profiles"} elem:{name:"access-profile" key:{key:"access-profile" value:"ap2"}} elem:{name:"access-profile"} target:"target1"`:
			assert.Equal(t, `string_val:"ap2"`, upd.Val.String())
		case `elem:{name:"subscriber"} elem:{name:"ue" key:{key:"id" value:"Ue1"}} elem:{name:"profiles"} elem:{name:"access-profile" key:{key:"access-profile" value:"ap2"}} elem:{name:"allowed"} target:"target1"`:
			assert.Equal(t, `bool_val:true`, upd.Val.String())
		case `elem:{name:"subscriber"} elem:{name:"ue" key:{key:"id" value:"Ue1"}} elem:{name:"serving-plmn"} elem:{name:"mcc"} target:"target1"`:
			assert.Equal(t, `uint_val:123`, upd.Val.String())
		case `elem:{name:"subscriber"} elem:{name:"ue" key:{key:"id" value:"Ue1"}} elem:{name:"serving-plmn"} elem:{name:"mnc"} target:"target1"`:
			assert.Equal(t, `uint_val:456`, upd.Val.String())
		case `elem:{name:"subscriber"} elem:{name:"ue" key:{key:"id" value:"Ue1"}} elem:{name:"serving-plmn"} elem:{name:"tac"} target:"target1"`:
			assert.Equal(t, `uint_val:789`, upd.Val.String())
		case `elem:{name:"subscriber"} elem:{name:"ue" key:{key:"id" value:"Ue1"}} elem:{name:"display-name"} target:"target1"`:
			assert.Equal(t, `string_val:"UE1 Displ Name"`, upd.Val.String())
		case `elem:{name:"subscriber"} elem:{name:"ue" key:{key:"id" value:"Ue1"}} elem:{name:"enabled"} target:"target1"`:
			assert.Equal(t, `bool_val:true`, upd.Val.String())
		case `elem:{name:"subscriber"} elem:{name:"ue" key:{key:"id" value:"Ue1"}} elem:{name:"enterprise"} target:"target1"`:
			assert.Equal(t, `string_val:"ent1"`, upd.Val.String())
		case `elem:{name:"subscriber"} elem:{name:"ue" key:{key:"id" value:"Ue1"}} elem:{name:"id"} target:"target1"`:
			assert.Equal(t, `string_val:"Ue1"`, upd.Val.String())
		case `elem:{name:"subscriber"} elem:{name:"ue" key:{key:"id" value:"Ue1"}} elem:{name:"imsi-range-from"} target:"target1"`:
			assert.Equal(t, uint64(9223372036854775798), upd.Val.GetUintVal())
		case `elem:{name:"subscriber"} elem:{name:"ue" key:{key:"id" value:"Ue1"}} elem:{name:"imsi-range-to"} target:"target1"`:
			assert.Equal(t, uint64(9223372036854775807), upd.Val.GetUintVal())
		case `elem:{name:"subscriber"} elem:{name:"ue" key:{key:"id" value:"Ue1"}} elem:{name:"priority"} target:"target1"`:
			assert.Equal(t, `uint_val:10`, upd.Val.String())
		case `elem:{name:"subscriber"} elem:{name:"ue" key:{key:"id" value:"Ue1"}} elem:{name:"requested-apn"} target:"target1"`:
			assert.Equal(t, `string_val:"UE1ReqApn"`, upd.Val.String())
		default:
			t.Fatalf("unexpected path %s", path)
		}
	}

}

// Test where the update of a UE is called directly - it will have the UE ID in place of {id},
//and so must be wrapped in brackets
func Test_encodeToGnmiSubscriberUe2(t *testing.T) {
	ueExampleJSON, err := ioutil.ReadFile("../testdata/SubscriberUeOapiExample.json")
	assert.NilError(t, err, "error loading testdata file")
	jsonObj := new(types.SubscriberUe)
	err = json.Unmarshal(ueExampleJSON, jsonObj)
	assert.NilError(t, err)

	gnmiUpdates, err := EncodeToGnmiSubscriberUe(jsonObj, false, false, "target1", "/subscriber/ue/64Ff4CB4-Cc5B-F91c-9ED6-4dc133bA0599", "64Ff4CB4-Cc5B-F91c-9ED6-4dc133bA0599")
	assert.NilError(t, err)
	assert.Equal(t, 16, len(gnmiUpdates))
}
