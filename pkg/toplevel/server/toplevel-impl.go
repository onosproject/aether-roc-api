// Code generated by oapi-codegen. DO NOT EDIT.
// Package server provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/ghodss/yaml"
	"github.com/labstack/echo/v4"
	aether_2_1_0 "github.com/onosproject/aether-roc-api/pkg/aether_2_1_0/server"
	aether_3_0_0 "github.com/onosproject/aether-roc-api/pkg/aether_3_0_0/server"
	externalRef0 "github.com/onosproject/aether-roc-api/pkg/toplevel/types"
	"github.com/openconfig/gnmi/proto/gnmi"
	"net/http"
)

// server-interface template override

import (
	"github.com/onosproject/aether-roc-api/pkg/southbound"
	"github.com/onosproject/aether-roc-api/pkg/utils"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"reflect"
)

const authorization = "Authorization"

// Implement the Server Interface for access to gNMI
var log = logging.GetLogger("toplevel")

// gnmiGetTargets returns a list of Targets.
func (i *ServerImpl) gnmiGetTargets(ctx context.Context) (*externalRef0.TargetsNames, error) {
	gnmiGet := new(gnmi.GetRequest)
	gnmiGet.Encoding = gnmi.Encoding_PROTO
	gnmiGet.Path = make([]*gnmi.Path, 1)
	gnmiGet.Path[0] = &gnmi.Path{
		Target: "*",
	}

	log.Infof("gnmiGetRequest %s", gnmiGet.String())
	gnmiVal, err := utils.GetResponseUpdate(i.GnmiClient.Get(ctx, gnmiGet))
	if err != nil {
		return nil, err
	}
	gnmiLeafListStr, ok := gnmiVal.Value.(*gnmi.TypedValue_LeaflistVal)
	if !ok {
		return nil, fmt.Errorf("expecting a leaf list. Got %s", gnmiVal.String())
	}

	log.Infof("gNMI %s", gnmiLeafListStr.LeaflistVal.String())
	targetsNames := make(externalRef0.TargetsNames, 0)
	for _, elem := range gnmiLeafListStr.LeaflistVal.Element {
		targetName := elem.GetStringVal()
		targetsNames = append(targetsNames, externalRef0.TargetName{
			Name: &targetName,
		})
	}
	return &targetsNames, nil
}

// ServerImpl -
type ServerImpl struct {
	GnmiClient southbound.GnmiClient
}

// PatchAetherRocApi impl of gNMI access at /aether-roc-api
func (i *ServerImpl) PatchAetherRocApi(ctx echo.Context) error {

	var response interface{}
	var err error

	// Response patched
	body, err := utils.ReadRequestBody(ctx.Request().Body)
	if err != nil {
		return err
	}
	response, err = i.gnmiPatchAetherRocAPI(utils.NewGnmiContext(ctx), body, "/aether-roc-api")
	if err != nil {
		return utils.ConvertGrpcError(err)
	}
	// It's not enough to check if response==nil - see https://medium.com/@glucn/golang-an-interface-holding-a-nil-value-is-not-nil-bb151f472cc7
	if reflect.ValueOf(response).Kind() == reflect.Ptr && reflect.ValueOf(response).IsNil() {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	log.Infof("PatchAetherRocApi")
	return ctx.JSON(http.StatusOK, response)
}

func (i *ServerImpl) GetTargets(ctx echo.Context) error {
	var response interface{}
	var err error

	// Response GET OK 200
	response, err = i.gnmiGetTargets(utils.NewGnmiContext(ctx))
	if err != nil {
		return utils.ConvertGrpcError(err)
	}
	log.Infof("GetTargets")
	return ctx.JSON(http.StatusOK, response)
}

func (i *ServerImpl) GetSpec(ctx echo.Context) error {
	response, err := GetSwagger()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	log.Infof("GetSpec")
	return acceptTypes(ctx, response)
}

func (i *ServerImpl) GetAether210Spec(ctx echo.Context) error {
	response, err := aether_2_1_0.GetSwagger()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return acceptTypes(ctx, response)
}

func (i *ServerImpl) GetAether300Spec(ctx echo.Context) error {
	response, err := aether_3_0_0.GetSwagger()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return acceptTypes(ctx, response)
}

func acceptTypes(ctx echo.Context, response *openapi3.T) error {
	acceptType := ctx.Request().Header.Get("Accept")

	if acceptType == "application/json" {
		return ctx.JSONPretty(http.StatusOK, response, "  ")
	} else if acceptType == "application/yaml" || acceptType == "*/*"{
		jsonFirst, err := json.Marshal(response)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		yamlResp, err := yaml.JSONToYAML(jsonFirst)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		ctx.Response().Header().Set("Content-Type", "application/yaml")
		return ctx.HTMLBlob(http.StatusOK, yamlResp)
	}
	return echo.NewHTTPError(http.StatusNotImplemented,
		fmt.Sprintf("encoding %s not yet implemented", ctx.Request().Header.Get("Accept")))
}

// register template override
