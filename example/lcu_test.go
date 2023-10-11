package example

import (
	"fmt"
	"github.com/Vikyanite/lcu-driver"
	"testing"
)

func TestSugar(t *testing.T) {
	if keepalive, err := lcuapi.Start(); err != nil {
		fmt.Printf("lcuapi.Start() err: %+v\n", err)
		return
	} else {
		go func() {
			err := <-keepalive
			fmt.Printf("lcuapi.Start() keepalive err: %+v\n", err)
			return
		}()
	}
	data, err := lcuapi.GetCurrentSummoner()
	fmt.Printf("data: %v, err: %v\n", data, err)
}
