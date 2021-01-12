// Code generated by oapi-codegen. DO NOT EDIT.
// Package server provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package server

import (
	"fmt"
)

import (
	"github.com/onosproject/aether-roc-api/pkg/rbac_1_0_0/types"
	"github.com/onosproject/aether-roc-api/pkg/utils"
	modelplugin "github.com/onosproject/config-models/modelplugin/rbac-1.0.0/rbac_1_0_0"
)

// ModelPluginDevice - a wrapper for the model plugin
type ModelPluginDevice struct {
	device modelplugin.Device
}

// toRbac converts gNMI to OAPI.
func (d *ModelPluginDevice) toRbac(params ...string) (*types.Rbac, error) {
	resource := new(types.Rbac)

	//Property: { Group {[]RbacGroup  map[] [] false <nil> [] false} false false}
	// Handle []Object
	groups := make([]types.RbacGroup, 0)
	reflectRbacGroup, err := utils.FindModelPluginObject(d.device, "RbacGroup", params...)
	if err != nil {
		return nil, err
	}
	for _, key := range reflectRbacGroup.MapKeys() {
		v := reflectRbacGroup.MapIndex(key).Interface()
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
		group, err := d.toRbacGroup(childParams...)
		if err != nil {
			return nil, err
		}
		groups = append(groups, *group)
	}
	resource.Group = &groups

	//Property: { Role {[]RbacRole  map[] [] false <nil> [] false} false false}
	// Handle []Object
	roles := make([]types.RbacRole, 0)
	reflectRbacRole, err := utils.FindModelPluginObject(d.device, "RbacRole", params...)
	if err != nil {
		return nil, err
	}
	for _, key := range reflectRbacRole.MapKeys() {
		v := reflectRbacRole.MapIndex(key).Interface()
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
		role, err := d.toRbacRole(childParams...)
		if err != nil {
			return nil, err
		}
		roles = append(roles, *role)
	}
	resource.Role = &roles

	return resource, nil
}

// toRbacGroup converts gNMI to OAPI.
func (d *ModelPluginDevice) toRbacGroup(params ...string) (*types.RbacGroup, error) {
	resource := new(types.RbacGroup)

	//Property: { Role {[]RbacGroupRole  map[] [] false <nil> [] false} false false}
	// Handle []Object
	roles := make([]types.RbacGroupRole, 0)
	reflectRbacGroupRole, err := utils.FindModelPluginObject(d.device, "RbacGroupRole", params...)
	if err != nil {
		return nil, err
	}
	for _, key := range reflectRbacGroupRole.MapKeys() {
		v := reflectRbacGroupRole.MapIndex(key).Interface()
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
		role, err := d.toRbacGroupRole(childParams...)
		if err != nil {
			return nil, err
		}
		roles = append(roles, *role)
	}
	resource.Role = &roles

	//Property: { description {string  map[] [] false <nil> [] false} false false}
	//encoding gNMI attribute to OAPI
	reflectDescription, err := utils.FindModelPluginObject(d.device, "RbacGroupDescription", params...)
	if err != nil {
		return nil, err
	}
	if reflectDescription != nil {
		attrDescription := reflectDescription.Interface().(string)
		resource.Description = &attrDescription
	}

	//Property: { groupid {string  map[] [] false <nil> [] false} false false}
	//encoding gNMI attribute to OAPI
	reflectGroupid, err := utils.FindModelPluginObject(d.device, "RbacGroupGroupid", params...)
	if err != nil {
		return nil, err
	}
	if reflectGroupid != nil {
		attrGroupid := reflectGroupid.Interface().(string)
		resource.Groupid = &attrGroupid
	}

	return resource, nil
}

