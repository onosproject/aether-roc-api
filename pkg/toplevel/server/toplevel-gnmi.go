// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
//

package server

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/onosproject/aether-roc-api/pkg/toplevel/types"
)

import (
	"github.com/onosproject/aether-roc-api/pkg/utils"
)

// gnmiPatchAetherRocAPI patches an existing configuration with PatchBody.
func (i *TopLevelServer) gnmiPatchAetherRocAPI(ctx context.Context, body []byte, dummy string) (*string, error) {

	var jsonObj types.PatchBody
	dec := json.NewDecoder(bytes.NewReader(body))
	dec.DisallowUnknownFields() // Force errors

	if err := dec.Decode(&jsonObj); err != nil {
		return nil, fmt.Errorf("unable to unmarshal JSON as types.PatchBody: %s", err.Error())
	}

	patchBody, err := encodeToGnmiPatchBody(&jsonObj)
	if err != nil {
		return nil, err
	}
	gnmiSet, err := utils.NewGnmiSetRequest(patchBody.Updates, patchBody.Deletes,
		patchBody.Ext100Name, patchBody.Ext101Version, patchBody.Ext102Type, patchBody.Ext111Strategy)
	if err != nil {
		return nil, err
	}
	log.Infof("gnmiSetRequest %s", gnmiSet.String())
	gnmiSetResponse, err := i.GnmiClient.Set(ctx, gnmiSet)
	if err != nil {
		return nil, fmt.Errorf(" %v", err)
	}
	return utils.ExtractResponseID(gnmiSetResponse)
}
