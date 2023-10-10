package lcuapi

import "sync/atomic"

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

func (c *Driver) Run(startFunc ...func() error) (err error) {
	// get lcu process commandline map
	mp, err := GetUxProcessCommandlineMapByCmd()
	if err != nil {
		return
	}

	c.authKey = mp["remoting-auth-token"]
	c.port = mp["app-port"]

	c.Inquirer = NewInquirer(c.authKey, c.port)

	c.Watcher, err = NewWatcher(c.authKey, c.port)
	if err != nil {
		return
	}

	atomic.StoreUint32(&c.isRunning, 1)
	defer atomic.StoreUint32(&c.isRunning, 0)

	for i := range startFunc {
		err = startFunc[i]()
		if err != nil {
			return
		}
	}

	err = c.Watcher.watch()

	return
}

func (c *Driver) IsRunning() bool {
	return atomic.LoadUint32(&c.isRunning) == 1
}
