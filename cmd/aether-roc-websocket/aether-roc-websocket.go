// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package main

import (
	"flag"
	"github.com/onosproject/aether-roc-api/pkg/websocket"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"os"
)

var log = logging.GetLogger("main")

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

// Start a web server to handle both a webhook and websockets
func main() {
	var allowCorsOrigins arrayFlags
	flag.Var(&allowCorsOrigins, "allowCorsOrigin", "URLs of CORS origins (repeated)")
	port := flag.Uint("port", 8120, "http port")

	mgr, err := websocket.NewManager(allowCorsOrigins)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
	mgr.Run(*port)
}
