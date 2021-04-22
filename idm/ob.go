package idm

import (
	"io/ioutil"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/secureBankingAcceleratorToolkit/securebanking-openbanking-uk-fidc-initialiszer/am"
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

	path := "/openidm/config/managed"
	s := am.Client.Patch(path, b, map[string]string{
		"Accept":       "*/*",
		"Content-Type": "application/json",
		"Connection":   "keep-alive",
	})

	zap.S().Infow("OpenBanking Managed Objects", "statusCode", s)
}

func AddAdditionalCDKObjects(cookie *http.Cookie, accessToken string) {
	zap.L().Debug("Adding OB managed objects")
	b, err := ioutil.ReadFile(viper.GetString("REQUEST_BODY_PATH") + "cdk-additional-objects.json")
	if err != nil {
		panic(err)
	}

	path := "/openidm/config/managed"
	s := am.Client.Patch(path, b, map[string]string{
		"Accept":       "*/*",
		"Content-Type": "application/json",
		"Connection":   "keep-alive",
	})

	zap.S().Infow("OpenBanking Managed Objects", "statusCode", s)
}

func CreateApiJwksEndpoint(cookie *http.Cookie, accessToken string) {
	zap.L().Debug("Adding OB managed objects")
	b, err := ioutil.ReadFile(viper.GetString("REQUEST_BODY_PATH") + "create-jwks-endpoint.json")
	if err != nil {
		panic(err)
	}

	path := "/openidm/config/endpoint/apiclientjwks"
	s := am.Client.Put(path, b, map[string]string{
		"Accept":       "*/*",
		"Content-Type": "application/json",
		"Connection":   "keep-alive",
	})

	zap.S().Infow("JWKS endpoint", "statusCode", s)
}

// CreateUser will create a user that will allow us to create new identities
//    in the alpha realm
func CreateUser(cookie *http.Cookie, accessToken string) {
	zap.L().Debug("Creating new user")
	b, err := ioutil.ReadFile(viper.GetString("REQUEST_BODY_PATH") + "create-user.json")
	if err != nil {
		panic(err)
	}

	path := "/openidm/config/managed"
	s := am.Client.Patch(path, b, map[string]string{
		"Accept":       "*/*",
		"Content-Type": "application/json",
		"Connection":   "keep-alive",
	})

	zap.S().Infow("User created", "statusCode", s)
}
