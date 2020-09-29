// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package server

import (
	"fmt"
	"github.com/onosproject/aether-roc-api/pkg/rbac_1_0_0/types"
	"github.com/onosproject/aether-roc-api/pkg/utils"
	"github.com/onosproject/config-models/modelplugin/rbac-1.0.0/rbac_1_0_0"
	"github.com/openconfig/gnmi/proto/gnmi"
	"strings"
)

func encodeToGnmiRbacV100targetRbac(
	jsonObj *types.RbacV100targetRbac, parentPath string, params ...string) (
	[]*gnmi.Update, error) {

	updates := make([]*gnmi.Update, 0)
	if jsonObj.ListRbacV100targetRbacGroup != nil {
		for _, group := range *jsonObj.ListRbacV100targetRbacGroup {
			group := group //Pinning
			updateGroup, err := encodeToGnmiRbacV100targetRbacGroup(&group, parentPath, params...)
			if err != nil {
				return nil, err
			}
			updates = append(updates, updateGroup...)
		}
	}
	if jsonObj.ListRbacV100targetRbacRole != nil {
		for _, role := range *jsonObj.ListRbacV100targetRbacRole {
			role := role //Pinning
			updateRole, err := encodeToGnmiRbacV100targetRbacRole(&role, parentPath, params...)
			if err != nil {
				return nil, err
			}
			updates = append(updates, updateRole...)
		}
	}

	return updates, nil
}

func encodeToGnmiRbacV100targetRbacGroup(
	jsonObj *types.RbacV100targetRbacGroup, parentPath string, params ...string) (
	[]*gnmi.Update, error) {

	if len(params) < 1 || params[0] == "" {
		return nil, fmt.Errorf("error groupid is empty")
	}

	if jsonObj.Groupid != nil && *jsonObj.Groupid != params[0] {
		return nil, fmt.Errorf("error groupid in body is different to param %s != %s", *jsonObj.Groupid, params[0])
	}
	updateID, err := utils.UpdateForElement(rbac_1_0_0.Rbac_Rbac_Group{Groupid: &params[0]}.Groupid, "/groupid")
	if err != nil {
		return nil, err
	}
	updates := make([]*gnmi.Update, 0)
	updates = append(updates, updateID)

	if jsonObj.Description != nil {
		updateDesc, err := utils.UpdateForElement(rbac_1_0_0.Rbac_Rbac_Group{Description: jsonObj.Description}.Description, "/description")
		if err != nil {
			return nil, err
		}
		updates = append(updates, updateDesc)
	}

	if jsonObj.ListRbacV100targetRbacGroupRole != nil {
		for _, groupRole := range *jsonObj.ListRbacV100targetRbacGroupRole {
			groupRole := groupRole //Pinning
			updatesPerm, err := encodeToGnmiRbacV100targetRbacGroupRole(&groupRole, parentPath, params...)
			if err != nil {
				return nil, err
			}
			updates = append(updates, updatesPerm...)
		}
	}

	return updates, nil
}

func encodeToGnmiRbacV100targetRbacGroupRole(
	jsonObj *types.RbacV100targetRbacGroupRole, parentPath string, params ...string) (
	[]*gnmi.Update, error) {

	if jsonObj.Description == nil || jsonObj.Roleid == nil {
		return nil, fmt.Errorf("A role ID and a description must be given")
	}
	updates := make([]*gnmi.Update, 0)
	updateID, err := utils.UpdateForElement(rbac_1_0_0.Rbac_Rbac_Group_Role{Roleid: jsonObj.Roleid}.Roleid, "/roleid")
	if err != nil {
		return nil, err
	}
	updates = append(updates, updateID)

	updateDesc, err := utils.UpdateForElement(rbac_1_0_0.Rbac_Rbac_Group_Role{Description: jsonObj.Description}.Roleid, "/description")
	if err != nil {
		return nil, err
	}
	updates = append(updates, updateDesc)

	return updates, nil
}

