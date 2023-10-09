package lcuapi

import (
	"encoding/json"
	"fmt"
	"github.com/Vikyanite/lcu-driver/model"
)

type PrintHandler struct{}

func (h *PrintHandler) Handle(data []byte) error {
	println(string(data))
	return nil
}

type EventHandler struct {
	handlers map[string]func(data []byte) error
}

func NewEventHandler() *EventHandler {
	return &EventHandler{
		handlers: make(map[string]func(data []byte) error),
	}
}

func (h *EventHandler) Register(name string, f func([]byte) error) {
	if _, exist := h.handlers[name]; exist {
		// TODO warn log
	}
	h.handlers[name] = f
}

func (h *EventHandler) Handle(binData []byte) (err error) {
	msg := &model.WatcherMsg{}
	err = json.Unmarshal(binData[OnJsonApiEventPrefixLen:len(binData)-1], msg)
	if err != nil {
		return
	}
	f, exist := h.handlers[msg.URI]
	if !exist {
		fmt.Printf("no handler for %s\n", msg.URI)
		return
	}
	return f(msg.Data)
}
