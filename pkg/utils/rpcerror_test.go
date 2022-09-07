// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
//

package utils

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gotest.tools/assert"
	"net/http"
	"testing"
)

func Test_ConvertGRPCError(t *testing.T) {
	grpcErr := status.Error(codes.NotFound, "not found")
	err := ConvertGrpcError(grpcErr)
	assert.Error(t, err, "code=404, message=Not Found")
}

// if we receive an error that is already an HTTP error
// simply return that without converting
func Test_ConvertHttpError(t *testing.T) {
	err := echo.NewHTTPError(http.StatusBadRequest, "http-test-err")
	httpError := ConvertGrpcError(err)
	assert.Error(t, httpError, "code=400, message=http-test-err")
}

func Test_ConvertGrpcError_NoContent(t *testing.T) {
	validationErrMsg := respInternalInvalid + ` test1`
	httpError := ConvertGrpcError(fmt.Errorf(validationErrMsg))
	assert.Error(t, httpError, "code=204, message=rpc error: code = Internal desc = rpc error: code = InvalidArgument test1")
}

func Test_ConvertGrpcError_Leafref(t *testing.T) {
	validationErrMsg := respInvalidValidation + ` Application value starbucks-nvr (string ptr) schema path /a/b/c has leafref path /d/e/f not equal to any target nodes, field name Application value starbucks-nvr (string ptr) schema path /a/b/c has leafref path /d/e/f not equal to any target nodes`
	httpError := ConvertGrpcError(fmt.Errorf(validationErrMsg))
	assert.Error(t, httpError, "code=400, message=Change gives LeafRef error on /d/e/f. Application value starbucks-nvr not present. From path:  /a/b/c")
}

func Test_ConvertGrpcError_ValidationOther(t *testing.T) {
	validationErrMsg := respInvalidValidation + ` test1234` + notEqual
	httpError := ConvertGrpcError(fmt.Errorf(validationErrMsg))
	assert.Error(t, httpError, "code=400, message=test1234")
}

func Test_ConvertGrpcError_ValidationOther1(t *testing.T) {
	validationErrMsg := respInvalidValidation + ` test1234 (test abc)` + " " + schemaPath + " /a/b/c " + hasLrPath + " /d/e/f " + notEqual + " " + notEqual
	httpError := ConvertGrpcError(fmt.Errorf(validationErrMsg))
	assert.Error(t, httpError, "code=400, message=Change gives LeafRef error on /d/e/f. test1234 not present. From path:  /a/b/c")
}

func Test_ConvertGrpcError_ValidationNotRepeated(t *testing.T) {
	validationErrMsg := respInvalidValidation + ` test1234 (test abc)` + " " + schemaPath + " /a/b/c " + hasLrPath + " /d/e/f " + notEqual
	httpError := ConvertGrpcError(fmt.Errorf(validationErrMsg))
	assert.Error(t, httpError, "code=400, message=Change gives LeafRef error on /d/e/f. test1234 not present. From path:  /a/b/c")
}

func Test_ConvertGrpcError_ValidationNoBracket(t *testing.T) {
	validationErrMsg := respInvalidValidation + ` test1234 test abc)` + " " + schemaPath + " /a/b/c " + hasLrPath + " /d/e/f " + notEqual
	httpError := ConvertGrpcError(fmt.Errorf(validationErrMsg))
	assert.Error(t, httpError, "code=400, message=test1234 test abc) schema path /a/b/c has leafref path /d/e/f ")
}

func Test_ConvertGrpcError_InvalidOther(t *testing.T) {
	validationErrMsg := respInvalidBase + ` test1234`
	httpError := ConvertGrpcError(fmt.Errorf(validationErrMsg))
	assert.Error(t, httpError, "code=400, message=rpc error: code = InvalidArgument desc = test1234")
}

func Test_ConvertGrpcError_Unauthorized(t *testing.T) {
	validationErrMsg := respUnauthorized + ` test1234`
	httpError := ConvertGrpcError(fmt.Errorf(validationErrMsg))
	assert.Error(t, httpError, "code=401, message=rpc error: code = Unauthenticated desc = test1234")
}

func Test_ConvertGrpcError_Internal(t *testing.T) {
	validationErrMsg := `rpc error: code = test1234`
	httpError := ConvertGrpcError(fmt.Errorf(validationErrMsg))
	assert.Error(t, httpError, "code=500, message=rpc error: code = test1234")
}
