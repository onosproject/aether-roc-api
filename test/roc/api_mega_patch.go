// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package roc

import (
	"github.com/onosproject/aether-roc-api/test/utils/config"
	"github.com/onosproject/aether-roc-api/test/utils/proto"

	"testing"
)

const (
	sdCoreV3Device = "connectivity-service-v3"
	csPath         = "/connectivity-service/connectivity-service[id=cs5gtest]/display-name"
	csValue        = "ROC 5G Test Connectivity Service"
)

// TestRocMegaPatch that the MEGA Patch can be applied through the API with no error.
func (s *TestSuite) TestRocMegaPatch(t *testing.T) {

	// Make a GNMI client to use for requests
	gnmiClient := config.GetGNMIClientOrFail(t)

	devicePath := config.GetDevicePathWithValue(sdCoreV3Device, csPath, csValue, proto.StringVal)

	// Check that the value was set correctly
	config.CheckGNMIValue(t, gnmiClient, devicePath, csValue, 0, "Query after set returned the wrong value")

}
