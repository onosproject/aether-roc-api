// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package proto

import (
	"github.com/onosproject/onos-config/pkg/utils"
	"strings"
)

const (
	// StringVal :
	StringVal = "string_val"

	// IntVal :
	IntVal = "int_val"

	// BoolVal :
	BoolVal = "bool_val"
)

// DevicePath describes the results of a get operation for a single Path
// It specifies the device, Path, and value
type DevicePath struct {
	DeviceName    string
	Path          string
	PathDataType  string
	PathDataValue string
}

// MakeProtoTarget returns a GNMI proto path for a given target
func MakeProtoTarget(target string, path string) string {
	var protoBuilder strings.Builder
	var pathElements []string

	protoBuilder.WriteString("<target: '")
	protoBuilder.WriteString(target)
	protoBuilder.WriteString("', ")

	pathElements = utils.SplitPath(path)

	for _, pathElement := range pathElements {
		protoBuilder.WriteString("elem: <name: '")
		protoBuilder.WriteString(pathElement)
		protoBuilder.WriteString("'>")
	}
	protoBuilder.WriteString(">")
	return protoBuilder.String()
}

// MakeProtoPath returns a path: element for a given target and path
func MakeProtoPath(target string, path string) string {
	var protoBuilder strings.Builder

	protoBuilder.WriteString("path: ")
	gnmiPath := MakeProtoTarget(target, path)
	protoBuilder.WriteString(gnmiPath)
	return protoBuilder.String()
}

func makeProtoValue(value string, valueType string) string {
	var protoBuilder strings.Builder

	var valueString string

	if valueType == StringVal {
		valueString = "'" + value + "'"
	} else {
		valueString = value
	}
	protoBuilder.WriteString(" val: <")
	protoBuilder.WriteString(valueType)
	protoBuilder.WriteString(":")
	protoBuilder.WriteString(valueString)
	protoBuilder.WriteString(">")
	return protoBuilder.String()
}

// MakeProtoUpdatePath returns an update: element for a target, path, and new value
func MakeProtoUpdatePath(devicePath DevicePath) string {
	var protoBuilder strings.Builder

	protoBuilder.WriteString("update: <")
	protoBuilder.WriteString(MakeProtoPath(devicePath.DeviceName, devicePath.Path))
	protoBuilder.WriteString(makeProtoValue(devicePath.PathDataValue, devicePath.PathDataType))
	protoBuilder.WriteString(">")
	return protoBuilder.String()
}

// MakeProtoDeletePath returns a delete: element for a given target and path
func MakeProtoDeletePath(target string, path string) string {
	var protoBuilder strings.Builder

	protoBuilder.WriteString("delete: ")
	protoBuilder.WriteString(MakeProtoTarget(target, path))
	return protoBuilder.String()
}
