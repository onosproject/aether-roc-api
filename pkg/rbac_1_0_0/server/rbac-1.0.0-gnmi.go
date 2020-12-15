// Code generated by oapi-codegen. DO NOT EDIT.
// Package server provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package server

import (
	"context"
	"encoding/json"
	"fmt"
)

import (
	"github.com/onosproject/aether-roc-api/pkg/rbac_1_0_0/types"
	"github.com/onosproject/aether-roc-api/pkg/utils"
	modelplugin "github.com/onosproject/config-models/modelplugin/rbac-1.0.0/rbac_1_0_0"
)

// gnmiDeleteRbac deletes an instance of Rbac.
func (i *ServerImpl) gnmiDeleteRbac(ctx context.Context,
	openApiPath string, target types.Target, args ...string) error {

	gnmiSet, err := utils.NewGnmiSetDeleteRequest(openApiPath, string(target), args...)
	if err != nil {
		return err
	}
	log.Infof("gnmiSetRequest %s", gnmiSet.String())
	_, err = i.GnmiClient.Set(ctx, gnmiSet)

	return err
}

// gnmiGetRbac returns an instance of Rbac.
func (i *ServerImpl) gnmiGetRbac(ctx context.Context,
	openApiPath string, target types.Target, args ...string) (*types.Rbac, error) {

	gnmiGet, err := utils.NewGnmiGetRequest(openApiPath, string(target), args...)
	if err != nil {
		return nil, err
	}
	log.Infof("gnmiGetRequest %s", gnmiGet.String())
	gnmiJsonVal, err := utils.GetResponseUpdate(i.GnmiClient.Get(ctx, gnmiGet))
	if err != nil {
		return nil, err
	}
	if gnmiJsonVal == nil {
		return nil, nil
	}

	log.Infof("gNMI Json %s", string(gnmiJsonVal.JsonVal))
	var gnmiResponse modelplugin.Device
	if err = modelplugin.Unmarshal(gnmiJsonVal.JsonVal, &gnmiResponse); err != nil {
		return nil, fmt.Errorf("error unmarshalling gnmiResponse %v", err)
	}
	mpd := ModelPluginDevice{
		device: gnmiResponse,
	}

	return mpd.toRbac(args...)
}

// gnmiPostRbac adds an instance of Rbac.
func (i *ServerImpl) gnmiPostRbac(ctx context.Context, body []byte,
	openApiPath string, target types.Target, args ...string) (*string, error) {

	jsonObj := new(types.Rbac)
	if err := json.Unmarshal(body, jsonObj); err != nil {
		return nil, fmt.Errorf("unable to unmarshal JSON as types.Rbac %v", err)
	}
	gnmiUpdates, err := encodeToGnmiRbac(jsonObj, "", args...)
	if err != nil {
		return nil, fmt.Errorf("unable to convert types.Rbac to gNMI %v", err)
	}
	gnmiSet, err := utils.NewGnmiSetUpdateRequestUpdates(openApiPath, string(target), gnmiUpdates, args...)
	if err != nil {
		return nil, err
	}
	log.Infof("gnmiSetRequest %s", gnmiSet.String())
	gnmiSetResponse, err := i.GnmiClient.Set(ctx, gnmiSet)
	if err != nil {
		return nil, fmt.Errorf(" %v", err)
	}
	return utils.ExtractExtension100(gnmiSetResponse), nil
}

// gnmiDeleteRbacGroup deletes an instance of Rbac_Group.
func (i *ServerImpl) gnmiDeleteRbacGroup(ctx context.Context,
	openApiPath string, target types.Target, args ...string) error {

	gnmiSet, err := utils.NewGnmiSetDeleteRequest(openApiPath, string(target), args...)
	if err != nil {
		return err
	}
	log.Infof("gnmiSetRequest %s", gnmiSet.String())
	_, err = i.GnmiClient.Set(ctx, gnmiSet)

	return err
}

