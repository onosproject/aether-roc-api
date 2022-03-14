// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
//

package server

import (
	"fmt"
	externalRef0 "github.com/onosproject/aether-roc-api/pkg/aether_2_0_0/types"
	"github.com/onosproject/aether-roc-api/pkg/utils"
	externalRef1 "github.com/onosproject/config-models/modelplugin/aether-2.0.0/aether_2_0_0"
)

// ModelPluginDevice - a wrapper for the model plugin
type ModelPluginDevice struct {
	device externalRef1.Device
}

// toEnterprisesEnterpriseSite converts gNMI to OAPI.
func (d *ModelPluginDevice) toEnterprisesEnterpriseSite(params ...string) (*externalRef0.EnterprisesEnterpriseSite, error) {
	resource := new(externalRef0.EnterprisesEnterpriseSite)

	// Property: description string
	//encoding gNMI attribute to OAPI
	reflectDescription, err := utils.FindModelPluginObject(d.device, "EnterprisesEnterpriseSiteDescription", params...)
	if err != nil {
		return nil, err
	}
	if reflectDescription != nil {
		attrDescription := reflectDescription.Interface().(string)
		resource.Description = &attrDescription
	}

	// Property: device []EnterprisesEnterpriseSiteDevice
	// Handle []Object
	devices := make([]externalRef0.EnterprisesEnterpriseSiteDevice, 0)
	reflectEnterprisesEnterpriseSiteDevice, err := utils.FindModelPluginObject(d.device, "EnterprisesEnterpriseSiteDevice", params...)
	if err != nil {
		return nil, err
	}
	if reflectEnterprisesEnterpriseSiteDevice != nil {
		for _, key := range reflectEnterprisesEnterpriseSiteDevice.MapKeys() {
			v := reflectEnterprisesEnterpriseSiteDevice.MapIndex(key).Interface()
			// Pass down all top level properties as we don't know which one(s) is key
			attribs, err := utils.ExtractGnmiListKeyMap(v)
			if err != nil {
				return nil, err
			}
			childParams := make([]string, len(params))
			copy(childParams, params)
			for _, attribVal := range attribs {
				childParams = append(childParams, fmt.Sprintf("%v", attribVal))
			}
			device, err := d.toEnterprisesEnterpriseSiteDevice(childParams...)
			if err != nil {
				return nil, err
			}
			devices = append(devices, *device)
		}
	}
	resource.Device = &devices

	// Property: device-group []EnterprisesEnterpriseSiteDeviceGroup
	// Handle []Object
	deviceGroups := make([]externalRef0.EnterprisesEnterpriseSiteDeviceGroup, 0)
	reflectEnterprisesEnterpriseSiteDeviceGroup, err := utils.FindModelPluginObject(d.device, "EnterprisesEnterpriseSiteDeviceGroup", params...)
	if err != nil {
		return nil, err
	}
	if reflectEnterprisesEnterpriseSiteDeviceGroup != nil {
		for _, key := range reflectEnterprisesEnterpriseSiteDeviceGroup.MapKeys() {
			v := reflectEnterprisesEnterpriseSiteDeviceGroup.MapIndex(key).Interface()
			// Pass down all top level properties as we don't know which one(s) is key
			attribs, err := utils.ExtractGnmiListKeyMap(v)
			if err != nil {
				return nil, err
			}
			childParams := make([]string, len(params))
			copy(childParams, params)
			for _, attribVal := range attribs {
				childParams = append(childParams, fmt.Sprintf("%v", attribVal))
			}
			deviceGroup, err := d.toEnterprisesEnterpriseSiteDeviceGroup(childParams...)
			if err != nil {
				return nil, err
			}
			deviceGroups = append(deviceGroups, *deviceGroup)
		}
	}
	resource.DeviceGroup = &deviceGroups

	// Property: display-name string
	//encoding gNMI attribute to OAPI
	reflectDisplayName, err := utils.FindModelPluginObject(d.device, "EnterprisesEnterpriseSiteDisplayName", params...)
	if err != nil {
		return nil, err
	}
	if reflectDisplayName != nil {
		attrDisplayName := reflectDisplayName.Interface().(string)
		resource.DisplayName = &attrDisplayName
	}

	// Property: imsi-definition EnterprisesEnterpriseSiteImsiDefinition
	//Handle object
	attrImsiDefinition, err := d.toEnterprisesEnterpriseSiteImsiDefinition(params...)
	if err != nil {
		return nil, err
	}
	resource.ImsiDefinition = attrImsiDefinition

	// Property: sim-card []EnterprisesEnterpriseSiteSimCard
	// Handle []Object
	simCards := make([]externalRef0.EnterprisesEnterpriseSiteSimCard, 0)
	reflectEnterprisesEnterpriseSiteSimCard, err := utils.FindModelPluginObject(d.device, "EnterprisesEnterpriseSiteSimCard", params...)
	if err != nil {
		return nil, err
	}
	if reflectEnterprisesEnterpriseSiteSimCard != nil {
		for _, key := range reflectEnterprisesEnterpriseSiteSimCard.MapKeys() {
			v := reflectEnterprisesEnterpriseSiteSimCard.MapIndex(key).Interface()
			// Pass down all top level properties as we don't know which one(s) is key
			attribs, err := utils.ExtractGnmiListKeyMap(v)
			if err != nil {
				return nil, err
			}
			childParams := make([]string, len(params))
			copy(childParams, params)
			for _, attribVal := range attribs {
				childParams = append(childParams, fmt.Sprintf("%v", attribVal))
			}
			simCard, err := d.toEnterprisesEnterpriseSiteSimCard(childParams...)
			if err != nil {
				return nil, err
			}
			simCards = append(simCards, *simCard)
		}
	}
	resource.SimCard = &simCards

	// Property: site-id string
	//encoding gNMI attribute to OAPI
	reflectSiteID, err := utils.FindModelPluginObject(d.device, "EnterprisesEnterpriseSiteSiteId", params...)
	if err != nil {
		return nil, err
	}
	if reflectSiteID != nil {
		attrSiteID := reflectSiteID.Interface().(string)
		resource.SiteId = attrSiteID
	}

	return resource, nil
}

