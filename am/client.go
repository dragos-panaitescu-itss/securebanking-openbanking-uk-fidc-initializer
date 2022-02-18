package am

import (
	"log"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/secureBankingAccessToolkit/securebanking-openbanking-uk-fidc-initialiszer/common"
	"github.com/spf13/viper"
)

type RestReaderWriter interface {
	Get(string, map[string]string) ([]byte, int)
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

func (r *RestClient) Get(path string, headers map[string]string) ([]byte, int) {
	route := r.FQDN + path
	resp, err := r.request(headers).
		Get(route)
	log.Println("Route:", route, resp.Status())
	common.RaiseForStatus(err, resp.Error(), resp.StatusCode())

	return resp.Body(), resp.StatusCode()
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
  
	common.RaiseForStatus(err, resp.Error(), resp.StatusCode())

	return resp.StatusCode()
}

func (r *RestClient) Patch(path string, ob interface{}, headers map[string]string) int {
	resp, err := r.request(headers).
		SetBody(ob).
		Patch(r.FQDN + path)

	common.RaiseForStatus(err, resp.Error(), resp.StatusCode())

	return resp.StatusCode()
}

func (r *RestClient) Put(path string, ob interface{}, headers map[string]string) int {
	resp, err := r.request(headers).
		SetBody(ob).
		SetContentLength(true).
		Put(r.FQDN + path)

	common.RaiseForStatus(err, resp.Error(), resp.StatusCode())

	return resp.StatusCode()
}
