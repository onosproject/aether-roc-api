// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
//

package main

import (
	"flag"
	"fmt"
	"github.com/onosproject/aether-roc-api/pkg/manager"
	"github.com/onosproject/onos-lib-go/pkg/certs"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"os"
	"strings"
	"time"
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

func stringToLogLevel(l string) logging.Level {
	switch strings.ToLower(l) {
	case "debug":
		return logging.DebugLevel
	case "info":
		return logging.InfoLevel
	case "warn":
		return logging.WarnLevel
	case "error":
		return logging.ErrorLevel
	default:
		log.Warnf("Provided log level (%s) is unknown, defaulting to INFO. Valid values are: DEBUG, INFO, WARN, ERROR")
		return logging.InfoLevel
	}
}

// Start a web server with REST interface proxying the gNMI interface to onos-config
func main() {
	var allowCorsOrigins arrayFlags
	flag.Var(&allowCorsOrigins, "allowCorsOrigin", "URLs of CORS origins (repeated)")
	caPath := flag.String("caPath", "", "path to CA certificate")
	keyPath := flag.String("keyPath", "", "path to client private key")
	certPath := flag.String("certPath", "", "path to client certificate")
	gnmiEndpoint := flag.String("gnmiEndpoint", "onos-config:5150", "address of onos-config")
	gnmiTimeout := flag.Duration("gnmiTimeout", 10*time.Second, "timeout for the gnmi requests")
	port := flag.Uint("port", 8181, "http port")
	validateResp := flag.Bool("validateResp", true, "Validate response are compliant with OpenAPI3 schema")
	logLevel := flag.String("logLevel", "INFO", "Set the log level (DEBUG, INFO, WARN, ERROR)")
	flag.Parse()

	log.SetLevel(stringToLogLevel(*logLevel))

	log.Infow("Starting aether-roc-api",
		"gnmiEndpoint", *gnmiEndpoint,
		"allowCorsOrigin", allowCorsOrigins,
		"caPath", *caPath,
		"keyPath", *keyPath,
		"gnmiTimeout", fmt.Sprintf("%gs", gnmiTimeout.Seconds()),
		"port", *port,
		"validateResp", *validateResp,
		"logLevel", *logLevel)

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

	mgr, err := manager.NewManager(*gnmiEndpoint, allowCorsOrigins, *validateResp, authorization, *gnmiTimeout, opts...)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
	mgr.Run(*port)
}
