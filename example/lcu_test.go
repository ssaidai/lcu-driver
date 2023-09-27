package example

import (
	"fmt"
	"lcu"
	"testing"
)

func TestGetCurrentSummoner(t *testing.T) {
	err := lcuapi.Init()
	if err != nil {
		t.Error(err)
		return
	}
	if data, err := lcuapi.GetCurrentSummoner(); err != nil {
		t.Error(err)
		return
	} else {
		fmt.Println(data)
	}

}
