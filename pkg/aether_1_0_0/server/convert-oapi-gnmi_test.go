// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package server

import (
	"github.com/onosproject/aether-roc-api/pkg/aether_1_0_0/types"
	"gotest.tools/assert"
	"testing"
)

func Test_encodeToGnmiAetherV100targetAccessProfile(t *testing.T) {
	jsonObj := new(types.AetherV100targetAccessProfile)
	jsonList := make([]types.AetherV100targetAccessProfileAccessProfile, 0)
	jsonObj.ListAetherV100targetAccessProfileAccessProfile = &jsonList
	desc := "test description"
	filter := "test filter"
	id := "test-id"
	testType := "testType"

	ap1 := types.AetherV100targetAccessProfileAccessProfile{
		Description: &desc,
		Filter:      &filter,
		Id:          &id,
		Type:        &testType,
	}
	jsonList = append(jsonList, ap1)

	gnmiObj, err := encodeToGnmiAetherV100targetAccessProfile(jsonObj)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(gnmiObj.AccessProfile))
	gnmiAp1 := gnmiObj.AccessProfile[id]
	assert.Equal(t, desc, *gnmiAp1.Description)
	assert.Equal(t, filter, *gnmiAp1.Filter)
	assert.Equal(t, id, *gnmiAp1.Id)
	assert.Equal(t, testType, *gnmiAp1.Type)
}
