// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package server

import (
	"github.com/onosproject/aether-roc-api/pkg/aether_2_0_0/types"
	"gotest.tools/assert"
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

	gnmiUpdates, err := encodeToGnmiApnProfile(&jsonObj, false, "/apn-profile")
	assert.NilError(t, err)
	assert.Equal(t, 15, len(gnmiUpdates))
	for _, gnmiUpdate := range gnmiUpdates {
		switch gnmiUpdate.String() {
		case
			`path:{elem:{name:"apn-profile"} elem:{name:"apn-profile" key:{key:"id" value:"apn1"}} elem:{name:"apn-name"}} val:{string_val:"APN1 Name"}`,
			`path:{elem:{name:"apn-profile"} elem:{name:"apn-profile" key:{key:"id" value:"apn1"}} elem:{name:"description"}} val:{string_val:"APN1 Desc"}`,
			`path:{elem:{name:"apn-profile"} elem:{name:"apn-profile" key:{key:"id" value:"apn1"}} elem:{name:"display-name"}} val:{string_val:"APN1 Display Name"}`,
			`path:{elem:{name:"apn-profile"} elem:{name:"apn-profile" key:{key:"id" value:"apn1"}} elem:{name:"dns-primary"}} val:{string_val:"1.1.1.1"}`,
			`path:{elem:{name:"apn-profile"} elem:{name:"apn-profile" key:{key:"id" value:"apn1"}} elem:{name:"dns-secondary"}} val:{string_val:"1.1.1.0"}`,
			`path:{elem:{name:"apn-profile"} elem:{name:"apn-profile" key:{key:"id" value:"apn1"}} elem:{name:"gx-enabled"}} val:{bool_val:true}`,
			`path:{elem:{name:"apn-profile"} elem:{name:"apn-profile" key:{key:"id" value:"apn1"}} elem:{name:"id"}} val:{string_val:"apn1"}`,
			`path:{elem:{name:"apn-profile"} elem:{name:"apn-profile" key:{key:"id" value:"apn1"}} elem:{name:"mtu"}} val:{uint_val:9601}`,
			// And with double spacing for some reason
			`path:{elem:{name:"apn-profile"}  elem:{name:"apn-profile"  key:{key:"id"  value:"apn1"}}  elem:{name:"apn-name"}}  val:{string_val:"APN1 Name"}`,
			`path:{elem:{name:"apn-profile"}  elem:{name:"apn-profile"  key:{key:"id"  value:"apn1"}}  elem:{name:"description"}}  val:{string_val:"APN1 Desc"}`,
			`path:{elem:{name:"apn-profile"}  elem:{name:"apn-profile"  key:{key:"id"  value:"apn1"}}  elem:{name:"display-name"}}  val:{string_val:"APN1 Display Name"}`,
			`path:{elem:{name:"apn-profile"}  elem:{name:"apn-profile"  key:{key:"id"  value:"apn1"}}  elem:{name:"dns-primary"}}  val:{string_val:"1.1.1.1"}`,
			`path:{elem:{name:"apn-profile"}  elem:{name:"apn-profile"  key:{key:"id"  value:"apn1"}}  elem:{name:"dns-secondary"}}  val:{string_val:"1.1.1.0"}`,
			`path:{elem:{name:"apn-profile"}  elem:{name:"apn-profile"  key:{key:"id"  value:"apn1"}}  elem:{name:"gx-enabled"}}  val:{bool_val:true}`,
			`path:{elem:{name:"apn-profile"}  elem:{name:"apn-profile"  key:{key:"id"  value:"apn1"}}  elem:{name:"id"}}  val:{string_val:"apn1"}`,
			`path:{elem:{name:"apn-profile"}  elem:{name:"apn-profile"  key:{key:"id"  value:"apn1"}}  elem:{name:"mtu"}}  val:{uint_val:9601}`,
			// And for second instance
			`path:{elem:{name:"apn-profile"} elem:{name:"apn-profile" key:{key:"id" value:"apn2"}} elem:{name:"apn-name"}} val:{string_val:"APN2 Name"}`,
			`path:{elem:{name:"apn-profile"} elem:{name:"apn-profile" key:{key:"id" value:"apn2"}} elem:{name:"description"}} val:{string_val:"APN2 Desc"}`,
			`path:{elem:{name:"apn-profile"} elem:{name:"apn-profile" key:{key:"id" value:"apn2"}} elem:{name:"dns-primary"}} val:{string_val:"2.2.2.2"}`,
			`path:{elem:{name:"apn-profile"} elem:{name:"apn-profile" key:{key:"id" value:"apn2"}} elem:{name:"dns-secondary"}} val:{string_val:"2.2.2.0"}`,
			`path:{elem:{name:"apn-profile"} elem:{name:"apn-profile" key:{key:"id" value:"apn2"}} elem:{name:"gx-enabled"}} val:{bool_val:false}`,
			`path:{elem:{name:"apn-profile"} elem:{name:"apn-profile" key:{key:"id" value:"apn2"}} elem:{name:"id"}} val:{string_val:"apn2"}`,
			`path:{elem:{name:"apn-profile"} elem:{name:"apn-profile" key:{key:"id" value:"apn2"}} elem:{name:"mtu"}} val:{uint_val:9602}`,
			// And with double spacing for some reason
			`path:{elem:{name:"apn-profile"}  elem:{name:"apn-profile"  key:{key:"id"  value:"apn2"}}  elem:{name:"apn-name"}}  val:{string_val:"APN2 Name"}`,
			`path:{elem:{name:"apn-profile"}  elem:{name:"apn-profile"  key:{key:"id"  value:"apn2"}}  elem:{name:"description"}}  val:{string_val:"APN2 Desc"}`,
			`path:{elem:{name:"apn-profile"}  elem:{name:"apn-profile"  key:{key:"id"  value:"apn2"}}  elem:{name:"dns-primary"}}  val:{string_val:"2.2.2.2"}`,
			`path:{elem:{name:"apn-profile"}  elem:{name:"apn-profile"  key:{key:"id"  value:"apn2"}}  elem:{name:"dns-secondary"}}  val:{string_val:"2.2.2.0"}`,
			`path:{elem:{name:"apn-profile"}  elem:{name:"apn-profile"  key:{key:"id"  value:"apn2"}}  elem:{name:"gx-enabled"}}  val:{bool_val:false}`,
			`path:{elem:{name:"apn-profile"}  elem:{name:"apn-profile"  key:{key:"id"  value:"apn2"}}  elem:{name:"id"}}  val:{string_val:"apn2"}`,
			`path:{elem:{name:"apn-profile"}  elem:{name:"apn-profile"  key:{key:"id"  value:"apn2"}}  elem:{name:"mtu"}}  val:{uint_val:9602}`:

		default:
			t.Logf("unexpected: %s", gnmiUpdate.String())
		}
	}

}

func Test_encodeToGnmiSubscriberUe(t *testing.T) {
	ue1Id := "Ue1"
	ue1DispName := "UE1 Displ Name"
	ue1Enabled := true
	ue1Priority := int32(10)
	ue1RequestedApn := "UE1ReqApn"

	ue1RangeFrom := int64(123)
	ue1RangeTo := int64(321)
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
	gnmiUpdates, err := encodeToGnmiSubscriber(&jsonObj, false, "/subscriber")
	assert.NilError(t, err)
	assert.Equal(t, 19, len(gnmiUpdates))

}
