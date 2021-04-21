package idm

import (
	"io/ioutil"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/secureBankingAcceleratorToolkit/securebanking-openbanking-uk-fidc-initialiszer/common"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var client = resty.New().SetRedirectPolicy(resty.NoRedirectPolicy()).SetError(common.RestError{})

// AddOBManagedObjects -
func AddOBManagedObjects(cookie *http.Cookie, accessToken string) {
	zap.L().Debug("Adding OB managed objects")
	b, err := ioutil.ReadFile(viper.GetString("REQUEST_BODY_PATH") + "ob-managed-objects.json")
	if err != nil {
		panic(err)
	}

	path := "https://" + viper.GetString("IAM_FQDN") + "/openidm/config/managed"
	resp, err := client.R().
		SetHeader("Accept", "*/*").
		SetHeader("Content-Type", "application/json").
		SetHeader("Connection", "keep-alive").
		SetContentLength(true).
		SetAuthToken(accessToken).
		SetCookie(cookie).
		SetBody(b).
		Patch(path)

	common.RaiseForStatus(err, resp.Error())

	zap.S().Infow("OpenBanking Managed Objects", "statusCode", resp.StatusCode())
}

func AddAdditionalCDKObjects(cookie *http.Cookie, accessToken string) {
	zap.L().Debug("Adding OB managed objects")
	b, err := ioutil.ReadFile(viper.GetString("REQUEST_BODY_PATH") + "cdk-additional-objects.json")
	if err != nil {
		panic(err)
	}

	path := "https://" + viper.GetString("IAM_FQDN") + "/openidm/config/managed"
	resp, err := client.R().
		SetHeader("Accept", "*/*").
		SetHeader("Content-Type", "application/json").
		SetHeader("Connection", "keep-alive").
		SetContentLength(true).
		SetAuthToken(accessToken).
		SetCookie(cookie).
		SetBody(b).
		Patch(path)

	common.RaiseForStatus(err, resp.Error())

	zap.S().Infow("OpenBanking Managed Objects", "statusCode", resp.StatusCode())
}

func CreateApiJwksEndpoint(cookie *http.Cookie, accessToken string) {
	zap.L().Debug("Adding OB managed objects")
	b, err := ioutil.ReadFile(viper.GetString("REQUEST_BODY_PATH") + "create-jwks-endpoint.json")
	if err != nil {
		panic(err)
	}

	path := "https://" + viper.GetString("IAM_FQDN") + "/openidm/config/endpoint/apiclientjwks"
	resp, err := client.R().
		SetHeader("Accept", "*/*").
		SetHeader("Content-Type", "application/json").
		SetHeader("Connection", "keep-alive").
		SetContentLength(true).
		SetAuthToken(accessToken).
		SetCookie(cookie).
		SetBody(b).
		Put(path)

	common.RaiseForStatus(err, resp.Error())

	zap.S().Infow("JWKS endpoint", "statusCode", resp.StatusCode())
}

// CreateUser will create a user that will allow us to create new identities
//    in the alpha realm
func CreateUser(cookie *http.Cookie, accessToken string) {
	zap.L().Debug("Creating new user")
	b, err := ioutil.ReadFile(viper.GetString("REQUEST_BODY_PATH") + "create-user.json")
	if err != nil {
		panic(err)
	}

	path := "https://" + viper.GetString("IAM_FQDN") + "/openidm/config/managed"
	resp, err := client.R().
		SetHeader("Accept", "*/*").
		SetHeader("Content-Type", "application/json").
		SetHeader("Connection", "keep-alive").
		SetContentLength(true).
		SetAuthToken(accessToken).
		SetCookie(cookie).
		SetBody(b).
		Patch(path)

	common.RaiseForStatus(err, resp.Error())

	zap.S().Infow("User created", "statusCode", resp.StatusCode())
}
