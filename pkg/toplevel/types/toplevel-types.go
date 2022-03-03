// Code generated by GENERATOR. DO NOT EDIT.
// Package types provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.7.0 DO NOT EDIT.
package types

import (
	"time"

	externalRef0 "github.com/onosproject/aether-roc-api/pkg/aether_2_0_0/types"
	externalRef1 "github.com/onosproject/aether-roc-api/pkg/aether_4_0_0/types"
)

// Defines values for TransactionStatusFailureType.
const (
	TransactionStatusFailureTypeALREADYEXISTS TransactionStatusFailureType = "ALREADY_EXISTS"

	TransactionStatusFailureTypeCANCELED TransactionStatusFailureType = "CANCELED"

	TransactionStatusFailureTypeCONFLICT TransactionStatusFailureType = "CONFLICT"

	TransactionStatusFailureTypeFORBIDDEN TransactionStatusFailureType = "FORBIDDEN"

	TransactionStatusFailureTypeINTERNAL TransactionStatusFailureType = "INTERNAL"

	TransactionStatusFailureTypeINVALID TransactionStatusFailureType = "INVALID"

	TransactionStatusFailureTypeNOTFOUND TransactionStatusFailureType = "NOT_FOUND"

	TransactionStatusFailureTypeNOTSUPPORTED TransactionStatusFailureType = "NOT_SUPPORTED"

	TransactionStatusFailureTypeTIMEOUT TransactionStatusFailureType = "TIMEOUT"

	TransactionStatusFailureTypeUNAUTHORIZED TransactionStatusFailureType = "UNAUTHORIZED"

	TransactionStatusFailureTypeUNAVAILABLE TransactionStatusFailureType = "UNAVAILABLE"

	TransactionStatusFailureTypeUNKNOWN TransactionStatusFailureType = "UNKNOWN"
)

// Defines values for TransactionStatusPhasesAbortState.
const (
	TransactionStatusPhasesAbortStateABORTED TransactionStatusPhasesAbortState = "ABORTED"

	TransactionStatusPhasesAbortStateABORTING TransactionStatusPhasesAbortState = "ABORTING"
)

// Defines values for TransactionStatusPhasesApplyFailureType.
const (
	TransactionStatusPhasesApplyFailureTypeALREADYEXISTS TransactionStatusPhasesApplyFailureType = "ALREADY_EXISTS"

	TransactionStatusPhasesApplyFailureTypeCANCELED TransactionStatusPhasesApplyFailureType = "CANCELED"

	TransactionStatusPhasesApplyFailureTypeCONFLICT TransactionStatusPhasesApplyFailureType = "CONFLICT"

	TransactionStatusPhasesApplyFailureTypeFORBIDDEN TransactionStatusPhasesApplyFailureType = "FORBIDDEN"

	TransactionStatusPhasesApplyFailureTypeINTERNAL TransactionStatusPhasesApplyFailureType = "INTERNAL"

	TransactionStatusPhasesApplyFailureTypeINVALID TransactionStatusPhasesApplyFailureType = "INVALID"

	TransactionStatusPhasesApplyFailureTypeNOTFOUND TransactionStatusPhasesApplyFailureType = "NOT_FOUND"

	TransactionStatusPhasesApplyFailureTypeNOTSUPPORTED TransactionStatusPhasesApplyFailureType = "NOT_SUPPORTED"

	TransactionStatusPhasesApplyFailureTypeTIMEOUT TransactionStatusPhasesApplyFailureType = "TIMEOUT"

	TransactionStatusPhasesApplyFailureTypeUNAUTHORIZED TransactionStatusPhasesApplyFailureType = "UNAUTHORIZED"

	TransactionStatusPhasesApplyFailureTypeUNAVAILABLE TransactionStatusPhasesApplyFailureType = "UNAVAILABLE"

	TransactionStatusPhasesApplyFailureTypeUNKNOWN TransactionStatusPhasesApplyFailureType = "UNKNOWN"
)

// Defines values for TransactionStatusPhasesApplyState.
const (
	TransactionStatusPhasesApplyStateAPPLIED TransactionStatusPhasesApplyState = "APPLIED"

	TransactionStatusPhasesApplyStateAPPLYING TransactionStatusPhasesApplyState = "APPLYING"

	TransactionStatusPhasesApplyStateFAILED TransactionStatusPhasesApplyState = "FAILED"
)

// Defines values for TransactionStatusPhasesCommitState.
const (
	TransactionStatusPhasesCommitStateCOMMITTED TransactionStatusPhasesCommitState = "COMMITTED"

	TransactionStatusPhasesCommitStateCOMMITTING TransactionStatusPhasesCommitState = "COMMITTING"
)

