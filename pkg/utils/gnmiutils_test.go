// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package utils

import (
	"github.com/onosproject/config-models/modelplugin/aether-2.0.0/aether_2_0_0"
	"github.com/openconfig/gnmi/proto/gnmi"
	"gotest.tools/assert"
	"testing"
)

func Test_NewGnmiGetRequest(t *testing.T) {
	gnmiGet, err := NewGnmiGetRequest("/rbac/v1.0.0/{target}/rbac/group/role/{roleid}", "internal", "r1")
	assert.NilError(t, err, "unexpected error handling path")

	assert.Equal(t, 1, len(gnmiGet.Path), "expected only one path")
	path0 := gnmiGet.Path[0]
	assert.Equal(t, "internal", path0.Target)
	assert.Equal(t, 3, len(path0.Elem), "expected 3 path elems")
	assert.Equal(t, "rbac", path0.Elem[0].Name)
	assert.Equal(t, "group", path0.Elem[1].Name)
	assert.Equal(t, "role", path0.Elem[2].Name)
	assert.Equal(t, 1, len(path0.Elem[2].Key))
	key2, ok := path0.Elem[2].Key["roleid"]
	assert.Assert(t, ok)
	assert.Equal(t, "r1", key2)
}

func Test_GetResponseUpdate(t *testing.T) {
	path0Elems := make([]*gnmi.PathElem, 0)
	path0Elems = append(path0Elems, &gnmi.PathElem{Name: "pe1"})
	path0Elems = append(path0Elems, &gnmi.PathElem{Name: "pe2"})
	path0Elems = append(path0Elems, &gnmi.PathElem{Name: "pe3"})
	path0 := gnmi.Path{
		Elem:   path0Elems,
		Target: "internal",
	}
	u0 := gnmi.Update{
		Path: &path0,
		Val: &gnmi.TypedValue{
			Value: &gnmi.TypedValue_JsonVal{JsonVal: []byte("{testvalue: 't'}")},
		},
	}
	n0 := gnmi.Notification{
		Update: []*gnmi.Update{
			&u0,
		},
	}

	gr := gnmi.GetResponse{
		Notification: []*gnmi.Notification{
			&n0,
		},
	}

	typedVal, err := GetResponseUpdate(&gr, nil)
	assert.NilError(t, err, "unexpected error")
	jsonVal, ok := typedVal.Value.(*gnmi.TypedValue_JsonVal)
	assert.Assert(t, ok, "expecting to cast to JsonVal")
	assert.Equal(t, "{testvalue: 't'}", string(jsonVal.JsonVal))
}

func Test_buildElems(t *testing.T) {
	pathElems, err := BuildElems(
		"/rbac/v1.0.0/{target}/rbac/role/{roleid}", 4, "role-1")
	assert.NilError(t, err)
	assert.Equal(t, 2, len(pathElems))
	elem0 := pathElems[0]
	assert.Equal(t, "rbac", elem0.Name)
	assert.Equal(t, 0, len(elem0.Key))
	elem1 := pathElems[1]
	assert.Equal(t, "role", elem1.Name)
	assert.Equal(t, 1, len(elem1.Key))
	key1, ok := elem1.Key["roleid"]
	assert.Assert(t, ok)
	assert.Equal(t, "role-1", key1)
}

func Test_updateForElement(t *testing.T) {
	desc := "this is a description"
	gnmiUpdate, err := UpdateForElement(
		aether_2_0_0.AccessProfile_AccessProfile_AccessProfile{Description: &desc}.Description,
		"/test1/test2/{name}", "t1")
	assert.NilError(t, err, "unexpected error")
	assert.Assert(t, gnmiUpdate != nil)
	if gnmiUpdate != nil {
		assert.Equal(t, 2, len(gnmiUpdate.Path.Elem))
		elem0 := gnmiUpdate.Path.Elem[0]
		assert.Equal(t, "test1", elem0.Name)
		assert.Equal(t, 0, len(elem0.Key))
		elem1 := gnmiUpdate.Path.Elem[1]
		assert.Equal(t, "test2", elem1.Name)
		assert.Equal(t, 1, len(elem1.Key))
		key1, ok := elem1.Key["name"]
		assert.Assert(t, ok)
		assert.Equal(t, "t1", key1)
		assert.Equal(t, desc, gnmiUpdate.Val.GetStringVal())
	}
}

func Test_ReplaceUnknownKey(t *testing.T) {
	desc := "this is a description"
	gnmiUpdate, err := UpdateForElement(
		aether_2_0_0.AccessProfile_AccessProfile_AccessProfile{Description: &desc}.Description,
		"/test1/test2/{"+UnknownKey+"}", UnknownID)
	assert.NilError(t, err)
	assert.Assert(t, gnmiUpdate != nil)
	if gnmiUpdate != nil {
		keyID, ok := gnmiUpdate.Path.Elem[1].Key[UnknownKey]
		assert.Equal(t, true, ok)
		assert.Equal(t, UnknownID, keyID)
		err = ReplaceUnknownKey(gnmiUpdate, "known_key", "known_value", UnknownKey, UnknownID)
		assert.NilError(t, err, "unexpected error")
		keyID, ok = gnmiUpdate.Path.Elem[1].Key["known_key"]
		assert.Equal(t, true, ok)
		assert.Equal(t, "known_value", keyID)
	}
}
