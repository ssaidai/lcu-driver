package lcuapi

type Connection struct {
	Inquirer *Inquirer
	Watcher  *Watcher
	exitChan chan struct{}
}

func NewConnection() (c *Connection, err error) {
	// get lcu process commandline map
	mp, err := GetUxProcessCommandlineMapByCmd()
	if err != nil {
		return
	}

	authKey := mp["remoting-auth-token"]
	port := mp["app-port"]
	if err != nil {
		return
	}

	c = &Connection{
		exitChan: make(chan struct{}),
		Inquirer: NewInquirer(authKey, port),
	}

	c.Watcher, err = NewWatcher(authKey, port)
	if err != nil {
		return
	}

	return
}
