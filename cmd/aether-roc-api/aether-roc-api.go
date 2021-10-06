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

// OIDCServerURL - address of an OpenID Connect server
const OIDCServerURL = "OIDC_SERVER_URL"

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
	port := flag.Uint("port", 8181, "http port")
	validateResp := flag.Bool("validateResp", true, "Validate response are compliant with OpenAPI3 schema")
	flag.Parse()

	log.Infof("Starting aether-roc-api - connecting to %s", *gnmiEndpoint)

	opts, err := certs.HandleCertPaths(*caPath, *keyPath, *certPath, true)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	authorization := false
	if oidcURL := os.Getenv(OIDCServerURL); oidcURL != "" {
		authorization = true
		log.Infof("Authorization enabled. %s=%s", OIDCServerURL, oidcURL)
		// OIDCServerURL is also referenced in jwt.go (from onos-lib-go)
		// It only applies to /sdcore/synchronize/:id (all gnmi requests are passed
		// down to onos-config for authorization)
	} else {
		log.Infof("Authorization not enabled %s", os.Getenv(OIDCServerURL))
	}

	mgr, err := manager.NewManager(*gnmiEndpoint, allowCorsOrigins, *validateResp, authorization, opts...)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
	mgr.Run(*port)
}
