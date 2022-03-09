// Code generated by oapi-codegen. DO NOT EDIT.
// Package server provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package server

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/ghodss/yaml"
	"github.com/labstack/echo/v4"
	aether_2_0_0 "github.com/onosproject/aether-roc-api/pkg/aether_2_0_0/server"
	aether_4_0_0 "github.com/onosproject/aether-roc-api/pkg/aether_4_0_0/server"
	externalRef0 "github.com/onosproject/aether-roc-api/pkg/toplevel/types"
	"github.com/onosproject/onos-api/go/onos/config/admin"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/openconfig/gnmi/proto/gnmi"
	htmltemplate "html/template"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// server-interface template override

import (
	"github.com/onosproject/aether-roc-api/pkg/southbound"
	"github.com/onosproject/aether-roc-api/pkg/utils"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"reflect"
)

type HtmlData struct {
	File        string
	Description string
}

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

// grpcGetTransactions returns a list of Transactions.
func (i *ServerImpl) grpcGetTransactions(ctx context.Context) (*externalRef0.TransactionList, error) {
	log.Infof("grpcGetTransactions - subscribe=false")

	// At present (Jan '22) ListTransactions is not implemented - use ListNetworkChanges
	//stream, err := i.ConfigClient.ListNetworkChanges(ctx, &diags.ListNetworkChangeRequest{
	//	Subscribe: false,
	//})

	stream, err := i.ConfigClient.ListTransactions(ctx, &admin.ListTransactionsRequest{
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	})
	if err != nil {
		return nil, errors.FromGRPC(err)
	}
	transactionList := make(externalRef0.TransactionList, 0)
	for {
		networkChange, err := stream.Recv()
		if err == io.EOF || networkChange == nil {
			break
		}
		created := networkChange.GetTransaction().GetCreated()
		updated := networkChange.GetTransaction().GetUpdated()
		deleted := networkChange.GetTransaction().GetDeleted()
		username := networkChange.GetTransaction().GetUsername()
		key := networkChange.GetTransaction().GetKey()
		//version := networkChange.GetTransaction().GetVersion()

		objMeta := struct {
			Created  *time.Time `json:"created,omitempty"`
			Deleted  *time.Time `json:"deleted,omitempty"`
			Key      *string    `json:"key,omitempty"`
			Revision *externalRef0.Revision `json:"revision,omitempty"`
			Updated *time.Time `json:"updated,omitempty"`
			Version *int64    `json:"version,omitempty"`
		}{
			Created:  &created,
			Deleted:  deleted,
			Key:      &key,
			Revision: nil, //TODO: need to implement
			Updated:  &updated,
			Version:  nil,
		}

		transaction := externalRef0.Transaction{
			Details:  nil,
			Id:       string(networkChange.GetTransaction().GetID()),
			Index:    int64(networkChange.GetTransaction().GetIndex()),
			Meta:     objMeta,
			Status:   nil,
			Strategy: nil,
			Username: &username,
		}

		transactionList = append(transactionList, transaction)
	}

	return &transactionList, nil
}

// ServerImpl -
type ServerImpl struct {
	GnmiClient    southbound.GnmiClient
	ConfigClient  admin.TransactionServiceClient
	GnmiTimeout   time.Duration
	Authorization bool
}

// PatchAetherRocApi impl of gNMI access at /aether-roc-api
func (i *ServerImpl) PatchAetherRocApi(ctx echo.Context) error {

	var response interface{}
	var err error

	gnmiCtx, cancel := utils.NewGnmiContext(ctx, i.GnmiTimeout)
	defer cancel()

	// Response patched
	body, err := utils.ReadRequestBody(ctx.Request().Body)
	if err != nil {
		return err
	}
	response, err = i.gnmiPatchAetherRocAPI(gnmiCtx, body, "/aether-roc-api")
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

	gnmiCtx, cancel := utils.NewGnmiContext(ctx, i.GnmiTimeout)
	defer cancel()

	// Response GET OK 200
	response, err = i.gnmiGetTargets(gnmiCtx)
	if err != nil {
		return utils.ConvertGrpcError(err)
	}
	log.Infof("GetTargets")
	return ctx.JSON(http.StatusOK, response)
}

func (i *ServerImpl) GetTransactions(ctx echo.Context) error {
	var response interface{}
	var err error

	gnmiCtx, cancel := utils.NewGnmiContext(ctx, i.GnmiTimeout)
	defer cancel()

	// Response GET OK 200
	response, err = i.grpcGetTransactions(gnmiCtx)
	if err != nil {
		return utils.ConvertGrpcError(err)
	}
	log.Infof("GetTransactions")
	return ctx.JSON(http.StatusOK, response)
}

func (i *ServerImpl) PostSdcoreSynchronize(httpContext echo.Context) error {

	// Response GET OK 200
	if i.Authorization {
		if err := checkAuthorization(httpContext, "AetherROCAdmin"); err != nil {
			return err
		}
	}

	address := fmt.Sprintf("http://%s:8080/synchronize", httpContext.Param("service"))
	resp, err := http.Post(address, "application/json", nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("error calling %s. %v", address, err))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("error reading body %s. %v", address, err))
	}

	log.Infof("PostSdcoreSynchronize to %s %s %s", httpContext.Param("service"), resp.Status, string(body))
	respStruct := struct {
		Response string `json:"response"`
	}{Response: string(body)}
	return httpContext.JSON(resp.StatusCode, &respStruct)
}

func (i *ServerImpl) GetSpec(ctx echo.Context) error {
	response, err := GetSwagger()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	log.Infof("GetSpec")
	return acceptTypes(ctx, response)
}

func (i *ServerImpl) GetAether200Spec(ctx echo.Context) error {
	response, err := aether_2_0_0.GetSwagger()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return acceptTypes(ctx, response)
}

func (i *ServerImpl) GetAether400Spec(ctx echo.Context) error {
	response, err := aether_4_0_0.GetSwagger()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return acceptTypes(ctx, response)
}

func acceptTypes(ctx echo.Context, response *openapi3.T) error {
	acceptType := ctx.Request().Header.Get("Accept")

	if strings.Contains(acceptType, "application/json") {
		return ctx.JSONPretty(http.StatusOK, response, "  ")
	} else if strings.Contains(acceptType, "text/html") {
		templateText, err := ioutil.ReadFile("assets/html-page.tpl")
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "unable to load template %s", err)
		}
		specTemplate, err := htmltemplate.New("spectemplate").Parse(string(templateText))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "error parsing template %s", err)
		}
		var b bytes.Buffer
		specTemplate.Execute(&b, HtmlData{
			File:        ctx.Request().RequestURI[1:],
			Description: "Aether ROC API",
		})
		ctx.Response().Header().Set("Content-Type", "text/html")
		return ctx.HTMLBlob(http.StatusOK, b.Bytes())
	} else if strings.Contains(acceptType, "application/yaml") || strings.Contains(acceptType, "*/*") {
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
		fmt.Sprintf("only application/yaml, application/json and text/html encoding supported. "+
			"No match for %s", acceptType))
}

// register template override
