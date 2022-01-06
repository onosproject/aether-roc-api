// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package server

import (
	"encoding/json"
	types2 "github.com/onosproject/aether-roc-api/pkg/aether_2_0_0/types"
	"github.com/onosproject/aether-roc-api/pkg/toplevel/types"
	"gotest.tools/assert"
	"io/ioutil"
	"strings"
	"testing"
)

func Test_encodeToGnmiPatchBody(t *testing.T) {
	patchBodyExampleJSON, err := ioutil.ReadFile("../testdata/PatchBody_Example.json")
	assert.NilError(t, err, "error loading testdata file")
	jsonObj := new(types.PatchBody)
	err = json.Unmarshal(patchBodyExampleJSON, jsonObj)
	assert.NilError(t, err)

	testMap := make(map[string]interface{})
	err = json.Unmarshal(patchBodyExampleJSON, &testMap)
	assert.NilError(t, err)

	updates, deletes, ext100Name, ext101Version, ext102Type, defaultTarget, err :=
		encodeToGnmiPatchBody(jsonObj)
	assert.NilError(t, err)
	assert.Assert(t, ext101Version != nil)
	assert.Assert(t, ext102Type != nil)
	if ext100Name != nil {
		assert.Equal(t, "test-name", *ext100Name)
	}
	if ext101Version != nil {
		assert.Equal(t, "4.0.0", *ext101Version)
	}
	if ext102Type != nil {
		assert.Equal(t, "Aether", *ext102Type)
	}
	assert.Equal(t, "connectivity-service-v4", defaultTarget)
	assert.Equal(t, 94, len(updates))
	for _, upd := range updates {
		switch path := strings.ReplaceAll(upd.Path.String(), "  ", " "); path {
		case `elem:{name:"connectivity-service"} elem:{name:"connectivity-service" key:{key:"id" value:"cs5gtest"}} elem:{name:"acc-prometheus-url"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"http://onf-acc-cluster-prometheus:9090"`, upd.Val.String())
		case `elem:{name:"connectivity-service"} elem:{name:"connectivity-service" key:{key:"id" value:"cs5gtest"}} elem:{name:"core-5g-endpoint"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"http://aether-roc-umbrella-sdcore-test-dummy/v1/config/5g"`, upd.Val.String())
		case `elem:{name:"connectivity-service"} elem:{name:"connectivity-service" key:{key:"id" value:"cs5gtest"}} elem:{name:"description"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"5G Test"`, upd.Val.String())
		case `elem:{name:"connectivity-service"} elem:{name:"connectivity-service" key:{key:"id" value:"cs5gtest"}} elem:{name:"display-name"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"ROC 5G Test Connectivity Service"`, upd.Val.String())
		case `elem:{name:"connectivity-service"} elem:{name:"connectivity-service" key:{key:"id" value:"cs5gtest"}} elem:{name:"id"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"cs5gtest"`, upd.Val.String())

		case `elem:{name:"enterprise"} elem:{name:"enterprise" key:{key:"id" value:"starbucks"}} elem:{name:"id"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"starbucks"`, upd.Val.String())
		case `elem:{name:"enterprise"} elem:{name:"enterprise" key:{key:"id" value:"starbucks"}} elem:{name:"description"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"Starbucks Corporation"`, upd.Val.String())
		case `elem:{name:"enterprise"} elem:{name:"enterprise" key:{key:"id" value:"starbucks"}} elem:{name:"display-name"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"Starbucks Inc."`, upd.Val.String())
		case `elem:{name:"enterprise"} elem:{name:"enterprise" key:{key:"id" value:"starbucks"}} elem:{name:"connectivity-service" key:{key:"connectivity-service" value:"cs5gtest"}} elem:{name:"connectivity-service"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"cs5gtest"`, upd.Val.String())
		case `elem:{name:"enterprise"} elem:{name:"enterprise" key:{key:"id" value:"starbucks"}} elem:{name:"connectivity-service" key:{key:"connectivity-service" value:"cs5gtest"}} elem:{name:"enabled"} target:"connectivity-service-v4"`:
			assert.Equal(t, `bool_val:true`, upd.Val.String())

		case `elem:{name:"site"} elem:{name:"site" key:{key:"id" value:"starbucks-newyork"}} elem:{name:"id"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"starbucks-newyork"`, upd.Val.String())
		case `elem:{name:"site"} elem:{name:"site" key:{key:"id" value:"starbucks-newyork"}} elem:{name:"description"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"Starbucks New York"`, upd.Val.String())
		case `elem:{name:"site"} elem:{name:"site" key:{key:"id" value:"starbucks-newyork"}} elem:{name:"enterprise"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"starbucks"`, upd.Val.String())
		case `elem:{name:"site"} elem:{name:"site" key:{key:"id" value:"starbucks-newyork"}} elem:{name:"display-name"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"New York"`, upd.Val.String())
		case `elem:{name:"site"} elem:{name:"site" key:{key:"id" value:"starbucks-newyork"}} elem:{name:"imsi-definition"} elem:{name:"mcc"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"021"`, upd.Val.String())
		case `elem:{name:"site"} elem:{name:"site" key:{key:"id" value:"starbucks-newyork"}} elem:{name:"imsi-definition"} elem:{name:"mnc"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"32"`, upd.Val.String())
		case `elem:{name:"site"} elem:{name:"site" key:{key:"id" value:"starbucks-newyork"}} elem:{name:"imsi-definition"} elem:{name:"enterprise"} target:"connectivity-service-v4"`:
			assert.Equal(t, `uint_val:2`, upd.Val.String())
		case `elem:{name:"site"} elem:{name:"site" key:{key:"id" value:"starbucks-newyork"}} elem:{name:"imsi-definition"} elem:{name:"format"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"CCCNNNEEESSSSSS"`, upd.Val.String())
		case `elem:{name:"site"} elem:{name:"site" key:{key:"id" value:"starbucks-newyork"}} elem:{name:"monitoring"} elem:{name:"edge-cluster-prometheus-url"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"http://starbucks-ny-edge-cluster-prometheus:9090"`, upd.Val.String())
		case `elem:{name:"site"} elem:{name:"site" key:{key:"id" value:"starbucks-newyork"}} elem:{name:"monitoring"} elem:{name:"edge-monitoring-prometheus-url"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"http://rancher-monitoring-prometheus.cattle-monitoring-system.svc:9090"`, upd.Val.String())
		case `elem:{name:"site"} elem:{name:"site" key:{key:"id" value:"starbucks-newyork"}} elem:{name:"monitoring"} elem:{name:"edge-device" key:{key:"edge-device-id" value:"monitoring-pi-1"}} elem:{name:"description"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"monitoring device placed near the cash registers"`, upd.Val.String())
		case `elem:{name:"site"} elem:{name:"site" key:{key:"id" value:"starbucks-newyork"}} elem:{name:"monitoring"} elem:{name:"edge-device" key:{key:"edge-device-id" value:"monitoring-pi-1"}} elem:{name:"display-name"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"cash registers"`, upd.Val.String())
		case `elem:{name:"site"} elem:{name:"site" key:{key:"id" value:"starbucks-newyork"}} elem:{name:"monitoring"} elem:{name:"edge-device" key:{key:"edge-device-id" value:"monitoring-pi-1"}} elem:{name:"edge-device-id"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"monitoring-pi-1"`, upd.Val.String())
		case `elem:{name:"site"} elem:{name:"site" key:{key:"id" value:"starbucks-newyork"}} elem:{name:"small-cell" key:{key:"small-cell-id" value:"cell4"}} elem:{name:"small-cell-id"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"cell4"`, upd.Val.String())
		case `elem:{name:"site"} elem:{name:"site" key:{key:"id" value:"starbucks-newyork"}} elem:{name:"small-cell" key:{key:"small-cell-id" value:"cell4"}} elem:{name:"address"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"ap2.newyork.starbucks.com"`, upd.Val.String())
		case `elem:{name:"site"} elem:{name:"site" key:{key:"id" value:"starbucks-newyork"}} elem:{name:"small-cell" key:{key:"small-cell-id" value:"cell4"}} elem:{name:"enable"} target:"connectivity-service-v4"`:
			assert.Equal(t, `bool_val:true`, upd.Val.String())
		case `elem:{name:"site"} elem:{name:"site" key:{key:"id" value:"starbucks-newyork"}} elem:{name:"small-cell" key:{key:"small-cell-id" value:"cell4"}} elem:{name:"tac"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"8002"`, upd.Val.String())

		case `elem:{name:"application"} elem:{name:"application" key:{key:"id" value:"starbucks-nvr"}} elem:{name:"id"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"starbucks-nvr"`, upd.Val.String())
		case `elem:{name:"application"} elem:{name:"application" key:{key:"id" value:"starbucks-nvr"}} elem:{name:"description"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"Network Video Recorder"`, upd.Val.String())
		case `elem:{name:"application"} elem:{name:"application" key:{key:"id" value:"starbucks-nvr"}} elem:{name:"display-name"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"NVR"`, upd.Val.String())
		case `elem:{name:"application"} elem:{name:"application" key:{key:"id" value:"starbucks-nvr"}} elem:{name:"address"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"nvr.starbucks.com"`, upd.Val.String())
		case `elem:{name:"application"} elem:{name:"application" key:{key:"id" value:"starbucks-nvr"}} elem:{name:"enterprise"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"starbucks"`, upd.Val.String())
		case `elem:{name:"application"} elem:{name:"application" key:{key:"id" value:"starbucks-nvr"}} elem:{name:"endpoint" key:{key:"endpoint-id" value:"rtsp"}} elem:{name:"endpoint-id"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"rtsp"`, upd.Val.String())
		case `elem:{name:"application"} elem:{name:"application" key:{key:"id" value:"starbucks-nvr"}} elem:{name:"endpoint" key:{key:"endpoint-id" value:"rtsp"}} elem:{name:"port-end"} target:"connectivity-service-v4"`:
			assert.Equal(t, `uint_val:3330`, upd.Val.String())
		case `elem:{name:"application"} elem:{name:"application" key:{key:"id" value:"starbucks-nvr"}} elem:{name:"endpoint" key:{key:"endpoint-id" value:"rtsp"}} elem:{name:"port-start"} target:"connectivity-service-v4"`:
			assert.Equal(t, `uint_val:3316`, upd.Val.String())
		case `elem:{name:"application"} elem:{name:"application" key:{key:"id" value:"starbucks-nvr"}} elem:{name:"endpoint" key:{key:"endpoint-id" value:"rtsp"}} elem:{name:"protocol"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"UDP"`, upd.Val.String())
		case `elem:{name:"application"} elem:{name:"application" key:{key:"id" value:"starbucks-nvr"}} elem:{name:"endpoint" key:{key:"endpoint-id" value:"rtsp"}} elem:{name:"traffic-class"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"class-1"`, upd.Val.String())
		case `elem:{name:"application"} elem:{name:"application" key:{key:"id" value:"starbucks-nvr"}} elem:{name:"endpoint" key:{key:"endpoint-id" value:"rtsp"}} elem:{name:"mbr"} elem:{name:"downlink"} target:"connectivity-service-v4"`:
			assert.Equal(t, `uint_val:1000000`, upd.Val.String())
		case `elem:{name:"application"} elem:{name:"application" key:{key:"id" value:"starbucks-nvr"}} elem:{name:"endpoint" key:{key:"endpoint-id" value:"rtsp"}} elem:{name:"mbr"} elem:{name:"uplink"} target:"connectivity-service-v4"`:
			assert.Equal(t, `uint_val:1000000`, upd.Val.String())

		case `elem:{name:"device-group"} elem:{name:"device-group" key:{key:"id" value:"starbucks-newyork-cameras"}} elem:{name:"display-name"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"New York Cameras"`, upd.Val.String())
		case `elem:{name:"device-group"} elem:{name:"device-group" key:{key:"id" value:"starbucks-newyork-cameras"}} elem:{name:"id"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"starbucks-newyork-cameras"`, upd.Val.String())
		case `elem:{name:"device-group"} elem:{name:"device-group" key:{key:"id" value:"starbucks-newyork-cameras"}} elem:{name:"ip-domain"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"starbucks-newyork"`, upd.Val.String())
		case `elem:{name:"device-group"} elem:{name:"device-group" key:{key:"id" value:"starbucks-newyork-cameras"}} elem:{name:"site"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"starbucks-newyork"`, upd.Val.String())
		case `elem:{name:"device-group"} elem:{name:"device-group" key:{key:"id" value:"starbucks-newyork-cameras"}} elem:{name:"device"} elem:{name:"mbr"} elem:{name:"downlink"} target:"connectivity-service-v4"`:
			assert.Equal(t, `uint_val:0`, upd.Val.String())
		case `elem:{name:"device-group"} elem:{name:"device-group" key:{key:"id" value:"starbucks-newyork-cameras"}} elem:{name:"device"} elem:{name:"mbr"} elem:{name:"uplink"} target:"connectivity-service-v4"`:
			assert.Equal(t, `uint_val:1000000`, upd.Val.String())
		case `elem:{name:"device-group"} elem:{name:"device-group" key:{key:"id" value:"starbucks-newyork-cameras"}} elem:{name:"device"} elem:{name:"traffic-class"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"class-1"`, upd.Val.String())
		case `elem:{name:"device-group"} elem:{name:"device-group" key:{key:"id" value:"starbucks-newyork-cameras"}} elem:{name:"imsis" key:{key:"imsi-id" value:"front"}} elem:{name:"imsi-id"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"front"`, upd.Val.String())
		case `elem:{name:"device-group"} elem:{name:"device-group" key:{key:"id" value:"starbucks-newyork-cameras"}} elem:{name:"imsis" key:{key:"imsi-id" value:"front"}} elem:{name:"imsi-range-from"} target:"connectivity-service-v4"`:
			assert.Equal(t, `uint_val:40`, upd.Val.String())
		case `elem:{name:"device-group"} elem:{name:"device-group" key:{key:"id" value:"starbucks-newyork-cameras"}} elem:{name:"imsis" key:{key:"imsi-id" value:"front"}} elem:{name:"imsi-range-to"} target:"connectivity-service-v4"`:
			assert.Equal(t, `uint_val:41`, upd.Val.String())

		case `elem:{name:"ip-domain"} elem:{name:"ip-domain" key:{key:"id" value:"starbucks-newyork"}} elem:{name:"id"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"starbucks-newyork"`, upd.Val.String())
		case `elem:{name:"ip-domain"} elem:{name:"ip-domain" key:{key:"id" value:"starbucks-newyork"}} elem:{name:"admin-status"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"ENABLE"`, upd.Val.String())
		case `elem:{name:"ip-domain"} elem:{name:"ip-domain" key:{key:"id" value:"starbucks-newyork"}} elem:{name:"description"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"New York IP Domain"`, upd.Val.String())
		case `elem:{name:"ip-domain"} elem:{name:"ip-domain" key:{key:"id" value:"starbucks-newyork"}} elem:{name:"display-name"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"New York"`, upd.Val.String())
		case `elem:{name:"ip-domain"} elem:{name:"ip-domain" key:{key:"id" value:"starbucks-newyork"}} elem:{name:"dns-primary"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"8.8.8.1"`, upd.Val.String())
		case `elem:{name:"ip-domain"} elem:{name:"ip-domain" key:{key:"id" value:"starbucks-newyork"}} elem:{name:"dns-secondary"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"8.8.8.2"`, upd.Val.String())
		case `elem:{name:"ip-domain"} elem:{name:"ip-domain" key:{key:"id" value:"starbucks-newyork"}} elem:{name:"mtu"} target:"connectivity-service-v4"`:
			assert.Equal(t, `uint_val:57600`, upd.Val.String())
		case `elem:{name:"ip-domain"} elem:{name:"ip-domain" key:{key:"id" value:"starbucks-newyork"}} elem:{name:"subnet"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"254.186.117.251/31"`, upd.Val.String())
		case `elem:{name:"ip-domain"} elem:{name:"ip-domain" key:{key:"id" value:"starbucks-newyork"}} elem:{name:"enterprise"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"starbucks"`, upd.Val.String())
		case `elem:{name:"ip-domain"} elem:{name:"ip-domain" key:{key:"id" value:"starbucks-newyork"}} elem:{name:"dnn"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"somednn"`, upd.Val.String())

		case `elem:{name:"traffic-class"} elem:{name:"traffic-class" key:{key:"id" value:"class-1"}} elem:{name:"description"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"High Priority TC"`, upd.Val.String())
		case `elem:{name:"traffic-class"} elem:{name:"traffic-class" key:{key:"id" value:"class-1"}} elem:{name:"display-name"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"Class 1"`, upd.Val.String())
		case `elem:{name:"traffic-class"} elem:{name:"traffic-class" key:{key:"id" value:"class-1"}} elem:{name:"id"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"class-1"`, upd.Val.String())
		case `elem:{name:"traffic-class"} elem:{name:"traffic-class" key:{key:"id" value:"class-1"}} elem:{name:"pdb"} target:"connectivity-service-v4"`:
			assert.Equal(t, `uint_val:100`, upd.Val.String())
		case `elem:{name:"traffic-class"} elem:{name:"traffic-class" key:{key:"id" value:"class-1"}} elem:{name:"pelr"} target:"connectivity-service-v4"`:
			assert.Equal(t, `int_val:10`, upd.Val.String())
		case `elem:{name:"traffic-class"} elem:{name:"traffic-class" key:{key:"id" value:"class-1"}} elem:{name:"qci"} target:"connectivity-service-v4"`:
			assert.Equal(t, `uint_val:10`, upd.Val.String())
		case `elem:{name:"traffic-class"} elem:{name:"traffic-class" key:{key:"id" value:"class-1"}} elem:{name:"arp"} target:"connectivity-service-v4"`:
			assert.Equal(t, `uint_val:1`, upd.Val.String())

		case `elem:{name:"template"} elem:{name:"template" key:{key:"id" value:"template-1"}} elem:{name:"id"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"template-1"`, upd.Val.String())
		case `elem:{name:"template"} elem:{name:"template" key:{key:"id" value:"template-1"}} elem:{name:"sd"} target:"connectivity-service-v4"`:
			assert.Equal(t, `uint_val:10886763`, upd.Val.String())
		case `elem:{name:"template"} elem:{name:"template" key:{key:"id" value:"template-1"}} elem:{name:"sst"} target:"connectivity-service-v4"`:
			assert.Equal(t, `uint_val:158`, upd.Val.String())
		case `elem:{name:"template"} elem:{name:"template" key:{key:"id" value:"template-1"}} elem:{name:"default-behavior"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"DENY-ALL"`, upd.Val.String())
		case `elem:{name:"template"} elem:{name:"template" key:{key:"id" value:"template-1"}} elem:{name:"slice"} elem:{name:"mbr"} elem:{name:"uplink"} target:"connectivity-service-v4"`:
			assert.Equal(t, `uint_val:10000000`, upd.Val.String())
		case `elem:{name:"template"} elem:{name:"template" key:{key:"id" value:"template-1"}} elem:{name:"slice"} elem:{name:"mbr"} elem:{name:"downlink"} target:"connectivity-service-v4"`:
			assert.Equal(t, `uint_val:5000000`, upd.Val.String())
		case `elem:{name:"template"} elem:{name:"template" key:{key:"id" value:"template-1"}} elem:{name:"slice"} elem:{name:"mbr"} elem:{name:"uplink-burst-size"} target:"connectivity-service-v4"`:
			assert.Equal(t, `uint_val:600000`, upd.Val.String())
		case `elem:{name:"template"} elem:{name:"template" key:{key:"id" value:"template-1"}} elem:{name:"slice"} elem:{name:"mbr"} elem:{name:"downlink-burst-size"} target:"connectivity-service-v4"`:
			assert.Equal(t, `uint_val:600000`, upd.Val.String())

		case `elem:{name:"upf"} elem:{name:"upf" key:{key:"id" value:"starbucks-newyork-pool-entry1"}} elem:{name:"id"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"starbucks-newyork-pool-entry1"`, upd.Val.String())
		case `elem:{name:"upf"} elem:{name:"upf" key:{key:"id" value:"starbucks-newyork-pool-entry1"}} elem:{name:"address"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"entry1.upfpool.newyork.starbucks.com"`, upd.Val.String())
		case `elem:{name:"upf"} elem:{name:"upf" key:{key:"id" value:"starbucks-newyork-pool-entry1"}} elem:{name:"enterprise"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"starbucks"`, upd.Val.String())
		case `elem:{name:"upf"} elem:{name:"upf" key:{key:"id" value:"starbucks-newyork-pool-entry1"}} elem:{name:"site"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"starbucks-newyork"`, upd.Val.String())
		case `elem:{name:"upf"} elem:{name:"upf" key:{key:"id" value:"starbucks-newyork-pool-entry1"}} elem:{name:"port"} target:"connectivity-service-v4"`:
			assert.Equal(t, `uint_val:6161`, upd.Val.String())

		case `elem:{name:"vcs"} elem:{name:"vcs" key:{key:"id" value:"starbucks-newyork-cameras"}} elem:{name:"id"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"starbucks-newyork-cameras"`, upd.Val.String())
		case `elem:{name:"vcs"} elem:{name:"vcs" key:{key:"id" value:"starbucks-newyork-cameras"}} elem:{name:"default-behavior"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"DENY-ALL"`, upd.Val.String())
		case `elem:{name:"vcs"} elem:{name:"vcs" key:{key:"id" value:"starbucks-newyork-cameras"}} elem:{name:"enterprise"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"starbucks"`, upd.Val.String())
		case `elem:{name:"vcs"} elem:{name:"vcs" key:{key:"id" value:"starbucks-newyork-cameras"}} elem:{name:"upf"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"starbucks-newyork-pool-entry1"`, upd.Val.String())
		case `elem:{name:"vcs"} elem:{name:"vcs" key:{key:"id" value:"starbucks-newyork-cameras"}} elem:{name:"site"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"starbucks-newyork"`, upd.Val.String())
		case `elem:{name:"vcs"} elem:{name:"vcs" key:{key:"id" value:"starbucks-newyork-cameras"}} elem:{name:"sd"} target:"connectivity-service-v4"`:
			assert.Equal(t, `uint_val:8284729`, upd.Val.String())
		case `elem:{name:"vcs"} elem:{name:"vcs" key:{key:"id" value:"starbucks-newyork-cameras"}} elem:{name:"sst"} target:"connectivity-service-v4"`:
			assert.Equal(t, `uint_val:127`, upd.Val.String())
		case `elem:{name:"vcs"} elem:{name:"vcs" key:{key:"id" value:"starbucks-newyork-cameras"}} elem:{name:"slice"} elem:{name:"mbr"} elem:{name:"uplink"} target:"connectivity-service-v4"`:
			assert.Equal(t, `uint_val:10000000`, upd.Val.String())
		case `elem:{name:"vcs"} elem:{name:"vcs" key:{key:"id" value:"starbucks-newyork-cameras"}} elem:{name:"slice"} elem:{name:"mbr"} elem:{name:"downlink"} target:"connectivity-service-v4"`:
			assert.Equal(t, `uint_val:5000000`, upd.Val.String())
		case `elem:{name:"vcs"} elem:{name:"vcs" key:{key:"id" value:"starbucks-newyork-cameras"}} elem:{name:"slice"} elem:{name:"mbr"} elem:{name:"uplink-burst-size"} target:"connectivity-service-v4"`:
			assert.Equal(t, `uint_val:600000`, upd.Val.String())
		case `elem:{name:"vcs"} elem:{name:"vcs" key:{key:"id" value:"starbucks-newyork-cameras"}} elem:{name:"slice"} elem:{name:"mbr"} elem:{name:"downlink-burst-size"} target:"connectivity-service-v4"`:
			assert.Equal(t, `uint_val:600000`, upd.Val.String())
		case `elem:{name:"vcs"} elem:{name:"vcs" key:{key:"id" value:"starbucks-newyork-cameras"}} elem:{name:"device-group" key:{key:"device-group" value:"starbucks-newyork-cameras"}} elem:{name:"device-group"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"starbucks-newyork-cameras"`, upd.Val.String())
		case `elem:{name:"vcs"} elem:{name:"vcs" key:{key:"id" value:"starbucks-newyork-cameras"}} elem:{name:"device-group" key:{key:"device-group" value:"starbucks-newyork-cameras"}} elem:{name:"enable"} target:"connectivity-service-v4"`:
			assert.Equal(t, `bool_val:true`, upd.Val.String())
		case `elem:{name:"vcs"} elem:{name:"vcs" key:{key:"id" value:"starbucks-newyork-cameras"}} elem:{name:"filter" key:{key:"application" value:"starbucks-nvr"}} elem:{name:"allow"} target:"connectivity-service-v4"`:
			assert.Equal(t, `bool_val:true`, upd.Val.String())
		case `elem:{name:"vcs"} elem:{name:"vcs" key:{key:"id" value:"starbucks-newyork-cameras"}} elem:{name:"filter" key:{key:"application" value:"starbucks-nvr"}} elem:{name:"application"} target:"connectivity-service-v4"`:
			assert.Equal(t, `string_val:"starbucks-nvr"`, upd.Val.String())

		default:
			t.Fatalf("unexpected path %s", path)
		}
	}

	assert.Equal(t, 5, len(deletes))
	for _, del := range deletes {
		switch path := strings.ReplaceAll(del.String(), "  ", " "); path {
		case `elem:{name:"access-profile"} elem:{name:"access-profile" key:{key:"id" value:"ap3d"}} elem:{name:"id"} target:"connectivity-service-v2"`:
		case `elem:{name:"vcs"} elem:{name:"vcs" key:{key:"id" value:"vcs-to-delete-from"}} elem:{name:"filter" key:{key:"application" value:"application-to-delete-the-allow-from"}} elem:{name:"allow"} target:"connectivity-service-v4"`:
		case `elem:{name:"vcs"} elem:{name:"vcs" key:{key:"id" value:"vcs-to-delete-from-2"}} elem:{name:"filter" key:{key:"application" value:"application-to-delete"}} elem:{name:"application"} target:"connectivity-service-v4"`:
		case `elem:{name:"vcs"} elem:{name:"vcs" key:{key:"id" value:"vcs-to-delete-from-3"}} elem:{name:"device-group" key:{key:"device-group" value:"device-group-to-delete-the-enable-from"}} elem:{name:"enable"} target:"connectivity-service-v4"`:
		case `elem:{name:"vcs"} elem:{name:"vcs" key:{key:"id" value:"vcs-to-delete-from-4"}} elem:{name:"device-group" key:{key:"device-group" value:"device-group-to-delete"}} elem:{name:"device-group"} target:"connectivity-service-v4"`:
		case `elem:{name:"vcs"} elem:{name:"vcs" key:{key:"id" value:"vcs-to-delete-from-5"}} elem:{name:"description"} target:"connectivity-service-v4"`:
		default:
			t.Fatalf("unexpected path %s", path)
		}
	}
}

