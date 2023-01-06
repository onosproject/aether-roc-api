// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
//

package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	externalRef0Svr "github.com/onosproject/aether-roc-api/pkg/aether_2_0_0/server"
	externalRef0 "github.com/onosproject/aether-roc-api/pkg/aether_2_0_0/types"
	externalRef1Svr "github.com/onosproject/aether-roc-api/pkg/aether_2_1_0/server"
	externalRef1 "github.com/onosproject/aether-roc-api/pkg/aether_2_1_0/types"
	"github.com/onosproject/aether-roc-api/pkg/toplevel/types"
	"github.com/openconfig/gnmi/proto/gnmi"
	"net/http"
	"regexp"
)

var re = regexp.MustCompile(`[0-9a-z\-\._]+`)

const undefined = "undefined"

// GnmiPatchBody contains all the info required to build a gNMI requests
type GnmiPatchBody struct {
	Updates       []*gnmi.Update
	Deletes       []*gnmi.Path
	DefaultTarget string
	Ext100Name    *string
	Ext101Version *string
	Ext102Type    *string
	Ext110Info    *struct {
		ID    *string `json:"ID,omitempty"`
		Index *int    `json:"index,omitempty"`
	}
	Ext111Strategy *int
}

func encodeToGnmiPatchBody(jsonObj *types.PatchBody) (*GnmiPatchBody, error) {
	updates := make([]*gnmi.Update, 0)
	deletes := make([]*gnmi.Path, 0)

	pb := &GnmiPatchBody{}

	if jsonObj.Extensions != nil {
		pb.Ext100Name = jsonObj.Extensions.ChangeName100
		pb.Ext101Version = jsonObj.Extensions.ModelVersion101
		pb.Ext102Type = jsonObj.Extensions.ModelType102
		//pb.Ext110Info = jsonObj.Extensions.TransactionInfo110
		pb.Ext111Strategy = jsonObj.Extensions.TransactionStrategy111
	}

	if !re.MatchString(jsonObj.DefaultTarget) {
		return nil, fmt.Errorf("default-target cannot be blank")
	}
	pb.DefaultTarget = jsonObj.DefaultTarget

	gnmiUpdates, err := encodeToGnmiElements(jsonObj.Updates, jsonObj.DefaultTarget, false)
	if err != nil {
		return nil, err
	}
	updates = append(updates, gnmiUpdates...)
	pb.Updates = updates

	gnmiDeletes, err := encodeToGnmiElements(jsonObj.Deletes, jsonObj.DefaultTarget, true)
	if err != nil {
		return nil, err
	}
	for _, gd := range gnmiDeletes {
		deletes = append(deletes, gd.Path)
	}
	pb.Deletes = deletes

	return pb, nil
}

