package example

import (
	"fmt"
	"github.com/Vikyanite/lcu-driver"
	"testing"
	"time"
)

func TestSugar(t *testing.T) {
	go func() {
		if err := lcuapi.Run(); err != nil {
			fmt.Printf("lcuapi.Run() err: %+v\n", err)
			return
		}
	}()

	time.Sleep(time.Second * 1)

	data, err := lcuapi.GetCurrentSummoner()
	fmt.Printf("data: %v, err: %v\n", data, err)
}
