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
	jsonObj *types.AetherV100targetAccessProfile, parentPath string, params ...string) (
	[]*gnmi.Update, error) {

	updates := make([]*gnmi.Update, 0)
	if jsonObj.ListAetherV100targetAccessProfileAccessProfile == nil || len(*jsonObj.ListAetherV100targetAccessProfileAccessProfile) == 0 {
		return nil, fmt.Errorf("nothing to commit - please add a profile")
	}
	for _, ap := range *jsonObj.ListAetherV100targetAccessProfileAccessProfile {
		ap := ap //Pinning
		apUpdates, err := encodeToGnmiAetherV100targetAccessProfileAccessProfile(&ap, "/access-profile/{id}", *ap.Id)
		if err != nil {
			return nil, err
		}
		updates = append(updates, apUpdates...)
	}

	return updates, nil
}

func encodeToGnmiAetherV100targetAccessProfileAccessProfile(
	jsonObj *types.AetherV100targetAccessProfileAccessProfile, parentPath string, params ...string) (
	[]*gnmi.Update, error) {
	updates := make([]*gnmi.Update, 0)

	if len(parentPath) > 0 && (len(params) < 1 || params[0] == "") {
		return nil, fmt.Errorf("error: id is empty")
	}
	if jsonObj.Id != nil && *jsonObj.Id != params[0] {
		return nil, fmt.Errorf("error. id in body is different to param %s != %s", *jsonObj.Id, params[0])
	}
	if jsonObj.Id != nil {
		updateID, err := utils.UpdateForElement(aether_1_0_0.AccessProfile_AccessProfile_AccessProfile{Id: &params[0]}.Id,
			fmt.Sprintf("%s/%s", parentPath, "id"), params...)
		if err != nil {
			return nil, err
		}
		updates = append(updates, updateID)
	}

	if jsonObj.Description != nil {
		updateDesc, err := utils.UpdateForElement(aether_1_0_0.AccessProfile_AccessProfile_AccessProfile{Description: jsonObj.Description}.Description,
			fmt.Sprintf("%s/%s", parentPath, "description"), params...)
		if err != nil {
			return nil, err
		}
		updates = append(updates, updateDesc)
	}

	if jsonObj.Type != nil {
		updateType, err := utils.UpdateForElement(aether_1_0_0.AccessProfile_AccessProfile_AccessProfile{Type: jsonObj.Type}.Type,
			fmt.Sprintf("%s/%s", parentPath, "type"), params...)
		if err != nil {
			return nil, err
		}
		updates = append(updates, updateType)
	}

	if jsonObj.Filter != nil {
		updateFilter, err := utils.UpdateForElement(aether_1_0_0.AccessProfile_AccessProfile_AccessProfile{Filter: jsonObj.Filter}.Filter,
			fmt.Sprintf("%s/%s", parentPath, "filter"), params...)
		if err != nil {
			return nil, err
		}
		updates = append(updates, updateFilter)
	}

	return updates, nil
}

func encodeToGnmiAetherV100targetApnProfile(
	jsonObj *types.AetherV100targetApnProfile, parentPath string, params ...string) (
	[]*gnmi.Update, error) {

	updates := make([]*gnmi.Update, 0)
	if jsonObj.ListAetherV100targetApnProfileApnProfile == nil || len(*jsonObj.ListAetherV100targetApnProfileApnProfile) == 0 {
		return nil, fmt.Errorf("nothing to commit - please add a profile")
	}

	for _, apn := range *jsonObj.ListAetherV100targetApnProfileApnProfile {
		apn := apn //Pinning
		apUpdates, err := encodeToGnmiAetherV100targetApnProfileApnProfile(&apn, "/apn-profile/{id}", *apn.Id)
		if err != nil {
			return nil, err
		}
		updates = append(updates, apUpdates...)
	}

	return updates, nil
}

