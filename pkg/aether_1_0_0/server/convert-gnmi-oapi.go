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

func (d *ModelPluginDevice) toAetherV100targetAccessProfile(params ...string) (*types.AetherV100targetAccessProfile, error) {
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

func (d *ModelPluginDevice) toAetherV100targetAccessProfileAccessProfile(params ...string) (*types.AetherV100targetAccessProfileAccessProfile, error) {
	ap, ok := d.device.AccessProfile.AccessProfile[params[0]]
	if !ok {
		return nil, fmt.Errorf("error. could not find access-profile %s", params[0])
	}
	jsonObj := new(types.AetherV100targetAccessProfileAccessProfile)
	jsonObj.Id = ap.Id
	if ap.Description != nil {
		jsonObj.Description = ap.Description
	}
	if ap.Type != nil {
		jsonObj.Type = ap.Type
	}
	if ap.Filter != nil {
		jsonObj.Filter = ap.Filter
	}

	return jsonObj, nil
}

func (d *ModelPluginDevice) toAetherV100targetApnProfile(params ...string) (*types.AetherV100targetApnProfile, error) {
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

func (d *ModelPluginDevice) toAetherV100targetApnProfileApnProfile(params ...string) (*types.AetherV100targetApnProfileApnProfile, error) {
	return nil, fmt.Errorf("toAetherV100targetApnProfileApnProfile() not yet implemented")
}

func (d *ModelPluginDevice) toAetherV100targetQosProfile(params ...string) (*types.AetherV100targetQosProfile, error) {
	jsonObj := new(types.AetherV100targetQosProfile)
	jsonList := make([]types.AetherV100targetQosProfileQosProfile, 0)
	jsonObj.ListAetherV100targetQosProfileQosProfile = &jsonList

	for id, qosp := range d.device.QosProfile.QosProfile {
		var uplinkInt int32
		if qosp.ApnAmbr.Uplink != nil {
			uplinkInt = int32(*qosp.ApnAmbr.Uplink)
		}
		var downlinkInt int32
		if qosp.ApnAmbr.Uplink != nil {
			downlinkInt = int32(*qosp.ApnAmbr.Downlink)
		}
		id := id // Pinning
		jsonApn := types.AetherV100targetQosProfileQosProfile{
			Description: qosp.Description,
			Id:          &id,
			AetherV100targetQosProfileQosProfileApnAmbr: &types.AetherV100targetQosProfileQosProfileApnAmbr{
				Downlink: &uplinkInt,
				Uplink:   &downlinkInt,
			},
		}
		jsonList = append(jsonList, jsonApn)
	}

	return jsonObj, nil
}

func (d *ModelPluginDevice) toAetherV100targetQosProfileQosProfile(params ...string) (*types.AetherV100targetQosProfileQosProfile, error) {
	return nil, fmt.Errorf("toAetherV100targetQosProfileQosProfile() not yet implemented")
}

func (d *ModelPluginDevice) toAetherV100targetQosProfileQosProfileApnAmbr(params ...string) (*types.AetherV100targetQosProfileQosProfileApnAmbr, error) {
	qosp, ok := d.device.QosProfile.QosProfile[params[0]]
	if !ok {
		return nil, fmt.Errorf("error. QOS profile %s not found", params[0])
	}

	apnAmbr := new(types.AetherV100targetQosProfileQosProfileApnAmbr)
	var uplinkInt int32
	if qosp.ApnAmbr.Uplink != nil {
		uplinkInt = int32(*qosp.ApnAmbr.Uplink)
		apnAmbr.Uplink = &uplinkInt
	}
	if qosp.ApnAmbr.Uplink != nil {
		downlinkInt := int32(*qosp.ApnAmbr.Downlink)
		apnAmbr.Downlink = &downlinkInt
	}

	return apnAmbr, nil
}

func (d *ModelPluginDevice) toAetherV100targetSubscriber(params ...string) (*types.AetherV100targetSubscriber, error) {
	jsonObj := new(types.AetherV100targetSubscriber)
	jsonList := make([]types.AetherV100targetSubscriberUe, 0)
	jsonObj.ListAetherV100targetSubscriberUe = &jsonList

	for id := range d.device.Subscriber.Ue {
		id := id // Pinning
		updates, _ := d.toAetherV100targetSubscriberUe(id)
		jsonList = append(jsonList, *updates)
	}
	return jsonObj, nil
}

func (d *ModelPluginDevice) toAetherV100targetSubscriberUe(params ...string) (*types.AetherV100targetSubscriberUe, error) {
	ue, ok := d.device.Subscriber.Ue[params[0]]
	if !ok {
		return nil, fmt.Errorf("error. could not find up-profile %s", params[0])
	}
	jsonObj := new(types.AetherV100targetSubscriberUe)
	jsonObj.Ueid = ue.Ueid
	if ue.RequestedApn != nil {
		jsonObj.RequestedApn = ue.RequestedApn
	}
	if ue.Priority != nil {
		var prio = int32(*ue.Priority)
		jsonObj.Priority = &prio
	}
	if ue.Enabled != nil {
		jsonObj.Enabled = ue.Enabled
	}
	if ue.ServingPlmn != nil {
		splmn, _ := d.toAetherV100targetSubscriberUeServingPlmn(params...)
		jsonObj.AetherV100targetSubscriberUeServingPlmn = splmn
	}

	if ue.Profiles != nil {
		profiles, _ := d.toAetherV100targetSubscriberUeProfiles(params...)
		jsonObj.AetherV100targetSubscriberUeProfiles = profiles
	}

	return jsonObj, nil
}

func (d *ModelPluginDevice) toAetherV100targetSubscriberUeProfiles(params ...string) (*types.AetherV100targetSubscriberUeProfiles, error) {
	ue, ok := d.device.Subscriber.Ue[params[0]]
	if !ok {
		return nil, fmt.Errorf("error. Subscriber Ue %s not found", params[0])
	}
	profiles := new(types.AetherV100targetSubscriberUeProfiles)

	if ue.Profiles.ApnProfile != nil {
		profiles.ApnProfile = ue.Profiles.ApnProfile
	}
	if ue.Profiles.UpProfile != nil {
		profiles.UpProfile = ue.Profiles.UpProfile
	}
	if ue.Profiles.QosProfile != nil {
		profiles.QosProfile = ue.Profiles.QosProfile
	}
	if ue.Profiles.AccessProfile != nil {
		jsonList := make([]types.AetherV100targetSubscriberUeProfilesAccessProfile, 0)
		profiles.ListAetherV100targetSubscriberUeProfilesAccessProfile = &jsonList
		for id := range ue.Profiles.AccessProfile {
			aps, _ := d.toAetherV100targetSubscriberUeProfilesAccessProfile(append(params, id)...)
			jsonList = append(jsonList, *aps)
		}
	}

	return profiles, nil
}

func (d *ModelPluginDevice) toAetherV100targetSubscriberUeProfilesAccessProfile(params ...string) (*types.AetherV100targetSubscriberUeProfilesAccessProfile, error) {
	ue, ok := d.device.Subscriber.Ue[params[0]]
	if !ok {
		return nil, fmt.Errorf("error. Subscriber Ue %s not found", params[0])
	}
	ap, ok := ue.Profiles.AccessProfile[params[1]]
	if !ok {
		return nil, fmt.Errorf("error. Subscriber Ue Profiles Access Profile %s not found", params[0])
	}
	jsonAp := types.AetherV100targetSubscriberUeProfilesAccessProfile{
		AccessProfile: ap.AccessProfile,
		Allowed:       ap.Allowed,
	}
	return &jsonAp, nil
}

func (d *ModelPluginDevice) toAetherV100targetSubscriberUeServingPlmn(params ...string) (*types.AetherV100targetSubscriberUeServingPlmn, error) {
	ue, ok := d.device.Subscriber.Ue[params[0]]
	if !ok {
		return nil, fmt.Errorf("error. Subscriber Ue %s not found", params[0])
	}
	splmn := new(types.AetherV100targetSubscriberUeServingPlmn)

	if ue.ServingPlmn.Mcc != nil {
		mcc := int32(*ue.ServingPlmn.Mcc)
		splmn.Mcc = &mcc
	}
	if ue.ServingPlmn.Mnc != nil {
		mnc := int32(*ue.ServingPlmn.Mnc)
		splmn.Mnc = &mnc
	}
	if ue.ServingPlmn.Tac != nil {
		tac := int32(*ue.ServingPlmn.Tac)
		splmn.Tac = &tac
	}

	return splmn, nil
}

func (d *ModelPluginDevice) toAetherV100targetUpProfile(params ...string) (*types.AetherV100targetUpProfile, error) {
	jsonObj := new(types.AetherV100targetUpProfile)
	jsonList := make([]types.AetherV100targetUpProfileUpProfile, 0)
	jsonObj.ListAetherV100targetUpProfileUpProfile = &jsonList

	for id := range d.device.UpProfile.UpProfile {
		id := id // Pinning
		updates, _ := d.toAetherV100targetUpProfileUpProfile(id)
		jsonList = append(jsonList, *updates)
	}
	return jsonObj, nil
}

func (d *ModelPluginDevice) toAetherV100targetUpProfileUpProfile(params ...string) (*types.AetherV100targetUpProfileUpProfile, error) {
	up, ok := d.device.UpProfile.UpProfile[params[0]]
	if !ok {
		return nil, fmt.Errorf("error. could not find up-profile %s", params[0])
	}
	jsonObj := new(types.AetherV100targetUpProfileUpProfile)
	jsonObj.Id = up.Id
	if up.Description != nil {
		jsonObj.Description = up.Description
	}
	if up.UserPlane != nil {
		jsonObj.UserPlane = up.UserPlane
	}
	if up.AccessControl != nil {
		jsonObj.AccessControl = up.AccessControl
	}

	return jsonObj, nil
}

func (d *ModelPluginDevice) toTarget(params ...string) (*types.Target, error) {
	return nil, fmt.Errorf("toTarget() should not be called directly")
}
