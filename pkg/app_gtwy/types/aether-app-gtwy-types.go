// Code generated by GENERATOR. DO NOT EDIT.
// Package types provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.7.0 DO NOT EDIT.
package types

// AppGtwyDevice defines model for App_Gtwy_Device.
type AppGtwyDevice struct {

	// attached
	Attached *string `json:"attached,omitempty"`

	// long description field
	Description *string `json:"description,omitempty"`

	// Link to device
	DeviceId string `json:"device-id,omitempty"`

	// The list of device groups
	DeviceGroups *[]string `json:"device_groups,omitempty"`

	// display name to use in GUI or CLI
	DisplayName *string `json:"display-name,omitempty"`

	// imei
	Imei *string `json:"imei,omitempty"`

	// ip
	Ip *string `json:"ip,omitempty"`

	// attached
	SimIccid *string `json:"sim_iccid,omitempty"`
}

// The top level app gateway devices container
type AppGtwyDevices struct {

	// List of app gateway devices
	Devices *[]AppGtwyDevice `json:"devices,omitempty"`
}

// target (device in onos-config)
type Target string
