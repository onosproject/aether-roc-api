// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package server

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/onosproject/aether-roc-api/pkg/southbound"
	externalRef0 "github.com/onosproject/aether-roc-api/pkg/toplevel/types"
	"github.com/onosproject/onos-api/go/onos/config/admin"
	v2 "github.com/onosproject/onos-api/go/onos/config/v2"
	"github.com/openconfig/gnmi/proto/gnmi"
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

	details := &v2.Transaction_Change{
		Change: &v2.ChangeTransaction{
			Values: map[v2.TargetID]*v2.PathValues{
				"target1": {
					Values: map[string]*v2.PathValue{
						"/a/b/c": {
							Path: "/a/b/c",
							Value: v2.TypedValue{
								Bytes:    []byte("some value"),
								Type:     v2.ValueType_STRING,
								TypeOpts: nil,
							},
							Deleted: false,
						},
					},
				},
				"target2": {
					Values: map[string]*v2.PathValue{
						"/d/e/f": {
							Path:    "/d/e/f",
							Value:   *v2.NewTypedValueInt(1234, 16),
							Deleted: false,
						},
						"/d/e/g": {
							Path:    "/d/e/g",
							Value:   *v2.NewTypedValueBool(true),
							Deleted: false,
						},
						"/d/e/h": {
							Path:    "/d/e/h",
							Value:   *v2.NewTypedValueDecimal(12345, 2),
							Deleted: true,
						},
					},
				},
			},
		},
	}

	v1 := v2.Transaction{
		ObjectMeta:          meta,
		ID:                  "acbfu-23fgtj",
		Index:               2,
		Username:            "",
		TransactionStrategy: ts,
		Details:             details,
		Status:              status,
	}

	t1 := admin.ListTransactionsResponse{
		Transaction: &v1,
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
	assert.Equal(t, (externalRef0.End)(appEt), *ct.Status.Phases.Apply.Status.End)
	for _, cTransaction := range *ct.Details.Change {
		tName := *cTransaction.TargetName
		switch tName {
		case "target1", "target2":
			if tName == "target1" {
				for _, pV := range *cTransaction.PathValues {
					p := "/a/b/c"
					assert.Equal(t, (*externalRef0.Path)(&p), pV.PathValue.Path)
					assert.Equal(t, (externalRef0.Deleted)(false), *pV.PathValue.Deleted)
					assert.Equal(t, externalRef0.ValueTypeSTRING, *pV.PathValue.Value.Type)
					assert.Equal(t, "some value", *pV.PathValue.Value.Value)
				}
			} else if tName == "target2" {
				for _, pV := range *cTransaction.PathValues {
					pth := *pV.PathValue.Path
					switch pth {
					case "/d/e/f":
						assert.Equal(t, (externalRef0.Deleted)(false), *pV.PathValue.Deleted)
						assert.Equal(t, externalRef0.ValueTypeINT, *pV.PathValue.Value.Type)
						assert.Equal(t, "1234", *pV.PathValue.Value.Value)
					case "/d/e/g":
						assert.Equal(t, (externalRef0.Deleted)(false), *pV.PathValue.Deleted)
						assert.Equal(t, externalRef0.ValueTypeBOOL, *pV.PathValue.Value.Type)
						assert.Equal(t, "true", *pV.PathValue.Value.Value)
					case "/d/e/h":
						assert.Equal(t, (externalRef0.Deleted)(true), *pV.PathValue.Deleted)
						assert.Equal(t, externalRef0.ValueTypeDECIMAL, *pV.PathValue.Value.Type)
						assert.Equal(t, "123.45", *pV.PathValue.Value.Value)
					default:
						t.Errorf("unexpected Path: %s", pth)
					}
				}
			}
		default:
			t.Errorf("unexpected TargetID: %s", tName)
		}
	}

	meta2 := v2.ObjectMeta{
		Version: 12,
	}

	details2 := &v2.Transaction_Rollback{
		Rollback: &v2.RollbackTransaction{
			RollbackIndex: v2.Index(22)},
	}

	v3 := v2.Transaction{
		ObjectMeta: meta2,
		ID:         "acbfu-323fgtj",
		Index:      4,
		Details:    details2,
	}

	t2 := admin.ListTransactionsResponse{
		Transaction: &v3,
	}

	ct1 := convertTrasaction(&t2)
	assert.NotNil(t, ct1)
	assert.Equal(t, "acbfu-323fgtj", ct1.Id)
	assert.Equal(t, int64(12), *ct1.Meta.Version)
	assert.Equal(t, externalRef0.Index(22), *ct1.Details.Rollback.RollbackIndex)

	t3 := admin.ListTransactionsResponse{}

	ct2 := convertTrasaction(&t3)
	assert.Equal(t, externalRef0.Transaction{}, ct2)
	assert.Len(t, ct2.Id, 0)
	assert.Nil(t, ct2.Status)
}

func Test_GetTargets(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := southbound.NewMockGnmiClient(ctrl)
	mockClient.EXPECT().Get(gomock.Any(), gomock.Any()).DoAndReturn(
		func(ctx context.Context, request *gnmi.GetRequest) (*gnmi.GetResponse, error) {
			gr := gnmi.GetResponse{
				Notification: []*gnmi.Notification{
					{
						Update: []*gnmi.Update{
							{
								Val: &gnmi.TypedValue{
									Value: &gnmi.TypedValue_LeaflistVal{
										LeaflistVal: &gnmi.ScalarArray{
											Element: []*gnmi.TypedValue{
												{
													Value: &gnmi.TypedValue_StringVal{
														StringVal: "target-1",
													},
												},
												{
													Value: &gnmi.TypedValue_StringVal{
														StringVal: "target-2",
													},
												},
												{
													Value: &gnmi.TypedValue_StringVal{
														StringVal: "connectivity-service-v2",
													},
												},
											},
										},
									}},
							},
						},
					},
				},
			}
			return &gr, nil
		},
	).AnyTimes()

	topLevel := &TopLevelServer{
		GnmiClient:    mockClient,
		ConfigClient:  nil,
		GnmiTimeout:   5,
		Authorization: false,
	}

	targetNames, err := topLevel.gnmiGetTargets(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, 2, len(*targetNames))
	assert.Equal(t, "target-1", *(*targetNames)[0].Name)
	assert.Equal(t, "target-2", *(*targetNames)[1].Name)
	// connectivity-service-v2 was removed
}
