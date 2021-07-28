// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package websocket

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

// serveWS - Serve the WebSocket caller - one per browser page
// Each Websocket has its own channel - as webhooks are made, the body is passed
// to the channel and send out over the websocket here
// If something is sent to us on the web socket we quietly ignore it at present
// TODO enhance this so that the context of the caller can be established
//  This might be through an Auth header with JWT that tells us which enterprise
//  the caller belongs to
//  It might be possible to receive an opening message from the websocket that lets
//  us know what the client is interested in e.g. VCS, UE etc
func serveWS(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		receiverChan := make(chan string)
		openWebSockets[ws] = receiverChan
		defer ws.Close()
		defer close(receiverChan)
		defer delete(openWebSockets, ws)
		defer log.Warnf("Closed WebSocket %p", ws)
		go func() {
			for msg := range receiverChan {
				log.Infof("Receieved msg for WS %s", msg)
				if err := websocket.Message.Send(ws, msg); err != nil {
					log.Warnf("error sending message through websocket %p. %s", ws, err)
				}
			}
		}()
		for {
			// Read
			msg := ""
			err := websocket.Message.Receive(ws, &msg)
			if err != nil {
				c.Logger().Error(err)
				return
			}
			log.Infof("%s\n", msg)
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}
