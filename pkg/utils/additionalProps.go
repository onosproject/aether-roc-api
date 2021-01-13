// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
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
