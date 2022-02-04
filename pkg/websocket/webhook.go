// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
//

package websocket

import (
	"github.com/labstack/echo/v4"
	"github.com/onosproject/aether-roc-api/pkg/utils"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger("webhook")

// webhook - serve the webhook endpoint - distribute to each open websocket
func webhook(ctx echo.Context) error {
	log.Infof("Alert %s", ctx.Request().RequestURI)

	if len(openWebSockets) == 0 {
		log.Infof("No web sockets to send to")
		return nil
	}
	body, err := utils.ReadRequestBody(ctx.Request().Body)
	if err != nil {
		return err
	}
	for ws, c := range openWebSockets {
		log.Infof("Sending to Websocket %p", ws)
		c <- string(body)
	}
	return nil
}