func encodeToGnmiAetherV100targetApnProfileApnProfile(
	jsonObj *types.AetherV100targetApnProfileApnProfile, parentPath string, params ...string) (
	[]*gnmi.Update, error) {
	updates := make([]*gnmi.Update, 0)

	if len(parentPath) > 0 && (len(params) < 1 || params[0] == "") {
		return nil, fmt.Errorf("error: id is empty")
	}
	if jsonObj.Id != nil && *jsonObj.Id != params[0] {
		return nil, fmt.Errorf("error id in body is different to param %s != %s", *jsonObj.Id, params[0])
	}
	if jsonObj.Id != nil {
		updateID, err := utils.UpdateForElement(aether_1_0_0.ApnProfile_ApnProfile_ApnProfile{Id: &params[0]}.Id,
			fmt.Sprintf("%s/%s", parentPath, "id"), params...)
		if err != nil {
			return nil, err
		}
		updates = append(updates, updateID)
	}

	if jsonObj.Description != nil {
		updateDesc, err := utils.UpdateForElement(aether_1_0_0.ApnProfile_ApnProfile_ApnProfile{Description: jsonObj.Description}.Description,
			fmt.Sprintf("%s/%s", parentPath, "description"), params...)
		if err != nil {
			return nil, err
		}
		updates = append(updates, updateDesc)
	}

	if jsonObj.ApnName != nil {
		updateApnName, err := utils.UpdateForElement(aether_1_0_0.ApnProfile_ApnProfile_ApnProfile{ApnName: jsonObj.ApnName}.ApnName,
			fmt.Sprintf("%s/%s", parentPath, "apn-name"), params...)
		if err != nil {
			return nil, err
		}
		updates = append(updates, updateApnName)
	}

	if jsonObj.DnsPrimary != nil {
		updateDNSPrimary, err := utils.UpdateForElement(aether_1_0_0.ApnProfile_ApnProfile_ApnProfile{DnsPrimary: jsonObj.DnsPrimary}.DnsPrimary,
			fmt.Sprintf("%s/%s", parentPath, "dns-primary"), params...)
		if err != nil {
			return nil, err
		}
		updates = append(updates, updateDNSPrimary)
	}

	if jsonObj.DnsSecondary != nil {
		updateDNSSecondary, err := utils.UpdateForElement(aether_1_0_0.ApnProfile_ApnProfile_ApnProfile{DnsSecondary: jsonObj.DnsSecondary}.DnsSecondary,
			fmt.Sprintf("%s/%s", parentPath, "dns-secondary"), params...)
		if err != nil {
			return nil, err
		}
		updates = append(updates, updateDNSSecondary)
	}

	if jsonObj.Mtu != nil {
		var mtu uint32 = uint32(*jsonObj.Mtu)
		updateMtu, err := utils.UpdateForElement(aether_1_0_0.ApnProfile_ApnProfile_ApnProfile{Mtu: &mtu}.Mtu,
			fmt.Sprintf("%s/%s", parentPath, "mtu"), params...)
		if err != nil {
			return nil, err
		}
		updates = append(updates, updateMtu)
	}

	if jsonObj.GxEnabled != nil {
		updateGxEnabled, err := utils.UpdateForElement(aether_1_0_0.ApnProfile_ApnProfile_ApnProfile{GxEnabled: jsonObj.GxEnabled}.GxEnabled,
			fmt.Sprintf("%s/%s", parentPath, "gx-enabled"), params...)
		if err != nil {
			return nil, err
		}
		updates = append(updates, updateGxEnabled)
	}

	return updates, nil
}

