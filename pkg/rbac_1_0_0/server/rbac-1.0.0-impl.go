// Code generated by oapi-codegen. DO NOT EDIT.
// Package server provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"reflect"

	"github.com/labstack/echo/v4"
	"github.com/onosproject/aether-roc-api/pkg/rbac_1_0_0/types"
	"github.com/onosproject/aether-roc-api/pkg/southbound"
	"github.com/onosproject/aether-roc-api/pkg/utils"
	externalRef0 "github.com/onosproject/config-models/modelplugin/rbac-1.0.0/rbac_1_0_0"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

//Ignoring AdditionalPropertyTarget

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
	var gnmiResponse externalRef0.Device
	if err = externalRef0.Unmarshal(gnmiJsonVal.JsonVal, &gnmiResponse); err != nil {
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
	gnmiUpdates, err := EncodeToGnmiRbac(jsonObj, false, target, "", args...)
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
	var gnmiResponse externalRef0.Device
	if err = externalRef0.Unmarshal(gnmiJsonVal.JsonVal, &gnmiResponse); err != nil {
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
	gnmiUpdates, err := EncodeToGnmiRbacGroup(jsonObj, false, target, "", args...)
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
	var gnmiResponse externalRef0.Device
	if err = externalRef0.Unmarshal(gnmiJsonVal.JsonVal, &gnmiResponse); err != nil {
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
	gnmiUpdates, err := EncodeToGnmiRbacGroupRole(jsonObj, false, target, "", args...)
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
	var gnmiResponse externalRef0.Device
	if err = externalRef0.Unmarshal(gnmiJsonVal.JsonVal, &gnmiResponse); err != nil {
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
	gnmiUpdates, err := EncodeToGnmiRbacRole(jsonObj, false, target, "", args...)
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
	var gnmiResponse externalRef0.Device
	if err = externalRef0.Unmarshal(gnmiJsonVal.JsonVal, &gnmiResponse); err != nil {
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
	gnmiUpdates, err := EncodeToGnmiRbacRolePermission(jsonObj, false, target, "", args...)
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
	var gnmiResponse externalRef0.Device
	if err = externalRef0.Unmarshal(gnmiJsonVal.JsonVal, &gnmiResponse); err != nil {
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

//Ignoring RequestBodyRbac

//Ignoring RequestBodyRbacGroup

//Ignoring RequestBodyRbacGroupRole

//Ignoring RequestBodyRbacRole

//Ignoring RequestBodyRbacRolePermission

type Translator interface {
	toAdditionalPropertyTarget(args ...string) (*types.AdditionalPropertyTarget, error)
	toRbac(args ...string) (*types.Rbac, error)
	toRbacGroup(args ...string) (*types.RbacGroup, error)
	toRbacGroupRole(args ...string) (*types.RbacGroupRole, error)
	toRbacRole(args ...string) (*types.RbacRole, error)
	toRbacRolePermission(args ...string) (*types.RbacRolePermission, error)
	toTarget(args ...string) (*types.Target, error)
}

// Not generating param-types
// Not generating request-bodies

// Not generating additional-properties
// Not generating additional-properties
// server-interface template override

const authorization = "Authorization"

// Implement the Server Interface for access to gNMI
var log = logging.GetLogger("model_0_0_0")

// ServerImpl -
type ServerImpl struct {
	GnmiClient southbound.GnmiClient
}

// DeleteRbac impl of gNMI access at /rbac/v1.0.0/{target}/rbac
func (i *ServerImpl) DeleteRbac(ctx echo.Context, target types.Target) error {

	var response interface{}
	var err error

	// Response
	err = i.gnmiDeleteRbac(utils.NewGnmiContext(ctx), "/rbac/v1.0.0/{target}/rbac", target)

	if err != nil {
		return utils.ConvertGrpcError(err)
	}
	// It's not enough to check if response==nil - see https://medium.com/@glucn/golang-an-interface-holding-a-nil-value-is-not-nil-bb151f472cc7
	if reflect.ValueOf(response).Kind() == reflect.Ptr && reflect.ValueOf(response).IsNil() {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	log.Infof("DeleteRbac")
	return ctx.JSON(http.StatusOK, response)
}

// GetRbac impl of gNMI access at /rbac/v1.0.0/{target}/rbac
func (i *ServerImpl) GetRbac(ctx echo.Context, target types.Target) error {

	var response interface{}
	var err error

	// Response GET OK 200
	response, err = i.gnmiGetRbac(utils.NewGnmiContext(ctx), "/rbac/v1.0.0/{target}/rbac", target)

	if err != nil {
		return utils.ConvertGrpcError(err)
	}
	// It's not enough to check if response==nil - see https://medium.com/@glucn/golang-an-interface-holding-a-nil-value-is-not-nil-bb151f472cc7
	if reflect.ValueOf(response).Kind() == reflect.Ptr && reflect.ValueOf(response).IsNil() {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	log.Infof("GetRbac")
	return ctx.JSON(http.StatusOK, response)
}

// PostRbac impl of gNMI access at /rbac/v1.0.0/{target}/rbac
func (i *ServerImpl) PostRbac(ctx echo.Context, target types.Target) error {

	var response interface{}
	var err error

	// Response created

	body, err := utils.ReadRequestBody(ctx.Request().Body)
	if err != nil {
		return err
	}
	extension100, err := i.gnmiPostRbac(utils.NewGnmiContext(ctx), body, "/rbac/v1.0.0/{target}/rbac", target)
	if err == nil {
		log.Infof("Post succeded %s", *extension100)
		return ctx.JSON(http.StatusOK, extension100)
	}

	if err != nil {
		return utils.ConvertGrpcError(err)
	}
	// It's not enough to check if response==nil - see https://medium.com/@glucn/golang-an-interface-holding-a-nil-value-is-not-nil-bb151f472cc7
	if reflect.ValueOf(response).Kind() == reflect.Ptr && reflect.ValueOf(response).IsNil() {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	log.Infof("PostRbac")
	return ctx.JSON(http.StatusOK, response)
}

// DeleteRbacGroup impl of gNMI access at /rbac/v1.0.0/{target}/rbac/group/{groupid}
func (i *ServerImpl) DeleteRbacGroup(ctx echo.Context, target types.Target, groupid string) error {

	var response interface{}
	var err error

	// Response
	err = i.gnmiDeleteRbacGroup(utils.NewGnmiContext(ctx), "/rbac/v1.0.0/{target}/rbac/group/{groupid}", target, groupid)

	if err != nil {
		return utils.ConvertGrpcError(err)
	}
	// It's not enough to check if response==nil - see https://medium.com/@glucn/golang-an-interface-holding-a-nil-value-is-not-nil-bb151f472cc7
	if reflect.ValueOf(response).Kind() == reflect.Ptr && reflect.ValueOf(response).IsNil() {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	log.Infof("DeleteRbacGroup")
	return ctx.JSON(http.StatusOK, response)
}

// GetRbacGroup impl of gNMI access at /rbac/v1.0.0/{target}/rbac/group/{groupid}
func (i *ServerImpl) GetRbacGroup(ctx echo.Context, target types.Target, groupid string) error {

	var response interface{}
	var err error

	// Response GET OK 200
	response, err = i.gnmiGetRbacGroup(utils.NewGnmiContext(ctx), "/rbac/v1.0.0/{target}/rbac/group/{groupid}", target, groupid)

	if err != nil {
		return utils.ConvertGrpcError(err)
	}
	// It's not enough to check if response==nil - see https://medium.com/@glucn/golang-an-interface-holding-a-nil-value-is-not-nil-bb151f472cc7
	if reflect.ValueOf(response).Kind() == reflect.Ptr && reflect.ValueOf(response).IsNil() {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	log.Infof("GetRbacGroup")
	return ctx.JSON(http.StatusOK, response)
}

// PostRbacGroup impl of gNMI access at /rbac/v1.0.0/{target}/rbac/group/{groupid}
func (i *ServerImpl) PostRbacGroup(ctx echo.Context, target types.Target, groupid string) error {

	var response interface{}
	var err error

	// Response created

	body, err := utils.ReadRequestBody(ctx.Request().Body)
	if err != nil {
		return err
	}
	extension100, err := i.gnmiPostRbacGroup(utils.NewGnmiContext(ctx), body, "/rbac/v1.0.0/{target}/rbac/group/{groupid}", target, groupid)
	if err == nil {
		log.Infof("Post succeded %s", *extension100)
		return ctx.JSON(http.StatusOK, extension100)
	}

	if err != nil {
		return utils.ConvertGrpcError(err)
	}
	// It's not enough to check if response==nil - see https://medium.com/@glucn/golang-an-interface-holding-a-nil-value-is-not-nil-bb151f472cc7
	if reflect.ValueOf(response).Kind() == reflect.Ptr && reflect.ValueOf(response).IsNil() {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	log.Infof("PostRbacGroup")
	return ctx.JSON(http.StatusOK, response)
}

// DeleteRbacGroupRole impl of gNMI access at /rbac/v1.0.0/{target}/rbac/group/{groupid}/role/{roleid}
func (i *ServerImpl) DeleteRbacGroupRole(ctx echo.Context, target types.Target, groupid string, roleid string) error {

	var response interface{}
	var err error

	// Response
	err = i.gnmiDeleteRbacGroupRole(utils.NewGnmiContext(ctx), "/rbac/v1.0.0/{target}/rbac/group/{groupid}/role/{roleid}", target, groupid, roleid)

	if err != nil {
		return utils.ConvertGrpcError(err)
	}
	// It's not enough to check if response==nil - see https://medium.com/@glucn/golang-an-interface-holding-a-nil-value-is-not-nil-bb151f472cc7
	if reflect.ValueOf(response).Kind() == reflect.Ptr && reflect.ValueOf(response).IsNil() {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	log.Infof("DeleteRbacGroupRole")
	return ctx.JSON(http.StatusOK, response)
}

// GetRbacGroupRole impl of gNMI access at /rbac/v1.0.0/{target}/rbac/group/{groupid}/role/{roleid}
func (i *ServerImpl) GetRbacGroupRole(ctx echo.Context, target types.Target, groupid string, roleid string) error {

	var response interface{}
	var err error

	// Response GET OK 200
	response, err = i.gnmiGetRbacGroupRole(utils.NewGnmiContext(ctx), "/rbac/v1.0.0/{target}/rbac/group/{groupid}/role/{roleid}", target, groupid, roleid)

	if err != nil {
		return utils.ConvertGrpcError(err)
	}
	// It's not enough to check if response==nil - see https://medium.com/@glucn/golang-an-interface-holding-a-nil-value-is-not-nil-bb151f472cc7
	if reflect.ValueOf(response).Kind() == reflect.Ptr && reflect.ValueOf(response).IsNil() {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	log.Infof("GetRbacGroupRole")
	return ctx.JSON(http.StatusOK, response)
}

// PostRbacGroupRole impl of gNMI access at /rbac/v1.0.0/{target}/rbac/group/{groupid}/role/{roleid}
func (i *ServerImpl) PostRbacGroupRole(ctx echo.Context, target types.Target, groupid string, roleid string) error {

	var response interface{}
	var err error

	// Response created

	body, err := utils.ReadRequestBody(ctx.Request().Body)
	if err != nil {
		return err
	}
	extension100, err := i.gnmiPostRbacGroupRole(utils.NewGnmiContext(ctx), body, "/rbac/v1.0.0/{target}/rbac/group/{groupid}/role/{roleid}", target, groupid, roleid)
	if err == nil {
		log.Infof("Post succeded %s", *extension100)
		return ctx.JSON(http.StatusOK, extension100)
	}

	if err != nil {
		return utils.ConvertGrpcError(err)
	}
	// It's not enough to check if response==nil - see https://medium.com/@glucn/golang-an-interface-holding-a-nil-value-is-not-nil-bb151f472cc7
	if reflect.ValueOf(response).Kind() == reflect.Ptr && reflect.ValueOf(response).IsNil() {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	log.Infof("PostRbacGroupRole")
	return ctx.JSON(http.StatusOK, response)
}

// DeleteRbacRole impl of gNMI access at /rbac/v1.0.0/{target}/rbac/role/{roleid}
func (i *ServerImpl) DeleteRbacRole(ctx echo.Context, target types.Target, roleid string) error {

	var response interface{}
	var err error

	// Response
	err = i.gnmiDeleteRbacRole(utils.NewGnmiContext(ctx), "/rbac/v1.0.0/{target}/rbac/role/{roleid}", target, roleid)

	if err != nil {
		return utils.ConvertGrpcError(err)
	}
	// It's not enough to check if response==nil - see https://medium.com/@glucn/golang-an-interface-holding-a-nil-value-is-not-nil-bb151f472cc7
	if reflect.ValueOf(response).Kind() == reflect.Ptr && reflect.ValueOf(response).IsNil() {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	log.Infof("DeleteRbacRole")
	return ctx.JSON(http.StatusOK, response)
}

// GetRbacRole impl of gNMI access at /rbac/v1.0.0/{target}/rbac/role/{roleid}
func (i *ServerImpl) GetRbacRole(ctx echo.Context, target types.Target, roleid string) error {

	var response interface{}
	var err error

	// Response GET OK 200
	response, err = i.gnmiGetRbacRole(utils.NewGnmiContext(ctx), "/rbac/v1.0.0/{target}/rbac/role/{roleid}", target, roleid)

	if err != nil {
		return utils.ConvertGrpcError(err)
	}
	// It's not enough to check if response==nil - see https://medium.com/@glucn/golang-an-interface-holding-a-nil-value-is-not-nil-bb151f472cc7
	if reflect.ValueOf(response).Kind() == reflect.Ptr && reflect.ValueOf(response).IsNil() {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	log.Infof("GetRbacRole")
	return ctx.JSON(http.StatusOK, response)
}

// PostRbacRole impl of gNMI access at /rbac/v1.0.0/{target}/rbac/role/{roleid}
func (i *ServerImpl) PostRbacRole(ctx echo.Context, target types.Target, roleid string) error {

	var response interface{}
	var err error

	// Response created

	body, err := utils.ReadRequestBody(ctx.Request().Body)
	if err != nil {
		return err
	}
	extension100, err := i.gnmiPostRbacRole(utils.NewGnmiContext(ctx), body, "/rbac/v1.0.0/{target}/rbac/role/{roleid}", target, roleid)
	if err == nil {
		log.Infof("Post succeded %s", *extension100)
		return ctx.JSON(http.StatusOK, extension100)
	}

	if err != nil {
		return utils.ConvertGrpcError(err)
	}
	// It's not enough to check if response==nil - see https://medium.com/@glucn/golang-an-interface-holding-a-nil-value-is-not-nil-bb151f472cc7
	if reflect.ValueOf(response).Kind() == reflect.Ptr && reflect.ValueOf(response).IsNil() {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	log.Infof("PostRbacRole")
	return ctx.JSON(http.StatusOK, response)
}

// DeleteRbacRolePermission impl of gNMI access at /rbac/v1.0.0/{target}/rbac/role/{roleid}/permission
func (i *ServerImpl) DeleteRbacRolePermission(ctx echo.Context, target types.Target, roleid string) error {

	var response interface{}
	var err error

	// Response
	err = i.gnmiDeleteRbacRolePermission(utils.NewGnmiContext(ctx), "/rbac/v1.0.0/{target}/rbac/role/{roleid}/permission", target, roleid)

	if err != nil {
		return utils.ConvertGrpcError(err)
	}
	// It's not enough to check if response==nil - see https://medium.com/@glucn/golang-an-interface-holding-a-nil-value-is-not-nil-bb151f472cc7
	if reflect.ValueOf(response).Kind() == reflect.Ptr && reflect.ValueOf(response).IsNil() {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	log.Infof("DeleteRbacRolePermission")
	return ctx.JSON(http.StatusOK, response)
}

// GetRbacRolePermission impl of gNMI access at /rbac/v1.0.0/{target}/rbac/role/{roleid}/permission
func (i *ServerImpl) GetRbacRolePermission(ctx echo.Context, target types.Target, roleid string) error {

	var response interface{}
	var err error

	// Response GET OK 200
	response, err = i.gnmiGetRbacRolePermission(utils.NewGnmiContext(ctx), "/rbac/v1.0.0/{target}/rbac/role/{roleid}/permission", target, roleid)

	if err != nil {
		return utils.ConvertGrpcError(err)
	}
	// It's not enough to check if response==nil - see https://medium.com/@glucn/golang-an-interface-holding-a-nil-value-is-not-nil-bb151f472cc7
	if reflect.ValueOf(response).Kind() == reflect.Ptr && reflect.ValueOf(response).IsNil() {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	log.Infof("GetRbacRolePermission")
	return ctx.JSON(http.StatusOK, response)
}

// PostRbacRolePermission impl of gNMI access at /rbac/v1.0.0/{target}/rbac/role/{roleid}/permission
func (i *ServerImpl) PostRbacRolePermission(ctx echo.Context, target types.Target, roleid string) error {

	var response interface{}
	var err error

	// Response created

	body, err := utils.ReadRequestBody(ctx.Request().Body)
	if err != nil {
		return err
	}
	extension100, err := i.gnmiPostRbacRolePermission(utils.NewGnmiContext(ctx), body, "/rbac/v1.0.0/{target}/rbac/role/{roleid}/permission", target, roleid)
	if err == nil {
		log.Infof("Post succeded %s", *extension100)
		return ctx.JSON(http.StatusOK, extension100)
	}

	if err != nil {
		return utils.ConvertGrpcError(err)
	}
	// It's not enough to check if response==nil - see https://medium.com/@glucn/golang-an-interface-holding-a-nil-value-is-not-nil-bb151f472cc7
	if reflect.ValueOf(response).Kind() == reflect.Ptr && reflect.ValueOf(response).IsNil() {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	log.Infof("PostRbacRolePermission")
	return ctx.JSON(http.StatusOK, response)
}

// register template override
