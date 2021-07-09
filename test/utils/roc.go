// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package utils

import (
	"context"
	"github.com/onosproject/helmit/pkg/helm"
	"github.com/onosproject/helmit/pkg/input"
	"github.com/onosproject/helmit/pkg/kubernetes"
	"github.com/onosproject/onos-test/pkg/onostest"
	"time"
)

func getCredentials() (string, string, error) {
	kubClient, err := kubernetes.New()
	if err != nil {
		return "", "", err
	}
	secrets, err := kubClient.CoreV1().Secrets().Get(context.Background(), onostest.SecretsName)
	if err != nil {
		return "", "", err
	}
	username := string(secrets.Object.Data["sd-ran-username"])
	password := string(secrets.Object.Data["sd-ran-password"])

	return username, password, nil
}

// CreateRocRelease creates a helm release for a ROC instance
func CreateRocRelease(c *input.Context) (*helm.HelmRelease, error) {
	username, password, err := getCredentials()
	registry := c.GetArg("registry").String("")
	if err != nil {
		return nil, err
	}

	roc := helm.Chart("aether-roc-umbrella", onostest.SdranChartRepo).
		Release("aether-roc-umbrella").
		SetUsername(username).
		SetPassword(password).
		WithTimeout(6*time.Minute).
		Set("global.storage.consensus.enabled", "true").
		Set("aether-roc-api.image.tag", "latest").
		Set("onos-config.image.tag", "latest").
		Set("onos-topo.image.tag", "latest").
		Set("sdcore-adapter-v3.image.tag", "latest").
		Set("aether-roc-gui-v3.ingress.enabled", false).
		Set("aether-roc-gui-v3.service.enabled", false).
		Set("import.sdcore-adapter.v2_1.enabled", false).
		Set("import.aether-roc-gui.v2_1.enabled", false).
		Set("import.aether-roc-gui.v3.enabled", false).
		Set("import.grafana.enabled", false).
		Set("import.prometheus.enabled", false).
		Set("import.onos-gui.enabled", false).
		Set("import.onos-cli.enabled", false). // not needed - can enabled be through Helm for investigations
		Set("global.image.registry", registry)

	return roc, nil
}
