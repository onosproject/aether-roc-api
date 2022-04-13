// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
//

package server

import (
	"github.com/onosproject/aether-roc-api/pkg/aether_2_1_0/types"
	"gotest.tools/assert"
	"strings"
	"testing"
)

// Test the ability to remove index updates when requested (needed for delete of non index attributes)
func Test_encodeToGnmiIpDomain(t *testing.T) {
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

	enterpriseTest := "enterprise-test"

	addPropsUnch1 := "subnet"

	ipDomains := types.SiteIpDomainList{
		{
			Description: &ipd1Desc,
			IpDomainId:  types.ListKey(ipd1ID), // With the ID in the middle
			Mtu:         &ipd1Mtu,
			Dnn:         ipd1Dnn,
			AdditionalProperties: map[string]types.AdditionalPropertyUnchanged{
				"unused": {Unchanged: &addPropsUnch1},
			},
		},
		{
			IpDomainId: types.ListKey(ipd2ID), // With only the ID - should not remove
			Dnn:        ipd2Dnn,
			AdditionalProperties: map[string]types.AdditionalPropertyUnchanged{
				"unused": {Unchanged: &addPropsUnch1},
			},
		},
		{
			Description: &ipd3Desc,
			Dnn:         ipd3Dnn,
			IpDomainId:  types.ListKey(ipd3ID), // With the ID last
			AdditionalProperties: map[string]types.AdditionalPropertyUnchanged{
				"unused": {Unchanged: &addPropsUnch1},
			},
		},
		{
			IpDomainId:  types.ListKey(ipd4ID), // With the ID first
			Dnn:         ipd4Dnn,
			Description: &ipd4Desc,
			AdditionalProperties: map[string]types.AdditionalPropertyUnchanged{
				"unused": {Unchanged: &addPropsUnch1},
			},
		},
	}

	jsonObj := types.Site{
		IpDomain: &ipDomains,
		SiteId:   "site-1",
		AdditionalProperties: map[string]types.AdditionalPropertyEnterpriseId{
			"enterprise-id": {
				EnterpriseId: &enterpriseTest,
			},
		},
	}

	gnmiUpdates, err := EncodeToGnmiSite(&jsonObj, false, false, types.EnterpriseId(enterpriseTest), "/site/{site-id}", "site-1")
	assert.NilError(t, err)
	assert.Equal(t, 13, len(gnmiUpdates))
	for _, gnmiUpdate := range gnmiUpdates {
		switch path := strings.ReplaceAll(gnmiUpdate.String(), "  ", " "); path {
		case
			`path:{elem:{name:"site" key:{key:"site-id" value:"site-1"}} elem:{name:"ip-domain" key:{key:"ip-domain-id" value:"ipd1"}} elem:{name:"description"} target:"enterprise-test"} val:{string_val:"IpDomain1 Desc"}`,
			`path:{elem:{name:"site" key:{key:"site-id" value:"site-1"}} elem:{name:"ip-domain" key:{key:"ip-domain-id" value:"ipd1"}} elem:{name:"mtu"} target:"enterprise-test"} val:{uint_val:9601}`,
			`path:{elem:{name:"site" key:{key:"site-id" value:"site-1"}} elem:{name:"ip-domain" key:{key:"ip-domain-id" value:"ipd1"}} elem:{name:"dnn"} target:"enterprise-test"} val:{string_val:"dnn1"}`,
			`path:{elem:{name:"site" key:{key:"site-id" value:"site-1"}} elem:{name:"ip-domain" key:{key:"ip-domain-id" value:"ipd1"}} elem:{name:"ip-domain-id"} target:"enterprise-test"} val:{string_val:"ipd1"}`,
			// And for second instance
			`path:{elem:{name:"site" key:{key:"site-id" value:"site-1"}} elem:{name:"ip-domain" key:{key:"ip-domain-id" value:"ipd2"}} elem:{name:"ip-domain-id"} target:"enterprise-test"} val:{string_val:"ipd2"}`,
			`path:{elem:{name:"site" key:{key:"site-id" value:"site-1"}} elem:{name:"ip-domain" key:{key:"ip-domain-id" value:"ipd2"}} elem:{name:"dnn"} target:"enterprise-test"} val:{string_val:"dnn2"}`,
			`path:{elem:{name:"site" key:{key:"site-id" value:"site-1"}} elem:{name:"ip-domain" key:{key:"ip-domain-id" value:"ipd3"}} elem:{name:"description"} target:"enterprise-test"} val:{string_val:"IpDomain3 Desc"}`,
			`path:{elem:{name:"site" key:{key:"site-id" value:"site-1"}} elem:{name:"ip-domain" key:{key:"ip-domain-id" value:"ipd3"}} elem:{name:"ip-domain-id"} target:"enterprise-test"} val:{string_val:"ipd3"}`,
			`path:{elem:{name:"site" key:{key:"site-id" value:"site-1"}} elem:{name:"ip-domain" key:{key:"ip-domain-id" value:"ipd3"}} elem:{name:"dnn"} target:"enterprise-test"} val:{string_val:"dnn3"}`,
			`path:{elem:{name:"site" key:{key:"site-id" value:"site-1"}} elem:{name:"ip-domain" key:{key:"ip-domain-id" value:"ipd4"}} elem:{name:"description"} target:"enterprise-test"} val:{string_val:"IpDomain4 Desc"}`,
			`path:{elem:{name:"site" key:{key:"site-id" value:"site-1"}} elem:{name:"ip-domain" key:{key:"ip-domain-id" value:"ipd4"}} elem:{name:"dnn"} target:"enterprise-test"} val:{string_val:"dnn4"}`,
			`path:{elem:{name:"site" key:{key:"site-id" value:"site-1"}} elem:{name:"ip-domain" key:{key:"ip-domain-id" value:"ipd4"}} elem:{name:"ip-domain-id"} target:"enterprise-test"} val:{string_val:"ipd4"}`,
			`path:{elem:{name:"site" key:{key:"site-id" value:"site-1"}} elem:{name:"site-id"} target:"enterprise-test"} val:{string_val:"site-1"}`:

		default:
			t.Fatalf("unexpected: %s", path)
		}
	}
}