// Defines values for TransactionStatusPhasesInitializeFailureType.
const (
	TransactionStatusPhasesInitializeFailureTypeALREADYEXISTS TransactionStatusPhasesInitializeFailureType = "ALREADY_EXISTS"

	TransactionStatusPhasesInitializeFailureTypeCANCELED TransactionStatusPhasesInitializeFailureType = "CANCELED"

	TransactionStatusPhasesInitializeFailureTypeCONFLICT TransactionStatusPhasesInitializeFailureType = "CONFLICT"

	TransactionStatusPhasesInitializeFailureTypeFORBIDDEN TransactionStatusPhasesInitializeFailureType = "FORBIDDEN"

	TransactionStatusPhasesInitializeFailureTypeINTERNAL TransactionStatusPhasesInitializeFailureType = "INTERNAL"

	TransactionStatusPhasesInitializeFailureTypeINVALID TransactionStatusPhasesInitializeFailureType = "INVALID"

	TransactionStatusPhasesInitializeFailureTypeNOTFOUND TransactionStatusPhasesInitializeFailureType = "NOT_FOUND"

	TransactionStatusPhasesInitializeFailureTypeNOTSUPPORTED TransactionStatusPhasesInitializeFailureType = "NOT_SUPPORTED"

	TransactionStatusPhasesInitializeFailureTypeTIMEOUT TransactionStatusPhasesInitializeFailureType = "TIMEOUT"

	TransactionStatusPhasesInitializeFailureTypeUNAUTHORIZED TransactionStatusPhasesInitializeFailureType = "UNAUTHORIZED"

	TransactionStatusPhasesInitializeFailureTypeUNAVAILABLE TransactionStatusPhasesInitializeFailureType = "UNAVAILABLE"

	TransactionStatusPhasesInitializeFailureTypeUNKNOWN TransactionStatusPhasesInitializeFailureType = "UNKNOWN"
)

// Defines values for TransactionStatusPhasesInitializeState.
const (
	TransactionStatusPhasesInitializeStateFAILED TransactionStatusPhasesInitializeState = "FAILED"

	TransactionStatusPhasesInitializeStateINITIALIZED TransactionStatusPhasesInitializeState = "INITIALIZED"

	TransactionStatusPhasesInitializeStateINITIALIZING TransactionStatusPhasesInitializeState = "INITIALIZING"
)

// Defines values for TransactionStatusPhasesValidateFailureType.
const (
	TransactionStatusPhasesValidateFailureTypeALREADYEXISTS TransactionStatusPhasesValidateFailureType = "ALREADY_EXISTS"

	TransactionStatusPhasesValidateFailureTypeCANCELED TransactionStatusPhasesValidateFailureType = "CANCELED"

	TransactionStatusPhasesValidateFailureTypeCONFLICT TransactionStatusPhasesValidateFailureType = "CONFLICT"

	TransactionStatusPhasesValidateFailureTypeFORBIDDEN TransactionStatusPhasesValidateFailureType = "FORBIDDEN"

	TransactionStatusPhasesValidateFailureTypeINTERNAL TransactionStatusPhasesValidateFailureType = "INTERNAL"

	TransactionStatusPhasesValidateFailureTypeINVALID TransactionStatusPhasesValidateFailureType = "INVALID"

	TransactionStatusPhasesValidateFailureTypeNOTFOUND TransactionStatusPhasesValidateFailureType = "NOT_FOUND"

	TransactionStatusPhasesValidateFailureTypeNOTSUPPORTED TransactionStatusPhasesValidateFailureType = "NOT_SUPPORTED"

	TransactionStatusPhasesValidateFailureTypeTIMEOUT TransactionStatusPhasesValidateFailureType = "TIMEOUT"

	TransactionStatusPhasesValidateFailureTypeUNAUTHORIZED TransactionStatusPhasesValidateFailureType = "UNAUTHORIZED"

	TransactionStatusPhasesValidateFailureTypeUNAVAILABLE TransactionStatusPhasesValidateFailureType = "UNAVAILABLE"

	TransactionStatusPhasesValidateFailureTypeUNKNOWN TransactionStatusPhasesValidateFailureType = "UNKNOWN"
)

// Defines values for TransactionStatusPhasesValidateState.
const (
	TransactionStatusPhasesValidateStateFAILED TransactionStatusPhasesValidateState = "FAILED"

	TransactionStatusPhasesValidateStateINITIALIZED TransactionStatusPhasesValidateState = "INITIALIZED"

	TransactionStatusPhasesValidateStateINITIALIZING TransactionStatusPhasesValidateState = "INITIALIZING"
)

// Defines values for TransactionStatusState.
const (
	TransactionStatusStateAPPLIED TransactionStatusState = "APPLIED"

	TransactionStatusStateCOMMITTED TransactionStatusState = "COMMITTED"

	TransactionStatusStateFAILED TransactionStatusState = "FAILED"

	TransactionStatusStatePENDING TransactionStatusState = "PENDING"

	TransactionStatusStateVALIDATED TransactionStatusState = "VALIDATED"
)