// toEnterprisesEnterpriseSiteSimCard converts gNMI to OAPI.
func (d *ModelPluginDevice) toEnterprisesEnterpriseSiteSimCard(params ...string) (*externalRef0.EnterprisesEnterpriseSiteSimCard, error) {
	resource := new(externalRef0.EnterprisesEnterpriseSiteSimCard)

	// Property: description string
	//encoding gNMI attribute to OAPI
	reflectDescription, err := utils.FindModelPluginObject(d.device, "EnterprisesEnterpriseSiteSimCardDescription", params...)
	if err != nil {
		return nil, err
	}
	if reflectDescription != nil {
		attrDescription := reflectDescription.Interface().(string)
		resource.Description = &attrDescription
	}

	// Property: display-name string
	//encoding gNMI attribute to OAPI
	reflectDisplayName, err := utils.FindModelPluginObject(d.device, "EnterprisesEnterpriseSiteSimCardDisplayName", params...)
	if err != nil {
		return nil, err
	}
	if reflectDisplayName != nil {
		attrDisplayName := reflectDisplayName.Interface().(string)
		resource.DisplayName = &attrDisplayName
	}

	// Property: iccid string
	//encoding gNMI attribute to OAPI
	reflectIccid, err := utils.FindModelPluginObject(d.device, "EnterprisesEnterpriseSiteSimCardIccid", params...)
	if err != nil {
		return nil, err
	}
	if reflectIccid != nil {
		attrIccid := reflectIccid.Interface().(string)
		resource.Iccid = &attrIccid
	}

	// Property: imsi int64
	//encoding gNMI attribute to OAPI
	reflectImsi, err := utils.FindModelPluginObject(d.device, "EnterprisesEnterpriseSiteSimCardImsi", params...)
	if err != nil {
		return nil, err
	}
	if reflectImsi != nil {
		//OpenAPI does not have unsigned numbers.
		if resource.Imsi, err = utils.ToInt64Ptr(reflectImsi); err != nil {
			return nil, err
		}
	}

	// Property: sim-id string
	//encoding gNMI attribute to OAPI
	reflectSimID, err := utils.FindModelPluginObject(d.device, "EnterprisesEnterpriseSiteSimCardSimId", params...)
	if err != nil {
		return nil, err
	}
	if reflectSimID != nil {
		attrSimID := reflectSimID.Interface().(string)
		resource.SimId = attrSimID
	}

	return resource, nil
}

