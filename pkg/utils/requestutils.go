// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package utils

import (
	"fmt"
	"io"
)

//ReadRequestBody - read the bytes from the Request Body
func ReadRequestBody(bodyReader io.ReadCloser) ([]byte, error) {
	body := make([]byte, 0)
	buf := make([]byte, 100)
	for {
		count, err := bodyReader.Read(buf)
		body = append(body, buf[:count]...)
		if err == io.EOF {
			bodyReader.Close()
			break
		}
		if err != nil {
			return nil, fmt.Errorf("unable to read POST body %v", err)
		}
	}
	return body, nil
}
