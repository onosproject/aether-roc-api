// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
//

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
	aether_2_1_0 "github.com/onosproject/aether-roc-api/pkg/aether_2_1_0/server"
	aether_4_0_0 "github.com/onosproject/aether-roc-api/pkg/aether_4_0_0/server"
	app_gtwy "github.com/onosproject/aether-roc-api/pkg/app_gtwy/server"
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

// HTMLData -
type HTMLData struct {
	File        string
	Description string
}

const authorization = "Authorization"

// Implement the Server Interface for access to gNMI
var log = logging.GetLogger("toplevel")

// gnmiGetTargets returns a list of Targets.
func (i *TopLevelServer) gnmiGetTargets(ctx context.Context) (*externalRef0.TargetsNames, error) {
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
func (i *TopLevelServer) grpcGetTransactions(ctx context.Context) (*externalRef0.TransactionList, error) {
	log.Infof("grpcGetTransactions - subscribe=false")

	stream, err := i.ConfigClient.ListTransactions(ctx, &admin.ListTransactionsRequest{})
	if err != nil {
		return nil, errors.FromGRPC(err)
	}
	transactionList := make(externalRef0.TransactionList, 0)
	for {
		networkChange, err := stream.Recv()
		if err == io.EOF || networkChange == nil {
			break
		}
		transactionList = append(transactionList, convertTrasaction(networkChange))
	}

	return &transactionList, nil
}

func convertTrasaction(networkChange *admin.ListTransactionsResponse) externalRef0.Transaction {

	if networkChange.GetTransaction() == nil {
		return externalRef0.Transaction{}
	}
	created := networkChange.GetTransaction().GetCreated()
	updated := networkChange.GetTransaction().GetUpdated()
	deleted := networkChange.GetTransaction().GetDeleted()
	username := networkChange.GetTransaction().GetUsername()
	key := networkChange.GetTransaction().GetKey()
	version := (int64)(networkChange.GetTransaction().GetVersion())
	revision := (externalRef0.Revision)(networkChange.GetTransaction().GetRevision())

	objMeta := struct {
		Created  *time.Time             `json:"created,omitempty"`
		Deleted  *time.Time             `json:"deleted,omitempty"`
		Key      *string                `json:"key,omitempty"`
		Revision *externalRef0.Revision `json:"revision,omitempty"`
		Updated  *time.Time             `json:"updated,omitempty"`
		Version  *int64                 `json:"version,omitempty"`
	}{
		Created:  &created,
		Deleted:  deleted,
		Key:      &key,
		Revision: &revision,
		Updated:  &updated,
		Version:  &version,
	}

	changeTrasactions := make(externalRef0.ChangeTransaction, 0)
	if (networkChange.GetTransaction().GetChange() != nil) && (networkChange.GetTransaction().GetChange().Values != nil) {
		for targetID, pathValues := range networkChange.GetTransaction().GetChange().Values {
			pValues := make(externalRef0.PathValues, 0)
			for targetName, pValue := range pathValues.GetValues() {
				path := pValue.GetPath()
				bytes := pValue.GetValue().Bytes
				valueType := pValue.GetValue().Type.String()
				var typeOpts []externalRef0.TypeOpts
				for _, tOpts := range pValue.GetValue().TypeOpts {
					typeOpts = append(typeOpts, (externalRef0.TypeOpts)(tOpts))
				}
				pDeleted := pValue.GetDeleted()

				typedValue := new(externalRef0.TypedValue)
				typedValue.Bytes = (*externalRef0.Bytes)(&bytes)
				typedValue.TypeOpts = &typeOpts
				typedValue.Type = (*externalRef0.ValueType)(&valueType)

				pathValue := new(externalRef0.PathValue)
				pathValue.Path = (*externalRef0.Path)(&path)
				pathValue.Value = typedValue
				pathValue.Deleted = (*externalRef0.Deleted)(&pDeleted)

				tName := targetName
				pTarget := new(externalRef0.PathTarget)
				pTarget.Path = &tName
				pTarget.PathValue = pathValue
				pValues = append(pValues, *pTarget)
			}
			tID := targetID
			cTarget := new(externalRef0.ChangeTarget)
			cTarget.TargetName = (*string)(&tID)
			cTarget.PathValues = &pValues
			changeTrasactions = append(changeTrasactions, *cTarget)
		}
	}

	rollBackIndex := (externalRef0.Index)(networkChange.GetTransaction().GetRollback().GetRollbackIndex())
	rollback := externalRef0.RollbackTransaction{RollbackIndex: &rollBackIndex}

	details := externalRef0.Details{
		Change:   &changeTrasactions,
		Rollback: &rollback,
	}

	var abort externalRef0.TransactionAbortPhase
	var abortState string
	var abortStatus externalRef0.TransactionPhaseStatus
	if networkChange.GetTransaction().GetStatus().Phases.Abort != nil {
		abortState = networkChange.GetTransaction().GetStatus().Phases.Abort.GetState().String()
		abortStatus = externalRef0.TransactionPhaseStatus{
			End:   (*externalRef0.End)(networkChange.GetTransaction().GetStatus().Phases.Abort.GetEnd()),
			Start: (*externalRef0.Start)(networkChange.GetTransaction().GetStatus().Phases.Abort.GetStart()),
		}
		abort.State = (*externalRef0.AbortPhaseState)(&abortState)
		abort.Status = &abortStatus
	}

	var apply externalRef0.TransactionApplyPhase
	var applyState string
	var appFailDes string
	var appFailType string
	var appFailure externalRef0.Failure
	var appStatus externalRef0.TransactionPhaseStatus
	if (networkChange.GetTransaction().GetStatus().Phases.Apply != nil) && (networkChange.GetTransaction().GetStatus().Phases.Apply.Failure != nil) {
		appFailDes = networkChange.GetTransaction().GetStatus().Phases.Apply.Failure.GetDescription()
		appFailType = networkChange.GetTransaction().GetStatus().Phases.Apply.Failure.GetType().String()
		appFailure.Description = &appFailDes
		appFailure.Type = (*externalRef0.FailureType)(&appFailType)
	}

	if networkChange.GetTransaction().GetStatus().Phases.Apply != nil {
		applyState = networkChange.GetTransaction().GetStatus().Phases.Apply.GetState().String()
		appStatus.Start = (*externalRef0.Start)(networkChange.GetTransaction().GetStatus().Phases.Apply.GetStart())
		appStatus.End = (*externalRef0.End)(networkChange.GetTransaction().GetStatus().Phases.Apply.GetEnd())
		apply.Failure = &appFailure
		apply.State = (*externalRef0.ApplyPhaseState)(&applyState)
		apply.Status = &appStatus
	}

	var commit externalRef0.TransactionCommitPhase
	var commitState string
	var comStatus externalRef0.TransactionPhaseStatus
	if networkChange.GetTransaction().GetStatus().Phases.Commit != nil {
		commitState = networkChange.GetTransaction().GetStatus().Phases.Commit.GetState().String()
		comStatus.Start = (*externalRef0.Start)(networkChange.GetTransaction().GetStatus().Phases.Commit.GetStart())
		comStatus.End = (*externalRef0.End)(networkChange.GetTransaction().GetStatus().Phases.Commit.GetEnd())
		commit.State = (*externalRef0.CommitPhaseState)(&commitState)
		commit.Status = &comStatus
	}

	var initialize externalRef0.TransactionInitializePhase
	var initializeState string
	var iniFailDes string
	var iniFailType string
	var iniFailure externalRef0.Failure
	var iniStatus externalRef0.TransactionPhaseStatus
	if (networkChange.GetTransaction().GetStatus().Phases.Initialize != nil) && (networkChange.GetTransaction().GetStatus().Phases.Initialize.Failure != nil) {
		iniFailDes = networkChange.GetTransaction().GetStatus().Phases.Initialize.Failure.GetDescription()
		iniFailType = networkChange.GetTransaction().GetStatus().Phases.Initialize.Failure.GetType().String()
		iniFailure.Description = &iniFailDes
		iniFailure.Type = (*externalRef0.FailureType)(&iniFailType)
	}

	if networkChange.GetTransaction().GetStatus().Phases.Initialize != nil {
		initializeState = networkChange.GetTransaction().GetStatus().Phases.Initialize.GetState().String()
		iniStatus.Start = (*externalRef0.Start)(networkChange.GetTransaction().GetStatus().Phases.Initialize.GetStart())
		iniStatus.End = (*externalRef0.End)(networkChange.GetTransaction().GetStatus().Phases.Initialize.GetEnd())
		initialize.Failure = &iniFailure
		initialize.State = (*externalRef0.InitializePhaseState)(&initializeState)
		initialize.Status = &iniStatus
	}

	var validate externalRef0.TransactionValidatePhase
	var validateState string
	var valFailDes string
	var valFailType string
	var valFailure externalRef0.Failure
	var valStatus externalRef0.TransactionPhaseStatus
	if (networkChange.GetTransaction().GetStatus().Phases.Validate != nil) && (networkChange.GetTransaction().GetStatus().Phases.Validate.Failure != nil) {
		valFailDes = networkChange.GetTransaction().GetStatus().Phases.Validate.Failure.GetDescription()
		valFailType = networkChange.GetTransaction().GetStatus().Phases.Validate.Failure.GetType().String()
		valFailure.Description = &valFailDes
		valFailure.Type = (*externalRef0.FailureType)(&valFailType)
	}

	if networkChange.GetTransaction().GetStatus().Phases.Validate != nil {
		validateState = networkChange.GetTransaction().GetStatus().Phases.Validate.GetState().String()
		valStatus.Start = (*externalRef0.Start)(networkChange.GetTransaction().GetStatus().Phases.Validate.GetStart())
		valStatus.End = (*externalRef0.End)(networkChange.GetTransaction().GetStatus().Phases.Validate.GetEnd())
		validate.Failure = &valFailure
		validate.State = (*externalRef0.ValidatePhaseState)(&validateState)
		validate.Status = &valStatus
	}

	phases := externalRef0.TransactionPhases{
		Abort:      &abort,
		Apply:      &apply,
		Commit:     &commit,
		Initialize: &initialize,
		Validate:   &validate,
	}

	proposals := make([]externalRef0.ProposalID, 0)
	for _, pro := range networkChange.GetTransaction().GetStatus().Proposals {
		proposals = append(proposals, (externalRef0.ProposalID)(pro))
	}

	state := networkChange.GetTransaction().GetStatus().State.String()

	failure := externalRef0.Failure{}
	if networkChange.GetTransaction().GetStatus().Failure != nil {
		failureDescription := networkChange.GetTransaction().GetStatus().Failure.GetDescription()
		failureType := networkChange.GetTransaction().GetStatus().Failure.GetType().String()
		failure.Description = &failureDescription
		failure.Type = (*externalRef0.FailureType)(&failureType)
	}

	status := externalRef0.Status{
		Failure:   &failure,
		Phases:    &phases,
		Proposals: &proposals,
		State:     (*externalRef0.State)(&state),
	}

	isolation := networkChange.GetTransaction().Isolation.String()
	synchronicity := networkChange.GetTransaction().Synchronicity.String()
	strategy := externalRef0.Strategy{
		Isolation:     (*externalRef0.Isolation)(&isolation),
		Synchronicity: (*externalRef0.Synchronicity)(&synchronicity),
	}

	transaction := externalRef0.Transaction{
		Details:  &details,
		Id:       string(networkChange.GetTransaction().GetID()),
		Index:    int64(networkChange.GetTransaction().GetIndex()),
		Meta:     objMeta,
		Status:   &status,
		Strategy: &strategy,
		Username: &username,
	}
	return transaction
}

// TopLevelServer -
type TopLevelServer struct {
	GnmiClient    southbound.GnmiClient
	ConfigClient  admin.TransactionServiceClient
	GnmiTimeout   time.Duration
	Authorization bool
}

// PatchAetherRocAPI impl of gNMI access at /aether-roc-api
func (i *TopLevelServer) PatchAetherRocAPI(ctx echo.Context) error {

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

	log.Infof("PatchAetherRocAPI")
	return ctx.JSON(http.StatusOK, response)
}

// GetTargets -
func (i *TopLevelServer) GetTargets(ctx echo.Context) error {
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

// GetTransactions -
func (i *TopLevelServer) GetTransactions(ctx echo.Context) error {
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

// PostSdcoreSynchronize -
func (i *TopLevelServer) PostSdcoreSynchronize(httpContext echo.Context) error {

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

// GetSpec -
func (i *TopLevelServer) GetSpec(ctx echo.Context) error {
	response, err := GetSwagger()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	log.Infof("GetSpec")
	return acceptTypes(ctx, response)
}

// GetAether200Spec -
func (i *TopLevelServer) GetAether200Spec(ctx echo.Context) error {
	response, err := aether_2_0_0.GetSwagger()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return acceptTypes(ctx, response)
}

// GetAether210Spec -
func (i *TopLevelServer) GetAether210Spec(ctx echo.Context) error {
	response, err := aether_2_1_0.GetSwagger()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return acceptTypes(ctx, response)
}

// GetAether400Spec -
func (i *TopLevelServer) GetAether400Spec(ctx echo.Context) error {
	response, err := aether_4_0_0.GetSwagger()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return acceptTypes(ctx, response)
}

// GetAetherAppGtwySpec -
func (i *TopLevelServer) GetAetherAppGtwySpec(ctx echo.Context) error {
	response, err := app_gtwy.GetSwagger()
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
		_ = specTemplate.Execute(&b, HTMLData{
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
