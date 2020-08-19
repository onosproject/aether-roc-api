// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package schemawalker

import (
	"fmt"
	openapi "github.com/go-openapi/spec"
	"github.com/openconfig/goyang/pkg/yang"
)

// ExtractPaths is a recursive function to extract a list of read only paths from a YGOT schema
func ExtractPaths(deviceEntry *yang.Entry, parentState yang.TriState, parentPath string,
	subpathPrefix string, paths map[string]openapi.PathItem, defs openapi.Definitions) error {

	for _, dirEntry := range deviceEntry.Dir {
		if dirEntry.IsLeaf() || dirEntry.IsLeafList() {
			fmt.Printf("%v %s\n", dirEntry.Type, dirEntry.Name)
		} else if dirEntry.IsContainer() {
			fmt.Printf("%v %s\n", dirEntry.Type, dirEntry.Name)
		} else if dirEntry.IsList() {
			fmt.Printf("%v %s\n", dirEntry.Type, dirEntry.Name)
		} else if dirEntry.IsChoice() || dirEntry.IsCase() {
			fmt.Printf("%v %s\n", dirEntry.Type, dirEntry.Name)
		}
	}

	return nil
}
