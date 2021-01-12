// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package server

import (
	"github.com/onosproject/aether-roc-api/pkg/rbac_1_0_0/types"
	"gotest.tools/assert"
	"strings"
	"testing"
)

func Test_encodeToGnmiRbacGroup(t *testing.T) {

	testGroupDesc := "Test group"
	testGroupID := "test-group"
	role1Id := "role1"
	role1Desc := "First role"
	role2Id := "role2"
	role2Desc := "Second role"

	roles := []types.RbacGroupRole{
		{
			Description: &role2Desc,
			Roleid:      &role2Id,
		},
		{
			Description: &role1Desc,
			Roleid:      &role1Id,
		},
	}

	jsonObj := types.RbacGroup{
		Role:        &roles,
		Description: &testGroupDesc,
		Groupid:     &testGroupID,
	}

	gnmiUpdates, err := EncodeToGnmiRbacGroup(&jsonObj, true, "test1", "/rbac/group/{unknown_key}", "unknown_id")
	assert.NilError(t, err)
	assert.Equal(t, 6, len(gnmiUpdates))
	for _, gnmiUpdate := range gnmiUpdates {
		switch path := strings.ReplaceAll(gnmiUpdate.String(), "  ", " "); path {
		case
			`path:{elem:{name:"rbac"} elem:{name:"group" key:{key:"groupid" value:"test-group"}} elem:{name:"role" key:{key:"roleid" value:"role2"}} elem:{name:"description"} target:"test1"} val:{string_val:"Second role"}`,
			`path:{elem:{name:"rbac"} elem:{name:"group" key:{key:"groupid" value:"test-group"}} elem:{name:"role" key:{key:"roleid" value:"role2"}} elem:{name:"roleid"} target:"test1"} val:{string_val:"role2"}`,
			`path:{elem:{name:"rbac"} elem:{name:"group" key:{key:"groupid" value:"test-group"}} elem:{name:"role" key:{key:"roleid" value:"role1"}} elem:{name:"description"} target:"test1"} val:{string_val:"First role"}`,
			`path:{elem:{name:"rbac"} elem:{name:"group" key:{key:"groupid" value:"test-group"}} elem:{name:"role" key:{key:"roleid" value:"role1"}} elem:{name:"roleid"} target:"test1"} val:{string_val:"role1"}`,
			`path:{elem:{name:"rbac"} elem:{name:"group" key:{key:"groupid" value:"test-group"}} elem:{name:"description"} target:"test1"} val:{string_val:"Test group"}`,
			`path:{elem:{name:"rbac"} elem:{name:"group" key:{key:"groupid" value:"test-group"}} elem:{name:"groupid"} target:"test1"} val:{string_val:"test-group"}`,
			`path:{elem:{name:"rbac"}  elem:{name:"group"  key:{key:"groupid"  value:"test-group"}}  elem:{name:"role"  key:{key:"roleid"  value:"role2"}}  elem:{name:"description"} target:"test1"}  val:{string_val:"Second role"}`,
			`path:{elem:{name:"rbac"}  elem:{name:"group"  key:{key:"groupid"  value:"test-group"}}  elem:{name:"role"  key:{key:"roleid"  value:"role2"}}  elem:{name:"roleid"} target:"test1"}  val:{string_val:"role2"}`,
			`path:{elem:{name:"rbac"}  elem:{name:"group"  key:{key:"groupid"  value:"test-group"}}  elem:{name:"role"  key:{key:"roleid"  value:"role1"}}  elem:{name:"description"} target:"test1"}  val:{string_val:"First role"}`,
			`path:{elem:{name:"rbac"}  elem:{name:"group"  key:{key:"groupid"  value:"test-group"}}  elem:{name:"role"  key:{key:"roleid"  value:"role1"}}  elem:{name:"roleid"} target:"test1"}  val:{string_val:"role1"}`,
			`path:{elem:{name:"rbac"}  elem:{name:"group"  key:{key:"groupid"  value:"test-group"}}  elem:{name:"description"} target:"test1"}  val:{string_val:"Test group"}`,
			`path:{elem:{name:"rbac"}  elem:{name:"group"  key:{key:"groupid"  value:"test-group"}}  elem:{name:"groupid"} target:"test1"}  val:{string_val:"test-group"}`:
			// all ok
		default:
			t.Errorf("unexpected update %s", path)
		}

	}
}

func Test_encodeToGnmiUpdatesRbacRole(t *testing.T) {
	roleID := "role1"
	roleDesc := "Role 1"
	opRead := "READ"
	typeConfig := "CONFIG"

	jsonRole := types.RbacRole{
		Permission: &types.RbacRolePermission{
			LeafListNoun: &[]string{"noun1", "noun2"},
			Operation:    &opRead,
			Type:         &typeConfig,
		},
		Description: &roleDesc,
		Roleid:      &roleID,
	}

	gnmiUpdates, err := EncodeToGnmiRbacRole(&jsonRole, false, "test-target", "", roleID)
	assert.NilError(t, err)
	assert.Equal(t, 5, len(gnmiUpdates))

	update0Noun := gnmiUpdates[0]
	assert.Equal(t, 2, len(update0Noun.Path.Elem))
	update0Noun0 := update0Noun.Path.Elem[0]
	assert.Equal(t, "permission", update0Noun0.Name)
	update0Noun1 := update0Noun.Path.Elem[1]
	assert.Equal(t, "noun", update0Noun1.Name)
	assert.Equal(t, 2, len(update0Noun.Val.GetLeaflistVal().GetElement()))

	update1Operation := gnmiUpdates[1]
	assert.Equal(t, 2, len(update1Operation.Path.Elem))
	update1Operation0 := update1Operation.Path.Elem[0]
	assert.Equal(t, "permission", update1Operation0.Name)
	update1Operation1 := update1Operation.Path.Elem[1]
	assert.Equal(t, "operation", update1Operation1.Name)
	assert.Equal(t, opRead, update1Operation.Val.GetStringVal())

	update2Type := gnmiUpdates[2]
	assert.Equal(t, 2, len(update2Type.Path.Elem))
	update2Type0 := update2Type.Path.Elem[0]
	assert.Equal(t, "permission", update2Type0.Name)
	update2Type1 := update2Type.Path.Elem[1]
	assert.Equal(t, "type", update2Type1.Name)
	assert.Equal(t, typeConfig, update2Type.Val.GetStringVal())

	update3Description := gnmiUpdates[3]
	assert.Equal(t, 1, len(update3Description.Path.Elem))
	update3Description0 := update3Description.Path.Elem[0]
	assert.Equal(t, "description", update3Description0.Name)
	assert.Equal(t, "Role 1", update3Description.Val.GetStringVal())

	update4RoleID := gnmiUpdates[4]
	assert.Equal(t, 1, len(update4RoleID.Path.Elem))
	update4RoleID0 := update4RoleID.Path.Elem[0]
	assert.Equal(t, "roleid", update4RoleID0.Name)
	assert.Equal(t, "role1", update4RoleID.Val.GetStringVal())

}
