// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package server

import (
	"context"
	"encoding/json"
	"fmt"
)

import (
	"github.com/onosproject/aether-roc-api/pkg/toplevel/types"
	"github.com/onosproject/aether-roc-api/pkg/utils"
)

// gnmiPatchAetherRocAPI patches an existing configuration with PatchBody.
func (i *ServerImpl) gnmiPatchAetherRocAPI(ctx context.Context, body []byte,
	openAPIPath string, args ...string) (*string, error) {

	jsonObj := new(types.PatchBody)
	if err := json.Unmarshal(body, jsonObj); err != nil {
		return nil, fmt.Errorf("unable to unmarshal JSON as types.PatchBody %v", err)
	}
	gnmiUpdates, gnmiDeletes, err := encodeToGnmiPatchBody(jsonObj)
	if err != nil {
		return nil, fmt.Errorf("unable to convert types.PatchBody to gNMI %v", err)
	}
	gnmiSet, err := utils.NewGnmiSetRequest(gnmiUpdates, gnmiDeletes)
	if err != nil {
		return nil, err
	}
	log.Infof("gnmiSetRequest %s", gnmiSet.String())
	gnmiSetResponse, err := i.GnmiClient.Set(ctx, gnmiSet)
	if err != nil {
		return nil, fmt.Errorf(" %v", err)
	}
	return utils.ExtractExtension100(gnmiSetResponse), nil
}
