// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package server

import (
	"fmt"
	"github.com/onosproject/aether-roc-api/pkg/rbac_1_0_0/types"
	"github.com/onosproject/config-models/modelplugin/rbac-1.0.0/rbac_1_0_0"
	"strings"
)

func encodeToGnmiRbacV100targetRbac(
	jsonObj *types.RbacV100targetRbac) (
	*rbac_1_0_0.Rbac_Rbac_Role, error) {

	return nil, fmt.Errorf("not yet implemented")
}

func encodeToGnmiRbacV100targetRbacGroup(
	jsonObj *types.RbacV100targetRbacGroup) (
	*rbac_1_0_0.Rbac_Rbac_Group, error) {

	if jsonObj.Groupid == nil || *jsonObj.Groupid == "" {
		return nil, fmt.Errorf("error Groupid cannot be empty")
	}

	roles := make(map[string]*rbac_1_0_0.Rbac_Rbac_Group_Role)
	for _, jsonRoleLeafref := range *jsonObj.ListRbacV100targetRbacGroupRole {
		roles[*jsonRoleLeafref.Roleid] = &rbac_1_0_0.Rbac_Rbac_Group_Role{
			Description: jsonRoleLeafref.Description,
			Roleid:      jsonRoleLeafref.Roleid,
		}
	}
	group := rbac_1_0_0.Rbac_Rbac_Group{
		Groupid:     jsonObj.Groupid,
		Description: jsonObj.Description,
		Role:        roles,
	}

	return &group, nil
}

func encodeToGnmiRbacV100targetRbacGroupRole(
	jsonObj *types.RbacV100targetRbacGroupRole) (
	*rbac_1_0_0.Rbac_Rbac_Group_Role, error) {

	return nil, fmt.Errorf("not yet implemented")
}

func encodeToGnmiRbacV100targetRbacRole(
	jsonObj *types.RbacV100targetRbacRole) (
	*rbac_1_0_0.Rbac_Rbac_Role, error) {

	if jsonObj.Roleid == nil || *jsonObj.Roleid == "" {
		return nil, fmt.Errorf("error Roleid cannot be empty")
	}
	nouns := make([]string, 0)
	nouns = append(nouns, *jsonObj.RbacV100targetRbacRolePermission.LeafListNoun...)
	role := rbac_1_0_0.Rbac_Rbac_Role{
		Roleid:      jsonObj.Roleid,
		Description: jsonObj.Description,
		Permission: &rbac_1_0_0.Rbac_Rbac_Role_Permission{
			Operation: convertPermissionOperation(*jsonObj.RbacV100targetRbacRolePermission.Operation),
			Type:      convertPermissionType(*jsonObj.RbacV100targetRbacRolePermission.Type),
			Noun:      nouns,
		},
	}

	return &role, nil
}

func encodeToGnmiRbacV100targetRbacRolePermission(
	jsonObj *types.RbacV100targetRbacRolePermission) (
	*rbac_1_0_0.Rbac_Rbac_Role_Permission, error) {

	return nil, fmt.Errorf("not yet implemented")
}

func convertPermissionType(jsonPermissionType string) rbac_1_0_0.E_RbacIdentities_NOUNTYPE {
	switch strings.ToLower(jsonPermissionType) {
	case "config":
		return rbac_1_0_0.RbacIdentities_NOUNTYPE_CONFIG
	case "grpc":
		return rbac_1_0_0.RbacIdentities_NOUNTYPE_GRPC
	default:
		return rbac_1_0_0.RbacIdentities_NOUNTYPE_UNSET
	}
}

func convertPermissionOperation(jsonPermissionPermission string) rbac_1_0_0.E_RbacIdentities_PERMISSION {
	switch strings.ToLower(jsonPermissionPermission) {
	case "read":
		return rbac_1_0_0.RbacIdentities_PERMISSION_READ
	case "create":
		return rbac_1_0_0.RbacIdentities_PERMISSION_CREATE
	case "all":
		return rbac_1_0_0.RbacIdentities_PERMISSION_ALL
	default:
		return rbac_1_0_0.RbacIdentities_PERMISSION_UNSET
	}
}
