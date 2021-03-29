// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package server

import (
	"gotest.tools/assert"
	"os"
	"testing"
)

func Test_TopLevelSpec(t *testing.T) {
	err := os.Chdir("../../../api")
	assert.NilError(t, err)

	swagger, err := GetSwagger()
	assert.NilError(t, err)

	assert.Assert(t, swagger != nil)
}