func encodeToGnmiRbacV100targetRbacRole(
	jsonObj *types.RbacV100targetRbacRole, parentPath string, params ...string) (
	[]*gnmi.Update, error) {

	if len(params) < 1 || params[0] == "" {
		return nil, fmt.Errorf("error roleid is empty")
	}

	if jsonObj.Roleid != nil && *jsonObj.Roleid != params[0] {
		return nil, fmt.Errorf("error Roleid in body is different to param %s != %s", *jsonObj.Roleid, params[0])
	}
	updateID, err := utils.UpdateForElement(rbac_1_0_0.Rbac_Rbac_Role{Roleid: &params[0]}.Roleid, "/roleid")
	if err != nil {
		return nil, err
	}
	updates := make([]*gnmi.Update, 0)
	updates = append(updates, updateID)

	if jsonObj.Description != nil {
		updateDesc, err := utils.UpdateForElement(rbac_1_0_0.Rbac_Rbac_Role{Description: jsonObj.Description}.Description, "/description")
		if err != nil {
			return nil, err
		}
		updates = append(updates, updateDesc)
	}

	updatesPerm, err := encodeToGnmiRbacV100targetRbacRolePermission(jsonObj.RbacV100targetRbacRolePermission, parentPath, params...)
	if err != nil {
		return nil, err
	}
	updates = append(updates, updatesPerm...)

	return updates, nil
}

func encodeToGnmiRbacV100targetRbacRolePermission(
	jsonObjPermission *types.RbacV100targetRbacRolePermission, parentPath string, params ...string) (
	[]*gnmi.Update, error) {

	updates := make([]*gnmi.Update, 0)

	if jsonObjPermission != nil {
		if jsonObjPermission.Operation != nil {
			updatePermissionOp, err := updateForPermissionOperation(*jsonObjPermission.Operation, "/permission/operation")
			if err != nil {
				return nil, err
			}
			updates = append(updates, updatePermissionOp)
		}
		if jsonObjPermission.Type != nil {
			updatePermissionType, err := updateForPermissionType(*jsonObjPermission.Type, "/permission/type")
			if err != nil {
				return nil, err
			}
			updates = append(updates, updatePermissionType)
		}
		if jsonObjPermission.LeafListNoun != nil {
			updateNouns, err := utils.UpdateForElement(
				rbac_1_0_0.Rbac_Rbac_Role_Permission{Noun: *jsonObjPermission.LeafListNoun}.Noun, "/permission/noun")
			if err != nil {
				return nil, err
			}
			updates = append(updates, updateNouns)
		}
	}
	return updates, nil
}

func updateForPermissionType(value string, path string, pathParams ...string) (*gnmi.Update, error) {
	update := new(gnmi.Update)
	update.Path = new(gnmi.Path)
	var err error
	if update.Path.Elem, err = utils.BuildElems(path, 1, pathParams...); err != nil {
		return nil, err
	}
	update.Val = new(gnmi.TypedValue)
	switch strings.ToUpper(value) {
	case "CONFIG", "GRPC":
		update.Val.Value = &gnmi.TypedValue_StringVal{StringVal: strings.ToUpper(value)}
	default:
		return nil, fmt.Errorf("unexpected value for type %s", value)
	}

	return update, nil
}

func updateForPermissionOperation(value string, path string, pathParams ...string) (*gnmi.Update, error) {
	update := new(gnmi.Update)
	update.Path = new(gnmi.Path)
	var err error
	if update.Path.Elem, err = utils.BuildElems(path, 1, pathParams...); err != nil {
		return nil, err
	}
	update.Val = new(gnmi.TypedValue)
	switch strings.ToUpper(value) {
	case "READ", "CREATE", "ALL":
		update.Val.Value = &gnmi.TypedValue_StringVal{StringVal: strings.ToUpper(value)}
	default:
		return nil, fmt.Errorf("unexpected value for operation %s", value)
	}

	return update, nil
}
