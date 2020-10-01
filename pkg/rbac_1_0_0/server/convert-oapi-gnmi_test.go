// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package server

import (
	"github.com/onosproject/aether-roc-api/pkg/rbac_1_0_0/types"
	"gotest.tools/assert"
	"testing"
)

func Test_encodeToGnmiRbacV100targetRbacGroup(t *testing.T) {

	testGroupDesc := "Test group"
	testGroupID := "test-group"
	role1Id := "role1"
	role1Desc := "First role"
	role2Id := "role2"
	role2Desc := "Second role"

	roles := []types.RbacV100targetRbacGroupRole{
		{
			Description: &role2Desc,
			Roleid:      &role2Id,
		},
		{
			Description: &role1Desc,
			Roleid:      &role1Id,
		},
	}

	jsonObj := types.RbacV100targetRbacGroup{
		ListRbacV100targetRbacGroupRole: &roles,
		Description:                     &testGroupDesc,
		Groupid:                         &testGroupID,
	}

	gnmiUpdates, err := encodeToGnmiRbacV100targetRbacGroup(&jsonObj, "", testGroupID)
	assert.NilError(t, err)
	assert.Equal(t, 6, len(gnmiUpdates))
	for _, gnmiUpdate := range gnmiUpdates {
		switch gnmiUpdate.String() {
		case "path:<elem:<name:\"groupid\" > > val:<string_val:\"test-group\" > ",
			"path:<elem:<name:\"description\" > > val:<string_val:\"Test group\" > ",
			"path:<elem:<name:\"role\" key:<key:\"roleid\" value:\"role2\" > > elem:<name:\"roleid\" > > val:<string_val:\"role2\" > ",
			"path:<elem:<name:\"role\" key:<key:\"roleid\" value:\"role2\" > > elem:<name:\"description\" > > val:<string_val:\"Second role\" > ",
			"path:<elem:<name:\"role\" key:<key:\"roleid\" value:\"role1\" > > elem:<name:\"roleid\" > > val:<string_val:\"role1\" > ",
			"path:<elem:<name:\"role\" key:<key:\"roleid\" value:\"role1\" > > elem:<name:\"description\" > > val:<string_val:\"First role\" > ":
			// all ok
		default:
			t.Errorf("unexpected update %v", gnmiUpdate)
		}

	}
}

func Test_encodeToGnmiUpdatesRbacV100targetRbacRole(t *testing.T) {
	roleID := "role1"
	roleDesc := "Role 1"
	opRead := "read"
	typeConfig := "config"

	jsonRole := types.RbacV100targetRbacRole{
		RbacV100targetRbacRolePermission: &types.RbacV100targetRbacRolePermission{
			LeafListNoun: &[]string{"noun1", "noun2"},
			Operation:    &opRead,
			Type:         &typeConfig,
		},
		Description: &roleDesc,
		Roleid:      &roleID,
	}

	gnmiUpdates, err := encodeToGnmiRbacV100targetRbacRole(&jsonRole, "", roleID)
	assert.NilError(t, err)
	assert.Equal(t, 5, len(gnmiUpdates))

	update1RoleID := gnmiUpdates[0]
	assert.Equal(t, 1, len(update1RoleID.Path.Elem))
	update1RoleID0 := update1RoleID.Path.Elem[0]
	assert.Equal(t, "roleid", update1RoleID0.Name)
	assert.Equal(t, roleID, update1RoleID.Val.GetStringVal())

	update0Desc := gnmiUpdates[1]
	assert.Equal(t, 1, len(update0Desc.Path.Elem))
	update0Desc0 := update0Desc.Path.Elem[0]
	assert.Equal(t, "description", update0Desc0.Name)
	assert.Equal(t, roleDesc, update0Desc.Val.GetStringVal())

	update2PermissionOperation := gnmiUpdates[2]
	assert.Equal(t, 2, len(update2PermissionOperation.Path.Elem))
	update2PermissionOperation1 := update2PermissionOperation.Path.Elem[1]
	assert.Equal(t, "operation", update2PermissionOperation1.Name)
	assert.Equal(t, "READ", update2PermissionOperation.Val.GetStringVal())

	update2PermissionType := gnmiUpdates[3]
	assert.Equal(t, 2, len(update2PermissionType.Path.Elem))
	update2PermissionType1 := update2PermissionType.Path.Elem[1]
	assert.Equal(t, "type", update2PermissionType1.Name)
	assert.Equal(t, "CONFIG", update2PermissionType.Val.GetStringVal())

}
