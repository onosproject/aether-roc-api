// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package roc

import (
	"github.com/onosproject/aether-roc-api/test/utils"
	"github.com/onosproject/helmit/pkg/input"
	"github.com/onosproject/helmit/pkg/test"
)

// TestSuite is the primary aether-roc-api test suite
type TestSuite struct {
	test.Suite
	c *input.Context
}

// SetupTestSuite sets up the onos-e2t test suite
func (s *TestSuite) SetupTestSuite(c *input.Context) error {
	s.c = c

	roc, err := utils.CreateRocRelease(c)
	if err != nil {
		return err
	}

	registry := c.GetArg("registry").String("")

	return roc.Set("global.image.registry", registry).Install(true)
}