// toEnterprisesEnterpriseSiteDevice converts gNMI to OAPI.
func (d *ModelPluginDevice) toEnterprisesEnterpriseSiteDevice(params ...string) (*externalRef0.EnterprisesEnterpriseSiteDevice, error) {
	resource := new(externalRef0.EnterprisesEnterpriseSiteDevice)

	// Property: description string
	//encoding gNMI attribute to OAPI
	reflectDescription, err := utils.FindModelPluginObject(d.device, "EnterprisesEnterpriseSiteDeviceDescription", params...)
	if err != nil {
		return nil, err
	}
	if reflectDescription != nil {
		attrDescription := reflectDescription.Interface().(string)
		resource.Description = &attrDescription
	}

	// Property: device-id string
	//encoding gNMI attribute to OAPI
	reflectDeviceID, err := utils.FindModelPluginObject(d.device, "EnterprisesEnterpriseSiteDeviceDeviceId", params...)
	if err != nil {
		return nil, err
	}
	if reflectDeviceID != nil {
		attrDeviceID := reflectDeviceID.Interface().(string)
		resource.DeviceId = attrDeviceID
	}

	// Property: display-name string
	//encoding gNMI attribute to OAPI
	reflectDisplayName, err := utils.FindModelPluginObject(d.device, "EnterprisesEnterpriseSiteDeviceDisplayName", params...)
	if err != nil {
		return nil, err
	}
	if reflectDisplayName != nil {
		attrDisplayName := reflectDisplayName.Interface().(string)
		resource.DisplayName = &attrDisplayName
	}

	// Property: imei string
	//encoding gNMI attribute to OAPI
	reflectImei, err := utils.FindModelPluginObject(d.device, "EnterprisesEnterpriseSiteDeviceImei", params...)
	if err != nil {
		return nil, err
	}
	if reflectImei != nil {
		attrImei := reflectImei.Interface().(string)
		resource.Imei = &attrImei
	}

	// Property: sim-card string
	//encoding gNMI attribute to OAPI
	reflectSimCard, err := utils.FindModelPluginObject(d.device, "EnterprisesEnterpriseSiteDeviceSimCard", params...)
	if err != nil {
		return nil, err
	}
	if reflectSimCard != nil {
		attrSimCard := reflectSimCard.Interface().(string)
		resource.SimCard = &attrSimCard
	}

	return resource, nil
}

