// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"strings"
	"unicode"
)

// ToCamelCase accepted delimiter: -
func ToCamelCase(str string) string {
	n := ""
	capNext := true
	for _, v := range str {
		if unicode.IsUpper(v) {
			n += string(v)
		}
		if unicode.IsDigit(v) {
			n += string(v)
		}
		if unicode.IsLower(v) {
			if capNext {
				n += strings.ToUpper(string(v))
			} else {
				n += string(v)
			}
		}
		capNext = v == '-'
	}
	return n
}
