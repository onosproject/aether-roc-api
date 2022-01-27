// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package server

import (
	"github.com/onosproject/aether-roc-api/pkg/aether_2_0_0/types"
	"gotest.tools/assert"
	"strings"
	"testing"
)

// Test the ability to remove index updates when requested (needed for delete of non index attributes)
func Test_encodeToGnmiIpDomainRemoveIndex(t *testing.T) {
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

	ipDomains := []types.EnterprisesEnterpriseSiteIpDomain{
		{
			Description: &ipd1Desc,
			IpId:        ipd1ID, // With the ID in the middle
			Mtu:         &ipd1Mtu,
			Dnn:         ipd1Dnn,
			AdditionalProperties: map[string]types.AdditionalPropertyUnchanged{
				"unused": {Unchanged: &addPropsUnch1},
			},
		},
		{
			IpId: ipd2ID, // With only the ID - should not remove
			Dnn:  ipd2Dnn,
			AdditionalProperties: map[string]types.AdditionalPropertyUnchanged{
				"unused": {Unchanged: &addPropsUnch1},
			},
		},
		{
			Description: &ipd3Desc,
			Dnn:         ipd3Dnn,
			IpId:        ipd3ID, // With the ID last
			AdditionalProperties: map[string]types.AdditionalPropertyUnchanged{
				"unused": {Unchanged: &addPropsUnch1},
			},
		},
		{
			IpId:        ipd4ID, // With the ID first
			Dnn:         ipd4Dnn,
			Description: &ipd4Desc,
			AdditionalProperties: map[string]types.AdditionalPropertyUnchanged{
				"unused": {Unchanged: &addPropsUnch1},
			},
		},
	}

	jsonObj := types.Enterprises{
		Enterprise: &[]types.EnterprisesEnterprise{
			{
				EntId: "ent-1",
				Site: &[]types.EnterprisesEnterpriseSite{
					{
						IpDomain: &ipDomains,
						SiteId:   "site-1",
					},
				},
			},
		},
	}

	gnmiUpdates, err := EncodeToGnmiEnterprises(&jsonObj, false, true, "target1", "/enterprises")
	assert.NilError(t, err)
	assert.Equal(t, 8, len(gnmiUpdates))
	for _, gnmiUpdate := range gnmiUpdates {
		switch path := strings.ReplaceAll(gnmiUpdate.String(), "  ", " "); path {
		case
			`path:{elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"ent-1"}} elem:{name:"site" key:{key:"site-id" value:"site-1"}} elem:{name:"ip-domain" key:{key:"ip-id" value:"ipd1"}} elem:{name:"description"} target:"target1"} val:{string_val:"IpDomain1 Desc"}`,
			`path:{elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"ent-1"}} elem:{name:"site" key:{key:"site-id" value:"site-1"}} elem:{name:"ip-domain" key:{key:"ip-id" value:"ipd1"}} elem:{name:"mtu"} target:"target1"} val:{uint_val:9601}`,
			`path:{elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"ent-1"}} elem:{name:"site" key:{key:"site-id" value:"site-1"}} elem:{name:"ip-domain" key:{key:"ip-id" value:"ipd1"}} elem:{name:"dnn"} target:"target1"} val:{string_val:"dnn1"}`,
			// And for second instance
			`path:{elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"ent-1"}} elem:{name:"site" key:{key:"site-id" value:"site-1"}} elem:{name:"ip-domain" key:{key:"ip-id" value:"ipd2"}} elem:{name:"id"} target:"target1"} val:{string_val:"ipd2"}`,
			`path:{elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"ent-1"}} elem:{name:"site" key:{key:"site-id" value:"site-1"}} elem:{name:"ip-domain" key:{key:"ip-id" value:"ipd2"}} elem:{name:"dnn"} target:"target1"} val:{string_val:"dnn2"}`,
			`path:{elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"ent-1"}} elem:{name:"site" key:{key:"site-id" value:"site-1"}} elem:{name:"ip-domain" key:{key:"ip-id" value:"ipd3"}} elem:{name:"description"} target:"target1"} val:{string_val:"IpDomain3 Desc"}`,
			`path:{elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"ent-1"}} elem:{name:"site" key:{key:"site-id" value:"site-1"}} elem:{name:"ip-domain" key:{key:"ip-id" value:"ipd3"}} elem:{name:"dnn"} target:"target1"} val:{string_val:"dnn3"}`,
			`path:{elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"ent-1"}} elem:{name:"site" key:{key:"site-id" value:"site-1"}} elem:{name:"ip-domain" key:{key:"ip-id" value:"ipd4"}} elem:{name:"description"} target:"target1"} val:{string_val:"IpDomain4 Desc"}`,
			`path:{elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"ent-1"}} elem:{name:"site" key:{key:"site-id" value:"site-1"}} elem:{name:"ip-domain" key:{key:"ip-id" value:"ipd4"}} elem:{name:"dnn"} target:"target1"} val:{string_val:"dnn4"}`:

		default:
			t.Fatalf("unexpected: %s", path)
		}
	}
}