func encodeToGnmiAetherV100targetQosProfile(
	jsonObj *types.AetherV100targetQosProfile, parentPath string, params ...string) (
	[]*gnmi.Update, error) {

	updates := make([]*gnmi.Update, 0)
	if jsonObj.ListAetherV100targetQosProfileQosProfile == nil || len(*jsonObj.ListAetherV100targetQosProfileQosProfile) == 0 {
		return nil, fmt.Errorf("nothing to commit - please add a profile")
	}
	for _, ap := range *jsonObj.ListAetherV100targetQosProfileQosProfile {
		ap := ap //Pinning
		apUpdates, err := encodeToGnmiAetherV100targetQosProfileQosProfile(&ap, "/qos-profile/{id}", *ap.Id)
		if err != nil {
			return nil, err
		}
		updates = append(updates, apUpdates...)
	}

	return updates, nil
}

func encodeToGnmiAetherV100targetQosProfileQosProfile(
	jsonObj *types.AetherV100targetQosProfileQosProfile, parentPath string, params ...string) (
	[]*gnmi.Update, error) {

	updates := make([]*gnmi.Update, 0)

	if len(parentPath) > 0 && (len(params) < 1 || params[0] == "") {
		return nil, fmt.Errorf("error: id is empty")
	}
	if jsonObj.Id != nil && *jsonObj.Id != params[0] {
		return nil, fmt.Errorf("error id in body is different to param %s != %s", *jsonObj.Id, params[0])
	}
	if jsonObj.Id != nil {
		updateID, err := utils.UpdateForElement(aether_1_0_0.QosProfile_QosProfile_QosProfile{Id: &params[0]}.Id,
			fmt.Sprintf("%s/%s", parentPath, "id"), params...)
		if err != nil {
			return nil, err
		}
		updates = append(updates, updateID)
	}

	if jsonObj.Description != nil {
		updateDesc, err := utils.UpdateForElement(aether_1_0_0.QosProfile_QosProfile_QosProfile{Description: jsonObj.Description}.Description,
			fmt.Sprintf("%s/%s", parentPath, "description"), params...)
		if err != nil {
			return nil, err
		}
		updates = append(updates, updateDesc)
	}

	if jsonObj.AetherV100targetQosProfileQosProfileApnAmbr != nil {
		updateApnAmbr, err := encodeToGnmiAetherV100targetQosProfileQosProfileApnAmbr(
			jsonObj.AetherV100targetQosProfileQosProfileApnAmbr, fmt.Sprintf("%s/%s", parentPath, "apn-ambr"), params...)
		if err != nil {
			return nil, err
		}
		updates = append(updates, updateApnAmbr...)
	}

	return updates, nil
}

func encodeToGnmiAetherV100targetQosProfileQosProfileApnAmbr(
	jsonObj *types.AetherV100targetQosProfileQosProfileApnAmbr, parentPath string, params ...string) (
	[]*gnmi.Update, error) {

	updates := make([]*gnmi.Update, 0)
	if jsonObj.Uplink != nil {
		var uplink = uint32(*jsonObj.Uplink)
		updateType, err := utils.UpdateForElement(aether_1_0_0.QosProfile_QosProfile_QosProfile_ApnAmbr{Uplink: &uplink}.Uplink,
			fmt.Sprintf("%s/%s", parentPath, "uplink"), params...)
		if err != nil {
			return nil, err
		}
		updates = append(updates, updateType)
	}
	if jsonObj.Downlink != nil {
		var downlink = uint32(*jsonObj.Downlink)
		updateType, err := utils.UpdateForElement(aether_1_0_0.QosProfile_QosProfile_QosProfile_ApnAmbr{Downlink: &downlink}.Downlink,
			fmt.Sprintf("%s/%s", parentPath, "downlink"), params...)
		if err != nil {
			return nil, err
		}
		updates = append(updates, updateType)
	}

	return updates, nil
}

