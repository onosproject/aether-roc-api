// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"reflect"
)

// ToFloat32Ptr -
func ToFloat32Ptr(value *reflect.Value) (*float32, error) {
	switch val := value.Interface().(type) {
	case float32:
		floatVal := val
		return &floatVal, nil
	case float64:
		floatVal := float32(val)
		return &floatVal, nil
	default:
		return nil, errors.NewNotSupported("unhandled conversion %s to *float32", value.Kind().String())
	}
}

// ToFloat32 -
func ToFloat32(value *reflect.Value) (float32, error) {
	switch val := value.Interface().(type) {
	case float32:
		floatVal := val
		return floatVal, nil
	case float64:
		floatVal := float32(val)
		return floatVal, nil
	default:
		return 0, errors.NewNotSupported("unhandled conversion %s to float32", value.Kind().String())
	}
}

// ToFloat64Ptr -
func ToFloat64Ptr(value *reflect.Value) (*float64, error) {
	switch val := value.Interface().(type) {
	case float32:
		floatVal := float64(val)
		return &floatVal, nil
	case float64:
		floatVal := val
		return &floatVal, nil
	default:
		return nil, errors.NewNotSupported("unhandled conversion %s to *float64", value.Kind().String())
	}
}

// ToFloat64 -
func ToFloat64(value *reflect.Value) (float64, error) {
	switch val := value.Interface().(type) {
	case float32:
		floatVal := float64(val)
		return floatVal, nil
	case float64:
		floatVal := val
		return floatVal, nil
	default:
		return 0, errors.NewNotSupported("unhandled conversion %s to float64", value.Kind().String())
	}
}
