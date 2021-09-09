// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package utils

import (
	"fmt"
	"reflect"
)

// ToIntPtr -
func ToIntPtr(value *reflect.Value) (*int, error) {
	switch val := value.Interface().(type) {
	case int:
		intVal := val
		return &intVal, nil
	case uint:
		intVal := int(val)
		return &intVal, nil
	case uint16:
		intVal := int(val)
		return &intVal, nil
	case uint32:
		intVal := int(val)
		return &intVal, nil
	case uint64:
		intVal := int(val)
		return &intVal, nil
	case int64:
		intVal := int(val)
		return &intVal, nil
	default:
		return nil, fmt.Errorf("error converting %v to *int", value.Interface())
	}
}

// ToInt32Ptr -
func ToInt32Ptr(value *reflect.Value) (*int32, error) {
	switch val := value.Interface().(type) {
	case int:
		intVal := int32(val)
		return &intVal, nil
	case uint:
		intVal := int32(val)
		return &intVal, nil
	case uint16:
		intVal := int32(val)
		return &intVal, nil
	case uint32:
		intVal := int32(val)
		return &intVal, nil
	case uint64:
		intVal := int32(val)
		return &intVal, nil
	case int64:
		intVal := int32(val)
		return &intVal, nil
	default:
		return nil, fmt.Errorf("error converting %v to *int32", value.Interface())
	}
}

// ToInt64Ptr -
func ToInt64Ptr(value *reflect.Value) (*int64, error) {
	switch val := value.Interface().(type) {
	case int:
		intVal := int64(val)
		return &intVal, nil
	case uint:
		intVal := int64(val)
		return &intVal, nil
	case uint16:
		intVal := int64(val)
		return &intVal, nil
	case uint32:
		intVal := int64(val)
		return &intVal, nil
	case uint64:
		intVal := int64(val)
		return &intVal, nil
	case int64:
		intVal := val
		return &intVal, nil
	default:
		return nil, fmt.Errorf("error converting %v to *int64", value.Interface())
	}
}

// ToInt -
func ToInt(value *reflect.Value) (int, error) {
	switch val := value.Interface().(type) {
	case int:
		intVal := val
		return intVal, nil
	case uint:
		intVal := int(val)
		return intVal, nil
	case uint16:
		intVal := int(val)
		return intVal, nil
	case uint32:
		intVal := int(val)
		return intVal, nil
	case uint64:
		intVal := int(val)
		return intVal, nil
	case int64:
		intVal := int(val)
		return intVal, nil
	default:
		return 0, fmt.Errorf("error converting %v to int", value.Interface())
	}
}

// ToInt32 -
func ToInt32(value *reflect.Value) (int32, error) {
	switch val := value.Interface().(type) {
	case int:
		intVal := int32(val)
		return intVal, nil
	case uint:
		intVal := int32(val)
		return intVal, nil
	case uint16:
		intVal := int32(val)
		return intVal, nil
	case uint32:
		intVal := int32(val)
		return intVal, nil
	case uint64:
		intVal := int32(val)
		return intVal, nil
	case int64:
		intVal := int32(val)
		return intVal, nil
	default:
		return 0, fmt.Errorf("error converting %v to int32", value.Interface())
	}
}

// ToInt64 -
func ToInt64(value *reflect.Value) (int64, error) {
	switch val := value.Interface().(type) {
	case int:
		intVal := int64(val)
		return intVal, nil
	case uint:
		intVal := int64(val)
		return intVal, nil
	case uint16:
		intVal := int64(val)
		return intVal, nil
	case uint32:
		intVal := int64(val)
		return intVal, nil
	case uint64:
		intVal := int64(val)
		return intVal, nil
	case int64:
		intVal := val
		return intVal, nil
	default:
		return 0, fmt.Errorf("error converting %v to int64", value.Interface())
	}
}
