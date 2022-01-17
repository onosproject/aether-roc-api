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
	assert.Equal(t, 196, len(updates))
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

		// Aether 2.0.x
		case `elem:{name:"connectivity-services"} elem:{name:"connectivity-service" key:{key:"id" value:"cs5gtest"}} elem:{name:"acc-prometheus-url"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"./prometheus-acc"`, upd.Val.String())
		case `elem:{name:"connectivity-services"} elem:{name:"connectivity-service" key:{key:"id" value:"cs5gtest"}} elem:{name:"core-5g-endpoint"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"http://aether-roc-umbrella-sdcore-test-dummy/v1/config/5g"`, upd.Val.String())
		case `elem:{name:"connectivity-services"} elem:{name:"connectivity-service" key:{key:"id" value:"cs5gtest"}} elem:{name:"description"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"5G Test"`, upd.Val.String())
		case `elem:{name:"connectivity-services"} elem:{name:"connectivity-service" key:{key:"id" value:"cs5gtest"}} elem:{name:"display-name"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"ROC 5G Test Connectivity Service"`, upd.Val.String())
		case `elem:{name:"connectivity-services"} elem:{name:"connectivity-service" key:{key:"id" value:"cs5gtest"}} elem:{name:"id"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"cs5gtest"`, upd.Val.String())
		case `elem:{name:"connectivity-services"} elem:{name:"connectivity-service" key:{key:"id" value:"cs4gtest"}} elem:{name:"description"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"ROC 4G Test Connectivity Service"`, upd.Val.String())
		case `elem:{name:"connectivity-services"} elem:{name:"connectivity-service" key:{key:"id" value:"cs4gtest"}} elem:{name:"display-name"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"4G Test"`, upd.Val.String())
		case `elem:{name:"connectivity-services"} elem:{name:"connectivity-service" key:{key:"id" value:"cs4gtest"}} elem:{name:"id"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"cs4gtest"`, upd.Val.String())

		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"ent-id"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"acme"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"description"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"ACME Corporation"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"display-name"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"ACME Corp"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"application" key:{key:"app-id" value:"acme-dataacquisition"}} elem:{name:"app-id"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"acme-dataacquisition"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"application" key:{key:"app-id" value:"acme-dataacquisition"}} elem:{name:"address"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"da.acme.com"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"application" key:{key:"app-id" value:"acme-dataacquisition"}} elem:{name:"description"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"Data Acquisition"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"application" key:{key:"app-id" value:"acme-dataacquisition"}} elem:{name:"display-name"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"DA"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"application" key:{key:"app-id" value:"acme-dataacquisition"}} elem:{name:"endpoint" key:{key:"endpoint-id" value:"da"}} elem:{name:"endpoint-id"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"da"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"application" key:{key:"app-id" value:"acme-dataacquisition"}} elem:{name:"endpoint" key:{key:"endpoint-id" value:"da"}} elem:{name:"display-name"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"data acquisition endpoint"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"application" key:{key:"app-id" value:"acme-dataacquisition"}} elem:{name:"endpoint" key:{key:"endpoint-id" value:"da"}} elem:{name:"mbr"} elem:{name:"downlink"} target:"connectivity-service-v2"`:
			assert.Equal(t, `uint_val:1000000`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"application" key:{key:"app-id" value:"acme-dataacquisition"}} elem:{name:"endpoint" key:{key:"endpoint-id" value:"da"}} elem:{name:"mbr"} elem:{name:"uplink"} target:"connectivity-service-v2"`:
			assert.Equal(t, `uint_val:2000000`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"application" key:{key:"app-id" value:"acme-dataacquisition"}} elem:{name:"endpoint" key:{key:"endpoint-id" value:"da"}} elem:{name:"port-end"} target:"connectivity-service-v2"`:
			assert.Equal(t, `uint_val:7588`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"application" key:{key:"app-id" value:"acme-dataacquisition"}} elem:{name:"endpoint" key:{key:"endpoint-id" value:"da"}} elem:{name:"port-start"} target:"connectivity-service-v2"`:
			assert.Equal(t, `uint_val:7585`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"application" key:{key:"app-id" value:"acme-dataacquisition"}} elem:{name:"endpoint" key:{key:"endpoint-id" value:"da"}} elem:{name:"protocol"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"TCP"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"application" key:{key:"app-id" value:"acme-dataacquisition"}} elem:{name:"endpoint" key:{key:"endpoint-id" value:"da"}} elem:{name:"traffic-class"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"class-2"`, upd.Val.String())

		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"connectivity-service" key:{key:"connectivity-service" value:"cs5gtest"}} elem:{name:"connectivity-service"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"cs5gtest"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"connectivity-service" key:{key:"connectivity-service" value:"cs5gtest"}} elem:{name:"enabled"} target:"connectivity-service-v2"`:
			assert.Equal(t, `bool_val:true`, upd.Val.String())

		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"site-id"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"acme-chicago"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"description"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"ACME HQ"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"display-name"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"Chicago"`, upd.Val.String())

		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"imsi-definition"} elem:{name:"enterprise"} target:"connectivity-service-v2"`:
			assert.Equal(t, `uint_val:1`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"imsi-definition"} elem:{name:"mnc"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"456"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"imsi-definition"} elem:{name:"mcc"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"123"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"imsi-definition"} elem:{name:"format"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"CCCNNNEEESSSSSS"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"monitoring"} elem:{name:"edge-cluster-prometheus-url"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"prometheus-ace1"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"monitoring"} elem:{name:"edge-monitoring-prometheus-url"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"prometheus-amp"`, upd.Val.String())

		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"monitoring"} elem:{name:"edge-device" key:{key:"edge-device-id" value:"acme-chicago-monitoring-pi-1"}} elem:{name:"description"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"monitoring device placed near the sprocket manufacturing machine"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"monitoring"} elem:{name:"edge-device" key:{key:"edge-device-id" value:"acme-chicago-monitoring-pi-1"}} elem:{name:"display-name"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"sprocket monitoring pi"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"monitoring"} elem:{name:"edge-device" key:{key:"edge-device-id" value:"acme-chicago-monitoring-pi-1"}} elem:{name:"edge-device-id"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"acme-chicago-monitoring-pi-1"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"monitoring"} elem:{name:"edge-device" key:{key:"edge-device-id" value:"acme-chicago-monitoring-pi-2"}} elem:{name:"description"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"monitoring device placed near the widget refinisher"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"monitoring"} elem:{name:"edge-device" key:{key:"edge-device-id" value:"acme-chicago-monitoring-pi-2"}} elem:{name:"display-name"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"widget monitoring pi"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"monitoring"} elem:{name:"edge-device" key:{key:"edge-device-id" value:"acme-chicago-monitoring-pi-2"}} elem:{name:"edge-device-id"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"acme-chicago-monitoring-pi-2"`, upd.Val.String())

		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"device-group" key:{key:"dg-id" value:"acme-chicago-default"}} elem:{name:"dg-id"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"acme-chicago-default"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"device-group" key:{key:"dg-id" value:"acme-chicago-default"}} elem:{name:"display-name"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"ACME Chicago Inventory"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"device-group" key:{key:"dg-id" value:"acme-chicago-default"}} elem:{name:"ip-domain"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"acme-chicago"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"device-group" key:{key:"dg-id" value:"acme-chicago-robots"}} elem:{name:"device"} elem:{name:"mbr"} elem:{name:"downlink"} target:"connectivity-service-v2"`:
			assert.Equal(t, `uint_val:1000000`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"device-group" key:{key:"dg-id" value:"acme-chicago-robots"}} elem:{name:"device"} elem:{name:"mbr"} elem:{name:"uplink"} target:"connectivity-service-v2"`:
			assert.Equal(t, `uint_val:5000000`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"device-group" key:{key:"dg-id" value:"acme-chicago-robots"}} elem:{name:"device"} elem:{name:"traffic-class"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"class-1"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"device-group" key:{key:"dg-id" value:"acme-chicago-robots"}} elem:{name:"dg-id"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"acme-chicago-robots"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"device-group" key:{key:"dg-id" value:"acme-chicago-robots"}} elem:{name:"display-name"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"ACME Robots"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"device-group" key:{key:"dg-id" value:"acme-chicago-robots"}} elem:{name:"ip-domain"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"acme-chicago"`, upd.Val.String())

		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"device-group" key:{key:"dg-id" value:"acme-chicago-robots"}} elem:{name:"imsis" key:{key:"imsi-id" value:"production"}} elem:{name:"display-name"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"production robots"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"device-group" key:{key:"dg-id" value:"acme-chicago-robots"}} elem:{name:"imsis" key:{key:"imsi-id" value:"production"}} elem:{name:"imsi-id"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"production"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"device-group" key:{key:"dg-id" value:"acme-chicago-robots"}} elem:{name:"imsis" key:{key:"imsi-id" value:"production"}} elem:{name:"imsi-range-from"} target:"connectivity-service-v2"`:
			assert.Equal(t, `uint_val:0`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"device-group" key:{key:"dg-id" value:"acme-chicago-robots"}} elem:{name:"imsis" key:{key:"imsi-id" value:"production"}} elem:{name:"imsi-range-to"} target:"connectivity-service-v2"`:
			assert.Equal(t, `uint_val:3`, upd.Val.String())

		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"ip-domain" key:{key:"ip-id" value:"acme-chicago"}} elem:{name:"admin-status"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"DISABLE"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"ip-domain" key:{key:"ip-id" value:"acme-chicago"}} elem:{name:"description"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"Chicago IP Domain"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"ip-domain" key:{key:"ip-id" value:"acme-chicago"}} elem:{name:"display-name"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"Chicago"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"ip-domain" key:{key:"ip-id" value:"acme-chicago"}} elem:{name:"dnn"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"dnnacme"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"ip-domain" key:{key:"ip-id" value:"acme-chicago"}} elem:{name:"dns-primary"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"8.8.8.4"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"ip-domain" key:{key:"ip-id" value:"acme-chicago"}} elem:{name:"dns-secondary"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"8.8.8.8"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"ip-domain" key:{key:"ip-id" value:"acme-chicago"}} elem:{name:"ip-id"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"acme-chicago"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"ip-domain" key:{key:"ip-id" value:"acme-chicago"}} elem:{name:"mtu"} target:"connectivity-service-v2"`:
			assert.Equal(t, `uint_val:12690`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"ip-domain" key:{key:"ip-id" value:"acme-chicago"}} elem:{name:"subnet"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"163.25.44.0/31"`, upd.Val.String())

		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"small-cell" key:{key:"small-cell-id" value:"cell1"}} elem:{name:"small-cell-id"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"cell1"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"small-cell" key:{key:"small-cell-id" value:"cell1"}} elem:{name:"display-name"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"cell number one"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"small-cell" key:{key:"small-cell-id" value:"cell1"}} elem:{name:"address"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"ap2.chicago.acme.com"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"small-cell" key:{key:"small-cell-id" value:"cell1"}} elem:{name:"enable"} target:"connectivity-service-v2"`:
			assert.Equal(t, `bool_val:true`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"small-cell" key:{key:"small-cell-id" value:"cell1"}} elem:{name:"tac"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"8002"`, upd.Val.String())

		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"upf" key:{key:"upf-id" value:"acme-chicago-pool-entry2"}} elem:{name:"upf-id"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"acme-chicago-pool-entry2"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"upf" key:{key:"upf-id" value:"acme-chicago-pool-entry2"}} elem:{name:"description"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"Chicago UPF Pool - Entry 2"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"upf" key:{key:"upf-id" value:"acme-chicago-pool-entry2"}} elem:{name:"display-name"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"Chicago Pool 2"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"upf" key:{key:"upf-id" value:"acme-chicago-pool-entry2"}} elem:{name:"address"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"entry2.upfpool.chicago.acme.com"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"upf" key:{key:"upf-id" value:"acme-chicago-pool-entry2"}} elem:{name:"port"} target:"connectivity-service-v2"`:
			assert.Equal(t, `uint_val:6161`, upd.Val.String())

		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"vcs" key:{key:"vcs-id" value:"acme-chicago-robots"}} elem:{name:"vcs-id"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"acme-chicago-robots"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"vcs" key:{key:"vcs-id" value:"acme-chicago-robots"}} elem:{name:"display-name"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"Chicago Robots VCS"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"vcs" key:{key:"vcs-id" value:"acme-chicago-robots"}} elem:{name:"description"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"Chicago Robots"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"vcs" key:{key:"vcs-id" value:"acme-chicago-robots"}} elem:{name:"default-behavior"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"DENY-ALL"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"vcs" key:{key:"vcs-id" value:"acme-chicago-robots"}} elem:{name:"sd"} target:"connectivity-service-v2"`:
			assert.Equal(t, `uint_val:2973238`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"vcs" key:{key:"vcs-id" value:"acme-chicago-robots"}} elem:{name:"sst"} target:"connectivity-service-v2"`:
			assert.Equal(t, `uint_val:79`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"vcs" key:{key:"vcs-id" value:"acme-chicago-robots"}} elem:{name:"slice"} elem:{name:"mbr"} elem:{name:"downlink"} target:"connectivity-service-v2"`:
			assert.Equal(t, `uint_val:5000000`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"vcs" key:{key:"vcs-id" value:"acme-chicago-robots"}} elem:{name:"slice"} elem:{name:"mbr"} elem:{name:"downlink-burst-size"} target:"connectivity-service-v2"`:
			assert.Equal(t, `uint_val:600000`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"vcs" key:{key:"vcs-id" value:"acme-chicago-robots"}} elem:{name:"slice"} elem:{name:"mbr"} elem:{name:"uplink"} target:"connectivity-service-v2"`:
			assert.Equal(t, `uint_val:5000000`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"vcs" key:{key:"vcs-id" value:"acme-chicago-robots"}} elem:{name:"slice"} elem:{name:"mbr"} elem:{name:"uplink-burst-size"} target:"connectivity-service-v2"`:
			assert.Equal(t, `uint_val:5000000`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"vcs" key:{key:"vcs-id" value:"acme-chicago-robots"}} elem:{name:"upf"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"acme-chicago-pool-entry1"`, upd.Val.String())

		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"vcs" key:{key:"vcs-id" value:"acme-chicago-robots"}} elem:{name:"device-group" key:{key:"device-group" value:"acme-chicago-robots"}} elem:{name:"device-group"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"acme-chicago-robots"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"vcs" key:{key:"vcs-id" value:"acme-chicago-robots"}} elem:{name:"device-group" key:{key:"device-group" value:"acme-chicago-robots"}} elem:{name:"enable"} target:"connectivity-service-v2"`:
			assert.Equal(t, `bool_val:true`, upd.Val.String())

		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"vcs" key:{key:"vcs-id" value:"acme-chicago-robots"}} elem:{name:"filter" key:{key:"application" value:"acme-dataacquisition"}} elem:{name:"allow"} target:"connectivity-service-v2"`:
			assert.Equal(t, `bool_val:false`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"site" key:{key:"site-id" value:"acme-chicago"}} elem:{name:"vcs" key:{key:"vcs-id" value:"acme-chicago-robots"}} elem:{name:"filter" key:{key:"application" value:"acme-dataacquisition"}} elem:{name:"application"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"acme-dataacquisition"`, upd.Val.String())

		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"template" key:{key:"tp-id" value:"template-1"}} elem:{name:"tp-id"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"template-1"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"template" key:{key:"tp-id" value:"template-1"}} elem:{name:"display-name"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"Template 1"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"template" key:{key:"tp-id" value:"template-1"}} elem:{name:"description"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"VCS Template 1"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"template" key:{key:"tp-id" value:"template-1"}} elem:{name:"default-behavior"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"DENY-ALL"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"template" key:{key:"tp-id" value:"template-1"}} elem:{name:"sd"} target:"connectivity-service-v2"`:
			assert.Equal(t, `uint_val:10886763`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"template" key:{key:"tp-id" value:"template-1"}} elem:{name:"sst"} target:"connectivity-service-v2"`:
			assert.Equal(t, `uint_val:158`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"template" key:{key:"tp-id" value:"template-1"}} elem:{name:"slice"} elem:{name:"mbr"} elem:{name:"downlink"} target:"connectivity-service-v2"`:
			assert.Equal(t, `uint_val:5000000`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"template" key:{key:"tp-id" value:"template-1"}} elem:{name:"slice"} elem:{name:"mbr"} elem:{name:"downlink-burst-size"} target:"connectivity-service-v2"`:
			assert.Equal(t, `uint_val:600000`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"template" key:{key:"tp-id" value:"template-1"}} elem:{name:"slice"} elem:{name:"mbr"} elem:{name:"uplink"} target:"connectivity-service-v2"`:
			assert.Equal(t, `uint_val:10000000`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"template" key:{key:"tp-id" value:"template-1"}} elem:{name:"slice"} elem:{name:"mbr"} elem:{name:"uplink-burst-size"} target:"connectivity-service-v2"`:
			assert.Equal(t, `uint_val:600000`, upd.Val.String())

		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"traffic-class" key:{key:"tc-id" value:"class-1"}} elem:{name:"tc-id"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"class-1"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"traffic-class" key:{key:"tc-id" value:"class-1"}} elem:{name:"display-name"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"Class 1"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"traffic-class" key:{key:"tc-id" value:"class-1"}} elem:{name:"description"} target:"connectivity-service-v2"`:
			assert.Equal(t, `string_val:"High Priority TC"`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"traffic-class" key:{key:"tc-id" value:"class-1"}} elem:{name:"qci"} target:"connectivity-service-v2"`:
			assert.Equal(t, `uint_val:10`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"traffic-class" key:{key:"tc-id" value:"class-1"}} elem:{name:"arp"} target:"connectivity-service-v2"`:
			assert.Equal(t, `uint_val:1`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"traffic-class" key:{key:"tc-id" value:"class-1"}} elem:{name:"pelr"} target:"connectivity-service-v2"`:
			assert.Equal(t, `int_val:10`, upd.Val.String())
		case `elem:{name:"enterprises"} elem:{name:"enterprise" key:{key:"ent-id" value:"acme"}} elem:{name:"traffic-class" key:{key:"tc-id" value:"class-1"}} elem:{name:"pdb"} target:"connectivity-service-v2"`:
			assert.Equal(t, `uint_val:100`, upd.Val.String())

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

	ap1 := types2.EnterprisesEnterpriseSite{
		Description: &desc1,
		DisplayName: &disp1,
		SiteId:      id1,
		ImsiDefinition: &types2.EnterprisesEnterpriseSiteImsiDefinition{
			Enterprise:           enterprise,
			Format:               "CCCNNNEEESSSSSS",
			Mnc:                  mnc,
			AdditionalProperties: addPropsImsi,
		},
	}

	bytes, err := json.Marshal(ap1)
	assert.NilError(t, err)
	assert.Equal(t,
		`{"description":"desc1","display-name":"display 1","imsi-definition":{"additional-properties":{"unchanged":"mcc"},"enterprise":789,"format":"CCCNNNEEESSSSSS","mcc":"","mnc":"456"},"site-id":"id1"}`,
		string(bytes))
}
