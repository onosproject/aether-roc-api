// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package utils

import (
	"gotest.tools/assert"
	"reflect"
	"testing"
)

type TestStructInt struct {
	value  interface{}
	expect int
}

type TestStructInt64 struct {
	value  interface{}
	expect int64
}

func Test_ToIntPtr(t *testing.T) {
	testSuite := []TestStructInt{
		{
			value:  -10,
			expect: -10,
		},
		{
			value:  uint(11),
			expect: 11,
		},
		{
			value:  int8(-12),
			expect: -12,
		},
		{
			value:  uint8(13),
			expect: 13,
		},
		{
			value:  int16(-14),
			expect: -14,
		},
		{
			value:  uint16(15),
			expect: 15,
		},
		{
			value:  int32(-16),
			expect: -16,
		},
		{
			value:  uint32(17),
			expect: 17,
		},
		{
			value:  int64(-18),
			expect: -18,
		},
		{
			value:  uint64(19),
			expect: 19,
		},
	}

	for _, test := range testSuite {
		iReflect := reflect.ValueOf(test.value)
		asPtr, err := ToIntPtr(&iReflect)
		assert.NilError(t, err)
		assert.Assert(t, asPtr != nil)
		assert.Equal(t, test.expect, *asPtr)
	}

	// Check the error handler
	iReflect := reflect.ValueOf("a string")
	asPtr, err := ToIntPtr(&iReflect)
	assert.Error(t, err, "unhandled conversion string to *int")
	assert.Assert(t, asPtr == nil)
}

func Test_ToInt64(t *testing.T) {
	testSuite := []TestStructInt64{
		{
			value:  -20,
			expect: -20,
		},
		{
			value:  uint(21),
			expect: 21,
		},
		{
			value:  int8(-22),
			expect: -22,
		},
		{
			value:  uint8(23),
			expect: 23,
		},
		{
			value:  int16(-24),
			expect: -24,
		},
		{
			value:  uint16(25),
			expect: 25,
		},
		{
			value:  int32(-26),
			expect: -26,
		},
		{
			value:  uint32(27),
			expect: 27,
		},
		{
			value:  int64(-28),
			expect: -28,
		},
		{
			value:  uint64(29),
			expect: 29,
		},
	}

	for _, test := range testSuite {
		iReflect := reflect.ValueOf(test.value)
		asInt64, err := ToInt64(&iReflect)
		assert.NilError(t, err)
		assert.Equal(t, test.expect, asInt64)
	}

	// Check the error handler
	iReflect := reflect.ValueOf("a string")
	asPtr, err := ToInt64(&iReflect)
	assert.Error(t, err, "unhandled conversion string to int64")
	assert.Assert(t, asPtr == 0)
}
