// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package server

import (
	"fmt"
	"github.com/onosproject/aether-roc-api/pkg/aether_1_0_0/types"
	"github.com/onosproject/aether-roc-api/pkg/utils"
	"github.com/onosproject/config-models/modelplugin/aether-1.0.0/aether_1_0_0"
	"github.com/openconfig/gnmi/proto/gnmi"
)

func encodeToGnmiAetherV100targetAccessProfile(
	jsonObj *types.AetherV100targetAccessProfile, params ...string) (
	[]*gnmi.Update, error) {

	updates := make([]*gnmi.Update, 0)
	if jsonObj.ListAetherV100targetAccessProfileAccessProfile == nil || len(*jsonObj.ListAetherV100targetAccessProfileAccessProfile) == 0 {
		return nil, fmt.Errorf("nothing to commit - please add a profile")
	}
	for _, ap := range *jsonObj.ListAetherV100targetAccessProfileAccessProfile {
		ap := ap //Pinning
		apUpdates, err := encodeToGnmiAetherV100targetAccessProfileAccessProfile(&ap, *ap.Id)
		if err != nil {
			return nil, err
		}
		updates = append(updates, apUpdates...)
	}

	return updates, nil
}

func encodeToGnmiAetherV100targetAccessProfileAccessProfile(
	jsonObj *types.AetherV100targetAccessProfileAccessProfile, params ...string) (
	[]*gnmi.Update, error) {
	updates := make([]*gnmi.Update, 0)

	if len(params) < 1 || params[0] == "" {
		return nil, fmt.Errorf("error id is empty")
	}
	if jsonObj.Id != nil && *jsonObj.Id != params[0] {
		return nil, fmt.Errorf("error id in body is different to param %s != %s", *jsonObj.Id, params[0])
	}
	updateID, err := utils.UpdateForElement(aether_1_0_0.AccessProfile_AccessProfile_AccessProfile{Id: &params[0]}.Id, "/access-profile/id")
	if err != nil {
		return nil, err
	}
	updates = append(updates, updateID)

	if jsonObj.Description != nil {
		updateDesc, err := utils.UpdateForElement(aether_1_0_0.AccessProfile_AccessProfile_AccessProfile{Description: jsonObj.Description}.Description, "/access-profile/description")
		if err != nil {
			return nil, err
		}
		updates = append(updates, updateDesc)
	}

	if jsonObj.Type != nil {
		updateType, err := utils.UpdateForElement(aether_1_0_0.AccessProfile_AccessProfile_AccessProfile{Type: jsonObj.Type}.Type, "/access-profile/type")
		if err != nil {
			return nil, err
		}
		updates = append(updates, updateType)
	}

	if jsonObj.Filter != nil {
		updateFilter, err := utils.UpdateForElement(aether_1_0_0.AccessProfile_AccessProfile_AccessProfile{Filter: jsonObj.Filter}.Filter, "/access-profile/type")
		if err != nil {
			return nil, err
		}
		updates = append(updates, updateFilter)
	}

	return updates, nil
}

func encodeToGnmiAetherV100targetApnProfile(
	jsonObj *types.AetherV100targetApnProfile, params ...string) (
	[]*gnmi.Update, error) {

	return nil, fmt.Errorf("not yet implemented")
}

func encodeToGnmiAetherV100targetApnProfileApnProfile(
	jsonObj *types.AetherV100targetApnProfileApnProfile, params ...string) (
	[]*gnmi.Update, error) {

	return nil, fmt.Errorf("not yet implemented")
}

func encodeToGnmiAetherV100targetQosProfile(
	jsonObj *types.AetherV100targetQosProfile, params ...string) (
	[]*gnmi.Update, error) {

	return nil, fmt.Errorf("not yet implemented")
}

func encodeToGnmiAetherV100targetQosProfileQosProfile(
	jsonObj *types.AetherV100targetQosProfileQosProfile, params ...string) (
	[]*gnmi.Update, error) {

	return nil, fmt.Errorf("not yet implemented")
}

func encodeToGnmiAetherV100targetQosProfileQosProfileApnAmbr(
	jsonObj *types.AetherV100targetQosProfileQosProfileApnAmbr, params ...string) (
	[]*gnmi.Update, error) {

	return nil, fmt.Errorf("not yet implemented")
}

func encodeToGnmiAetherV100targetSubscriber(
	jsonObj *types.AetherV100targetSubscriber, params ...string) (
	[]*gnmi.Update, error) {

	return nil, fmt.Errorf("not yet implemented")
}

func encodeToGnmiAetherV100targetSubscriberUe(
	jsonObj *types.AetherV100targetSubscriberUe, params ...string) (
	[]*gnmi.Update, error) {

	return nil, fmt.Errorf("not yet implemented")
}

func encodeToGnmiAetherV100targetSubscriberUeProfiles(
	jsonObj *types.AetherV100targetSubscriberUeProfiles, params ...string) (
	[]*gnmi.Update, error) {

	return nil, fmt.Errorf("not yet implemented")
}

func encodeToGnmiAetherV100targetSubscriberUeProfilesAccessProfile(
	jsonObj *types.AetherV100targetSubscriberUeProfilesAccessProfile, params ...string) (
	[]*gnmi.Update, error) {

	return nil, fmt.Errorf("not yet implemented")
}

func encodeToGnmiAetherV100targetSubscriberUeServingPlmn(
	jsonObj *types.AetherV100targetSubscriberUeServingPlmn, params ...string) (
	[]*gnmi.Update, error) {

	return nil, fmt.Errorf("not yet implemented")
}

func encodeToGnmiAetherV100targetUpProfile(
	jsonObj *types.AetherV100targetUpProfile, params ...string) (
	[]*gnmi.Update, error) {

	return nil, fmt.Errorf("not yet implemented")
}

func encodeToGnmiAetherV100targetUpProfileUpProfile(
	jsonObj *types.AetherV100targetUpProfileUpProfile, params ...string) (
	[]*gnmi.Update, error) {

	return nil, fmt.Errorf("not yet implemented")
}