// gnmiGetRbacGroup returns an instance of Rbac_Group.
func (i *ServerImpl) gnmiGetRbacGroup(ctx context.Context,
	openApiPath string, target types.Target, args ...string) (*types.RbacGroup, error) {

	gnmiGet, err := utils.NewGnmiGetRequest(openApiPath, string(target), args...)
	if err != nil {
		return nil, err
	}
	log.Infof("gnmiGetRequest %s", gnmiGet.String())
	gnmiJsonVal, err := utils.GetResponseUpdate(i.GnmiClient.Get(ctx, gnmiGet))
	if err != nil {
		return nil, err
	}
	if gnmiJsonVal == nil {
		return nil, nil
	}

	log.Infof("gNMI Json %s", string(gnmiJsonVal.JsonVal))
	var gnmiResponse modelplugin.Device
	if err = modelplugin.Unmarshal(gnmiJsonVal.JsonVal, &gnmiResponse); err != nil {
		return nil, fmt.Errorf("error unmarshalling gnmiResponse %v", err)
	}
	mpd := ModelPluginDevice{
		device: gnmiResponse,
	}

	return mpd.toRbacGroup(args...)
}

// gnmiPostRbacGroup adds an instance of Rbac_Group.
func (i *ServerImpl) gnmiPostRbacGroup(ctx context.Context, body []byte,
	openApiPath string, target types.Target, args ...string) (*string, error) {

	jsonObj := new(types.RbacGroup)
	if err := json.Unmarshal(body, jsonObj); err != nil {
		return nil, fmt.Errorf("unable to unmarshal JSON as types.Rbac_Group %v", err)
	}
	gnmiUpdates, err := encodeToGnmiRbacGroup(jsonObj, "", args...)
	if err != nil {
		return nil, fmt.Errorf("unable to convert types.RbacGroup to gNMI %v", err)
	}
	gnmiSet, err := utils.NewGnmiSetUpdateRequestUpdates(openApiPath, string(target), gnmiUpdates, args...)
	if err != nil {
		return nil, err
	}
	log.Infof("gnmiSetRequest %s", gnmiSet.String())
	gnmiSetResponse, err := i.GnmiClient.Set(ctx, gnmiSet)
	if err != nil {
		return nil, fmt.Errorf(" %v", err)
	}
	return utils.ExtractExtension100(gnmiSetResponse), nil
}

// gnmiDeleteRbacGroupRole deletes an instance of Rbac_Group_Role.
func (i *ServerImpl) gnmiDeleteRbacGroupRole(ctx context.Context,
	openApiPath string, target types.Target, args ...string) error {

	gnmiSet, err := utils.NewGnmiSetDeleteRequest(openApiPath, string(target), args...)
	if err != nil {
		return err
	}
	log.Infof("gnmiSetRequest %s", gnmiSet.String())
	_, err = i.GnmiClient.Set(ctx, gnmiSet)

	return err
}

// gnmiGetRbacGroupRole returns an instance of Rbac_Group_Role.
func (i *ServerImpl) gnmiGetRbacGroupRole(ctx context.Context,
	openApiPath string, target types.Target, args ...string) (*types.RbacGroupRole, error) {

	gnmiGet, err := utils.NewGnmiGetRequest(openApiPath, string(target), args...)
	if err != nil {
		return nil, err
	}
	log.Infof("gnmiGetRequest %s", gnmiGet.String())
	gnmiJsonVal, err := utils.GetResponseUpdate(i.GnmiClient.Get(ctx, gnmiGet))
	if err != nil {
		return nil, err
	}
	if gnmiJsonVal == nil {
		return nil, nil
	}

	log.Infof("gNMI Json %s", string(gnmiJsonVal.JsonVal))
	var gnmiResponse modelplugin.Device
	if err = modelplugin.Unmarshal(gnmiJsonVal.JsonVal, &gnmiResponse); err != nil {
		return nil, fmt.Errorf("error unmarshalling gnmiResponse %v", err)
	}
	mpd := ModelPluginDevice{
		device: gnmiResponse,
	}

	return mpd.toRbacGroupRole(args...)
}

// gnmiPostRbacGroupRole adds an instance of Rbac_Group_Role.
func (i *ServerImpl) gnmiPostRbacGroupRole(ctx context.Context, body []byte,
	openApiPath string, target types.Target, args ...string) (*string, error) {

	jsonObj := new(types.RbacGroupRole)
	if err := json.Unmarshal(body, jsonObj); err != nil {
		return nil, fmt.Errorf("unable to unmarshal JSON as types.Rbac_Group_Role %v", err)
	}
	gnmiUpdates, err := encodeToGnmiRbacGroupRole(jsonObj, "", args...)
	if err != nil {
		return nil, fmt.Errorf("unable to convert types.RbacGroupRole to gNMI %v", err)
	}
	gnmiSet, err := utils.NewGnmiSetUpdateRequestUpdates(openApiPath, string(target), gnmiUpdates, args...)
	if err != nil {
		return nil, err
	}
	log.Infof("gnmiSetRequest %s", gnmiSet.String())
	gnmiSetResponse, err := i.GnmiClient.Set(ctx, gnmiSet)
	if err != nil {
		return nil, fmt.Errorf(" %v", err)
	}
	return utils.ExtractExtension100(gnmiSetResponse), nil
}

