// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package gnmiutils

import (
	"fmt"
	"github.com/openconfig/gnmi/proto/gnmi"
	"strings"
)

// NewGnmiGetRequest creates a GetRequest from a REST call
func NewGnmiGetRequest(openapiPath string, target string, pathParams ...string) (*gnmi.GetRequest, error) {
	gnmiGet := new(gnmi.GetRequest)
	gnmiGet.Path = make([]*gnmi.Path, 1)
	gnmiGet.Path[0] = &gnmi.Path{
		Elem: make([]*gnmi.PathElem, 0),
	}

	oapiParts := strings.Split(openapiPath, "/")
	if len(oapiParts) < 5 {
		return nil, fmt.Errorf("expected path to have >=4 parts e.g. api,ver,device,path Got %v", oapiParts)
	}
	gnmiGet.Path[0].Target = target
	elemCount := 0
	paramCount := 0
	for i := 4; i < len(oapiParts); i++ {
		if strings.Contains(oapiParts[i], "{") { // Is a key
			if gnmiGet.Path[0].Elem[elemCount-1].Key == nil {
				gnmiGet.Path[0].Elem[elemCount-1].Key = make(map[string]string)
			}
			gnmiGet.Path[0].Elem[elemCount-1].Key[oapiParts[i]] = pathParams[paramCount]
			paramCount++
		} else {
			pathElem := gnmi.PathElem{
				Name: oapiParts[i],
			}
			gnmiGet.Path[0].Elem = append(gnmiGet.Path[0].Elem, &pathElem)
			elemCount++
		}
	}

	return gnmiGet, nil
}

// GetResponseUpdate -- extract the single Update from the GetResponse
func GetResponseUpdate(gr *gnmi.GetResponse, err error) (*gnmi.TypedValue_JsonVal, error) {
	if err != nil {
		return nil, err
	}
	if len(gr.Notification) != 1 {
		return nil, fmt.Errorf("unexpected number of GetResponse notifications %d", len(gr.Notification))
	}
	n0 := gr.Notification[0]
	if len(n0.Update) != 1 {
		return nil, fmt.Errorf("unexpected number of GetResponse notification updates %d", len(n0.Update))
	}
	u0 := n0.Update[0]
	jsonVal, ok := u0.Val.Value.(*gnmi.TypedValue_JsonVal)
	if !ok {
		return nil, fmt.Errorf("expected type jsonvalue")
	}
	return jsonVal, nil
}
