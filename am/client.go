package am

import (
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/secureBankingAcceleratorToolkit/securebanking-openbanking-uk-fidc-initialiszer/common"
)

type RestReaderWriter interface {
	Get(string, map[string]string, interface{})
	Patch(string, map[string]string, interface{})
	Post(string, map[string]string, interface{})
}

type RestClient struct {
	Resty    *resty.Client
	Cookie   *http.Cookie
	AuthCode string
}

var Client RestReaderWriter

func InitRestReaderWriter(cookie *http.Cookie, authCode string) {
	Client = &RestClient{
		Resty:    resty.New().SetRedirectPolicy(resty.NoRedirectPolicy()).SetError(common.RestError{}),
		Cookie:   cookie,
		AuthCode: authCode,
	}
}

func (r *RestClient) Get(path string, headers map[string]string, ob interface{}) {
	_, err := r.constructRestRequest(headers, ob).Get(path)

	if err != nil {
		panic(err)
	}
}

func (r *RestClient) constructRestRequest(headers map[string]string, ob interface{}) *resty.Request {
	return r.Resty.R().
		SetHeaders(headers).
		SetCookie(r.Cookie).
		SetAuthToken(r.AuthCode).
		SetResult(ob)
}

func (r *RestClient) Post(path string, headers map[string]string, ob interface{}) {
	_, err := r.constructRestRequest(headers, ob).Post(path)

	if err != nil {
		panic(err)
	}
}

func (r *RestClient) Patch(path string, headers map[string]string, ob interface{}) {
	_, err := r.constructRestRequest(headers, ob).Patch(path)

	if err != nil {
		panic(err)
	}
}
