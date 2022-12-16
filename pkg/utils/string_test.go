// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"gotest.tools/assert"
	"testing"
)

func Test_ToCamelCase(t *testing.T) {
	assert.Equal(t, "Cont1aName", ToCamelCase("cont1a-name"))
	assert.Equal(t, "Cont1Name2a", ToCamelCase("cont1-name2a"))
	assert.Equal(t, "Cont1aname", ToCamelCase("cont1aname"))
	assert.Equal(t, "Cont1234aname", ToCamelCase("cont1234aname"))
	assert.Equal(t, "Cont12aname32", ToCamelCase("cont12aname32"))
	assert.Equal(t, "ConT1Name", ToCamelCase("conT1Name"))
	assert.Equal(t, "ConT1b2AName", ToCamelCase("conT1b2AName"))
}
