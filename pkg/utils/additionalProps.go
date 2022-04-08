// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
//

package utils

import (
	"fmt"
	"reflect"
	"strings"
)

// CheckAdditionalProps - if attributes are not spelt properly they will be
// added to AdditionalProperties map of the JSON Object instead
// This checks that the only Additional Property is "target" (optional)
// Otherwise it give the suggested spellings
func CheckAdditionalProps(additionaProperties map[string]interface{}, ok bool, jsonObj interface{}) error {
	if (len(additionaProperties) > 0 && !ok) || (len(additionaProperties) > 1 && ok) {
		keys := make([]string, 0, len(additionaProperties))
		for k := range additionaProperties {
			keys = append(keys, k)
		}
		joType := reflect.TypeOf(jsonObj).Elem()
		options := make([]string, 0)
		for i := 0; i < joType.NumField(); i++ {
			if joType.Field(i).Name == "AdditionalProperties" {
				continue
			}
			jsonTag := joType.Field(i).Tag.Get("json")
			options = append(options, strings.Split(jsonTag, ",")[0])
		}
		return fmt.Errorf("unexpected properties '%s'. Choose from '%s'",
			strings.Join(keys, ", "), strings.Join(options, ", "))
	}
	return nil
}

// CheckForAdditionalProps - general function to deal with Additional Properties
func CheckForAdditionalProps(jsonObj interface{}) (unchanged map[string]interface{}, target *string) {
	unchanged = make(map[string]interface{})
	jsonObjType := reflect.TypeOf(jsonObj).Elem()
	if jsonObjType.Kind() != reflect.Struct {
		return
	}
	apField, ok := jsonObjType.FieldByName("AdditionalProperties")
	if !ok || apField.Type.Kind().String() != "map" {
		return
	}
	apValue := reflect.ValueOf(jsonObj).Elem().FieldByName("AdditionalProperties")

	for _, addPropName := range apValue.MapKeys() {
		addProp := apValue.MapIndex(addPropName)
		switch addProp.Type().Name() {
		case "AdditionalPropertyTarget":
			targetV := addProp.FieldByName("Target")
			if targetV.Pointer() != 0 {
				targetStr := targetV.Elem().String()
				target = &targetStr
			}
		case "AdditionalPropertyEnterpriseId":
			targetV := addProp.FieldByName("EnterpriseId")
			if targetV.Pointer() != 0 {
				targetStr := targetV.Elem().String()
				target = &targetStr
			}
		case "AdditionalPropertyUnchanged":
			unchV := addProp.FieldByName("Unchanged")
			unchangedProps := unchV.Elem().String()
			for _, p := range strings.Split(unchangedProps, ",") {
				unchanged[p] = struct{}{}
			}
		case "AdditionalPropertiesUnchTarget":
			targetV := addProp.FieldByName("EnterpriseId")
			if targetV.Pointer() != 0 {
				targetStr := targetV.Elem().String()
				target = &targetStr
			}
			unchV := addProp.FieldByName("Unchanged")
			unchangedProps := unchV.Elem().String()
			for _, p := range strings.Split(unchangedProps, ",") {
				unchanged[p] = struct{}{}
			}
		default:
			return
		}
	}
	return unchanged, target
}
