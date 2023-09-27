package lcuapi

import "github.com/go-resty/resty/v2"

type IInquirer interface {
	Put(uri string, body interface{}) (resp *resty.Response, err error)
	Patch(uri string, body interface{}) (resp *resty.Response, err error)
	Delete(uri string) (resp *resty.Response, err error)
	Get(uri string) (resp *resty.Response, err error)
	Post(uri string, body interface{}) (resp *resty.Response, err error)
	Request(method, uri string, body interface{}) (resp *resty.Response, err error)
}

type IWatcher interface {
	SetHandler(handler IHandler)
	Start() error
	Close()
}

type Connection struct {
	authKey string
	port    string
	IInquirer
	IWatcher
}

func NewConnection() (c *Connection, err error) {
	// get lcu process commandline map
	mp, err := GetUxProcessCommandlineMapByCmd()
	if err != nil {
		return
	}
	c = &Connection{
		authKey:   mp["remoting-auth-token"],
		port:      mp["app-port"],
		IInquirer: &NilInquirer{},
		IWatcher:  &NilWatcher{},
	}

	return
}

func (c *Connection) Init() (err error) {
	c.IInquirer = NewInquirer(c.authKey, c.port)
	if err != nil {
		return
	}
	c.IWatcher, err = NewWatcher(c.authKey, c.port)
	if err != nil {
		return
	}

	// TODO use goroutine to start watcher
	return
}
