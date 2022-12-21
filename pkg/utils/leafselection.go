/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package utils

import (
	"context"
	"github.com/onosproject/onos-api/go/onos/config/admin"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"strings"
)

// LeafSelection - used by roc-api
func LeafSelection(ctx context.Context, configAdminServiceClient admin.ConfigAdminServiceClient,
	modelType string, modelVersion string,
	queryPath string, enterpriseID string, args ...string) ([]string, error) {

	if strings.Count(queryPath, "{") != len(args) {
		return nil, errors.NewInvalid("unexpected number of args. Expect %d. Got %d. queryPath=%s. Args %v",
			strings.Count(queryPath, "{"), len(args), queryPath, args,
		)
	}
	queryPathParts := strings.Split(queryPath, "/")
	nextArgIdx := 0
	for idx, qpp := range queryPathParts {
		if strings.HasPrefix(qpp, "{") {
			queryPathParts[idx] = args[nextArgIdx]
			nextArgIdx++
		}
	}

	resp, err := configAdminServiceClient.LeafSelectionQuery(ctx, &admin.LeafSelectionQueryRequest{
		Target:        enterpriseID,
		Type:          modelType,
		Version:       modelVersion,
		SelectionPath: strings.Join(queryPathParts, "/"),
		ChangeContext: nil,
	})
	if err != nil {
		log.Warnf("LeafSelectionQuery error: %v", err)
		return nil, err
	}

	return resp.Selection, nil
}
