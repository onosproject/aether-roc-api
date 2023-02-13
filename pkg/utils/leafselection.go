/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package utils

import (
	"context"
	"fmt"
	"github.com/onosproject/onos-api/go/onos/config/admin"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/openconfig/gnmi/proto/gnmi"
	"strings"
)

// PathID - a tuple for the path name and value
type PathID struct {
	Name  string
	Value string
}

// LeafSelection - used by roc-api
func LeafSelection(ctx context.Context, configAdminServiceClient admin.ConfigAdminServiceClient,
	gnmiSet *gnmi.SetRequest,
	modelType string, modelVersion string,
	queryPath string, enterpriseID string, args ...PathID) ([]string, error) {

	if strings.Count(queryPath, "{") != len(args) {
		return nil, errors.NewInvalid("unexpected number of args. Expect %d. Got %d. queryPath=%s. Args %v",
			strings.Count(queryPath, "{"), len(args), queryPath, args,
		)
	}
	queryPathParts := strings.Split(queryPath, "/")
	nextArgIdx := 0
	newQueryParts := make([]string, 0)
	for _, qpp := range queryPathParts {
		if strings.HasPrefix(qpp, "{") {
			pathID := args[nextArgIdx]
			prevPart := newQueryParts[len(newQueryParts)-1]
			if pathID.Value != "new" {
				newQueryParts[len(newQueryParts)-1] =
					fmt.Sprintf("%s[%s=%s]", prevPart, pathID.Name, pathID.Value)
			} else {
				newQueryParts[len(newQueryParts)-1] =
					prevPart
			}
			nextArgIdx++
			continue
		}
		// else
		newQueryParts = append(newQueryParts, qpp)
	}

	resp, err := configAdminServiceClient.LeafSelectionQuery(ctx, &admin.LeafSelectionQueryRequest{
		Target:        enterpriseID,
		Type:          modelType,
		Version:       modelVersion,
		SelectionPath: strings.Join(newQueryParts, "/"),
		ChangeContext: gnmiSet,
	})
	if err != nil {
		log.Warnf("LeafSelectionQuery error: %v", err)
		return nil, err
	}

	return resp.Selection, nil
}