func Test_encodeToGnmiSlice(t *testing.T) {
	slice1Desc := "Slice1 Desc"
	slice1ID := types.ListKey("slice1")

	slice2Desc := "Slice2 Desc"
	slice2ID := types.ListKey("slice2")
	cs5g := types.SiteSliceConnectivityService("5g")

	enterpriseTest := "enterprise-test"

	addPropsUnch1 := "sd,sst"
	addPropsUnch2 := "default-behavior,sd,sst"

	slices := types.SiteSliceList{
		{
			SliceId:         slice1ID,
			Description:     &slice1Desc,
			DefaultBehavior: "DENY-ALL",
			AdditionalProperties: map[string]types.AdditionalPropertyUnchanged{
				"unused": {Unchanged: &addPropsUnch1},
			},
		},
		{
			SliceId:             slice2ID,
			Description:         &slice2Desc,
			ConnectivityService: &cs5g,
			AdditionalProperties: map[string]types.AdditionalPropertyUnchanged{
				"unused": {Unchanged: &addPropsUnch2},
			},
		},
	}

	jsonObj := types.Site{
		Slice:  &slices,
		SiteId: "site-1",
		AdditionalProperties: map[string]types.AdditionalPropertyEnterpriseId{
			"enterprise-id": {
				EnterpriseId: &enterpriseTest,
			},
		},
	}

	gnmiUpdates, err := EncodeToGnmiSite(&jsonObj, false, false, types.EnterpriseId(enterpriseTest), "/site/{site-id}", "site-1")
	assert.NilError(t, err)
	assert.Equal(t, 7, len(gnmiUpdates))
	for _, gnmiUpdate := range gnmiUpdates {
		switch path := strings.ReplaceAll(gnmiUpdate.String(), "  ", " "); path {
		case
			`path:{elem:{name:"site" key:{key:"site-id" value:"site-1"}} elem:{name:"slice" key:{key:"slice-id" value:"slice1"}} elem:{name:"default-behavior"} target:"enterprise-test"} val:{string_val:"DENY-ALL"}`,
			`path:{elem:{name:"site" key:{key:"site-id" value:"site-1"}} elem:{name:"slice" key:{key:"slice-id" value:"slice1"}} elem:{name:"description"} target:"enterprise-test"} val:{string_val:"Slice1 Desc"}`,
			`path:{elem:{name:"site" key:{key:"site-id" value:"site-1"}} elem:{name:"slice" key:{key:"slice-id" value:"slice1"}} elem:{name:"slice-id"} target:"enterprise-test"} val:{string_val:"slice1"}`,
			// and the 2nd one
			`path:{elem:{name:"site" key:{key:"site-id" value:"site-1"}} elem:{name:"slice" key:{key:"slice-id" value:"slice2"}} elem:{name:"description"} target:"enterprise-test"} val:{string_val:"Slice2 Desc"}`,
			`path:{elem:{name:"site" key:{key:"site-id" value:"site-1"}} elem:{name:"slice" key:{key:"slice-id" value:"slice2"}} elem:{name:"connectivity-service"} target:"enterprise-test"} val:{string_val:"5g"}`,
			`path:{elem:{name:"site" key:{key:"site-id" value:"site-1"}} elem:{name:"slice" key:{key:"slice-id" value:"slice2"}} elem:{name:"slice-id"} target:"enterprise-test"} val:{string_val:"slice2"}`,
			`path:{elem:{name:"site" key:{key:"site-id" value:"site-1"}} elem:{name:"site-id"} target:"enterprise-test"} val:{string_val:"site-1"}`:

		default:
			t.Fatalf("unexpected: %s", path)
		}
	}
}