// Defines values for TransactionStrategyIsolation.
const (
	TransactionStrategyIsolationDEFAULT TransactionStrategyIsolation = "DEFAULT"

	TransactionStrategyIsolationSERIALIZABLE TransactionStrategyIsolation = "SERIALIZABLE"
)

// Defines values for TransactionStrategySynchronicity.
const (
	TransactionStrategySynchronicityASYNCHRONOUS TransactionStrategySynchronicity = "ASYNCHRONOUS"

	TransactionStrategySynchronicitySYNCHRONOUS TransactionStrategySynchronicity = "SYNCHRONOUS"
)

// Elements defines model for Elements.
type Elements struct {

	// The top level container
	Application400 *externalRef1.Application `json:"application-4.0.0,omitempty"`

	// The top level container
	ConnectivityService400 *externalRef1.ConnectivityService `json:"connectivity-service-4.0.0,omitempty"`

	// The connectivity-services top level container
	ConnectivityServices200 *externalRef0.ConnectivityServices `json:"connectivity-services-2.0.0,omitempty"`

	// The top level container
	DeviceGroup400 *externalRef1.DeviceGroup `json:"device-group-4.0.0,omitempty"`

	// The top level container
	Enterprise400 *externalRef1.Enterprise `json:"enterprise-4.0.0,omitempty"`

	// The top level enterprises container
	Enterprises200 *externalRef0.Enterprises `json:"enterprises-2.0.0,omitempty"`

	// The top level container
	IpDomain400 *externalRef1.IpDomain `json:"ip-domain-4.0.0,omitempty"`

	// The top level container
	Site400 *externalRef1.Site `json:"site-4.0.0,omitempty"`

	// The top level container
	Template400 *externalRef1.Template `json:"template-4.0.0,omitempty"`

	// The top level container
	TrafficClass400 *externalRef1.TrafficClass `json:"traffic-class-4.0.0,omitempty"`

	// The top level container
	Upf400 *externalRef1.Upf `json:"upf-4.0.0,omitempty"`

	// The top level container
	Vcs400 *externalRef1.Vcs `json:"vcs-4.0.0,omitempty"`
}

