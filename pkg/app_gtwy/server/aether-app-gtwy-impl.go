// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
//

package server

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	externalRef0 "github.com/onosproject/aether-roc-api/pkg/aether_2_0_0/types"
	"github.com/onosproject/aether-roc-api/pkg/app_gtwy/types"
	"github.com/onosproject/aether-roc-api/pkg/southbound"
	"github.com/onosproject/aether-roc-api/pkg/utils"
	externalRef1 "github.com/onosproject/config-models/modelplugin/aether-2.0.0/aether_2_0_0"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/prometheus/common/model"
	"net/http"
	"strconv"
	"time"
)

var log = logging.GetLogger("app_gtwy")

const prometheusQuery = "subscribers_info"
const prometheusMetric = "mobile_ip"

// AppGtwy -
type AppGtwy struct {
	GnmiClient      southbound.GnmiClient
	GnmiTimeout     time.Duration
	AnalyticsClient AnalyticsClient
}

// GetDevices /appgtwy/v1/{target}/enterprises/{enterprise-id}/sites/{site-id}/devices
func (i *AppGtwy) GetDevices(ctx echo.Context, target externalRef0.Target, enterpriseID string, siteID string) error {

	var listDeviceResponse []types.AppGtwyDevice

	gnmiCtx, cancel := utils.NewGnmiContext(ctx, i.GnmiTimeout)
	defer cancel()

	site, err := i.gnmiGetEnterprisesEnterpriseSite(gnmiCtx, "/aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}", target, enterpriseID, siteID)
	if err != nil {
		return utils.ConvertGrpcError(err)
	}

	si, err := i.AnalyticsClient.Query(prometheusQuery)
	if err != nil {
		log.Errorf("error querying analytics %s", err)
		return ctx.JSON(http.StatusBadRequest, err)
	}

	simMap := mapSim(site)
	deviceGroupMap := mapDeviceToDeviceGroups(site)
	imsiMap := mapIMSI(si)

	for _, d := range *site.Device {
		s := simMap[*d.SimCard]
		newDev := types.AppGtwyDevice{
			DeviceId:     d.DeviceId,
			DisplayName:  d.DisplayName,
			Description:  d.Description,
			Imei:         d.Imei,
			SimIccid:     s.Iccid,
			DeviceGroups: deviceGroupMap[d.DeviceId],
		}

		imsi := strconv.Itoa(int(*s.Imsi))
		if _, ok := imsiMap[imsi]; ok {
			ip := string(imsiMap[imsi].Metric[prometheusMetric])
			newDev.Ip = &ip
			attached := imsiMap[imsi].Value.String()
			newDev.Attached = &attached
		}

		listDeviceResponse = append(listDeviceResponse, newDev)
	}

	log.Infof("GetAppGtwyDevices: returned %d devices", len(listDeviceResponse))
	return ctx.JSON(http.StatusOK, listDeviceResponse)
}

// GetDevice /appgtwy/v1/{target}/enterprises/{enterprise-id}/sites/{site-id}/devices/{device-id}
func (i *AppGtwy) GetDevice(ctx echo.Context, target externalRef0.Target, enterpriseID string, siteID string, deviceID string) error {

	gnmiCtx, cancel := utils.NewGnmiContext(ctx, i.GnmiTimeout)
	defer cancel()

	site, err := i.gnmiGetEnterprisesEnterpriseSite(gnmiCtx, "/aether/v2.0.0/{target}/enterprises/enterprise/{enterprise-id}/site/{site-id}", target, enterpriseID, siteID)
	if err != nil {
		return utils.ConvertGrpcError(err)
	}

	si, err := i.AnalyticsClient.Query(prometheusQuery)
	if err != nil {
		log.Errorf("error querying analytics %s", err)
		return ctx.JSON(http.StatusBadRequest, err)
	}

	simMap := mapSim(site)
	deviceGroupMap := mapDeviceToDeviceGroups(site)
	imsiMap := mapIMSI(si)

	var device types.AppGtwyDevice
	for _, d := range *site.Device {
		if d.DeviceId != deviceID {
			continue
		}
		s := simMap[*d.SimCard]
		device = types.AppGtwyDevice{
			DeviceId:     deviceID,
			DisplayName:  d.DisplayName,
			Description:  d.Description,
			Imei:         d.Imei,
			SimIccid:     s.Iccid,
			DeviceGroups: deviceGroupMap[deviceID],
		}

		imsi := strconv.Itoa(int(*s.Imsi))
		if _, ok := imsiMap[imsi]; ok {
			ip := string(imsiMap[imsi].Metric[prometheusMetric])
			device.Ip = &ip
			attached := imsiMap[imsi].Value.String()
			device.Attached = &attached
		}
		return ctx.JSON(http.StatusOK, device)
	}

	return ctx.JSON(http.StatusNotFound, "")
}

// gnmiGetEnterprisesEnterpriseSite returns an instance of Enterprises_Enterprise_Site.
func (i *AppGtwy) gnmiGetEnterprisesEnterpriseSite(ctx context.Context,
	openAPIPath string, target externalRef0.Target, args ...string) (*externalRef0.EnterprisesEnterpriseSite, error) {

	gnmiGet, err := utils.NewGnmiGetRequest(openAPIPath, string(target), args...)
	if err != nil {
		return nil, err
	}
	log.Infof("gnmiGetRequest %s", gnmiGet.String())
	gnmiVal, err := utils.GetResponseUpdate(i.GnmiClient.Get(ctx, gnmiGet))
	if err != nil {
		return nil, err
	}
	if gnmiVal == nil {
		return nil, nil
	}
	gnmiJSONVal, ok := gnmiVal.Value.(*gnmi.TypedValue_JsonVal)
	if !ok {
		return nil, fmt.Errorf("unexpected type of reply from server %v", gnmiVal.Value)
	}

	log.Debugf("gNMI Json %s", string(gnmiJSONVal.JsonVal))
	var gnmiResponse externalRef1.Device
	if err = externalRef1.Unmarshal(gnmiJSONVal.JsonVal, &gnmiResponse); err != nil {
		return nil, fmt.Errorf("error unmarshalling gnmiResponse %v", err)
	}
	mpd := ModelPluginDevice{
		device: gnmiResponse,
	}

	return mpd.toEnterprisesEnterpriseSite(args...)
}

func mapSim(site *externalRef0.EnterprisesEnterpriseSite) map[string]*externalRef0.EnterprisesEnterpriseSiteSimCard {
	simMap := make(map[string]*externalRef0.EnterprisesEnterpriseSiteSimCard)
	for _, s := range *site.SimCard {
		simMap[s.SimId] = &externalRef0.EnterprisesEnterpriseSiteSimCard{
			SimId:       s.SimId,
			Description: s.Description,
			DisplayName: s.DisplayName,
			Iccid:       s.Iccid,
			Imsi:        s.Imsi,
		}
	}
	return simMap
}

func mapDeviceToDeviceGroups(site *externalRef0.EnterprisesEnterpriseSite) map[string]*[]string {
	deviceGroupMap := make(map[string]*[]string)
	var dgs []string

	for _, dg := range *site.DeviceGroup {
		if len(*dg.Device) > 0 {
			dgs = append(dgs, dg.DeviceGroupId)
			for _, d := range *dg.Device {
				deviceGroupMap[d.DeviceId] = &dgs
			}
		}
	}

	return deviceGroupMap
}

func mapIMSI(si model.Value) map[string]*model.Sample {
	v := si.(model.Vector)
	imsiMap := make(map[string]*model.Sample)
	for _, val := range v {
		imsiMap[string(val.Metric["imsi"])] = val
	}
	return imsiMap
}
