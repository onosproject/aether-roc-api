// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package manager

import (
	"github.com/onosproject/aether-roc-api/pkg/southbound"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"google.golang.org/grpc"
)

var mgr Manager

var log = logging.GetLogger("manager")

// Manager single point of entry for the ROC system.
type Manager struct {
	gnmiProvisioner *southbound.GNMIProvisioner
}

// NewManager -
func NewManager(opts ...grpc.DialOption) (*Manager, error) {
	mgr = Manager{}

	var err error
	mgr.gnmiProvisioner = new(southbound.GNMIProvisioner)
	err = mgr.gnmiProvisioner.Init(opts...)
	if err != nil {
		log.Error("Unable to setup GNMI provisioner", err)
		return nil, err
	}

	return &mgr, nil
}

// Run starts the northbound services.
func (m *Manager) Run() {
	log.Info("Starting Manager")
	block := make(chan bool)
	<-block

	log.Info("Manager Stopping")
}
