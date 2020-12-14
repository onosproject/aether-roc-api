// Code generated by oapi-codegen. DO NOT EDIT.
// Package types provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package types

// Rbac defines model for Rbac.
type Rbac struct {
	Group *[]RbacGroup `json:"Group,omitempty"`
	Role  *[]RbacRole  `json:"Role,omitempty"`
}

// RbacGroup defines model for Rbac_Group.
type RbacGroup struct {
	Role        *[]RbacGroupRole `json:"Role,omitempty"`
	Description *string          `json:"description,omitempty"`
	Groupid     *string          `json:"groupid,omitempty"`
}

// RbacGroupRole defines model for Rbac_Group_Role.
type RbacGroupRole struct {
	Description *string `json:"description,omitempty"`
	Roleid      *string `json:"roleid,omitempty"`
}

// RbacRole defines model for Rbac_Role.
type RbacRole struct {
	Permission  *RbacRolePermission `json:"Permission,omitempty"`
	Description *string             `json:"description,omitempty"`
	Roleid      *string             `json:"roleid,omitempty"`
}

// RbacRolePermission defines model for Rbac_Role_Permission.
type RbacRolePermission struct {
	LeafListNoun *[]string `json:"leaf-list-noun,omitempty"`
	Operation    *string   `json:"operation,omitempty"`
	Type         *string   `json:"type,omitempty"`
}

// Target defines model for target.
type Target string

// RequestBodyRbac defines model for RequestBody_Rbac.
type RequestBodyRbac Rbac

// RequestBodyRbacGroup defines model for RequestBody_Rbac_Group.
type RequestBodyRbacGroup RbacGroup

// RequestBodyRbacGroupRole defines model for RequestBody_Rbac_Group_Role.
type RequestBodyRbacGroupRole RbacGroupRole

// RequestBodyRbacRole defines model for RequestBody_Rbac_Role.
type RequestBodyRbacRole RbacRole

// RequestBodyRbacRolePermission defines model for RequestBody_Rbac_Role_Permission.
type RequestBodyRbacRolePermission RbacRolePermission

// PostRbacRequestBody defines body for PostRbac for application/json ContentType.
type PostRbacJSONRequestBody RequestBodyRbac

// PostRbacGroupRequestBody defines body for PostRbacGroup for application/json ContentType.
type PostRbacGroupJSONRequestBody RequestBodyRbacGroup

// PostRbacGroupRoleRequestBody defines body for PostRbacGroupRole for application/json ContentType.
type PostRbacGroupRoleJSONRequestBody RequestBodyRbacGroupRole

// PostRbacRoleRequestBody defines body for PostRbacRole for application/json ContentType.
type PostRbacRoleJSONRequestBody RequestBodyRbacRole

// PostRbacRolePermissionRequestBody defines body for PostRbacRolePermission for application/json ContentType.
type PostRbacRolePermissionJSONRequestBody RequestBodyRbacRolePermission
