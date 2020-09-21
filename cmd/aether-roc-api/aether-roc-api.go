// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package main

import (
	"flag"
	"github.com/onosproject/aether-roc-api/pkg/manager"
	"github.com/onosproject/onos-lib-go/pkg/certs"
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

// Start a web server with REST interface proxying the gNMI interface to onos-config
func main() {
	var allowCorsOrigins arrayFlags
	flag.Var(&allowCorsOrigins, "allowCorsOrigin", "URLs of CORS origins (repeated)")
	caPath := flag.String("caPath", "", "path to CA certificate")
	keyPath := flag.String("keyPath", "", "path to client private key")
	certPath := flag.String("certPath", "", "path to client certificate")
	gnmiEndpoint := flag.String("gnmiendpoint", "onos-config:5150", "address of onos-config")
	flag.Parse()

	log.Info("Starting aether-roc-api - connecting to %s", *gnmiEndpoint)

	opts, err := certs.HandleCertPaths(*caPath, *keyPath, *certPath, true)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	mgr, err := manager.NewManager(*gnmiEndpoint, allowCorsOrigins, opts...)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
	mgr.Run()
}