func encodeToGnmiAetherV100targetSubscriber(
	jsonObj *types.AetherV100targetSubscriber, parentPath string, params ...string) (
	[]*gnmi.Update, error) {

	updates := make([]*gnmi.Update, 0)
	if jsonObj.ListAetherV100targetSubscriberUe == nil || len(*jsonObj.ListAetherV100targetSubscriberUe) == 0 {
		return nil, fmt.Errorf("nothing to commit - please add a profile")
	}
	for _, ue := range *jsonObj.ListAetherV100targetSubscriberUe {
		ue := ue //Pinning
		ueUpdates, err := encodeToGnmiAetherV100targetSubscriberUe(&ue, "/ue/{ueid}", *ue.Ueid)
		if err != nil {
			return nil, err
		}
		updates = append(updates, ueUpdates...)
	}

	return updates, nil
}

func encodeToGnmiAetherV100targetSubscriberUe(
	jsonObj *types.AetherV100targetSubscriberUe, parentPath string, params ...string) (
	[]*gnmi.Update, error) {

	updates := make([]*gnmi.Update, 0)

	if len(parentPath) > 0 && (len(params) < 1 || params[0] == "") {
		return nil, fmt.Errorf("error: ueid is empty")
	}
	if jsonObj.Ueid != nil && *jsonObj.Ueid != params[0] {
		return nil, fmt.Errorf("error. ueid in body is different to param %s != %s", *jsonObj.Ueid, params[0])
	}
	if jsonObj.Ueid != nil {
		updateUeID, err := utils.UpdateForElement(aether_1_0_0.AetherSubscriber_Subscriber_Ue{Ueid: &params[0]}.Ueid,
			fmt.Sprintf("%s/%s", parentPath, "ueid"), params...)
		if err != nil {
			return nil, err
		}
		updates = append(updates, updateUeID)
	}

	if jsonObj.RequestedApn != nil {
		updateDesc, err := utils.UpdateForElement(aether_1_0_0.AetherSubscriber_Subscriber_Ue{RequestedApn: jsonObj.RequestedApn}.RequestedApn,
			fmt.Sprintf("%s/%s", parentPath, "requested-apn"), params...)
		if err != nil {
			return nil, err
		}
		updates = append(updates, updateDesc)
	}

	if jsonObj.Priority != nil {
		var priority uint32 = uint32(*jsonObj.Priority)
		updateType, err := utils.UpdateForElement(aether_1_0_0.AetherSubscriber_Subscriber_Ue{Priority: &priority}.Priority,
			fmt.Sprintf("%s/%s", parentPath, "priority"), params...)
		if err != nil {
			return nil, err
		}
		updates = append(updates, updateType)
	}

	if jsonObj.Enabled != nil {
		updateFilter, err := utils.UpdateForElement(aether_1_0_0.AetherSubscriber_Subscriber_Ue{Enabled: jsonObj.Enabled}.Enabled,
			fmt.Sprintf("%s/%s", parentPath, "enabled"), params...)
		if err != nil {
			return nil, err
		}
		updates = append(updates, updateFilter)
	}

	if jsonObj.AetherV100targetSubscriberUeServingPlmn != nil {
		updateServingPlmn, err := encodeToGnmiAetherV100targetSubscriberUeServingPlmn(
			jsonObj.AetherV100targetSubscriberUeServingPlmn, fmt.Sprintf("%s/%s", parentPath, "serving-plmn"), params...)
		if err != nil {
			return nil, err
		}
		updates = append(updates, updateServingPlmn...)
	}

	if jsonObj.AetherV100targetSubscriberUeProfiles != nil {
		updateUeProfiles, err := encodeToGnmiAetherV100targetSubscriberUeProfiles(
			jsonObj.AetherV100targetSubscriberUeProfiles, fmt.Sprintf("%s/%s", parentPath, "profiles"), params...)
		if err != nil {
			return nil, err
		}
		updates = append(updates, updateUeProfiles...)
	}

	return updates, nil
}

