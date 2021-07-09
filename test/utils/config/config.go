// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package config

import (
	"context"
	"fmt"
	protoutils "github.com/onosproject/aether-roc-api/test/utils/proto"
	"github.com/onosproject/helmit/pkg/helm"
	"github.com/onosproject/helmit/pkg/kubernetes"
	v1 "github.com/onosproject/helmit/pkg/kubernetes/core/v1"
	"github.com/onosproject/onos-config/pkg/utils"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/openconfig/gnmi/client"
	gclient "github.com/openconfig/gnmi/client/gnmi"
	gpb "github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/gnmi/proto/gnmi_ext"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/proto"
	"strings"

	"testing"
	"time"
)

const aetherRocUmbrella = "aether-roc-umbrella"

// MakeContext returns a new context for use in GNMI requests
func MakeContext() context.Context {
	ctx := context.Background()
	return ctx
}

// GetGNMIClientOrFail makes a GNMI client to use for requests. If creating the client fails, the test is failed.
func GetGNMIClientOrFail(t *testing.T) client.Impl {
	t.Helper()
	release := helm.Chart(aetherRocUmbrella).Release(aetherRocUmbrella)
	conn, err := connectService(release, "onos-config")
	assert.NoError(t, err)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	dest, err := GetDestination()
	if !assert.NoError(t, err) {
		t.Fail()
	}
	client, err := gclient.NewFromConn(ctx, conn, dest)
	assert.NoError(t, err)
	assert.True(t, client != nil, "Fetching device client returned nil")
	return client
}

// GetGNMIValue generates a GET request on the given client for a Path on a device
func GetGNMIValue(ctx context.Context, c client.Impl, paths []protoutils.DevicePath, encoding gpb.Encoding) ([]protoutils.DevicePath, []*gnmi_ext.Extension, error) {
	protoString := ""
	for _, devicePath := range paths {
		protoString = protoString + MakeProtoPath(devicePath.DeviceName, devicePath.Path)
	}
	getTZRequest := &gpb.GetRequest{}
	if err := proto.Unmarshal([]byte(protoString), getTZRequest); err != nil {
		fmt.Printf("unable to parse gnmi.GetRequest from %q : %v\n", protoString, err)
		return nil, nil, err
	}
	getTZRequest.Encoding = encoding
	response, err := c.(*gclient.Client).Get(ctx, getTZRequest)
	if err != nil || response == nil {
		return nil, nil, err
	}
	return convertGetResults(response)
}

// CheckGNMIValue makes sure a value has been assigned properly by querying the onos-config northbound API
func CheckGNMIValue(t *testing.T, gnmiClient client.Impl, paths []protoutils.DevicePath, expectedValue string, expectedExtensions int, failMessage string) {
	t.Helper()
	value, extensions, err := GetGNMIValue(MakeContext(), gnmiClient, paths, gpb.Encoding_PROTO)
	assert.NoError(t, err)
	assert.Equal(t, expectedExtensions, len(extensions))
	assert.Equal(t, expectedValue, value[0].PathDataValue, "%s: %s", failMessage, value)
}

// GetDevicePathWithValue creates a device path with a value to set
func GetDevicePathWithValue(device string, path string, value string, valueType string) []protoutils.DevicePath {
	devicePath := make([]protoutils.DevicePath, 1)
	devicePath[0].DeviceName = device
	devicePath[0].Path = path
	devicePath[0].PathDataValue = value
	devicePath[0].PathDataType = valueType
	return devicePath
}

// GetDestination :
func GetDestination() (client.Destination, error) {
	creds, err := getClientCredentials()
	if err != nil {
		return client.Destination{}, err
	}
	configRelease := helm.Release(aetherRocUmbrella)
	configClient := kubernetes.NewForReleaseOrDie(configRelease)

	configService, err := configClient.CoreV1().Services().Get(MakeContext(), "onos-config")
	if err != nil || configService == nil {
		return client.Destination{}, errors.NewNotFound("can't find service for onos-config")
	}

	return client.Destination{
		Addrs:   []string{configService.Ports()[0].Address(true)},
		Target:  configService.Name,
		TLS:     creds,
		Timeout: 10 * time.Second,
	}, nil
}

// MakeProtoPath returns a Path: element for a given target and Path
func MakeProtoPath(target string, path string) string {
	var protoBuilder strings.Builder

	protoBuilder.WriteString("path: ")
	gnmiPath := protoutils.MakeProtoTarget(target, path)
	protoBuilder.WriteString(gnmiPath)
	return protoBuilder.String()
}

func connectService(release *helm.HelmRelease, deploymentName string) (*grpc.ClientConn, error) {
	service, err := getService(release, deploymentName)
	if err != nil {
		return nil, err
	}
	tlsConfig, err := getClientCredentials()
	if err != nil {
		return nil, err
	}
	return grpc.Dial(service.Ports()[0].Address(true), grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)))
}

func getService(release *helm.HelmRelease, serviceName string) (*v1.Service, error) {
	releaseClient := kubernetes.NewForReleaseOrDie(release)
	service, err := releaseClient.CoreV1().Services().Get(context.Background(), serviceName)
	if err != nil {
		return nil, err
	}

	return service, nil
}

func convertGetResults(response *gpb.GetResponse) ([]protoutils.DevicePath, []*gnmi_ext.Extension, error) {
	entryCount := len(response.Notification)
	result := make([]protoutils.DevicePath, entryCount)

	for index, notification := range response.Notification {
		value := notification.Update[0].Val

		result[index].DeviceName = notification.Update[0].Path.Target
		pathString := ""

		for _, elem := range notification.Update[0].Path.Elem {
			pathString = pathString + "/" + elem.Name
		}
		result[index].Path = pathString

		result[index].PathDataType = "string_val"
		if value != nil {
			result[index].PathDataValue = utils.StrVal(value)
		} else {
			result[index].PathDataValue = ""
		}
	}

	return result, response.Extension, nil
}
