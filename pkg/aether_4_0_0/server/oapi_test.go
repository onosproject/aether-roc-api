// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
//

package server

import (
	"github.com/onosproject/aether-roc-api/pkg/aether_4_0_0/types"
	"gotest.tools/assert"
	"strings"
	"testing"
)

func Test_encodeToGnmiAccessProfile(t *testing.T) {
	ipdID := "ipd1"
	ipd1Desc := "IpDomain1 Desc"
	ipd1Disp := "IpDomain1 Display Name"
	ipd1Mtu := 9601
	ipd1DnsPri := "1.1.1.1"
	ipd1DnsSec := "1.1.1.0"
	ipd1Dnn := "dnn1"
	ipd1Subnet := "sub-1"
	ipd1admin := "ENABLE"
	ipd1ent := "ent1"

	ipd2ID := "ipd2"
	ipd2Desc := "IpDomain2 Desc"
	ipd2Mtu := 9602
	ipd2DnsPri := "2.2.2.2"
	ipd2DnsSec := "2.2.2.0"
	ipd2Dnn := "dnn2"
	ipd2Subnet := "sub-2"
	ipd2ent := "ent2"

	ipDomains := []types.IpDomainIpDomain{
		{
			Description:  &ipd1Desc,
			DisplayName:  &ipd1Disp,
			DnsPrimary:   &ipd1DnsPri,
			DnsSecondary: &ipd1DnsSec,
			Id:           ipdID,
			Mtu:          &ipd1Mtu,
			Dnn:          ipd1Dnn,
			Subnet:       ipd1Subnet,
			AdminStatus:  &ipd1admin,
			Enterprise:   ipd1ent,
		},
		{
			Description:  &ipd2Desc,
			DnsPrimary:   &ipd2DnsPri,
			DnsSecondary: &ipd2DnsSec,
			Id:           ipd2ID,
			Mtu:          &ipd2Mtu,
			Dnn:          ipd2Dnn,
			Subnet:       ipd2Subnet,
			Enterprise:   ipd2ent,
		},
	}

	jsonObj := types.IpDomain{
		IpDomain: &ipDomains,
	}

	gnmiUpdates, err := EncodeToGnmiIpDomain(&jsonObj, false, false, "target1", "/ip-domain")
	assert.NilError(t, err)
	assert.Equal(t, 18, len(gnmiUpdates))
	for _, gnmiUpdate := range gnmiUpdates {
		switch path := strings.ReplaceAll(gnmiUpdate.String(), "  ", " "); path {
		case
			`path:{elem:{name:"ip-domain"} elem:{name:"ip-domain" key:{key:"id" value:"ipd1"}} elem:{name:"admin-status"} target:"target1"} val:{string_val:"ENABLE"}`,
			`path:{elem:{name:"ip-domain"} elem:{name:"ip-domain" key:{key:"id" value:"ipd1"}} elem:{name:"description"} target:"target1"} val:{string_val:"IpDomain1 Desc"}`,
			`path:{elem:{name:"ip-domain"} elem:{name:"ip-domain" key:{key:"id" value:"ipd1"}} elem:{name:"display-name"} target:"target1"} val:{string_val:"IpDomain1 Display Name"}`,
			`path:{elem:{name:"ip-domain"} elem:{name:"ip-domain" key:{key:"id" value:"ipd1"}} elem:{name:"dns-primary"} target:"target1"} val:{string_val:"1.1.1.1"}`,
			`path:{elem:{name:"ip-domain"} elem:{name:"ip-domain" key:{key:"id" value:"ipd1"}} elem:{name:"dns-secondary"} target:"target1"} val:{string_val:"1.1.1.0"}`,
			`path:{elem:{name:"ip-domain"} elem:{name:"ip-domain" key:{key:"id" value:"ipd1"}} elem:{name:"dnn"} target:"target1"} val:{string_val:"dnn1"}`,
			`path:{elem:{name:"ip-domain"} elem:{name:"ip-domain" key:{key:"id" value:"ipd1"}} elem:{name:"enterprise"} target:"target1"} val:{string_val:"ent1"}`,
			`path:{elem:{name:"ip-domain"} elem:{name:"ip-domain" key:{key:"id" value:"ipd1"}} elem:{name:"subnet"} target:"target1"} val:{string_val:"sub-1"}`,
			`path:{elem:{name:"ip-domain"} elem:{name:"ip-domain" key:{key:"id" value:"ipd1"}} elem:{name:"id"} target:"target1"} val:{string_val:"ipd1"}`,
			`path:{elem:{name:"ip-domain"} elem:{name:"ip-domain" key:{key:"id" value:"ipd1"}} elem:{name:"mtu"} target:"target1"} val:{uint_val:9601}`,
			// And for second instance
			`path:{elem:{name:"ip-domain"} elem:{name:"ip-domain" key:{key:"id" value:"ipd2"}} elem:{name:"apn-name"} target:"target1"} val:{string_val:"IpDomain2 Name"}`,
			`path:{elem:{name:"ip-domain"} elem:{name:"ip-domain" key:{key:"id" value:"ipd2"}} elem:{name:"description"} target:"target1"} val:{string_val:"IpDomain2 Desc"}`,
			`path:{elem:{name:"ip-domain"} elem:{name:"ip-domain" key:{key:"id" value:"ipd2"}} elem:{name:"dns-primary"} target:"target1"} val:{string_val:"2.2.2.2"}`,
			`path:{elem:{name:"ip-domain"} elem:{name:"ip-domain" key:{key:"id" value:"ipd2"}} elem:{name:"dns-secondary"} target:"target1"} val:{string_val:"2.2.2.0"}`,
			`path:{elem:{name:"ip-domain"} elem:{name:"ip-domain" key:{key:"id" value:"ipd2"}} elem:{name:"dnn"} target:"target1"} val:{string_val:"dnn2"}`,
			`path:{elem:{name:"ip-domain"} elem:{name:"ip-domain" key:{key:"id" value:"ipd2"}} elem:{name:"enterprise"} target:"target1"} val:{string_val:"ent2"}`,
			`path:{elem:{name:"ip-domain"} elem:{name:"ip-domain" key:{key:"id" value:"ipd2"}} elem:{name:"subnet"} target:"target1"} val:{string_val:"sub-2"}`,
			`path:{elem:{name:"ip-domain"} elem:{name:"ip-domain" key:{key:"id" value:"ipd2"}} elem:{name:"id"} target:"target1"} val:{string_val:"ipd2"}`,
			`path:{elem:{name:"ip-domain"} elem:{name:"ip-domain" key:{key:"id" value:"ipd2"}} elem:{name:"mtu"} target:"target1"} val:{uint_val:9602}`:

		default:
			t.Logf("unexpected: %s", path)
			t.Fail()
		}
	}
}