func encodeToGnmiAetherV100targetSubscriberUeProfiles(
	jsonObj *types.AetherV100targetSubscriberUeProfiles, parentPath string, params ...string) (
	[]*gnmi.Update, error) {

	updates := make([]*gnmi.Update, 0)
	if jsonObj.ApnProfile != nil {
		updateApnProfile, err := utils.UpdateForElement(aether_1_0_0.AetherSubscriber_Subscriber_Ue_Profiles{ApnProfile: jsonObj.ApnProfile}.ApnProfile,
			fmt.Sprintf("%s/%s", parentPath, "apn-profile"), params...)
		if err != nil {
			return nil, err
		}
		updates = append(updates, updateApnProfile)
	}
	if jsonObj.UpProfile != nil {
		updateUpProfile, err := utils.UpdateForElement(aether_1_0_0.AetherSubscriber_Subscriber_Ue_Profiles{UpProfile: jsonObj.UpProfile}.UpProfile,
			fmt.Sprintf("%s/%s", parentPath, "up-profile"), params...)
		if err != nil {
			return nil, err
		}
		updates = append(updates, updateUpProfile)
	}
	if jsonObj.QosProfile != nil {
		updateQosProfile, err := utils.UpdateForElement(aether_1_0_0.AetherSubscriber_Subscriber_Ue_Profiles{QosProfile: jsonObj.QosProfile}.QosProfile,
			fmt.Sprintf("%s/%s", parentPath, "qos-profile"), params...)
		if err != nil {
			return nil, err
		}
		updates = append(updates, updateQosProfile)
	}

	if jsonObj.ListAetherV100targetSubscriberUeProfilesAccessProfile != nil {
		for _, ap := range *jsonObj.ListAetherV100targetSubscriberUeProfilesAccessProfile {
			ap := ap //Pinning
			apUpdates, err := encodeToGnmiAetherV100targetSubscriberUeProfilesAccessProfile(&ap,
				fmt.Sprintf("%s/%s", parentPath, "access-profile/{access-profile}"), append(params, *ap.AccessProfile)...)
			if err != nil {
				return nil, err
			}
			updates = append(updates, apUpdates...)
		}
	}

	return updates, nil
}

func encodeToGnmiAetherV100targetSubscriberUeProfilesAccessProfile(
	jsonObj *types.AetherV100targetSubscriberUeProfilesAccessProfile, parentPath string, params ...string) (
	[]*gnmi.Update, error) {

	updates := make([]*gnmi.Update, 0)
	if jsonObj.AccessProfile != nil {
		updateApnProfile, err := utils.UpdateForElement(aether_1_0_0.AetherSubscriber_Subscriber_Ue_Profiles_AccessProfile{AccessProfile: jsonObj.AccessProfile}.AccessProfile,
			fmt.Sprintf("%s/%s", parentPath, "access-profile"), params...)
		if err != nil {
			return nil, err
		}
		updates = append(updates, updateApnProfile)
	}
	if jsonObj.Allowed != nil {
		updateUpProfile, err := utils.UpdateForElement(aether_1_0_0.AetherSubscriber_Subscriber_Ue_Profiles_AccessProfile{Allowed: jsonObj.Allowed}.Allowed,
			fmt.Sprintf("%s/%s", parentPath, "allowed"), params...)
		if err != nil {
			return nil, err
		}
		updates = append(updates, updateUpProfile)
	}
	return updates, nil
}