// toEnterprisesEnterpriseSiteDeviceGroup converts gNMI to OAPI.
func (d *ModelPluginDevice) toEnterprisesEnterpriseSiteDeviceGroup(params ...string) (*externalRef0.EnterprisesEnterpriseSiteDeviceGroup, error) {
	resource := new(externalRef0.EnterprisesEnterpriseSiteDeviceGroup)

	// Property: description string
	//encoding gNMI attribute to OAPI
	reflectDescription, err := utils.FindModelPluginObject(d.device, "EnterprisesEnterpriseSiteDeviceGroupDescription", params...)
	if err != nil {
		return nil, err
	}
	if reflectDescription != nil {
		attrDescription := reflectDescription.Interface().(string)
		resource.Description = &attrDescription
	}

	// Property: device []EnterprisesEnterpriseSiteDeviceGroupDevice
	// Handle []Object
	devices := make([]externalRef0.EnterprisesEnterpriseSiteDeviceGroupDevice, 0)
	reflectEnterprisesEnterpriseSiteDeviceGroupDevice, err := utils.FindModelPluginObject(d.device, "EnterprisesEnterpriseSiteDeviceGroupDevice", params...)
	if err != nil {
		return nil, err
	}
	if reflectEnterprisesEnterpriseSiteDeviceGroupDevice != nil {
		for _, key := range reflectEnterprisesEnterpriseSiteDeviceGroupDevice.MapKeys() {
			v := reflectEnterprisesEnterpriseSiteDeviceGroupDevice.MapIndex(key).Interface()
			// Pass down all top level properties as we don't know which one(s) is key
			attribs, err := utils.ExtractGnmiListKeyMap(v)
			if err != nil {
				return nil, err
			}
			childParams := make([]string, len(params))
			copy(childParams, params)
			for _, attribVal := range attribs {
				childParams = append(childParams, fmt.Sprintf("%v", attribVal))
			}
			device, err := d.toEnterprisesEnterpriseSiteDeviceGroupDevice(childParams...)
			if err != nil {
				return nil, err
			}
			devices = append(devices, *device)
		}
	}
	resource.Device = &devices

	// Property: device-group-id string
	//encoding gNMI attribute to OAPI
	reflectDeviceGroupID, err := utils.FindModelPluginObject(d.device, "EnterprisesEnterpriseSiteDeviceGroupDeviceGroupId", params...)
	if err != nil {
		return nil, err
	}
	if reflectDeviceGroupID != nil {
		attrDeviceGroupID := reflectDeviceGroupID.Interface().(string)
		resource.DeviceGroupId = attrDeviceGroupID
	}

	// Property: display-name string
	//encoding gNMI attribute to OAPI
	reflectDisplayName, err := utils.FindModelPluginObject(d.device, "EnterprisesEnterpriseSiteDeviceGroupDisplayName", params...)
	if err != nil {
		return nil, err
	}
	if reflectDisplayName != nil {
		attrDisplayName := reflectDisplayName.Interface().(string)
		resource.DisplayName = &attrDisplayName
	}

	// Property: ip-domain string
	//encoding gNMI attribute to OAPI
	reflectIPDomain, err := utils.FindModelPluginObject(d.device, "EnterprisesEnterpriseSiteDeviceGroupIpDomain", params...)
	if err != nil {
		return nil, err
	}
	if reflectIPDomain != nil {
		attrIPDomain := reflectIPDomain.Interface().(string)
		resource.IpDomain = &attrIPDomain
	}

	// Property: mbr EnterprisesEnterpriseSiteDeviceGroupMbr
	//Handle object
	attrMbr, err := d.toEnterprisesEnterpriseSiteDeviceGroupMbr(params...)
	if err != nil {
		return nil, err
	}
	resource.Mbr = attrMbr

	// Property: traffic-class string
	//encoding gNMI attribute to OAPI
	reflectTrafficClass, err := utils.FindModelPluginObject(d.device, "EnterprisesEnterpriseSiteDeviceGroupTrafficClass", params...)
	if err != nil {
		return nil, err
	}
	if reflectTrafficClass != nil {
		attrTrafficClass := reflectTrafficClass.Interface().(string)
		resource.TrafficClass = attrTrafficClass
	}

	return resource, nil
}