// gnmiDeleteRbacRole deletes an instance of Rbac_Role.
func (i *ServerImpl) gnmiDeleteRbacRole(ctx context.Context,
	openApiPath string, target types.Target, args ...string) error {

	gnmiSet, err := utils.NewGnmiSetDeleteRequest(openApiPath, string(target), args...)
	if err != nil {
		return err
	}
	log.Infof("gnmiSetRequest %s", gnmiSet.String())
	_, err = i.GnmiClient.Set(ctx, gnmiSet)

	return err
}

// gnmiGetRbacRole returns an instance of Rbac_Role.
func (i *ServerImpl) gnmiGetRbacRole(ctx context.Context,
	openApiPath string, target types.Target, args ...string) (*types.RbacRole, error) {

	gnmiGet, err := utils.NewGnmiGetRequest(openApiPath, string(target), args...)
	if err != nil {
		return nil, err
	}
	log.Infof("gnmiGetRequest %s", gnmiGet.String())
	gnmiJsonVal, err := utils.GetResponseUpdate(i.GnmiClient.Get(ctx, gnmiGet))
	if err != nil {
		return nil, err
	}
	if gnmiJsonVal == nil {
		return nil, nil
	}

	log.Infof("gNMI Json %s", string(gnmiJsonVal.JsonVal))
	var gnmiResponse modelplugin.Device
	if err = modelplugin.Unmarshal(gnmiJsonVal.JsonVal, &gnmiResponse); err != nil {
		return nil, fmt.Errorf("error unmarshalling gnmiResponse %v", err)
	}
	mpd := ModelPluginDevice{
		device: gnmiResponse,
	}

	return mpd.toRbacRole(args...)
}

// gnmiPostRbacRole adds an instance of Rbac_Role.
func (i *ServerImpl) gnmiPostRbacRole(ctx context.Context, body []byte,
	openApiPath string, target types.Target, args ...string) (*string, error) {

	jsonObj := new(types.RbacRole)
	if err := json.Unmarshal(body, jsonObj); err != nil {
		return nil, fmt.Errorf("unable to unmarshal JSON as types.Rbac_Role %v", err)
	}
	gnmiUpdates, err := encodeToGnmiRbacRole(jsonObj, "", args...)
	if err != nil {
		return nil, fmt.Errorf("unable to convert types.RbacRole to gNMI %v", err)
	}
	gnmiSet, err := utils.NewGnmiSetUpdateRequestUpdates(openApiPath, string(target), gnmiUpdates, args...)
	if err != nil {
		return nil, err
	}
	log.Infof("gnmiSetRequest %s", gnmiSet.String())
	gnmiSetResponse, err := i.GnmiClient.Set(ctx, gnmiSet)
	if err != nil {
		return nil, fmt.Errorf(" %v", err)
	}
	return utils.ExtractExtension100(gnmiSetResponse), nil
}

// gnmiDeleteRbacRolePermission deletes an instance of Rbac_Role_Permission.
func (i *ServerImpl) gnmiDeleteRbacRolePermission(ctx context.Context,
	openApiPath string, target types.Target, args ...string) error {

	gnmiSet, err := utils.NewGnmiSetDeleteRequest(openApiPath, string(target), args...)
	if err != nil {
		return err
	}
	log.Infof("gnmiSetRequest %s", gnmiSet.String())
	_, err = i.GnmiClient.Set(ctx, gnmiSet)

	return err
}

