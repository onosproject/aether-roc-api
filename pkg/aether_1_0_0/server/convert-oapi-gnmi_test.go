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

// test the addition of the whole Access Profile tree from the top
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

	gnmiUpdates, err := encodeToGnmiAetherV100targetAccessProfile(jsonObj, "")
	assert.NilError(t, err)
	assert.Equal(t, 4, len(gnmiUpdates))
	for _, update := range gnmiUpdates {
		assert.Equal(t, 2, len(update.Path.Elem), "expecting each update to have path len=2")
		pathElem0 := update.Path.Elem[0]
		assert.Equal(t, "access-profile", pathElem0.Name)
		assert.Equal(t, 1, len(pathElem0.Key), "expecting 1 key on access-profile")
		keyVal, ok := pathElem0.Key["id"]
		assert.Assert(t, ok, "expected key to have an 'id'")
		assert.Equal(t, id, keyVal, "expecting key val")

		pathElem1 := update.Path.Elem[1]
		assert.Equal(t, 0, len(pathElem1.Key)) // Should not have key
		switch pathElem1.Name {
		case "id":
			assert.Equal(t, id, update.Val.GetStringVal())
		case "description":
			assert.Equal(t, desc, update.Val.GetStringVal())
		case "filter":
			assert.Equal(t, filter, update.Val.GetStringVal())
		case "type":
			assert.Equal(t, testType, update.Val.GetStringVal())
		default:
			t.Fatalf("unexpected element name %s", pathElem1.Name)
		}
	}
}

// test the update of a single Access Profile
func Test_encodeToGnmiAetherV100targetAccessProfileAccessProfile(t *testing.T) {
	testType := "testType2"
	filter := "test filter 2"

	ap2 := types.AetherV100targetAccessProfileAccessProfile{
		Type:   &testType,
		Filter: &filter,
	}

	gnmiUpdates, err := encodeToGnmiAetherV100targetAccessProfileAccessProfile(&ap2, "")
	assert.NilError(t, err)
	assert.Equal(t, 2, len(gnmiUpdates))
	for _, update := range gnmiUpdates {
		assert.Equal(t, 1, len(update.Path.Elem), "expecting each update to have path len=1")
		pathElem0 := update.Path.Elem[0]
		assert.Equal(t, 0, len(pathElem0.Key)) // Should not have key
		switch pathElem0.Name {
		case "filter":
			assert.Equal(t, filter, update.Val.GetStringVal())
		case "type":
			assert.Equal(t, testType, update.Val.GetStringVal())
		default:
			t.Fatalf("unexpected element name %s", pathElem0.Name)
		}
	}
}
