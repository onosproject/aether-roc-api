// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package server

import (
	"github.com/onosproject/config-models/modelplugin/aether-1.0.0/aether_1_0_0"
	"gotest.tools/assert"
	"testing"
)

func Test_handlePropListAetherV100targetAccessProfile(t *testing.T) {

	testMpd := ModelPluginDevice{
		device: aether_1_0_0.Device{
			AccessProfile: &aether_1_0_0.AccessProfile_AccessProfile{
				AccessProfile: make(map[string]*aether_1_0_0.AccessProfile_AccessProfile_AccessProfile),
			},
		},
	}

	desc := "test description"
	filter := "test filter"
	id := "test-id"
	testType := "testType"
	testMpd.device.AccessProfile.AccessProfile[id] =
		&aether_1_0_0.AccessProfile_AccessProfile_AccessProfile{
			Description: &desc,
			Filter:      &filter,
			Id:          &id,
			Type:        &testType,
		}
	jsonObj, err := testMpd.handlePropListAetherV100targetAccessProfile()
	assert.NilError(t, err)
	assert.Equal(t, 1, len(*jsonObj.ListAetherV100targetAccessProfileAccessProfile))
}