func Test_addProps(t *testing.T) {
	desc1 := "desc1"
	disp1 := "display 1"
	mnc := "456"
	id1 := "id1"
	enterprise := int32(789)
	unchangedSite := "enterprise"
	unchangedImsi := "mcc"
	addPropsSite := make(map[string]types2.AdditionalPropertyUnchanged)
	addPropsSite["additional-properties"] = types2.AdditionalPropertyUnchanged{Unchanged: &unchangedSite}
	addPropsImsi := make(map[string]types2.AdditionalPropertyUnchanged)
	addPropsImsi["additional-properties"] = types2.AdditionalPropertyUnchanged{Unchanged: &unchangedImsi}

	ap1 := types2.SiteSite{
		Description: &desc1,
		DisplayName: &disp1,
		Id:          id1,
		ImsiDefinition: &types2.SiteSiteImsiDefinition{
			Enterprise:           enterprise,
			Format:               "CCCNNNEEESSSSSS",
			Mnc:                  mnc,
			AdditionalProperties: addPropsImsi,
		},
		AdditionalProperties: addPropsSite,
	}

	bytes, err := json.Marshal(ap1)
	assert.NilError(t, err)
	assert.Equal(t,
		`{"additional-properties":{"unchanged":"enterprise"},"description":"desc1","display-name":"display 1","enterprise":"","id":"id1","imsi-definition":{"additional-properties":{"unchanged":"mcc"},"enterprise":789,"format":"CCCNNNEEESSSSSS","mcc":"","mnc":"456"}}`,
		string(bytes))
}
