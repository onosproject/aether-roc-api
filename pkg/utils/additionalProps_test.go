// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
//

package utils

import (
	"gotest.tools/assert"
	"testing"
)

type AdditionalPropertyTarget struct {
	Target *string
}

type TestObj struct {
	AdditionalProperties map[string]AdditionalPropertyTarget
}

type AdditionalPropertyUnchanged struct {
	Unchanged *string
}

type TestObj2 struct {
	AdditionalProperties map[string]AdditionalPropertyUnchanged
}

type AdditionalPropertiesUnchTarget struct {
	Unchanged *string
	Target    *string
}

type TestObj3 struct {
	AdditionalProperties map[string]AdditionalPropertiesUnchTarget
}

func Test_CheckForAdditionalProps(t *testing.T) {
	t1 := "t1"
	jsonObj := &TestObj{
		AdditionalProperties: map[string]AdditionalPropertyTarget{
			"additionalProperty": {
				Target: &t1,
			},
		},
	}

	unchanged, target := CheckForAdditionalProps(jsonObj)
	assert.Equal(t, "t1", *target)
	assert.Equal(t, 0, len(unchanged))
}

func Test_CheckForAdditionalPropsUnchanged(t *testing.T) {
	uc := "uc1,uc2"
	jsonObj := &TestObj2{
		AdditionalProperties: map[string]AdditionalPropertyUnchanged{
			"additionalProperty": {
				Unchanged: &uc,
			},
		},
	}

	unchanged, target := CheckForAdditionalProps(jsonObj)
	assert.Assert(t, target == nil)
	assert.Equal(t, 2, len(unchanged))
	_, uc1Present := unchanged["uc1"]
	assert.Equal(t, true, uc1Present)
	_, uc2Present := unchanged["uc2"]
	assert.Equal(t, true, uc2Present)
}

func Test_CheckForAdditionalPropsUnchanged_Empty(t *testing.T) {
	jsonObj := &TestObj3{
		AdditionalProperties: map[string]AdditionalPropertiesUnchTarget{
			"additionalProperty": {
				Unchanged: nil,
				Target:    nil,
			},
		},
	}

	unchanged, target := CheckForAdditionalProps(jsonObj)
	assert.Assert(t, target == nil)
	assert.Equal(t, 0, len(unchanged))
}

func Test_CheckForAdditionalPropsUnchangedTarget_Empty(t *testing.T) {
	jsonObj := &TestObj2{
		AdditionalProperties: map[string]AdditionalPropertyUnchanged{
			"additionalProperty": {
				Unchanged: nil,
			},
		},
	}

	unchanged, target := CheckForAdditionalProps(jsonObj)
	assert.Assert(t, target == nil)
	assert.Equal(t, 0, len(unchanged))
}
