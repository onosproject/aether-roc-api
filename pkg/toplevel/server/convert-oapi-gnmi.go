// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package server

import (
	"fmt"
	externalRef0Svr "github.com/onosproject/aether-roc-api/pkg/aether_2_0_0/server"
	externalRef0 "github.com/onosproject/aether-roc-api/pkg/aether_2_0_0/types"
	"github.com/onosproject/aether-roc-api/pkg/toplevel/types"
	"github.com/openconfig/gnmi/proto/gnmi"
	"regexp"
)

var re = regexp.MustCompile(`[0-9a-z\-\._]+`)

func encodeToGnmiPatchBody(jsonObj *types.PatchBody) ([]*gnmi.Update, []*gnmi.Path, *string, *string, string, error) {
	updates := make([]*gnmi.Update, 0)
	deletes := make([]*gnmi.Path, 0)
	var ext101Version *string
	var ext102Type *string

	if jsonObj.Extensions != nil {
		ext101Version = jsonObj.Extensions.ModelVersion101
		ext102Type = jsonObj.Extensions.ModelType102
	}

	if !re.MatchString(jsonObj.DefaultTarget) {
		return nil, nil, ext101Version, ext102Type, jsonObj.DefaultTarget, fmt.Errorf("default-target cannot be blank")
	}

	gnmiUpdates, err := encodeToGnmiElements(jsonObj.Updates, jsonObj.DefaultTarget)
	if err != nil {
		return nil, nil, ext101Version, ext102Type, jsonObj.DefaultTarget, fmt.Errorf("encodeToGnmiElements() %s", err.Error())
	}
	updates = append(updates, gnmiUpdates...)

	gnmiDeletes, err := encodeToGnmiElements(jsonObj.Deletes, jsonObj.DefaultTarget)
	if err != nil {
		return nil, nil, ext101Version, ext102Type, jsonObj.DefaultTarget, fmt.Errorf("encodeToGnmiElements() %s", err.Error())
	}
	for _, gd := range gnmiDeletes {
		deletes = append(deletes, gd.Path)
	}

	return updates, deletes, ext101Version, ext102Type, jsonObj.DefaultTarget, nil
}