// toEnterprisesEnterpriseSiteImsiDefinition converts gNMI to OAPI.
func (d *ModelPluginDevice) toEnterprisesEnterpriseSiteImsiDefinition(params ...string) (*externalRef0.EnterprisesEnterpriseSiteImsiDefinition, error) {
	resource := new(externalRef0.EnterprisesEnterpriseSiteImsiDefinition)

	// Property: enterprise int32
	//encoding gNMI attribute to OAPI
	reflectEnterprise, err := utils.FindModelPluginObject(d.device, "EnterprisesEnterpriseSiteImsiDefinitionEnterprise", params...)
	if err != nil {
		return nil, err
	}
	if reflectEnterprise != nil {
		//OpenAPI does not have unsigned numbers.
		if resource.Enterprise, err = utils.ToInt32(reflectEnterprise); err != nil {
			return nil, err
		}
	}

	// Property: format string
	//encoding gNMI attribute to OAPI
	reflectFormat, err := utils.FindModelPluginObject(d.device, "EnterprisesEnterpriseSiteImsiDefinitionFormat", params...)
	if err != nil {
		return nil, err
	}
	if reflectFormat != nil {
		attrFormat := reflectFormat.Interface().(string)
		resource.Format = attrFormat
	}

	// Property: mcc string
	//encoding gNMI attribute to OAPI
	reflectMcc, err := utils.FindModelPluginObject(d.device, "EnterprisesEnterpriseSiteImsiDefinitionMcc", params...)
	if err != nil {
		return nil, err
	}
	if reflectMcc != nil {
		attrMcc := reflectMcc.Interface().(string)
		resource.Mcc = attrMcc
	}

	// Property: mnc string
	//encoding gNMI attribute to OAPI
	reflectMnc, err := utils.FindModelPluginObject(d.device, "EnterprisesEnterpriseSiteImsiDefinitionMnc", params...)
	if err != nil {
		return nil, err
	}
	if reflectMnc != nil {
		attrMnc := reflectMnc.Interface().(string)
		resource.Mnc = attrMnc
	}

	return resource, nil
}

// toEnterprisesEnterpriseSiteDeviceGroupDevice converts gNMI to OAPI.
func (d *ModelPluginDevice) toEnterprisesEnterpriseSiteDeviceGroupDevice(params ...string) (*externalRef0.EnterprisesEnterpriseSiteDeviceGroupDevice, error) {
	resource := new(externalRef0.EnterprisesEnterpriseSiteDeviceGroupDevice)

	// Property: device-id string
	//encoding gNMI attribute to OAPI
	reflectDeviceID, err := utils.FindModelPluginObject(d.device, "EnterprisesEnterpriseSiteDeviceGroupDeviceDeviceId", params...)
	if err != nil {
		return nil, err
	}
	if reflectDeviceID != nil {
		attrDeviceID := reflectDeviceID.Interface().(string)
		resource.DeviceId = attrDeviceID
	}

	// Property: enable bool
	//encoding gNMI attribute to OAPI
	reflectEnable, err := utils.FindModelPluginObject(d.device, "EnterprisesEnterpriseSiteDeviceGroupDeviceEnable", params...)
	if err != nil {
		return nil, err
	}
	if reflectEnable != nil {
		boolEnable := reflectEnable.Interface().(bool)
		resource.Enable = &boolEnable
	}

	return resource, nil
}

// toEnterprisesEnterpriseSiteDeviceGroupMbr converts gNMI to OAPI.
func (d *ModelPluginDevice) toEnterprisesEnterpriseSiteDeviceGroupMbr(params ...string) (*externalRef0.EnterprisesEnterpriseSiteDeviceGroupMbr, error) {
	resource := new(externalRef0.EnterprisesEnterpriseSiteDeviceGroupMbr)

	// Property: downlink int64
	//encoding gNMI attribute to OAPI
	reflectDownlink, err := utils.FindModelPluginObject(d.device, "EnterprisesEnterpriseSiteDeviceGroupMbrDownlink", params...)
	if err != nil {
		return nil, err
	}
	if reflectDownlink != nil {
		//OpenAPI does not have unsigned numbers.
		if resource.Downlink, err = utils.ToInt64(reflectDownlink); err != nil {
			return nil, err
		}
	}

	// Property: uplink int64
	//encoding gNMI attribute to OAPI
	reflectUplink, err := utils.FindModelPluginObject(d.device, "EnterprisesEnterpriseSiteDeviceGroupMbrUplink", params...)
	if err != nil {
		return nil, err
	}
	if reflectUplink != nil {
		//OpenAPI does not have unsigned numbers.
		if resource.Uplink, err = utils.ToInt64(reflectUplink); err != nil {
			return nil, err
		}
	}

	return resource, nil
}
