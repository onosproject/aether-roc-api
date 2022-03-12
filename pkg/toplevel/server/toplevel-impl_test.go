package server

import (
	externalRef0 "github.com/onosproject/aether-roc-api/pkg/toplevel/types"
	"github.com/onosproject/onos-api/go/onos/config/admin"
	v2 "github.com/onosproject/onos-api/go/onos/config/v2"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_convertTrasaction(t *testing.T) {

	meta := v2.ObjectMeta{
		Key:      "",
		Version:  15,
		Revision: 5,
		Created:  time.Time{},
		Updated:  time.Time{},
		Deleted:  nil,
	}

	ts := v2.TransactionStrategy{
		Synchronicity: 1,
		Isolation:     0,
	}

	iniSt := time.Date(2022, time.January, 12, 10, 19, 0, 0, time.UTC)
	iniEt := time.Date(2022, time.February, 12, 11, 19, 0, 0, time.UTC)
	iniPhaseStatus := v2.TransactionPhaseStatus{
		Start: &iniSt,
		End:   &iniEt,
	}

	iniFailure := v2.Failure{
		Type:        4,
		Description: "abc",
	}

	initialize := v2.TransactionInitializePhase{
		TransactionPhaseStatus: iniPhaseStatus,
		State:                  1,
		Failure:                &iniFailure,
	}

	appSt := time.Date(2022, time.February, 12, 12, 12, 0, 0, time.UTC)
	appEt := time.Date(2022, time.February, 12, 15, 15, 0, 0, time.UTC)
	appPhaseStatus := v2.TransactionPhaseStatus{
		Start: &appSt,
		End:   &appEt,
	}

	appFailure := v2.Failure{
		Type:        9,
		Description: "xyz",
	}

	apply := v2.TransactionApplyPhase{
		TransactionPhaseStatus: appPhaseStatus,
		State:                  0,
		Failure:                &appFailure,
	}
	phase := v2.TransactionPhases{
		Initialize: &initialize,
		Validate:   nil,
		Commit:     nil,
		Apply:      &apply,
		Abort:      nil,
	}

	status := v2.TransactionStatus{
		Phases:    phase,
		Proposals: nil,
		State:     3,
		Failure:   nil,
	}

	//v2.ChangeTransaction{map[v2.TargetID]*v2.PathValues{}}
	//rt := v2.RollbackTransaction{RollbackIndex: 2}

	v1 := v2.Transaction{
		ObjectMeta:          meta,
		ID:                  "acbfu-23fgtj",
		Index:               2,
		Username:            "",
		TransactionStrategy: ts,
		Details:             nil, //TODO: need to discuss and then implement it.
		Status:              status,
	}

	t1 := admin.ListTransactionsResponse{
		Transaction:          &v1,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}

	ct := convertTrasaction(&t1)
	assert.NotNil(t, ct)
	assert.Equal(t, "acbfu-23fgtj", ct.Id)
	assert.Equal(t, int64(15), *ct.Meta.Version)
	assert.Equal(t, (externalRef0.State)("APPLIED"), *ct.Status.State)
	assert.Equal(t, (externalRef0.Synchronicity)("SYNCHRONOUS"), *ct.Strategy.Synchronicity)
	assert.Equal(t, (externalRef0.Isolation)("DEFAULT"), *ct.Strategy.Isolation)
	assert.Equal(t, (externalRef0.InitializePhaseState)("INITIALIZED"), *ct.Status.Phases.Initialize.State)
	assert.Equal(t, (externalRef0.ApplyPhaseState)("APPLYING"), *ct.Status.Phases.Apply.State)
	assert.Equal(t, "xyz", *ct.Status.Phases.Apply.Failure.Description)
	assert.Equal(t, "abc", *ct.Status.Phases.Initialize.Failure.Description)
	assert.Equal(t, (externalRef0.FailureType)("UNAUTHORIZED"), *ct.Status.Phases.Initialize.Failure.Type)
	assert.Equal(t, (externalRef0.FailureType)("NOT_SUPPORTED"), *ct.Status.Phases.Apply.Failure.Type)
	assert.Equal(t, (externalRef0.Start)(iniSt), *ct.Status.Phases.Initialize.Status.Start)
	assert.Equal(t, (externalRef0.Start)(appSt), *ct.Status.Phases.Apply.Status.Start)
}
