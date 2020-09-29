// Code generated by oapi-codegen. DO NOT EDIT.
// Package types provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package types

// RbacV100targetRbac defines model for RbacV100targetRbac.
type RbacV100targetRbac struct {
	ListRbacV100targetRbacGroup *[]RbacV100targetRbacGroup `json:"ListRbacV100targetRbacGroup,omitempty"`
	ListRbacV100targetRbacRole  *[]RbacV100targetRbacRole  `json:"ListRbacV100targetRbacRole,omitempty"`
}

// RbacV100targetRbacGroup defines model for RbacV100targetRbacGroup.
type RbacV100targetRbacGroup struct {
	ListRbacV100targetRbacGroupRole *[]RbacV100targetRbacGroupRole `json:"ListRbacV100targetRbacGroupRole,omitempty"`
	Description                     *string                        `json:"description,omitempty"`
	Groupid                         *string                        `json:"groupid,omitempty"`
}

// RbacV100targetRbacGroupRole defines model for RbacV100targetRbacGroupRole.
type RbacV100targetRbacGroupRole struct {
	Description *string `json:"description,omitempty"`
	Roleid      *string `json:"roleid,omitempty"`
}

// RbacV100targetRbacRole defines model for RbacV100targetRbacRole.
type RbacV100targetRbacRole struct {
	RbacV100targetRbacRolePermission *RbacV100targetRbacRolePermission `json:"RbacV100targetRbacRolePermission,omitempty"`
	Description                      *string                           `json:"description,omitempty"`
	Roleid                           *string                           `json:"roleid,omitempty"`
}

// RbacV100targetRbacRolePermission defines model for RbacV100targetRbacRolePermission.
type RbacV100targetRbacRolePermission struct {
	LeafListNoun *[]string `json:"leaf-list-noun,omitempty"`
	Operation    *string   `json:"operation,omitempty"`
	Type         *string   `json:"type,omitempty"`
}

// Target defines model for target.
type Target string

// PostRbacV100targetRbacJSONBody defines parameters for PostRbacV100targetRbac.
type PostRbacV100targetRbacJSONBody RbacV100targetRbac

// PostRbacV100targetRbacGroupJSONBody defines parameters for PostRbacV100targetRbacGroup.
type PostRbacV100targetRbacGroupJSONBody RbacV100targetRbacGroup

// PostRbacV100targetRbacGroupRoleJSONBody defines parameters for PostRbacV100targetRbacGroupRole.
type PostRbacV100targetRbacGroupRoleJSONBody RbacV100targetRbacGroupRole

// PostRbacV100targetRbacRoleJSONBody defines parameters for PostRbacV100targetRbacRole.
type PostRbacV100targetRbacRoleJSONBody RbacV100targetRbacRole

// PostRbacV100targetRbacRolePermissionJSONBody defines parameters for PostRbacV100targetRbacRolePermission.
type PostRbacV100targetRbacRolePermissionJSONBody RbacV100targetRbacRolePermission

// PostRbacV100targetRbacRequestBody defines body for PostRbacV100targetRbac for application/json ContentType.
type PostRbacV100targetRbacJSONRequestBody PostRbacV100targetRbacJSONBody

// PostRbacV100targetRbacGroupRequestBody defines body for PostRbacV100targetRbacGroup for application/json ContentType.
type PostRbacV100targetRbacGroupJSONRequestBody PostRbacV100targetRbacGroupJSONBody

// PostRbacV100targetRbacGroupRoleRequestBody defines body for PostRbacV100targetRbacGroupRole for application/json ContentType.
type PostRbacV100targetRbacGroupRoleJSONRequestBody PostRbacV100targetRbacGroupRoleJSONBody

// PostRbacV100targetRbacRoleRequestBody defines body for PostRbacV100targetRbacRole for application/json ContentType.
type PostRbacV100targetRbacRoleJSONRequestBody PostRbacV100targetRbacRoleJSONBody

// PostRbacV100targetRbacRolePermissionRequestBody defines body for PostRbacV100targetRbacRolePermission for application/json ContentType.
type PostRbacV100targetRbacRolePermissionJSONRequestBody PostRbacV100targetRbacRolePermissionJSONBody