// gnmiGetRbacRolePermission returns an instance of Rbac_Role_Permission.
func (i *ServerImpl) gnmiGetRbacRolePermission(ctx context.Context,
	openApiPath string, target types.Target, args ...string) (*types.RbacRolePermission, error) {

	gnmiGet, err := utils.NewGnmiGetRequest(openApiPath, string(target), args...)
	if err != nil {
		return nil, err
	}
	log.Infof("gnmiGetRequest %s", gnmiGet.String())
	gnmiJsonVal, err := utils.GetResponseUpdate(i.GnmiClient.Get(ctx, gnmiGet))
	if err != nil {
		return nil, err
	}
	if gnmiJsonVal == nil {
		return nil, nil
	}

	log.Infof("gNMI Json %s", string(gnmiJsonVal.JsonVal))
	var gnmiResponse modelplugin.Device
	if err = modelplugin.Unmarshal(gnmiJsonVal.JsonVal, &gnmiResponse); err != nil {
		return nil, fmt.Errorf("error unmarshalling gnmiResponse %v", err)
	}
	mpd := ModelPluginDevice{
		device: gnmiResponse,
	}

	return mpd.toRbacRolePermission(args...)
}

// gnmiPostRbacRolePermission adds an instance of Rbac_Role_Permission.
func (i *ServerImpl) gnmiPostRbacRolePermission(ctx context.Context, body []byte,
	openApiPath string, target types.Target, args ...string) (*string, error) {

	jsonObj := new(types.RbacRolePermission)
	if err := json.Unmarshal(body, jsonObj); err != nil {
		return nil, fmt.Errorf("unable to unmarshal JSON as types.Rbac_Role_Permission %v", err)
	}
	gnmiUpdates, err := encodeToGnmiRbacRolePermission(jsonObj, "", args...)
	if err != nil {
		return nil, fmt.Errorf("unable to convert types.RbacRolePermission to gNMI %v", err)
	}
	gnmiSet, err := utils.NewGnmiSetUpdateRequestUpdates(openApiPath, string(target), gnmiUpdates, args...)
	if err != nil {
		return nil, err
	}
	log.Infof("gnmiSetRequest %s", gnmiSet.String())
	gnmiSetResponse, err := i.GnmiClient.Set(ctx, gnmiSet)
	if err != nil {
		return nil, fmt.Errorf(" %v", err)
	}
	return utils.ExtractExtension100(gnmiSetResponse), nil
}

// gnmiDeleteTarget deletes an instance of target.
func (i *ServerImpl) gnmiDeleteTarget(ctx context.Context,
	openApiPath string, target types.Target, args ...string) error {

	gnmiSet, err := utils.NewGnmiSetDeleteRequest(openApiPath, string(target), args...)
	if err != nil {
		return err
	}
	log.Infof("gnmiSetRequest %s", gnmiSet.String())
	_, err = i.GnmiClient.Set(ctx, gnmiSet)

	return err
}

// gnmiGetTarget returns an instance of target.
func (i *ServerImpl) gnmiGetTarget(ctx context.Context,
	openApiPath string, target types.Target, args ...string) (*types.Target, error) {

	gnmiGet, err := utils.NewGnmiGetRequest(openApiPath, string(target), args...)
	if err != nil {
		return nil, err
	}
	log.Infof("gnmiGetRequest %s", gnmiGet.String())
	gnmiJsonVal, err := utils.GetResponseUpdate(i.GnmiClient.Get(ctx, gnmiGet))
	if err != nil {
		return nil, err
	}
	if gnmiJsonVal == nil {
		return nil, nil
	}

	log.Infof("gNMI Json %s", string(gnmiJsonVal.JsonVal))
	var gnmiResponse modelplugin.Device
	if err = modelplugin.Unmarshal(gnmiJsonVal.JsonVal, &gnmiResponse); err != nil {
		return nil, fmt.Errorf("error unmarshalling gnmiResponse %v", err)
	}
	mpd := ModelPluginDevice{
		device: gnmiResponse,
	}

	return mpd.toTarget(args...)
}

// gnmiPostTarget adds an instance of target.
func (i *ServerImpl) gnmiPostTarget(ctx context.Context, body []byte,
	openApiPath string, target types.Target, args ...string) (*string, error) {

	return nil, fmt.Errorf("Not implemented")

}

type Translator interface {
	toRbac(args ...string) (*types.Rbac, error)
	toRbacGroup(args ...string) (*types.RbacGroup, error)
	toRbacGroupRole(args ...string) (*types.RbacGroupRole, error)
	toRbacRole(args ...string) (*types.RbacRole, error)
	toRbacRolePermission(args ...string) (*types.RbacRolePermission, error)
	toTarget(args ...string) (*types.Target, error)
}

// Not generating param-types
// Not generating request-bodies
