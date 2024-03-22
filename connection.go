package lcuapi

import (
	"errors"
	"fmt"
	process2 "github.com/shirou/gopsutil/process"
	"sync/atomic"
)

type Driver struct {
	authKey   string
	port      string
	isRunning uint32
	*Inquirer
	*Watcher
}

func NewDriver() (c *Driver) {
	c = &Driver{}
	return
}

func (c *Driver) Start(cmd bool, startFunc ...func() error) (keepalive chan error, err error) {
	var mp map[string]string
	var process *process2.Process
	var cmdline string
	var riotClient bool = false

	process, err = GetUxProcessByPsutil()
	if process == nil {
		fmt.Printf("GetUxProcessByPsutil() err: League process not found\n")

		process, _ = GetRiotClientProcessByPsutil()
		if process == nil {
			fmt.Printf("GetRiotClientProcessByPsutil() err: Riot process not found\n")
			return
		}
		riotClient = true
	}
	cmdline, err = process.Cmdline()
	if err != nil {
		return
	}
	mp = flagsToMap(cmdline)
	if len(mp) == 0 {
		err = errors.New("need admin")
		return
	}

	if riotClient {
		c.authKey = mp["riotclient-auth-token"]
		c.port = mp["riotclient-app-port"]
	} else {
		c.authKey = mp["remoting-auth-token"]
		c.port = mp["app-port"]
	}
	c.Inquirer = NewInquirer(c.authKey, c.port)

	c.Watcher, err = NewWatcher(c.authKey, c.port)
	if err != nil {
		return
	}

	atomic.StoreUint32(&c.isRunning, 1)
	for i := range startFunc {
		err = startFunc[i]()
		if err != nil {
			return
		}
	}

	keepalive = make(chan error)
	go func() {
		defer atomic.StoreUint32(&c.isRunning, 0)
		keepalive <- c.Watcher.watch()
	}()
	return
}

func (c *Driver) IsRunning() bool {
	return atomic.LoadUint32(&c.isRunning) == 1
}