func encodeToGnmiElements(elements *types.Elements, target string) ([]*gnmi.Update, error) {
	if elements == nil {
		return nil, nil
	}
	updates := make([]*gnmi.Update, 0)

	if elements.AccessProfile200 != nil {
		accessProfileUpdates, err := externalRef0Svr.EncodeToGnmiAccessProfile(
			elements.AccessProfile200, false, externalRef0.Target(target),
			"/access-profile")
		if err != nil {
			return nil, fmt.Errorf("EncodeToGnmiAccessProfile() %s", err)
		}
		updates = append(updates, accessProfileUpdates...)
	}

	if elements.ApnProfile200 != nil {
		apnProfileUpdates, err := externalRef0Svr.EncodeToGnmiApnProfile(
			elements.ApnProfile200, false, externalRef0.Target(target),
			"/apn-profile")
		if err != nil {
			return nil, fmt.Errorf("EncodeToGnmiApnProfile() %s", err)
		}
		updates = append(updates, apnProfileUpdates...)
	}

	if elements.ConnectivityService200 != nil {
		connectivityServiceUpdates, err := externalRef0Svr.EncodeToGnmiConnectivityService(
			elements.ConnectivityService200, false, externalRef0.Target(target),
			"/connectivity-service")
		if err != nil {
			return nil, fmt.Errorf("EncodeToGnmiConnectivityService() %s", err)
		}
		updates = append(updates, connectivityServiceUpdates...)
	}

	if elements.Enterprise200 != nil {
		enterpriseUpdates, err := externalRef0Svr.EncodeToGnmiEnterprise(
			elements.Enterprise200, false, externalRef0.Target(target),
			"/enterprise")
		if err != nil {
			return nil, fmt.Errorf("EncodeToGnmiEnterprise() %s", err)
		}
		updates = append(updates, enterpriseUpdates...)
	}

	if elements.QosProfile200 != nil {
		qosProfileUpdates, err := externalRef0Svr.EncodeToGnmiQosProfile(
			elements.QosProfile200, false, externalRef0.Target(target),
			"/qos-profile")
		if err != nil {
			return nil, fmt.Errorf("EncodeToGnmiQosProfile() %s", err)
		}
		updates = append(updates, qosProfileUpdates...)
	}

	if elements.SecurityProfile200 != nil {
		securityProfileUpdates, err := externalRef0Svr.EncodeToGnmiSecurityProfile(
			elements.SecurityProfile200, false, externalRef0.Target(target),
			"/security-profile")
		if err != nil {
			return nil, fmt.Errorf("EncodeToGnmiSecurityProfile() %s", err)
		}
		updates = append(updates, securityProfileUpdates...)
	}

	if elements.Subscriber200 != nil {
		subscriberUpdates, err := externalRef0Svr.EncodeToGnmiSubscriber(
			elements.Subscriber200, false, externalRef0.Target(target),
			"/subscriber")
		if err != nil {
			return nil, fmt.Errorf("EncodeToGnmiSubscriber() %s", err)
		}
		updates = append(updates, subscriberUpdates...)
	}

	if elements.UpProfile200 != nil {
		upProfileUpdates, err := externalRef0Svr.EncodeToGnmiUpProfile(
			elements.UpProfile200, false, externalRef0.Target(target),
			"/up-profile")
		if err != nil {
			return nil, fmt.Errorf("EncodeToGnmiUpProfile() %s", err)
		}
		updates = append(updates, upProfileUpdates...)
	}

	//if elements.AccessProfile210 != nil {
	//	accessProfileUpdates, err := externalRef1Svr.EncodeToGnmiAccessProfile(
	//		elements.AccessProfile210, false, externalRef1.Target(target),
	//		"/access-profile")
	//	if err != nil {
	//		return nil, fmt.Errorf("EncodeToGnmiAccessProfile() %s", err)
	//	}
	//	updates = append(updates, accessProfileUpdates...)
	//}
	//
	//if elements.ApnProfile210 != nil {
	//	apnProfileUpdates, err := externalRef1Svr.EncodeToGnmiApnProfile(
	//		elements.ApnProfile210, false, externalRef1.Target(target),
	//		"/apn-profile")
	//	if err != nil {
	//		return nil, fmt.Errorf("EncodeToGnmiApnProfile() %s", err)
	//	}
	//	updates = append(updates, apnProfileUpdates...)
	//}
	//
	//if elements.ConnectivityService210 != nil {
	//	connectivityServiceUpdates, err := externalRef1Svr.EncodeToGnmiConnectivityService(
	//		elements.ConnectivityService210, false, externalRef1.Target(target),
	//		"/connectivity-service")
	//	if err != nil {
	//		return nil, fmt.Errorf("EncodeToGnmiConnectivityService() %s", err)
	//	}
	//	updates = append(updates, connectivityServiceUpdates...)
	//}
	//
	//if elements.Enterprise210 != nil {
	//	enterpriseUpdates, err := externalRef1Svr.EncodeToGnmiEnterprise(
	//		elements.Enterprise210, false, externalRef1.Target(target),
	//		"/enterprise")
	//	if err != nil {
	//		return nil, fmt.Errorf("EncodeToGnmiEnterprise() %s", err)
	//	}
	//	updates = append(updates, enterpriseUpdates...)
	//}
	//
	//if elements.QosProfile210 != nil {
	//	qosProfileUpdates, err := externalRef1Svr.EncodeToGnmiQosProfile(
	//		elements.QosProfile210, false, externalRef1.Target(target),
	//		"/qos-profile")
	//	if err != nil {
	//		return nil, fmt.Errorf("EncodeToGnmiQosProfile() %s", err)
	//	}
	//	updates = append(updates, qosProfileUpdates...)
	//}
	//
	//if elements.SecurityProfile210 != nil {
	//	securityProfileUpdates, err := externalRef1Svr.EncodeToGnmiSecurityProfile(
	//		elements.SecurityProfile210, false, externalRef1.Target(target),
	//		"/security-profile")
	//	if err != nil {
	//		return nil, fmt.Errorf("EncodeToGnmiSecurityProfile() %s", err)
	//	}
	//	updates = append(updates, securityProfileUpdates...)
	//}
	//
	//if elements.ServiceGroup210 != nil {
	//	serviceGroupUpdates, err := externalRef1Svr.EncodeToGnmiServiceGroup(
	//		elements.ServiceGroup210, false, externalRef1.Target(target),
	//		"/service-group")
	//	if err != nil {
	//		return nil, fmt.Errorf("EncodeToGnmiServiceGroup() %s", err)
	//	}
	//	updates = append(updates, serviceGroupUpdates...)
	//}
	//
	//if elements.ServicePolicy210 != nil {
	//	servicePolicyUpdates, err := externalRef1Svr.EncodeToGnmiServicePolicy(
	//		elements.ServicePolicy210, false, externalRef1.Target(target),
	//		"/service-policy")
	//	if err != nil {
	//		return nil, fmt.Errorf("EncodeToGnmiServicePolicy() %s", err)
	//	}
	//	updates = append(updates, servicePolicyUpdates...)
	//}
	//
	//if elements.ServiceRule210 != nil {
	//	serviceRuleUpdates, err := externalRef1Svr.EncodeToGnmiServiceRule(
	//		elements.ServiceRule210, false, externalRef1.Target(target),
	//		"/service-role")
	//	if err != nil {
	//		return nil, fmt.Errorf("EncodeToGnmiServiceRule() %s", err)
	//	}
	//	updates = append(updates, serviceRuleUpdates...)
	//}
	//
	//if elements.Subscriber210 != nil {
	//	subscriberUpdates, err := externalRef1Svr.EncodeToGnmiSubscriber(
	//		elements.Subscriber210, false, externalRef1.Target(target),
	//		"/subscriber")
	//	if err != nil {
	//		return nil, fmt.Errorf("EncodeToGnmiSubscriber() %s", err)
	//	}
	//	updates = append(updates, subscriberUpdates...)
	//}
	//
	//if elements.UpProfile210 != nil {
	//	upProfileUpdates, err := externalRef1Svr.EncodeToGnmiUpProfile(
	//		elements.UpProfile210, false, externalRef1.Target(target),
	//		"/up-profile")
	//	if err != nil {
	//		return nil, fmt.Errorf("EncodeToGnmiUpProfile() %s", err)
	//	}
	//	updates = append(updates, upProfileUpdates...)
	//}

	return updates, nil
}
