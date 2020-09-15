// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package server

import (
	"fmt"
	"github.com/onosproject/aether-roc-api/pkg/aether_1_0_0/types"
	modelplugin "github.com/onosproject/config-models/modelplugin/aether-1.0.0/aether_1_0_0"
)

// ModelPluginDevice - a wrapper for the model plugin
type ModelPluginDevice struct {
	device modelplugin.Device
}

func (d *ModelPluginDevice) handlePropListAetherV100targetAccessProfile() (*types.AetherV100targetAccessProfile, error) {
	jsonObj := new(types.AetherV100targetAccessProfile)
	jsonList := make([]types.AetherV100targetAccessProfileAccessProfile, 0)
	jsonObj.ListAetherV100targetAccessProfileAccessProfile = &jsonList

	for id, ap := range d.device.AccessProfile.AccessProfile {
		id := id // Pinning
		jsonAp := types.AetherV100targetAccessProfileAccessProfile{
			Description: ap.Description,
			Filter:      ap.Filter,
			Id:          &id,
			Type:        ap.Type,
		}
		jsonList = append(jsonList, jsonAp)
	}
	return jsonObj, nil
}

func (d *ModelPluginDevice) handlePropListAetherV100targetAccessProfileAccessProfile() (*types.AetherV100targetAccessProfileAccessProfile, error) {
	return nil, fmt.Errorf("handlePropListAetherV100targetAccessProfileAccessProfile() not yet implemented")
}

func (d *ModelPluginDevice) handlePropListAetherV100targetApnProfile() (*types.AetherV100targetApnProfile, error) {
	jsonObj := new(types.AetherV100targetApnProfile)
	jsonList := make([]types.AetherV100targetApnProfileApnProfile, 0)
	jsonObj.ListAetherV100targetApnProfileApnProfile = &jsonList

	for id, apnp := range d.device.ApnProfile.ApnProfile {
		var mtuInt int32
		if apnp.Mtu != nil {
			mtuInt = int32(*apnp.Mtu)
		}
		id := id // Pinning
		jsonApn := types.AetherV100targetApnProfileApnProfile{
			ApnName:      apnp.ApnName,
			Description:  apnp.Description,
			DnsPrimary:   apnp.DnsPrimary,
			DnsSecondary: apnp.DnsSecondary,
			GxEnabled:    apnp.GxEnabled,
			Id:           &id,
			Mtu:          &mtuInt,
		}
		jsonList = append(jsonList, jsonApn)
	}

	return jsonObj, nil
}

func (d *ModelPluginDevice) handlePropListAetherV100targetApnProfileApnProfile() (*types.AetherV100targetApnProfileApnProfile, error) {
	return nil, fmt.Errorf("handlePropListAetherV100targetApnProfileApnProfile() not yet implemented")
}

func (d *ModelPluginDevice) handlePropListAetherV100targetQosProfile() (*types.AetherV100targetQosProfile, error) {
	return nil, fmt.Errorf("handlePropListAetherV100targetQosProfile() not yet implemented")
}

func (d *ModelPluginDevice) handlePropListAetherV100targetQosProfileQosProfile() (*types.AetherV100targetQosProfileQosProfile, error) {
	return nil, fmt.Errorf("handlePropListAetherV100targetQosProfileQosProfile() not yet implemented")
}

func (d *ModelPluginDevice) handlePropListAetherV100targetQosProfileQosProfileApnAmbr() (*types.AetherV100targetQosProfileQosProfileApnAmbr, error) {
	return nil, fmt.Errorf("handlePropListAetherV100targetQosProfileQosProfileApnAmbr() not yet implemented")
}

func (d *ModelPluginDevice) handlePropListAetherV100targetSubscriber() (*types.AetherV100targetSubscriber, error) {
	return nil, fmt.Errorf("handlePropListAetherV100targetSubscriber() not yet implemented")
}

func (d *ModelPluginDevice) handlePropListAetherV100targetSubscriberUe() (*types.AetherV100targetSubscriberUe, error) {
	return nil, fmt.Errorf("handlePropListAetherV100targetSubscriberUe() not yet implemented")
}

func (d *ModelPluginDevice) handlePropListAetherV100targetSubscriberUeProfiles() (*types.AetherV100targetSubscriberUeProfiles, error) {
	return nil, fmt.Errorf("handlePropListAetherV100targetSubscriberUeProfiles() not yet implemented")
}

func (d *ModelPluginDevice) handlePropListAetherV100targetSubscriberUeProfilesAccessProfile() (*types.AetherV100targetSubscriberUeProfilesAccessProfile, error) {
	return nil, fmt.Errorf("handlePropListAetherV100targetSubscriberUeProfilesAccessProfile() not yet implemented")
}

func (d *ModelPluginDevice) handlePropListAetherV100targetSubscriberUeServingPlmn() (*types.AetherV100targetSubscriberUeServingPlmn, error) {
	return nil, fmt.Errorf("handlePropListAetherV100targetSubscriberUeServingPlmn() not yet implemented")
}

func (d *ModelPluginDevice) handlePropListAetherV100targetUpProfile() (*types.AetherV100targetUpProfile, error) {
	return nil, fmt.Errorf("handlePropListAetherV100targetUpProfile() not yet implemented")
}

func (d *ModelPluginDevice) handlePropListAetherV100targetUpProfileUpProfile() (*types.AetherV100targetUpProfileUpProfile, error) {
	return nil, fmt.Errorf("handlePropListAetherV100targetUpProfileUpProfile() not yet implemented")
}

func (d *ModelPluginDevice) handlePropListTarget() (*types.Target, error) {
	return nil, fmt.Errorf("handlePropListTarget() should not be called directly")
}