// PatchBody defines model for PatchBody.
type PatchBody struct {
	Deletes *Elements `json:"Deletes,omitempty"`

	// Model type and version of 'target' on first creation [link](https://docs.onosproject.org/onos-config/docs/gnmi_extensions/#use-of-extension-101-device-version-in-setrequest)
	Extensions *struct {
		ChangeName100   *string `json:"change-name-100,omitempty"`
		ModelType102    *string `json:"model-type-102,omitempty"`
		ModelVersion101 *string `json:"model-version-101,omitempty"`

		// Used in the responses, carries inforamtion about the transaction.
		TransactionInfo110 *struct {
			Id    *string `json:"id,omitempty"`
			Index *int    `json:"index,omitempty"`
		} `json:"transaction-info-110,omitempty"`

		// Identifies whether a request needs to be handles Asynchronously (val: 0) or Synchronously (val: 1)
		TransactionStrategy111 *int `json:"transaction-strategy-111,omitempty"`
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

// Transaction refers to a multi-target transactional change. Taken from https://github.com/onosproject/onos-api/tree/master/proto/onos/config/v2
type Transaction struct {

	// the transaction details
	Details *struct {
		Change *struct {

			// a set of changes to apply to targets
			Values *interface{} `json:"values,omitempty"`
		} `json:"change,omitempty"`

		// rollback
		Rollback *struct {

			// the index of the transaction to roll back
			RollbackIndex *int64 `json:"rollback_index,omitempty"`
		} `json:"rollback,omitempty"`
	} `json:"details,omitempty"`

	// the unique identifier of the transaction
	Id string `json:"id"`

	// a monotonically increasing, globally unique index of the change
	Index int64 `json:"index"`
	Meta  struct {

		// the time at which the transaction was created
		Created *time.Time `json:"created,omitempty"`

		// a flag indicating whether this transaction is being deleted by a snapshot
		Deleted *time.Time `json:"deleted,omitempty"`
		Key     *string    `json:"key,omitempty"`

		// the change revision number
		Revision *struct {

			// Revision is a revision number
			Revision *int64 `json:"revision,omitempty"`
		} `json:"revision,omitempty"`

		// the time at which the transaction was last updated
		Updated *time.Time `json:"updated,omitempty"`
		Version *uint64    `json:"version,omitempty"`
	} `json:"meta"`

	// the current lifecycle status of the transaction
	Status *struct {

		// the transaction failure
		Failure *struct {
			Description *string `json:"description,omitempty"`

			// transaction failure type
			Type *TransactionStatusFailureType `json:"type,omitempty"`
		} `json:"failure,omitempty"`

		// the transaction phases
		Phases *struct {
			Abort *struct {
				State  *TransactionStatusPhasesAbortState `json:"state,omitempty"`
				Status *struct {
					End   *time.Time `json:"end,omitempty"`
					Start *time.Time `json:"start,omitempty"`
				} `json:"status,omitempty"`
			} `json:"abort,omitempty"`

			// the transaction apply phase status
			Apply *struct {

				// the transaction failure
				Failure *struct {
					Description *string `json:"description,omitempty"`

					// transaction failure type
					Type *TransactionStatusPhasesApplyFailureType `json:"type,omitempty"`
				} `json:"failure,omitempty"`
				State  *TransactionStatusPhasesApplyState `json:"state,omitempty"`
				Status *struct {
					End   *time.Time `json:"end,omitempty"`
					Start *time.Time `json:"start,omitempty"`
				} `json:"status,omitempty"`
			} `json:"apply,omitempty"`

			// the transaction commit phase status
			Commit *struct {
				State  *TransactionStatusPhasesCommitState `json:"state,omitempty"`
				Status *struct {
					End   *time.Time `json:"end,omitempty"`
					Start *time.Time `json:"start,omitempty"`
				} `json:"status,omitempty"`
			} `json:"commit,omitempty"`

			// the transaction initialization phase status
			Initialize *struct {

				// the transaction failure
				Failure *struct {
					Description *string `json:"description,omitempty"`

					// transaction failure type
					Type *TransactionStatusPhasesInitializeFailureType `json:"type,omitempty"`
				} `json:"failure,omitempty"`
				State  *TransactionStatusPhasesInitializeState `json:"state,omitempty"`
				Status *struct {
					End   *time.Time `json:"end,omitempty"`
					Start *time.Time `json:"start,omitempty"`
				} `json:"status,omitempty"`
			} `json:"initialize,omitempty"`

			// the transaction validation phase status
			Validate *struct {

				// the transaction failure
				Failure *struct {
					Description *string `json:"description,omitempty"`

					// transaction failure type
					Type *TransactionStatusPhasesValidateFailureType `json:"type,omitempty"`
				} `json:"failure,omitempty"`
				State  *TransactionStatusPhasesValidateState `json:"state,omitempty"`
				Status *struct {
					End   *time.Time `json:"end,omitempty"`
					Start *time.Time `json:"start,omitempty"`
				} `json:"status,omitempty"`
			} `json:"validate,omitempty"`
		} `json:"phases,omitempty"`

		// the set of proposals managed by the transaction
		Proposals *[]interface{} `json:"proposals,omitempty"`

		// the overall transaction state
		State *TransactionStatusState `json:"state,omitempty"`
	} `json:"status,omitempty"`
	Strategy *struct {

		// indicates the transaction isolation level
		Isolation *TransactionStrategyIsolation `json:"isolation,omitempty"`

		// indicates the transaction synchronicity level
		Synchronicity *TransactionStrategySynchronicity `json:"synchronicity,omitempty"`
	} `json:"strategy,omitempty"`

	// the name of the user that made the transaction
	Username *string `json:"username,omitempty"`
}

// transaction failure type
type TransactionStatusFailureType string

// TransactionStatusPhasesAbortState defines model for Transaction.Status.Phases.Abort.State.
type TransactionStatusPhasesAbortState string

// transaction failure type
type TransactionStatusPhasesApplyFailureType string

// TransactionStatusPhasesApplyState defines model for Transaction.Status.Phases.Apply.State.
type TransactionStatusPhasesApplyState string

// TransactionStatusPhasesCommitState defines model for Transaction.Status.Phases.Commit.State.
type TransactionStatusPhasesCommitState string

// transaction failure type
type TransactionStatusPhasesInitializeFailureType string

// TransactionStatusPhasesInitializeState defines model for Transaction.Status.Phases.Initialize.State.
type TransactionStatusPhasesInitializeState string

// transaction failure type
type TransactionStatusPhasesValidateFailureType string

// TransactionStatusPhasesValidateState defines model for Transaction.Status.Phases.Validate.State.
type TransactionStatusPhasesValidateState string

// the overall transaction state
type TransactionStatusState string

// indicates the transaction isolation level
type TransactionStrategyIsolation string

// indicates the transaction synchronicity level
type TransactionStrategySynchronicity string

// TransactionList defines model for TransactionList.
type TransactionList []Transaction

// PatchTopLevelJSONBody defines parameters for PatchTopLevel.
type PatchTopLevelJSONBody PatchBody

// PatchTopLevelJSONRequestBody defines body for PatchTopLevel for application/json ContentType.
type PatchTopLevelJSONRequestBody PatchTopLevelJSONBody
