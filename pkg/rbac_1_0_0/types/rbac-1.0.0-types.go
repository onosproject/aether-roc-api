// Code generated by oapi-codegen. DO NOT EDIT.
// Package types provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package types

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
)

// Rbac defines model for Rbac.
type Rbac struct {
	Group *[]RbacGroup `json:"Group,omitempty"`
	Role  *[]RbacRole  `json:"Role,omitempty"`
}

// RbacGroup defines model for Rbac_Group.
type RbacGroup struct {
	Role                 *[]RbacGroupRole       `json:"Role,omitempty"`
	Description          *string                `json:"description,omitempty"`
	Groupid              *string                `json:"groupid,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// RbacGroupRole defines model for Rbac_Group_Role.
type RbacGroupRole struct {
	Description *string `json:"description,omitempty"`
	Roleid      *string `json:"roleid,omitempty"`
}

// RbacRole defines model for Rbac_Role.
type RbacRole struct {
	Permission           *RbacRolePermission    `json:"Permission,omitempty"`
	Description          *string                `json:"description,omitempty"`
	Roleid               *string                `json:"roleid,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`
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

// Getter for additional properties for RbacGroup. Returns the specified
// element and whether it was found
func (a RbacGroup) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for RbacGroup
func (a *RbacGroup) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for RbacGroup to handle AdditionalProperties
func (a *RbacGroup) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["Role"]; found {
		err = json.Unmarshal(raw, &a.Role)
		if err != nil {
			return errors.Wrap(err, "error reading 'Role'")
		}
		delete(object, "Role")
	}

	if raw, found := object["description"]; found {
		err = json.Unmarshal(raw, &a.Description)
		if err != nil {
			return errors.Wrap(err, "error reading 'description'")
		}
		delete(object, "description")
	}

	if raw, found := object["groupid"]; found {
		err = json.Unmarshal(raw, &a.Groupid)
		if err != nil {
			return errors.Wrap(err, "error reading 'groupid'")
		}
		delete(object, "groupid")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("error unmarshaling field %s", fieldName))
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for RbacGroup to handle AdditionalProperties
func (a RbacGroup) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.Role != nil {
		object["Role"], err = json.Marshal(a.Role)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling 'Role'"))
		}
	}

	if a.Description != nil {
		object["description"], err = json.Marshal(a.Description)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling 'description'"))
		}
	}

	if a.Groupid != nil {
		object["groupid"], err = json.Marshal(a.Groupid)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling 'groupid'"))
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling '%s'", fieldName))
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for RbacRole. Returns the specified
// element and whether it was found
func (a RbacRole) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for RbacRole
func (a *RbacRole) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for RbacRole to handle AdditionalProperties
func (a *RbacRole) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["Permission"]; found {
		err = json.Unmarshal(raw, &a.Permission)
		if err != nil {
			return errors.Wrap(err, "error reading 'Permission'")
		}
		delete(object, "Permission")
	}

	if raw, found := object["description"]; found {
		err = json.Unmarshal(raw, &a.Description)
		if err != nil {
			return errors.Wrap(err, "error reading 'description'")
		}
		delete(object, "description")
	}

	if raw, found := object["roleid"]; found {
		err = json.Unmarshal(raw, &a.Roleid)
		if err != nil {
			return errors.Wrap(err, "error reading 'roleid'")
		}
		delete(object, "roleid")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("error unmarshaling field %s", fieldName))
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for RbacRole to handle AdditionalProperties
func (a RbacRole) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.Permission != nil {
		object["Permission"], err = json.Marshal(a.Permission)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling 'Permission'"))
		}
	}

	if a.Description != nil {
		object["description"], err = json.Marshal(a.Description)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling 'description'"))
		}
	}

	if a.Roleid != nil {
		object["roleid"], err = json.Marshal(a.Roleid)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling 'roleid'"))
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling '%s'", fieldName))
		}
	}
	return json.Marshal(object)
}
