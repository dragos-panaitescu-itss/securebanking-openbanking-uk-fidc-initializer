package am

import (
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/secureBankingAccessToolkit/securebanking-openbanking-uk-fidc-initialiszer/common"
	"github.com/spf13/viper"
)

type RestReaderWriter interface {
	Get(string, map[string]string) []byte
	Patch(string, interface{}, map[string]string) int
	Post(string, interface{}, map[string]string) int
	Put(string, interface{}, map[string]string) int
}

type RestClient struct {
	Resty    *resty.Client
	Cookie   *http.Cookie
	AuthCode string
	FQDN     string
}

var Client RestReaderWriter

func InitRestReaderWriter(cookie *http.Cookie, authCode string) {
	Client = &RestClient{
		Resty:    resty.New().SetRedirectPolicy(resty.NoRedirectPolicy()).SetError(common.RestError{}),
		Cookie:   cookie,
		AuthCode: authCode,
		FQDN:     "https://" + viper.GetString("IAM_FQDN"),
	}
}

func (r *RestClient) Get(path string, headers map[string]string) []byte {
	resp, err := r.request(headers).
		Get(r.FQDN + path)

	common.RaiseForStatus(err, resp.Error())

	if err != nil {
		panic(err)
	}

	return resp.Body()
}

func (r *RestClient) request(headers map[string]string) *resty.Request {
	return r.Resty.R().
		SetHeaders(headers).
		SetCookie(r.Cookie).
		SetAuthToken(r.AuthCode)
}

func (r *RestClient) Post(path string, ob interface{}, headers map[string]string) int {
	resp, err := r.request(headers).
		SetBody(ob).
		SetContentLength(true).
		Post(r.FQDN + path)

	common.RaiseForStatus(err, resp.Error())
	if err != nil {
		panic(err)
	}
	return resp.StatusCode()
}

func (r *RestClient) Patch(path string, ob interface{}, headers map[string]string) int {
	resp, err := r.request(headers).
		SetBody(ob).
		Patch(r.FQDN + path)

	common.RaiseForStatus(err, resp.Error())
	if err != nil {
		panic(err)
	}
	return resp.StatusCode()
}

func (r *RestClient) Put(path string, ob interface{}, headers map[string]string) int {
	resp, err := r.request(headers).
		SetBody(ob).
		SetContentLength(true).
		Put(r.FQDN + path)

	common.RaiseForStatus(err, resp.Error())
	if err != nil {
		panic(err)
	}
	return resp.StatusCode()
}
