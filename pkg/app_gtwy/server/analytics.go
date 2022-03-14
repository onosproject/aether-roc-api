// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
//

package server

import (
	"context"
	"fmt"
	promApi "github.com/prometheus/client_golang/api"
	promApiV1 "github.com/prometheus/client_golang/api/prometheus/v1"
	promModel "github.com/prometheus/common/model"
	"time"
)

// AnalyticsClient -
type AnalyticsClient interface {
	Init() error
	Query(query string) (promModel.Value, error)
}

// AnalyticsConnection -
type AnalyticsConnection struct {
	Address string
	client  promApi.Client
	v1api   promApiV1.API
}

// Init -
func (a *AnalyticsConnection) Init() error {
	log.Infof("Initializing new AnalyticsConnection")

	var err error
	a.client, err = promApi.NewClient(promApi.Config{
		Address: a.Address,
	})
	if err != nil {
		return fmt.Errorf("error creating client: %v", err)
	}

	a.v1api = promApiV1.NewAPI(a.client)

	return nil
}

// Query executes a query
func (a *AnalyticsConnection) Query(query string) (promModel.Value, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, warnings, err := a.v1api.Query(ctx, query, time.Now())
	if err != nil {
		return nil, fmt.Errorf("error querying Prometheus: %v", err)
	}
	if len(warnings) > 0 {
		fmt.Printf("Prometheus Warnings: %v\n", warnings)
	}

	// result is a Value, which is an interface to ValueType and a String() method
	// Can cast to:
	//    Matrix, Vector, *Scalar, *String
	return result, nil
}
