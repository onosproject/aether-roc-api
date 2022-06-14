// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
//

package utils

import (
	"fmt"
	"github.com/openconfig/gnmi/proto/gnmi"
)

const (
	// UnknownKey - replaceable key
	UnknownKey = "unknown_key"
	// UnknownID - replaceable value
	UnknownID = "unknown_id"
)

// ReplaceUnknownKey - postfix on updates
func ReplaceUnknownKey(update *gnmi.Update, keyName string, keyValue interface{}, nameToReplace string, keyMap map[string]interface{}) error {
	length := len(update.GetPath().GetElem())
	for i := length; i >= 0; i-- {
		elem := update.GetPath().GetElem()[i-1]
		_, ok := elem.GetKey()[nameToReplace]
		if !ok {
			continue
		}
		delete(elem.GetKey(), nameToReplace)

		elem.GetKey()[keyName] = fmt.Sprintf("%v", keyValue)
		// if there are meant to be more than one key add another "to be replaced"
		if len(elem.GetKey()) < len(keyMap) {
			elem.GetKey()[nameToReplace] = UnknownID
		}
		return nil
	}
	return fmt.Errorf("no elements found")
}

// RemoveIndexAttributes - remove index attribute updates from a list
func RemoveIndexAttributes(updates []*gnmi.Update, indexPos []int) []*gnmi.Update {
	if len(indexPos) == 0 {
		return updates
	}
	newUpdates := append(updates[:indexPos[0]], updates[indexPos[len(indexPos)-1]+1:]...)
	updates = make([]*gnmi.Update, len(newUpdates))
	copy(updates, newUpdates)
	return updates
}
