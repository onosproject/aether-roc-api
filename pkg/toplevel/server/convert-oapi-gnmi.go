// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package server

import (
	"fmt"
	externalRef0Svr "github.com/onosproject/aether-roc-api/pkg/aether_2_0_0/server"
	"github.com/onosproject/aether-roc-api/pkg/toplevel/types"
	"github.com/openconfig/gnmi/proto/gnmi"
)

func encodeToGnmiPatchBody(jsonObj *types.PatchBody) ([]*gnmi.Update, []*gnmi.Path, error) {
	updates := make([]*gnmi.Update, 0)
	deletes := make([]*gnmi.Path, 0)

	gnmiUpdates, err := encodeToGnmiElements(jsonObj.Updates)
	if err != nil {
		return nil, nil, fmt.Errorf("encodeToGnmiElements() %s", err.Error())
	}
	updates = append(updates, gnmiUpdates...)

	gnmiDeletes, err := encodeToGnmiElements(jsonObj.Deletes)
	if err != nil {
		return nil, nil, fmt.Errorf("encodeToGnmiElements() %s", err.Error())
	}
	for _, gd := range gnmiDeletes {
		deletes = append(deletes, gd.Path)
	}

	return updates, deletes, nil
}

func encodeToGnmiElements(elements *types.Elements) ([]*gnmi.Update, error) {
	if elements == nil {
		return nil, nil
	}
	updates := make([]*gnmi.Update, 0)

	if elements.AccessProfile200 != nil {
		accessProfileUpdates, err := externalRef0Svr.EncodeToGnmiAccessProfile(
			elements.AccessProfile200, false,
			"/access-profile")
		if err != nil {
			return nil, fmt.Errorf("EncodeToGnmiAccessProfile() %s", err)
		}
		updates = append(updates, accessProfileUpdates...)
	}

	if elements.ApnProfile200 != nil {
		apnProfileUpdates, err := externalRef0Svr.EncodeToGnmiApnProfile(
			elements.ApnProfile200, false,
			"/apn-profile")
		if err != nil {
			return nil, fmt.Errorf("EncodeToGnmiApnProfile() %s", err)
		}
		updates = append(updates, apnProfileUpdates...)
	}

	return updates, nil
}