// Test the ability to remove index updates when requested (needed for delete of non index attributes)
func Test_encodeToGnmiAccessProfileRemoveIndex(t *testing.T) {
	ipd1ID := "ipd1"
	ipd1Desc := "IpDomain1 Desc"
	ipd1Mtu := 9601
	ipd1Dnn := "dnn1"

	ipd2ID := "ipd2"
	ipd2Dnn := "dnn2"
	ipd3ID := "ipd3"
	ipd3Dnn := "dnn3"
	ipd3Desc := "IpDomain3 Desc"
	ipd4ID := "ipd4"
	ipd4Dnn := "dnn4"
	ipd4Desc := "IpDomain4 Desc"

	addPropsUnch1 := "enterprise,subnet"

	ipDomains := []types.IpDomainIpDomain{
		{
			Description: &ipd1Desc,
			Id:          ipd1ID, // With the ID in the middle
			Mtu:         &ipd1Mtu,
			Dnn:         ipd1Dnn,
			AdditionalProperties: map[string]types.AdditionalPropertyUnchanged{
				"unused": {Unchanged: &addPropsUnch1},
			},
		},
		{
			Id:  ipd2ID, // With only the ID - should not remove
			Dnn: ipd2Dnn,
			AdditionalProperties: map[string]types.AdditionalPropertyUnchanged{
				"unused": {Unchanged: &addPropsUnch1},
			},
		},
		{
			Description: &ipd3Desc,
			Dnn:         ipd3Dnn,
			Id:          ipd3ID, // With the ID last
			AdditionalProperties: map[string]types.AdditionalPropertyUnchanged{
				"unused": {Unchanged: &addPropsUnch1},
			},
		},
		{
			Id:          ipd4ID, // With the ID first
			Dnn:         ipd4Dnn,
			Description: &ipd4Desc,
			AdditionalProperties: map[string]types.AdditionalPropertyUnchanged{
				"unused": {Unchanged: &addPropsUnch1},
			},
		},
	}

	jsonObj := types.IpDomain{
		IpDomain: &ipDomains,
	}

	gnmiUpdates, err := EncodeToGnmiIpDomain(&jsonObj, false, true, "target1", "/ip-domain")
	assert.NilError(t, err)
	assert.Equal(t, 8, len(gnmiUpdates))
	for _, gnmiUpdate := range gnmiUpdates {
		switch path := strings.ReplaceAll(gnmiUpdate.String(), "  ", " "); path {
		case
			`path:{elem:{name:"ip-domain"} elem:{name:"ip-domain" key:{key:"id" value:"ipd1"}} elem:{name:"description"} target:"target1"} val:{string_val:"IpDomain1 Desc"}`,
			`path:{elem:{name:"ip-domain"} elem:{name:"ip-domain" key:{key:"id" value:"ipd1"}} elem:{name:"mtu"} target:"target1"} val:{uint_val:9601}`,
			`path:{elem:{name:"ip-domain"} elem:{name:"ip-domain" key:{key:"id" value:"ipd1"}} elem:{name:"dnn"} target:"target1"} val:{string_val:"dnn1"}`,
			// And for second instance
			`path:{elem:{name:"ip-domain"} elem:{name:"ip-domain" key:{key:"id" value:"ipd2"}} elem:{name:"id"} target:"target1"} val:{string_val:"ipd2"}`,
			`path:{elem:{name:"ip-domain"} elem:{name:"ip-domain" key:{key:"id" value:"ipd2"}} elem:{name:"dnn"} target:"target1"} val:{string_val:"dnn2"}`,
			`path:{elem:{name:"ip-domain"} elem:{name:"ip-domain" key:{key:"id" value:"ipd3"}} elem:{name:"description"} target:"target1"} val:{string_val:"IpDomain3 Desc"}`,
			`path:{elem:{name:"ip-domain"} elem:{name:"ip-domain" key:{key:"id" value:"ipd3"}} elem:{name:"dnn"} target:"target1"} val:{string_val:"dnn3"}`,
			`path:{elem:{name:"ip-domain"} elem:{name:"ip-domain" key:{key:"id" value:"ipd4"}} elem:{name:"description"} target:"target1"} val:{string_val:"IpDomain4 Desc"}`,
			`path:{elem:{name:"ip-domain"} elem:{name:"ip-domain" key:{key:"id" value:"ipd4"}} elem:{name:"dnn"} target:"target1"} val:{string_val:"dnn4"}`:

		default:
			t.Fatalf("unexpected: %s", path)
		}
	}
}