func encodeToGnmiElements(elements *types.Elements, target string, forDelete bool) ([]*gnmi.Update, error) {
	if elements == nil {
		return nil, nil
	}
	updates := make([]*gnmi.Update, 0)

	// Aether 2.0.x

	if elements.ConnectivityServices200 != nil {
		connectivityServiceUpdates, err := externalRef0Svr.EncodeToGnmiConnectivityServices(
			elements.ConnectivityServices200, false, forDelete, externalRef0.Target(target),
			"/connectivity-services")
		if err != nil {
			return nil, fmt.Errorf("EncodeToGnmiConnectivityService() %s", err)
		}
		updates = append(updates, connectivityServiceUpdates...)
	}

	if elements.Enterprises200 != nil {
		for _, e := range *elements.Enterprises200.Enterprise {

			if e.EnterpriseId == undefined {
				log.Warnw("EnterpriseId is undefined", "enterprise", e)
				return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, "enterprise-id-cannot-be-undefined")
			}

			if e.Site != nil && len(*e.Site) > 0 {
				for _, s := range *e.Site {
					if s.SiteId == undefined {
						log.Warnw("SiteId is undefined", "site", s)
						return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, "site-id-cannot-be-undefined")
					}
				}
			}
		}

		enterpriseUpdates, err := externalRef0Svr.EncodeToGnmiEnterprises(
			elements.Enterprises200, false, forDelete, externalRef0.Target(target),
			"/enterprises")

		if err != nil {
			return nil, fmt.Errorf("EncodeToGnmiEnterprise() %s", err)
		}
		updates = append(updates, enterpriseUpdates...)
	}

	// Aether 2.1.x

	if elements.Application210 != nil && len(*elements.Application210) > 0 {
		for _, e := range *elements.Application210 {

			if e.ApplicationId == undefined {
				log.Warnw("ApplicationId is undefined", "application", e)
				return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, "application-id-cannot-be-undefined")
			}

			if e.Endpoint != nil && len(*e.Endpoint) > 0 {
				for _, s := range *e.Endpoint {
					if s.EndpointId == undefined {
						log.Warnw("EndpointId is undefined", "endpoint", s)
						return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, "endpoint-id-cannot-be-undefined")
					}
				}
			}
		}
		applicationUpdates, err := externalRef1Svr.EncodeToGnmiApplicationList(elements.Application210, false,
			forDelete, externalRef1.EnterpriseId(target), "/application")
		if err != nil {
			return nil, fmt.Errorf("EncodeToGnmiApplicationList() %s", err)
		}
		updates = append(updates, applicationUpdates...)
	}

	if elements.Site210 != nil && len(*elements.Site210) > 0 {
		for _, e := range *elements.Site210 {
			if e.SiteId == undefined {
				log.Warnw("SiteId is undefined", "site", e)
				return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, "site-id-cannot-be-undefined")
			}

			if e.Slice != nil && len(*e.Slice) > 0 {
				for _, s := range *e.Slice {
					if s.SliceId == undefined {
						log.Warnw("SliceId is undefined", "slice", s)
						return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, "slice-id-cannot-be-undefined")
					}
				}
			}
			if e.Device != nil && len(*e.Device) > 0 {
				for _, s := range *e.Device {
					if s.DeviceId == undefined {
						log.Warnw("DeviceId is undefined", "device", s)
						return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, "device-id-cannot-be-undefined")
					}
				}
			}
			if e.SimCard != nil && len(*e.SimCard) > 0 {
				for _, s := range *e.SimCard {
					if s.SimId == undefined {
						log.Warnw("SimId is undefined", "simcard", s)
						return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, "sim-id-cannot-be-undefined")
					}
				}
			}
			if e.DeviceGroup != nil && len(*e.DeviceGroup) > 0 {
				for _, s := range *e.DeviceGroup {
					if s.DeviceGroupId == undefined {
						log.Warnw("DeviceGroupId is undefined", "device-group", s)
						return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, "device-group-id-cannot-be-undefined")
					}
				}
			}
			if e.IpDomain != nil && len(*e.IpDomain) > 0 {
				for _, s := range *e.IpDomain {
					if s.IpDomainId == undefined {
						log.Warnw("IpDomainId is undefined", "ip-domain", s)
						return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, "ip-domain-id-cannot-be-undefined")
					}
				}
			}
			if e.Upf != nil && len(*e.Upf) > 0 {
				for _, s := range *e.Upf {
					if s.UpfId == undefined {
						log.Warnw("UpfId is undefined", "upf", s)
						return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, "upf-id-cannot-be-undefined")
					}
				}
			}
		}
		siteUpdates, err := externalRef1Svr.EncodeToGnmiSiteList(elements.Site210, false, forDelete,
			externalRef1.EnterpriseId(target), "/site")
		if err != nil {
			return nil, fmt.Errorf("EncodeToGnmiSiteList() %s", err)
		}
		updates = append(updates, siteUpdates...)
	}

	if elements.Template210 != nil && len(*elements.Template210) > 0 {
		for _, e := range *elements.Template210 {
			if e.TemplateId == undefined {
				log.Warnw("TemplateId is undefined", "template", e)
				return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, "template-id-cannot-be-undefined")
			}
		}
		templateUpdates, err := externalRef1Svr.EncodeToGnmiTemplateList(elements.Template210, false, forDelete,
			externalRef1.EnterpriseId(target), "/template")
		if err != nil {
			return nil, fmt.Errorf("EncodeToGnmiTemplateList() %s", err)
		}
		updates = append(updates, templateUpdates...)
	}

	if elements.TrafficClass210 != nil && len(*elements.TrafficClass210) > 0 {
		for _, e := range *elements.TrafficClass210 {
			if e.TrafficClassId == undefined {
				log.Warnw("TrafficClassId is undefined", "traffic-class", e)
				return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, "traffic-class-id-cannot-be-undefined")
			}
		}
		trafficClassUpdates, err := externalRef1Svr.EncodeToGnmiTrafficClassList(elements.TrafficClass210, false,
			forDelete, externalRef1.EnterpriseId(target), "/traffic-class")
		if err != nil {
			return nil, fmt.Errorf("EncodeToGnmiTrafficClassList() %s", err)
		}
		updates = append(updates, trafficClassUpdates...)
	}

	return updates, nil
}