// toRbacGroupRole converts gNMI to OAPI.
func (d *ModelPluginDevice) toRbacGroupRole(params ...string) (*types.RbacGroupRole, error) {
	resource := new(types.RbacGroupRole)

	//Property: { description {string  map[] [] false <nil> [] false} false false}
	//encoding gNMI attribute to OAPI
	reflectDescription, err := utils.FindModelPluginObject(d.device, "RbacGroupRoleDescription", params...)
	if err != nil {
		return nil, err
	}
	if reflectDescription != nil {
		attrDescription := reflectDescription.Interface().(string)
		resource.Description = &attrDescription
	}

	//Property: { roleid {string  map[] [] false <nil> [] false} false false}
	//encoding gNMI attribute to OAPI
	reflectRoleid, err := utils.FindModelPluginObject(d.device, "RbacGroupRoleRoleid", params...)
	if err != nil {
		return nil, err
	}
	if reflectRoleid != nil {
		attrRoleid := reflectRoleid.Interface().(string)
		resource.Roleid = &attrRoleid
	}

	return resource, nil
}

// toRbacRole converts gNMI to OAPI.
func (d *ModelPluginDevice) toRbacRole(params ...string) (*types.RbacRole, error) {
	resource := new(types.RbacRole)

	//Property: { Permission {RbacRolePermission  map[] [] false <nil> [] false} false false}
	//Handle object
	attrPermission, err := d.toRbacRolePermission(params...)
	if err != nil {
		return nil, err
	}
	resource.Permission = attrPermission

	//Property: { description {string  map[] [] false <nil> [] false} false false}
	//encoding gNMI attribute to OAPI
	reflectDescription, err := utils.FindModelPluginObject(d.device, "RbacRoleDescription", params...)
	if err != nil {
		return nil, err
	}
	if reflectDescription != nil {
		attrDescription := reflectDescription.Interface().(string)
		resource.Description = &attrDescription
	}

	//Property: { roleid {string  map[] [] false <nil> [] false} false false}
	//encoding gNMI attribute to OAPI
	reflectRoleid, err := utils.FindModelPluginObject(d.device, "RbacRoleRoleid", params...)
	if err != nil {
		return nil, err
	}
	if reflectRoleid != nil {
		attrRoleid := reflectRoleid.Interface().(string)
		resource.Roleid = &attrRoleid
	}

	return resource, nil
}

// toRbacRolePermission converts gNMI to OAPI.
func (d *ModelPluginDevice) toRbacRolePermission(params ...string) (*types.RbacRolePermission, error) {
	resource := new(types.RbacRolePermission)

	//Property: { leaf-list-noun {[]string  map[] [] false <nil> [] false} false false}
	//Leaf list handling
	reflectLeafListNoun, err := utils.FindModelPluginObject(d.device, "RbacRolePermissionNoun", params...)
	if err != nil {
		return nil, err
	}
	asArrayLeafListNoun := reflectLeafListNoun.Interface().([]string)
	resource.LeafListNoun = &asArrayLeafListNoun

	//Property: { operation {string  map[ALL:ALL CREATE:CREATE READ:READ] [] false <nil> [] false} false false}
	// Enums handling
	reflectOperation, err := utils.FindModelPluginObject(d.device, "RbacRolePermissionOperation", params...)
	if err != nil {
		return nil, err
	}
	attrOperation := reflectOperation.Interface()
	_, yangDefOperation, err := utils.ExtractGnmiEnumMap(&d.device, "RbacRolePermissionOperation", attrOperation)
	if err != nil {
		return nil, err
	}
	resource.Operation = &yangDefOperation.Name

	//Property: { type {string  map[CONFIG:CONFIG GRPC:GRPC] [] false <nil> [] false} false false}
	// Enums handling
	reflectType, err := utils.FindModelPluginObject(d.device, "RbacRolePermissionType", params...)
	if err != nil {
		return nil, err
	}
	attrType := reflectType.Interface()
	_, yangDefType, err := utils.ExtractGnmiEnumMap(&d.device, "RbacRolePermissionType", attrType)
	if err != nil {
		return nil, err
	}
	resource.Type = &yangDefType.Name

	return resource, nil
}

// toTarget converts gNMI to OAPI.
func (d *ModelPluginDevice) toTarget(params ...string) (*types.Target, error) {
	resource := new(types.Target)

	return resource, nil
}

//Ignoring RequestBodyRbac

//Ignoring RequestBodyRbacGroup

//Ignoring RequestBodyRbacGroupRole

//Ignoring RequestBodyRbacRole

//Ignoring RequestBodyRbacRolePermission

// Not generating param-types
// Not generating request-bodies

// Not generating additional-properties
// Not generating additional-properties
