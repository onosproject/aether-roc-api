// Code generated by GENERATOR. DO NOT EDIT.
package types

import (
	externalRef0 "github.com/onosproject/aether-roc-api/pkg/aether_2_1_0/types"
	externalRef1 "github.com/onosproject/aether-roc-api/pkg/aether_3_0_0/types"
)

// Elements defines model for Elements.
type Elements struct {
	AccessProfile210       *externalRef0.AccessProfile       `json:"access-profile-2.1.0,omitempty"`
	ApList300              *externalRef1.ApList              `json:"ap-list-3.0.0,omitempty"`
	ApnProfile210          *externalRef0.ApnProfile          `json:"apn-profile-2.1.0,omitempty"`
	Application300         *externalRef1.Application         `json:"application-3.0.0,omitempty"`
	ConnectivityService210 *externalRef0.ConnectivityService `json:"connectivity-service-2.1.0,omitempty"`
	ConnectivityService300 *externalRef1.ConnectivityService `json:"connectivity-service-3.0.0,omitempty"`
	DeviceGroup300         *externalRef1.DeviceGroup         `json:"device-group-3.0.0,omitempty"`
	Enterprise210          *externalRef0.Enterprise          `json:"enterprise-2.1.0,omitempty"`
	Enterprise300          *externalRef1.Enterprise          `json:"enterprise-3.0.0,omitempty"`
	IpDomain300            *externalRef1.IpDomain            `json:"ip-domain-3.0.0,omitempty"`
	QosProfile210          *externalRef0.QosProfile          `json:"qos-profile-2.1.0,omitempty"`
	SecurityProfile210     *externalRef0.SecurityProfile     `json:"security-profile-2.1.0,omitempty"`
	ServiceGroup210        *externalRef0.ServiceGroup        `json:"service-group-2.1.0,omitempty"`
	ServicePolicy210       *externalRef0.ServicePolicy       `json:"service-policy-2.1.0,omitempty"`
	ServiceRule210         *externalRef0.ServiceRule         `json:"service-rule-2.1.0,omitempty"`
	Site300                *externalRef1.Site                `json:"site-3.0.0,omitempty"`
	Subscriber210          *externalRef0.Subscriber          `json:"subscriber-2.1.0,omitempty"`
	Template300            *externalRef1.Template            `json:"template-3.0.0,omitempty"`
	TrafficClass300        *externalRef1.TrafficClass        `json:"traffic-class-3.0.0,omitempty"`
	UpProfile210           *externalRef0.UpProfile           `json:"up-profile-2.1.0,omitempty"`
	Upf300                 *externalRef1.Upf                 `json:"upf-3.0.0,omitempty"`
	Vcs300                 *externalRef1.Vcs                 `json:"vcs-3.0.0,omitempty"`
}

// PatchBody defines model for PatchBody.
type PatchBody struct {
	Deletes *Elements `json:"Deletes,omitempty"`

	// Model type and version of 'target' on first creation [link](https://docs.onosproject.org/onos-config/docs/gnmi_extensions/#use-of-extension-101-device-version-in-setrequest)
	Extensions *struct {
		ChangeName100   *string `json:"change-name-100,omitempty"`
		ModelType102    *string `json:"model-type-102,omitempty"`
		ModelVersion101 *string `json:"model-version-101,omitempty"`
	} `json:"Extensions,omitempty"`
	Updates *Elements `json:"Updates,omitempty"`

	// Target (device name) to use by default if not specified on indivdual updates/deletes as an additional property
	DefaultTarget string `json:"default-target"`
}

// TargetName defines model for TargetName.
type TargetName struct {
	Name *string `json:"name,omitempty"`
}

// TargetsNames defines model for TargetsNames.
type TargetsNames []TargetName

// PatchTopLevelJSONBody defines parameters for PatchTopLevel.
type PatchTopLevelJSONBody PatchBody

// PatchTopLevelJSONRequestBody defines body for PatchTopLevel for application/json ContentType.
type PatchTopLevelJSONRequestBody PatchTopLevelJSONBody
