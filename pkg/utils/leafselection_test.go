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
		SelectionPath: "/a/20/c/foo/e",
		ChangeContext: nil,
	}).Times(1).Return(&admin.LeafSelectionQueryResponse{
		Selection: []string{"value1", "value2"},
	}, nil)

	selection, err := LeafSelection(context.Background(), adminClient,
		"test-model", "1.0.x", "/a/{b}/c/{d}/e", "test-target", "20", "foo")
	assert.NoError(t, err)
	assert.NotNil(t, selection)
	assert.Equal(t, "value1", selection[0])

	_, err = LeafSelection(context.Background(), adminClient,
		"test-model", "1.0.x", "/a/{b}/c/{d}/e", "test-target", "20")
	assert.Error(t, err)
	assert.Equal(t, `unexpected number of args. Expect 2. Got 1. queryPath=/a/{b}/c/{d}/e. Args [20]`, err.Error())
}
