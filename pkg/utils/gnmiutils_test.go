// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package utils

import (
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

	jsonVal, err := GetResponseUpdate(&gr, nil)
	assert.NilError(t, err, "unexpected error")
	assert.Equal(t, "{testvalue: 't'}", string(jsonVal.JsonVal))
}
