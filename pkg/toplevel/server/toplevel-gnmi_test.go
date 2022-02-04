// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"context"
	"gotest.tools/assert"
	"testing"
)

func TestGnmiPachAetherRocApi_wrongFormat(t *testing.T) {

	server := &ServerImpl{
		GnmiClient:    nil,
		Authorization: false,
	}

	body := []byte(`{"foo":"bar"}`)
	_, err := server.gnmiPatchAetherRocAPI(context.Background(), body, "")
	assert.Error(t, err, `unable to unmarshal JSON as types.PatchBody: json: unknown field "foo"`)
}

func TestGnmiPachAetherRocApi_wrongFormat2(t *testing.T) {

	server := &ServerImpl{
		GnmiClient:    nil,
		Authorization: false,
	}

	body := []byte(`{"Updates":{}}`)
	_, err := server.gnmiPatchAetherRocAPI(context.Background(), body, "")
	assert.Error(t, err, `unable to convert types.PatchBody to gNMI default-target cannot be blank`)
}
