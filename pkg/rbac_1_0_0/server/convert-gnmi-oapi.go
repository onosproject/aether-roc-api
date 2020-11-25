// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package server

import (
	"fmt"
	"github.com/onosproject/aether-roc-api/pkg/rbac_1_0_0/types"
	modelplugin "github.com/onosproject/config-models/modelplugin/rbac-1.0.0/rbac_1_0_0"
)

// ModelPluginDevice - a wrapper for the model plugin
type ModelPluginDevice struct {
	device modelplugin.Device
}

// Hand coded for the moment until we generate a template
func (d *ModelPluginDevice) toRbacV100targetRbac(params ...string) (*types.RbacV100targetRbac, error) {
	groups := make([]types.RbacV100targetRbacGroup, 0)
	roles := make([]types.RbacV100targetRbacRole, 0)
	resource := types.RbacV100targetRbac{
		ListRbacV100targetRbacGroup: &groups,
		ListRbacV100targetRbacRole:  &roles,
	}
	for _, v := range d.device.Rbac.Group {
		groups = append(groups, *convertGroup(v))
	}
	for _, v := range d.device.Rbac.Role {

		roles = append(roles, *convertRole(v))
	}

	return &resource, nil
}

func (d *ModelPluginDevice) toRbacV100targetRbacGroup(params ...string) (*types.RbacV100targetRbacGroup, error) {
	for _, gnmiGroup := range d.device.Rbac.Group {
		return convertGroup(gnmiGroup), nil // Return the first, should only be 1
	}

	return nil, fmt.Errorf("no group element found")
}

func (d *ModelPluginDevice) toRbacV100targetRbacGroupgroupidRole(params ...string) (*types.RbacV100targetRbacGroupgroupidRole, error) {
	return nil, fmt.Errorf("not yet implemented")
}

func (d *ModelPluginDevice) toRbacV100targetRbacRole(params ...string) (*types.RbacV100targetRbacRole, error) {
	for _, gnmiRole := range d.device.Rbac.Role {
		return convertRole(gnmiRole), nil // Return the first, should only be 1
	}

	return nil, fmt.Errorf("no role element found")
}

func (d *ModelPluginDevice) toRbacV100targetRbacRoleroleidPermission(params ...string) (*types.RbacV100targetRbacRoleroleidPermission, error) {
	for _, gnmiRole := range d.device.Rbac.Role {
		return convertRolePermission(gnmiRole.Permission), nil // Return the first, should only be 1
	}
	return nil, fmt.Errorf("no role element found")
}

func (d *ModelPluginDevice) toTarget(params ...string) (*types.Target, error) {
	return nil, fmt.Errorf("toTarget() should not be called directly")
}

func convertOperation(op modelplugin.E_RbacIdentities_PERMISSION) string {
	switch op {
	case modelplugin.RbacIdentities_PERMISSION_ALL:
		return "all"
	case modelplugin.RbacIdentities_PERMISSION_CREATE:
		return "create"
	case modelplugin.RbacIdentities_PERMISSION_READ:
		return "read"
	default:
		return "unset"
	}
}

func convertNounType(op modelplugin.E_RbacIdentities_NOUNTYPE) string {
	switch op {
	case modelplugin.RbacIdentities_NOUNTYPE_CONFIG:
		return "config"
	case modelplugin.RbacIdentities_NOUNTYPE_GRPC:
		return "grpc"
	default:
		return "unset"
	}
}

func convertGroup(gnmiGroup *modelplugin.Rbac_Rbac_Group) *types.RbacV100targetRbacGroup {
	group := types.RbacV100targetRbacGroup{
		Description: gnmiGroup.Description,
		Groupid:     gnmiGroup.Groupid,
	}
	groupRoles := make([]types.RbacV100targetRbacGroupgroupidRole, 0)
	for _, v1 := range gnmiGroup.Role {
		groupRole := types.RbacV100targetRbacGroupgroupidRole{
			Description: v1.Description,
			Roleid:      v1.Roleid,
		}
		groupRoles = append(groupRoles, groupRole)
	}
	group.ListRbacV100targetRbacGroupgroupidRole = &groupRoles
	return &group
}

func convertRole(gnmiRole *modelplugin.Rbac_Rbac_Role) *types.RbacV100targetRbacRole {
	role := types.RbacV100targetRbacRole{
		Description: gnmiRole.Description,
		Roleid:      gnmiRole.Roleid,
	}
	rolePermission := convertRolePermission(gnmiRole.Permission)
	role.RbacV100targetRbacRoleroleidPermission = rolePermission

	return &role
}

func convertRolePermission(gnmiRolePermission *modelplugin.Rbac_Rbac_Role_Permission) *types.RbacV100targetRbacRoleroleidPermission {
	opStr := convertOperation(gnmiRolePermission.Operation)
	typeStr := convertNounType(gnmiRolePermission.Type)
	rolePermission := types.RbacV100targetRbacRoleroleidPermission{
		Operation: &opStr,
		Type:      &typeStr,
	}
	rolePermissionNouns := make([]string, 0)
	rolePermissionNouns = append(rolePermissionNouns, gnmiRolePermission.Noun...)
	rolePermission.LeafListNoun = &rolePermissionNouns

	return &rolePermission
}
