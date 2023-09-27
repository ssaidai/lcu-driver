package lcuapi

import (
	"encoding/base64"
	"github.com/go-resty/resty/v2"
	"net/http"
)

type Inquirer struct {
	*resty.Client
}

func NewInquirer(token, port string) (ret *Inquirer) {
	ret = &Inquirer{}

	LCUHeader := http.Header{}
	LCUHeader.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("riot:"+token)))
	LCUHeader.Add("Content-Type", "application/json")
	LCUHeader.Add("Accept", "application/json")

	// init http client
	ret.Client = resty.New()
	ret.Client.Header = LCUHeader
	ret.Client.SetBaseURL("https://127.0.0.1:" + port)

	return &Inquirer{
		Client: resty.New(),
	}
}

func (c *Inquirer) Put(uri string, body interface{}) (resp *resty.Response, err error) {
	return c.R().SetBody(body).Put(uri)
}

func (c *Inquirer) Patch(uri string, body interface{}) (resp *resty.Response, err error) {
	return c.R().SetBody(body).Patch(uri)
}

func (c *Inquirer) Delete(uri string) (resp *resty.Response, err error) {
	return c.R().Delete(uri)
}

func (c *Inquirer) Get(uri string) (resp *resty.Response, err error) {
	return c.R().Get(uri)
}

func (c *Inquirer) Post(uri string, body interface{}) (resp *resty.Response, err error) {
	return c.R().SetBody(body).Post(uri)
}
