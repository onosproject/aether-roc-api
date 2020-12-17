// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package utils

import (
	"fmt"
	"github.com/openconfig/gnmi/proto/gnmi"
)

// ReplaceUnknownKey - postfix on updates
func ReplaceUnknownKey(update *gnmi.Update, keyName string, keyValue interface{}, nameToReplace string, valueToReplace string) error {
	len := len(update.GetPath().GetElem())
	for i := len; i >= 0; i-- {
		elem := update.GetPath().GetElem()[i-1]
		k, ok := elem.GetKey()[nameToReplace]
		if !ok {
			continue
		}
		delete(elem.GetKey(), nameToReplace)
		if k == valueToReplace {
			elem.GetKey()[keyName] = keyValue.(string)
		} else {
			return fmt.Errorf("unexpected key value %s", k)
		}
		return nil
	}
	return fmt.Errorf("no elements found")
}
