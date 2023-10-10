package lcuapi

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

// Watcher is used to watch the LOL Client events
type Watcher struct {
	WSConn  *websocket.Conn
	handler IHandler
}

type IHandler interface {
	Handle(p []byte) error
}

func NewWatcher(token, port string) (w *Watcher, err error) {
	w = &Watcher{
		handler: &PrintHandler{},
	}

	LCUHeader := http.Header{}
	LCUHeader.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("riot:"+token)))
	LCUHeader.Add("Content-Type", "application/json")
	LCUHeader.Add("Accept", "application/json")

	wsDialer := &websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		TLSClientConfig:  &tls.Config{InsecureSkipVerify: true},
		HandshakeTimeout: 0,
	}
	w.WSConn, _, err = wsDialer.Dial("wss://127.0.0.1:"+port, LCUHeader)
	if err != nil {
		return
	}
	if err = w.WSConn.WriteMessage(websocket.TextMessage, []byte("[5, \"OnJsonApiEvent\"]")); err != nil {
		err = fmt.Errorf("write init msg to websocket err: %v\n", err)
		return
	}
	return
}

func (c *Watcher) SetHandler(handler IHandler) {
	c.handler = handler
}

func (c *Watcher) watch() (err error) {
	defer c.WSConn.Close()
	for {
		var (
			binData []byte
			msgType int
		)
		if msgType, binData, err = c.WSConn.ReadMessage(); err != nil {
			err = fmt.Errorf("read from websocket err: %v\n", err)
			return
		} else if msgType == websocket.TextMessage {
			if len(binData) < OnJsonApiEventPrefixLen+1 {
				continue
			}
			err = c.handler.Handle(binData)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
