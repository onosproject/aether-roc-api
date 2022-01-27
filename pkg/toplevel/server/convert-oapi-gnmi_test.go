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
	assert.Equal(t, 208, len(updates))
	for _, upd := range updates {
		switch tgt := upd.Path.Target; tgt {
		case "connectivity-service-v4":
			switch path0 := strings.ReplaceAll(upd.Path.Elem[0].String(), "  ", " "); path0 {
			case `name:"connectivity-service"`:
			case `name:"application"`:
			case `name:"device-group"`:
			case `name:"enterprise"`:
			case `name:"ip-domain"`:
			case `name:"site"`:
			case `name:"template"`:
			case `name:"traffic-class"`:
			case `name:"upf"`:
			case `name:"vcs"`:
			default:
				t.Logf("unhandled v4 update %s", path0)
			}
		case "connectivity-service-v2":
			switch path0 := upd.Path.Elem[0].Name; path0 {
			case "connectivity-services":
				switch path1 := strings.ReplaceAll(upd.Path.Elem[1].String(), "  ", " "); path1 {
				case `name:"connectivity-service" key:{key:"id" value:"cs5gtest"}`:
					switch path2 := upd.Path.Elem[2].String(); path2 {
					case `name:"acc-prometheus-url"`:
						assert.Equal(t, `string_val:"./prometheus-acc"`, upd.Val.String())
					case `name:"core-5g-endpoint"`:
						assert.Equal(t, `string_val:"http://aether-roc-umbrella-sdcore-test-dummy/v1/config/5g"`, upd.Val.String())
					case `name:"description"`:
						assert.Equal(t, `string_val:"5G Test"`, upd.Val.String())
					case `name:"display-name"`:
						assert.Equal(t, `string_val:"ROC 5G Test Connectivity Service"`, upd.Val.String())
					case `name:"id"`:
						assert.Equal(t, `string_val:"cs5gtest"`, upd.Val.String())
					default:
						t.Logf("unhandled v2 update %s %s %s", path0, path1, path2)
					}

				case `name:"connectivity-service" key:{key:"id" value:"cs4gtest"}`:
					switch path2 := upd.Path.Elem[2].String(); path2 {
					case `name:"acc-prometheus-url"`:
						assert.Equal(t, `string_val:"./prometheus-acc"`, upd.Val.String())
					case `name:"core-5g-endpoint"`:
						assert.Equal(t, `string_val:"http://aether-roc-umbrella-sdcore-test-dummy/v1/config/5g"`, upd.Val.String())
					case `name:"description"`:
						assert.Equal(t, `string_val:"ROC 4G Test Connectivity Service"`, upd.Val.String())
					case `name:"display-name"`:
						assert.Equal(t, `string_val:"4G Test"`, upd.Val.String())
					case `name:"id"`:
						assert.Equal(t, `string_val:"cs4gtest"`, upd.Val.String())
					default:
						t.Logf("unhandled v2 update %s %s %s", path0, path1, path2)
					}

				default:
					t.Logf("unhandled v2 update %s %s", path0, path1)
				}
			case "enterprises":
				switch path1 := strings.ReplaceAll(upd.Path.Elem[1].String(), "  ", " "); path1 {
				case `name:"enterprise" key:{key:"ent-id" value:"acme"}`:
					switch path2 := strings.ReplaceAll(upd.Path.Elem[2].String(), "  ", " "); path2 {
					case `name:"ent-id"`:
						assert.Equal(t, `string_val:"acme"`, upd.Val.String())
					case `name:"description"`:
						assert.Equal(t, `string_val:"ACME Corporation"`, upd.Val.String())
					case `name:"display-name"`:
						assert.Equal(t, `string_val:"ACME Corp"`, upd.Val.String())
					case `name:"application" key:{key:"app-id" value:"acme-dataacquisition"}`:
						switch path3 := strings.ReplaceAll(upd.Path.Elem[3].String(), "  ", " "); path3 {
						case `name:"app-id"`:
							assert.Equal(t, `string_val:"acme-dataacquisition"`, upd.Val.String())
						case `name:"description"`:
							assert.Equal(t, `string_val:"Data Acquisition"`, upd.Val.String())
						case `name:"display-name"`:
							assert.Equal(t, `string_val:"DA"`, upd.Val.String())
						case `name:"address"`:
							assert.Equal(t, `string_val:"da.acme.com"`, upd.Val.String())
						case `name:"endpoint" key:{key:"endpoint-id" value:"da"}`:
							switch path4 := strings.ReplaceAll(upd.Path.Elem[4].String(), "  ", " "); path4 {
							case `name:"endpoint-id"`:
								assert.Equal(t, `string_val:"da"`, upd.Val.String())
							case `name:"display-name"`:
								assert.Equal(t, `string_val:"data acquisition endpoint"`, upd.Val.String())
							case `name:"protocol"`:
								assert.Equal(t, `string_val:"TCP"`, upd.Val.String())
							case `name:"port-start"`:
								assert.Equal(t, `uint_val:7585`, upd.Val.String())
							case `name:"port-end"`:
								assert.Equal(t, `uint_val:7588`, upd.Val.String())
							case `name:"traffic-class"`:
								assert.Equal(t, `string_val:"class-2"`, upd.Val.String())
							case `name:"mbr"`:
								switch path5 := strings.ReplaceAll(upd.Path.Elem[5].String(), "  ", " "); path5 {
								case `name:"downlink"`:
									assert.Equal(t, `uint_val:1000000`, upd.Val.String())
								case `name:"uplink"`:
									assert.Equal(t, `uint_val:2000000`, upd.Val.String())
								default:
									t.Logf("unhandled v2 enterprises acme app endpoint update %s %s %s %s %s %s", path0, path1, path2, path3, path4, path5)
								}
							default:
								t.Logf("unhandled v2 enterprises acme app endpoint update %s %s %s %s %s", path0, path1, path2, path3, path4)
							}

						default:
							t.Logf("unhandled v2 enterprises acme app update %s %s %s %s", path0, path1, path2, path3)
						}
					case `name:"connectivity-service" key:{key:"connectivity-service" value:"cs5gtest"}`:
						switch path3 := strings.ReplaceAll(upd.Path.Elem[3].String(), "  ", " "); path3 {
						case `name:"connectivity-service"`:
							assert.Equal(t, `string_val:"cs5gtest"`, upd.Val.String())
						case `name:"enabled"`:
							assert.Equal(t, `bool_val:true`, upd.Val.String())
						default:
							t.Logf("unhandled v2 enterprises acme cs update %s %s %s %s", path0, path1, path2, path3)
						}
					case `name:"site" key:{key:"site-id" value:"acme-chicago"}`:
						switch path3 := strings.ReplaceAll(upd.Path.Elem[3].String(), "  ", " "); path3 {
						case `name:"site-id"`:
							assert.Equal(t, `string_val:"acme-chicago"`, upd.Val.String())
						case `name:"description"`:
							assert.Equal(t, `string_val:"ACME HQ"`, upd.Val.String())
						case `name:"display-name"`:
							assert.Equal(t, `string_val:"Chicago"`, upd.Val.String())
						case `name:"imsi-definition"`:
							switch path4 := strings.ReplaceAll(upd.Path.Elem[4].String(), "  ", " "); path4 {
							case `name:"enterprise"`:
								assert.Equal(t, `uint_val:1`, upd.Val.String())
							case `name:"mcc"`:
								assert.Equal(t, `string_val:"123"`, upd.Val.String())
							case `name:"mnc"`:
								assert.Equal(t, `string_val:"456"`, upd.Val.String())
							case `name:"format"`:
								assert.Equal(t, `string_val:"CCCNNNEEESSSSSS"`, upd.Val.String())
							default:
								t.Logf("unhandled v2 enterprises acme site imsidef update %s %s %s %s %s", path0, path1, path2, path3, path4)
							}
						case `name:"monitoring"`:
							switch path4 := strings.ReplaceAll(upd.Path.Elem[4].String(), "  ", " "); path4 {
							case `name:"edge-cluster-prometheus-url"`:
								assert.Equal(t, `string_val:"prometheus-ace1"`, upd.Val.String())
							case `name:"edge-monitoring-prometheus-url"`:
								assert.Equal(t, `string_val:"prometheus-amp"`, upd.Val.String())
							case `name:"edge-device" key:{key:"edge-device-id" value:"acme-chicago-monitoring-pi-1"}`:
								switch path5 := strings.ReplaceAll(upd.Path.Elem[5].String(), "  ", " "); path5 {
								case `name:"edge-device-id"`:
									assert.Equal(t, `string_val:"acme-chicago-monitoring-pi-1"`, upd.Val.String())
								case `name:"display-name"`:
									assert.Equal(t, `string_val:"sprocket monitoring pi"`, upd.Val.String())
								case `name:"description"`:
									assert.Equal(t, `string_val:"monitoring device placed near the sprocket manufacturing machine"`, upd.Val.String())
								default:
									t.Logf("unhandled v2 enterprises acme site mon edge update %s %s %s %s %s %s", path0, path1, path2, path3, path4, path5)
								}
							case `name:"edge-device" key:{key:"edge-device-id" value:"acme-chicago-monitoring-pi-2"}`:
								switch path5 := strings.ReplaceAll(upd.Path.Elem[5].String(), "  ", " "); path5 {
								case `name:"edge-device-id"`:
									assert.Equal(t, `string_val:"acme-chicago-monitoring-pi-2"`, upd.Val.String())
								case `name:"display-name"`:
									assert.Equal(t, `string_val:"widget monitoring pi"`, upd.Val.String())
								case `name:"description"`:
									assert.Equal(t, `string_val:"monitoring device placed near the widget refinisher"`, upd.Val.String())
								default:
									t.Logf("unhandled v2 enterprises acme site mon edge update %s %s %s %s %s %s", path0, path1, path2, path3, path4, path5)
								}
							default:
								t.Logf("unhandled v2 enterprises acme site mon update %s %s %s %s %s", path0, path1, path2, path3, path4)
							}
						case `name:"device" key:{key:"dev-id" value:"robot-1"}`:
							switch path4 := strings.ReplaceAll(upd.Path.Elem[4].String(), "  ", " "); path4 {
							case `name:"dev-id"`:
								assert.Equal(t, `string_val:"robot-1"`, upd.Val.String())
							case `name:"display-name"`:
								assert.Equal(t, `string_val:"Robot 1"`, upd.Val.String())
							case `name:"description"`:
								assert.Equal(t, `string_val:"The 1st Robot"`, upd.Val.String())
							case `name:"imei"`:
								assert.Equal(t, `uint_val:111222333`, upd.Val.String())
							case `name:"sim-card"`:
								assert.Equal(t, `string_val:"sim-1"`, upd.Val.String())
							default:
								t.Logf("unhandled v2 enterprises acme site device 1 update %s %s %s %s %s", path0, path1, path2, path3, path4)
							}
						case `name:"device" key:{key:"dev-id" value:"robot-2"}`:
							switch path4 := strings.ReplaceAll(upd.Path.Elem[4].String(), "  ", " "); path4 {
							case `name:"dev-id"`:
								assert.Equal(t, `string_val:"robot-2"`, upd.Val.String())
							case `name:"display-name"`:
								assert.Equal(t, `string_val:"Robot 2"`, upd.Val.String())
							case `name:"description"`:
								assert.Equal(t, `string_val:"The 2nd Robot"`, upd.Val.String())
							case `name:"imei"`:
								assert.Equal(t, `uint_val:111222334`, upd.Val.String())
							case `name:"sim-card"`:
								assert.Equal(t, `string_val:"sim-2"`, upd.Val.String())
							default:
								t.Logf("unhandled v2 enterprises acme site device 2 update %s %s %s %s %s", path0, path1, path2, path3, path4)
							}

						case `name:"sim-card" key:{key:"sim-id" value:"sim-1"}`:
							switch path4 := strings.ReplaceAll(upd.Path.Elem[4].String(), "  ", " "); path4 {
							case `name:"sim-id"`:
								assert.Equal(t, `string_val:"sim-1"`, upd.Val.String())
							case `name:"display-name"`:
								assert.Equal(t, `string_val:"Robot 1 Sim"`, upd.Val.String())
							case `name:"description"`:
								assert.Equal(t, `string_val:"Robot 1 Sim Card"`, upd.Val.String())
							case `name:"imsi"`:
								assert.Equal(t, `uint_val:1234011`, upd.Val.String())
							case `name:"iccid"`:
								assert.Equal(t, `uint_val:123401`, upd.Val.String())
							default:
								t.Logf("unhandled v2 enterprises acme site device 2 update %s %s %s %s %s", path0, path1, path2, path3, path4)
							}
						case `name:"sim-card" key:{key:"sim-id" value:"sim-2"}`:
							switch path4 := strings.ReplaceAll(upd.Path.Elem[4].String(), "  ", " "); path4 {
							case `name:"sim-id"`:
								assert.Equal(t, `string_val:"sim-2"`, upd.Val.String())
							case `name:"display-name"`:
								assert.Equal(t, `string_val:"Robot 2 Sim"`, upd.Val.String())
							case `name:"description"`:
								assert.Equal(t, `string_val:"Robot 2 Sim Card"`, upd.Val.String())
							case `name:"imsi"`:
								assert.Equal(t, `uint_val:1234021`, upd.Val.String())
							case `name:"iccid"`:
								assert.Equal(t, `uint_val:123402`, upd.Val.String())
							default:
								t.Logf("unhandled v2 enterprises acme site device 2 update %s %s %s %s %s", path0, path1, path2, path3, path4)
							}

						case `name:"device-group" key:{key:"dg-id" value:"acme-chicago-default"}`:
							switch path4 := strings.ReplaceAll(upd.Path.Elem[4].String(), "  ", " "); path4 {
							case `name:"dg-id"`:
								assert.Equal(t, `string_val:"acme-chicago-default"`, upd.Val.String())
							case `name:"display-name"`:
								assert.Equal(t, `string_val:"ACME Chicago Inventory"`, upd.Val.String())
							case `name:"ip-domain"`:
								assert.Equal(t, `string_val:"acme-chicago"`, upd.Val.String())
							default:
								t.Logf("unhandled v2 enterprises acme site device-group 1 update %s %s %s %s %s", path0, path1, path2, path3, path4)
							}
						case `name:"device-group" key:{key:"dg-id" value:"acme-chicago-robots"}`:
							switch path4 := strings.ReplaceAll(upd.Path.Elem[4].String(), "  ", " "); path4 {
							case `name:"dg-id"`:
								assert.Equal(t, `string_val:"acme-chicago-robots"`, upd.Val.String())
							case `name:"display-name"`:
								assert.Equal(t, `string_val:"ACME Robots"`, upd.Val.String())
							case `name:"ip-domain"`:
								assert.Equal(t, `string_val:"acme-chicago"`, upd.Val.String())
							case `name:"mbr"`:
								switch path5 := strings.ReplaceAll(upd.Path.Elem[5].String(), "  ", " "); path5 {
								case `name:"uplink"`:
									assert.Equal(t, `uint_val:5000000`, upd.Val.String())
								case `name:"downlink"`:
									assert.Equal(t, `uint_val:1000000`, upd.Val.String())
								case `name:"traffic-class"`:
									assert.Equal(t, `string_val:"class-1"`, upd.Val.String())
								default:
									t.Logf("unhandled v2 enterprises acme site device-group 1 mbr update %s %s %s %s %s %s", path0, path1, path2, path3, path4, path5)
								}
							case `name:"device" key:{key:"device-id" value:"robot-1"}`:
								switch path5 := strings.ReplaceAll(upd.Path.Elem[5].String(), "  ", " "); path5 {
								case `name:"device-id"`:
									assert.Equal(t, `string_val:"robot-1"`, upd.Val.String())
								case `name:"enable"`:
									assert.Equal(t, `bool_val:true`, upd.Val.String())
								default:
									t.Logf("unhandled v2 enterprises acme site device-group 1 device 1 update %s %s %s %s %s %s", path0, path1, path2, path3, path4, path5)
								}
							case `name:"device" key:{key:"device-id" value:"robot-2"}`:
								switch path5 := strings.ReplaceAll(upd.Path.Elem[5].String(), "  ", " "); path5 {
								case `name:"device-id"`:
									assert.Equal(t, `string_val:"robot-2"`, upd.Val.String())
								case `name:"enable"`:
									assert.Equal(t, `bool_val:true`, upd.Val.String())
								default:
									t.Logf("unhandled v2 enterprises acme site device-group 1 device 2 update %s %s %s %s %s %s", path0, path1, path2, path3, path4, path5)
								}
							default:
								t.Logf("unhandled v2 enterprises acme site device-group 1 update %s %s %s %s %s", path0, path1, path2, path3, path4)
							}
						case `name:"ip-domain" key:{key:"ip-id" value:"acme-chicago"}`:
							switch path4 := strings.ReplaceAll(upd.Path.Elem[4].String(), "  ", " "); path4 {
							case `name:"ip-id"`:
								assert.Equal(t, `string_val:"acme-chicago"`, upd.Val.String())
							case `name:"description"`:
								assert.Equal(t, `string_val:"Chicago IP Domain"`, upd.Val.String())
							case `name:"display-name"`:
								assert.Equal(t, `string_val:"Chicago"`, upd.Val.String())
							case `name:"dns-primary"`:
								assert.Equal(t, `string_val:"8.8.8.4"`, upd.Val.String())
							case `name:"dns-secondary"`:
								assert.Equal(t, `string_val:"8.8.8.8"`, upd.Val.String())
							case `name:"dnn"`:
								assert.Equal(t, `string_val:"dnnacme"`, upd.Val.String())
							case `name:"mtu"`:
								assert.Equal(t, `uint_val:12690`, upd.Val.String())
							case `name:"admin-status"`:
								assert.Equal(t, `string_val:"DISABLE"`, upd.Val.String())
							case `name:"subnet"`:
								assert.Equal(t, `string_val:"163.25.44.0/31"`, upd.Val.String())
							default:
								t.Logf("unhandled v2 enterprises acme site ipdomain 1 update %s %s %s %s %s", path0, path1, path2, path3, path4)
							}

						case `name:"small-cell" key:{key:"small-cell-id" value:"cell1"}`:
							switch path4 := strings.ReplaceAll(upd.Path.Elem[4].String(), "  ", " "); path4 {
							case `name:"small-cell-id"`:
								assert.Equal(t, `string_val:"cell1"`, upd.Val.String())
							case `name:"address"`:
								assert.Equal(t, `string_val:"ap2.chicago.acme.com"`, upd.Val.String())
							case `name:"display-name"`:
								assert.Equal(t, `string_val:"cell number one"`, upd.Val.String())
							case `name:"tac"`:
								assert.Equal(t, `string_val:"8002"`, upd.Val.String())
							case `name:"enable"`:
								assert.Equal(t, `bool_val:true`, upd.Val.String())

							default:
								t.Logf("unhandled v2 enterprises acme site smallcell 1 update %s %s %s %s %s", path0, path1, path2, path3, path4)
							}

						case `name:"upf" key:{key:"upf-id" value:"acme-chicago-pool-entry2"}`:
							switch path4 := strings.ReplaceAll(upd.Path.Elem[4].String(), "  ", " "); path4 {
							case `name:"upf-id"`:
								assert.Equal(t, `string_val:"acme-chicago-pool-entry2"`, upd.Val.String())
							case `name:"address"`:
								assert.Equal(t, `string_val:"entry2.upfpool.chicago.acme.com"`, upd.Val.String())
							case `name:"display-name"`:
								assert.Equal(t, `string_val:"Chicago Pool 2"`, upd.Val.String())
							case `name:"description"`:
								assert.Equal(t, `string_val:"Chicago UPF Pool - Entry 2"`, upd.Val.String())
							case `name:"port"`:
								assert.Equal(t, `uint_val:6161`, upd.Val.String())

							default:
								t.Logf("unhandled v2 enterprises acme site smallcell 1 update %s %s %s %s %s", path0, path1, path2, path3, path4)
							}

						case `name:"slice" key:{key:"slice-id" value:"acme-chicago-robots"}`:
							switch path4 := strings.ReplaceAll(upd.Path.Elem[4].String(), "  ", " "); path4 {
							case `name:"slice-id"`:
								assert.Equal(t, `string_val:"acme-chicago-robots"`, upd.Val.String())
							case `name:"display-name"`:
								assert.Equal(t, `string_val:"Chicago Robots Slice"`, upd.Val.String())
							case `name:"description"`:
								assert.Equal(t, `string_val:"Chicago Robots"`, upd.Val.String())
							case `name:"default-behavior"`:
								assert.Equal(t, `string_val:"DENY-ALL"`, upd.Val.String())
							case `name:"upf"`:
								assert.Equal(t, `string_val:"acme-chicago-pool-entry1"`, upd.Val.String())
							case `name:"sd"`:
								assert.Equal(t, `uint_val:2973238`, upd.Val.String())
							case `name:"sst"`:
								assert.Equal(t, `uint_val:79`, upd.Val.String())
							case `name:"mbr"`:
								switch path5 := strings.ReplaceAll(upd.Path.Elem[5].String(), "  ", " "); path5 {
								case `name:"downlink"`:
									assert.Equal(t, `uint_val:5000000`, upd.Val.String())
								case `name:"downlink-burst-size"`:
									assert.Equal(t, `uint_val:600000`, upd.Val.String())
								default:
									t.Logf("unhandled v2 enterprises acme site vcs slice mbr update %s %s %s %s %s %s", path0, path1, path2, path3, path4, path5)
								}

							case `name:"device-group" key:{key:"device-group" value:"acme-chicago-robots"}`:
								switch path5 := strings.ReplaceAll(upd.Path.Elem[5].String(), "  ", " "); path5 {
								case `name:"device-group"`:
									assert.Equal(t, `string_val:"acme-chicago-robots"`, upd.Val.String())
								case `name:"enable"`:
									assert.Equal(t, `bool_val:true`, upd.Val.String())
								default:
									t.Logf("unhandled v2 enterprises acme site slice devicegroup %s %s %s %s %s %s", path0, path1, path2, path3, path4, path5)
								}

							case `name:"filter" key:{key:"application" value:"acme-dataacquisition"}`:
								switch path5 := strings.ReplaceAll(upd.Path.Elem[5].String(), "  ", " "); path5 {
								case `name:"application"`:
									assert.Equal(t, `string_val:"acme-dataacquisition"`, upd.Val.String())
								case `name:"allow"`:
									assert.Equal(t, `bool_val:false`, upd.Val.String())
								default:
									t.Logf("unhandled v2 enterprises acme site slice devicegroup %s %s %s %s %s %s", path0, path1, path2, path3, path4, path5)
								}

							default:
								t.Logf("unhandled v2 enterprises acme site vcs 1 update %s %s %s %s %s", path0, path1, path2, path3, path4)
							}

						case `name:"priority-traffic-rule" key:{key:"ptr-id" value:"ptr-1"}`:
							switch path4 := strings.ReplaceAll(upd.Path.Elem[4].String(), "  ", " "); path4 {
							case `name:"ptr-id"`:
								assert.Equal(t, `string_val:"ptr-1"`, upd.Val.String())
							case `name:"display-name"`:
								assert.Equal(t, `string_val:"Priority Traffic Rule 1"`, upd.Val.String())
							case `name:"description"`:
								assert.Equal(t, `string_val:"Rule for priority traffic for robot-1 on da endpoint in acme-dataacquisition"`, upd.Val.String())
							case `name:"device"`:
								assert.Equal(t, `string_val:"robot-1"`, upd.Val.String())
							case `name:"application"`:
								assert.Equal(t, `string_val:"acme-dataacquisition"`, upd.Val.String())
							case `name:"endpoint"`:
								assert.Equal(t, `string_val:"da"`, upd.Val.String())
							case `name:"traffic-class"`:
								assert.Equal(t, `string_val:"class-1"`, upd.Val.String())
							case `name:"mbr"`:
								switch path5 := strings.ReplaceAll(upd.Path.Elem[5].String(), "  ", " "); path5 {
								case `name:"uplink"`:
									assert.Equal(t, `uint_val:1000000`, upd.Val.String())
								case `name:"downlink"`:
									assert.Equal(t, `uint_val:2000000`, upd.Val.String())
								default:
									t.Logf("unhandled v2 enterprises acme site ptr gbr update %s %s %s %s %s %s", path0, path1, path2, path3, path4, path5)
								}
							case `name:"gbr"`:
								switch path5 := strings.ReplaceAll(upd.Path.Elem[5].String(), "  ", " "); path5 {
								case `name:"uplink"`:
									assert.Equal(t, `uint_val:3000000`, upd.Val.String())
								case `name:"downlink"`:
									assert.Equal(t, `uint_val:4000000`, upd.Val.String())
								default:
									t.Logf("unhandled v2 enterprises acme site ptr gbr update %s %s %s %s %s %s", path0, path1, path2, path3, path4, path5)
								}

							default:
								t.Logf("unhandled v2 enterprises acme site ptr 1 update %s %s %s %s %s", path0, path1, path2, path3, path4)
							}

						default:
							t.Logf("unhandled v2 enterprises acme site update %s %s %s %s", path0, path1, path2, path3)
						}
					case `name:"template" key:{key:"tp-id" value:"template-1"}`:
						switch path3 := strings.ReplaceAll(upd.Path.Elem[3].String(), "  ", " "); path3 {
						case `name:"tp-id"`:
							assert.Equal(t, `string_val:"template-1"`, upd.Val.String())
						case `name:"description"`:
							assert.Equal(t, `string_val:"Slice Template 1"`, upd.Val.String())
						case `name:"display-name"`:
							assert.Equal(t, `string_val:"Template 1"`, upd.Val.String())
						case `name:"default-behavior"`:
							assert.Equal(t, `string_val:"DENY-ALL"`, upd.Val.String())
						case `name:"sd"`:
							assert.Equal(t, `uint_val:10886763`, upd.Val.String())
						case `name:"sst"`:
							assert.Equal(t, `uint_val:158`, upd.Val.String())
						case `name:"mbr"`:
							switch path4 := strings.ReplaceAll(upd.Path.Elem[4].String(), "  ", " "); path4 {
							case `name:"downlink"`:
								assert.Equal(t, `uint_val:5000000`, upd.Val.String())
							case `name:"downlink-burst-size"`:
								assert.Equal(t, `uint_val:600000`, upd.Val.String())
							case `name:"uplink"`:
								assert.Equal(t, `uint_val:10000000`, upd.Val.String())
							case `name:"uplink-burst-size"`:
								assert.Equal(t, `uint_val:600000`, upd.Val.String())
							default:
								t.Logf("unhandled v2 enterprises acme app endpoint update %s %s %s %s %s", path0, path1, path2, path3, path4)
							}

						default:
							t.Logf("unhandled v2 enterprises acme app update %s %s %s %s", path0, path1, path2, path3)
						}
					case `name:"traffic-class" key:{key:"tc-id" value:"class-1"}`:
						switch path3 := strings.ReplaceAll(upd.Path.Elem[3].String(), "  ", " "); path3 {
						case `name:"tc-id"`:
							assert.Equal(t, `string_val:"class-1"`, upd.Val.String())
						case `name:"description"`:
							assert.Equal(t, `string_val:"High Priority TC"`, upd.Val.String())
						case `name:"display-name"`:
							assert.Equal(t, `string_val:"Class 1"`, upd.Val.String())
						case `name:"arp"`:
							assert.Equal(t, `uint_val:1`, upd.Val.String())
						case `name:"pelr"`:
							assert.Equal(t, `int_val:10`, upd.Val.String())
						case `name:"pdb"`:
							assert.Equal(t, `uint_val:100`, upd.Val.String())
						case `name:"qci"`:
							assert.Equal(t, `uint_val:10`, upd.Val.String())

						default:
							t.Logf("unhandled v2 enterprises acme app update %s %s %s %s", path0, path1, path2, path3)
						}

					default:
						t.Logf("unhandled v2 enterprises acme update %s %s %s", path0, path1, path2)
					}

				default:
					t.Logf("unhandled v2 enterprises update %s %s", path0, path1)
				}
			}
		default:
			t.Logf("unhandled target %s", tgt)
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