func encodeToGnmiAetherV100targetSubscriberUeServingPlmn(
	jsonObj *types.AetherV100targetSubscriberUeServingPlmn, parentPath string, params ...string) (
	[]*gnmi.Update, error) {

	updates := make([]*gnmi.Update, 0)
	if jsonObj.Mcc != nil {
		var mcc = uint32(*jsonObj.Mcc)
		updateType, err := utils.UpdateForElement(aether_1_0_0.AetherSubscriber_Subscriber_Ue_ServingPlmn{Mcc: &mcc}.Mcc,
			fmt.Sprintf("%s/%s", parentPath, "mcc"), params...)
		if err != nil {
			return nil, err
		}
		updates = append(updates, updateType)
	}
	if jsonObj.Mnc != nil {
		var mnc = uint32(*jsonObj.Mnc)
		updateType, err := utils.UpdateForElement(aether_1_0_0.AetherSubscriber_Subscriber_Ue_ServingPlmn{Mnc: &mnc}.Mnc,
			fmt.Sprintf("%s/%s", parentPath, "mnc"), params...)
		if err != nil {
			return nil, err
		}
		updates = append(updates, updateType)
	}
	if jsonObj.Tac != nil {
		var tac = uint32(*jsonObj.Tac)
		updateType, err := utils.UpdateForElement(aether_1_0_0.AetherSubscriber_Subscriber_Ue_ServingPlmn{Tac: &tac}.Tac,
			fmt.Sprintf("%s/%s", parentPath, "tac"), params...)
		if err != nil {
			return nil, err
		}
		updates = append(updates, updateType)
	}

	return updates, nil
}

func encodeToGnmiAetherV100targetUpProfile(
	jsonObj *types.AetherV100targetUpProfile, parentPath string, params ...string) (
	[]*gnmi.Update, error) {

	updates := make([]*gnmi.Update, 0)
	if jsonObj.ListAetherV100targetUpProfileUpProfile == nil || len(*jsonObj.ListAetherV100targetUpProfileUpProfile) == 0 {
		return nil, fmt.Errorf("nothing to commit - please add a profile")
	}
	for _, up := range *jsonObj.ListAetherV100targetUpProfileUpProfile {
		up := up //Pinning
		apUpdates, err := encodeToGnmiAetherV100targetUpProfileUpProfile(&up, "/up-profile/{id}", *up.Id)
		if err != nil {
			return nil, err
		}
		updates = append(updates, apUpdates...)
	}

	return updates, nil
}

func encodeToGnmiAetherV100targetUpProfileUpProfile(
	jsonObj *types.AetherV100targetUpProfileUpProfile, parentPath string, params ...string) (
	[]*gnmi.Update, error) {

	updates := make([]*gnmi.Update, 0)

	if len(parentPath) > 0 && (len(params) < 1 || params[0] == "") {
		return nil, fmt.Errorf("error: id is empty")
	}
	if jsonObj.Id != nil && *jsonObj.Id != params[0] {
		return nil, fmt.Errorf("error id in body is different to param %s != %s", *jsonObj.Id, params[0])
	}
	if jsonObj.Id != nil {
		updateID, err := utils.UpdateForElement(aether_1_0_0.UpProfile_UpProfile_UpProfile{Id: &params[0]}.Id,
			fmt.Sprintf("%s/%s", parentPath, "id"), params...)
		if err != nil {
			return nil, err
		}
		updates = append(updates, updateID)
	}

	if jsonObj.Description != nil {
		updateDesc, err := utils.UpdateForElement(aether_1_0_0.UpProfile_UpProfile_UpProfile{Description: jsonObj.Description}.Description,
			fmt.Sprintf("%s/%s", parentPath, "description"), params...)
		if err != nil {
			return nil, err
		}
		updates = append(updates, updateDesc)
	}

	if jsonObj.AccessControl != nil {
		updateAccessControl, err := utils.UpdateForElement(aether_1_0_0.UpProfile_UpProfile_UpProfile{AccessControl: jsonObj.AccessControl}.AccessControl,
			fmt.Sprintf("%s/%s", parentPath, "access-control"), params...)
		if err != nil {
			return nil, err
		}
		updates = append(updates, updateAccessControl)
	}

	if jsonObj.UserPlane != nil {
		updateUserPlane, err := utils.UpdateForElement(aether_1_0_0.UpProfile_UpProfile_UpProfile{UserPlane: jsonObj.UserPlane}.UserPlane,
			fmt.Sprintf("%s/%s", parentPath, "user-plane"), params...)
		if err != nil {
			return nil, err
		}
		updates = append(updates, updateUserPlane)
	}

	return updates, nil

}
