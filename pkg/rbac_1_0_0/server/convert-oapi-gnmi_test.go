// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package server

import (
	"github.com/onosproject/aether-roc-api/pkg/rbac_1_0_0/types"
	"github.com/onosproject/config-models/modelplugin/rbac-1.0.0/rbac_1_0_0"
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

	gnmiGroup, err := encodeToGnmiRbacV100targetRbacGroup(&jsonObj)
	assert.NilError(t, err)
	assert.Assert(t, gnmiGroup != nil)
	assert.Equal(t, testGroupDesc, *gnmiGroup.Description)
	assert.Equal(t, testGroupID, *gnmiGroup.Groupid)
	assert.Equal(t, 2, len(gnmiGroup.Role))
	for roleID, role := range gnmiGroup.Role {
		switch roleID {
		case role1Id:
			assert.Equal(t, role1Desc, *role.Description)
		case role2Id:
			assert.Equal(t, role2Desc, *role.Description)
		default:
			t.Fatalf("unexpected role id %s", roleID)
		}
	}
}

func Test_encodeToGnmiRbacV100targetRbacRole(t *testing.T) {
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

	gnmiRole, err := encodeToGnmiRbacV100targetRbacRole(&jsonRole)
	assert.NilError(t, err)
	assert.Assert(t, gnmiRole != nil)
	assert.Equal(t, roleDesc, *gnmiRole.Description)
	assert.Equal(t, roleID, *gnmiRole.Roleid)
	assert.Equal(t, rbac_1_0_0.RbacIdentities_PERMISSION_READ, gnmiRole.Permission.Operation)
	assert.Equal(t, rbac_1_0_0.RbacIdentities_NOUNTYPE_CONFIG, gnmiRole.Permission.Type)
	assert.Equal(t, 2, len(gnmiRole.Permission.Noun))
	for _, noun := range gnmiRole.Permission.Noun {
		switch noun {
		case "noun1", "noun2":
		default:
			t.Fatalf("unexpected value %s", noun)
		}
	}
}
