package example

import (
	"lcu"
	"testing"
)

func TestGetUxProcess(t *testing.T) {
	conn, err := lcuapi.NewConnection()
	if err != nil {
		t.Error(err)
		return
	}
	if err := conn.Watcher.Start(); err != nil {
		t.Error(err)
		return
	}
}
