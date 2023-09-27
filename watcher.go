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
	WSConn   *websocket.Conn
	handler  IHandler
	exitChan chan struct{}
}

type IHandler interface {
	Handle(p []byte) error
}

func NewWatcher(token, port string) (w *Watcher, err error) {
	w = &Watcher{
		exitChan: make(chan struct{}),
		handler:  &PrintHandler{},
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
	return
}

func (c *Watcher) SetHandler(handler IHandler) {
	c.handler = handler
}

func (c *Watcher) Start() (err error) {
	defer c.WSConn.Close()

	if err = c.WSConn.WriteMessage(websocket.TextMessage, []byte("[5, \"OnJsonApiEvent\"]")); err != nil {
		err = fmt.Errorf("write to websocket err: %v\n", err)
		return
	}

	for {
		select {
		case <-c.exitChan:
			err = fmt.Errorf("watcher exit locally")
			return
		default:
			var (
				binData []byte
				msgType int
			)
			fmt.Println("waiting for message")
			if msgType, binData, err = c.WSConn.ReadMessage(); err != nil {
				err = fmt.Errorf("read from websocket err: %v\n", err)
				return
			} else if msgType == websocket.TextMessage {
				if len(binData) < OnJsonApiEventPrefixLen+1 {
					continue
				}
				err = c.handler.Handle(binData)
				fmt.Println(err)
			}
		}
	}
}

func (c *Watcher) Close() {
	close(c.exitChan)
}
