/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package utils

import (
	"context"
	"github.com/golang/mock/gomock"
	testadmin "github.com/onosproject/aether-roc-api/pkg/utils/test"
	"github.com/onosproject/onos-api/go/onos/config/admin"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_LeafSelection(t *testing.T) {
	mctl := gomock.NewController(t)
	adminClient := testadmin.NewMockConfigAdminServiceClient(mctl)
	adminClient.EXPECT().LeafSelectionQuery(gomock.Any(), &admin.LeafSelectionQueryRequest{
		Target:        "test-target",
		Type:          "test-model",
		Version:       "1.0.x",
		SelectionPath: "/a[a-id=20]/c[c-id=foo]/e",
		ChangeContext: nil,
	}).Times(1).Return(&admin.LeafSelectionQueryResponse{
		Selection: []string{"value1", "value2"},
	}, nil)
	adminClient.EXPECT().LeafSelectionQuery(gomock.Any(), &admin.LeafSelectionQueryRequest{
		Target:        "test-target",
		Type:          "test-model",
		Version:       "1.0.x",
		SelectionPath: "/a/p/c[c-id=foo]/e",
		ChangeContext: nil,
	}).Times(1).Return(&admin.LeafSelectionQueryResponse{
		Selection: []string{"value3", "value4", "value5"},
	}, nil)

	selection, err := LeafSelection(context.Background(), adminClient,
		"test-model", "1.0.x", "/a/{b}/c/{d}/e", "test-target",
		[]PathID{{"a-id", "20"}, {"c-id", "foo"}})
	assert.NoError(t, err)
	assert.NotNil(t, selection)
	assert.Equal(t, "value1", selection[0])

	selection2, err := LeafSelection(context.Background(), adminClient,
		"test-model", "1.0.x", "/a/p/c/{d}/e", "test-target",
		PathID{"c-id", "foo"})
	assert.NoError(t, err)
	assert.NotNil(t, selection2)
	assert.Equal(t, "value3", selection2[0])

	_, err = LeafSelection(context.Background(), adminClient,
		"test-model", "1.0.x", "/a/{b}/c/{d}/e", "test-target", PathID{"a-id", "20"})
	assert.Error(t, err)
	assert.Equal(t, `unexpected number of args. Expect 2. Got 1. queryPath=/a/{b}/c/{d}/e. Args [{a-id 20}]`, err.Error())
}
