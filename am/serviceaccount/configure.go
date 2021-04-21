package serviceaccount

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/secureBankingAcceleratorToolkit/securebanking-openbanking-uk-fidc-initialiszer/common"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var client = resty.New().SetRedirectPolicy(resty.NoRedirectPolicy()).SetError(common.RestError{})

// CreateIGServiceUser -
func CreateIGServiceUser(cookie *http.Cookie, accessToken string) {
	zap.L().Debug("Creating IG service user")

	user := &common.ServiceUser{
		UserName:  viper.GetString("IG_IDM_USER"),
		SN:        "Service Account",
		GivenName: "IG",
		Mail:      "ig@acme.com",
		Password:  viper.GetString("OPEN_AM_PASSWORD"),
		AuthzRole: []common.AuthzRole{
			{
				Ref: "internal/role/openidm-admin",
			},
		},
	}
	path := "https://" + viper.GetString("IAM_FQDN") + "/openidm/managed/user/?_action=create"
	resp, err := client.R().
		SetHeader("Accept", "*/*").
		SetHeader("Content-Type", "application/json").
		SetHeader("Connection", "keep-alive").
		SetContentLength(true).
		SetAuthToken(accessToken).
		SetCookie(cookie).
		SetBody(user).
		Post(path)

	common.RaiseForStatus(err, resp.Error())

	zap.S().Infow("IG Service User", "statusCode", resp.StatusCode())
}

// CreateIGOAuth2Client -
func CreateIGOAuth2Client(cookie *http.Cookie) {
	zap.L().Debug("Creating IG OAuth2 client")
	b, err := ioutil.ReadFile(viper.GetString("REQUEST_BODY_PATH") + "ig-oauth2-client.json")
	if err != nil {
		panic(err)
	}

	oauth2Client := &OAuth2Client{}
	err = json.Unmarshal(b, oauth2Client)
	if err != nil {
		panic(err)
	}
	oauth2Client.CoreOAuth2ClientConfig.Userpassword = "password"
	path := "https://" + viper.GetString("IAM_FQDN") + "/am/json/alpha/realm-config/agents/OAuth2Client/" + viper.GetString("IG_CLIENT_ID")
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetHeader("Connection", "keep-alive").
		SetHeader("X-Requested-With", "ForgeRock Identity Cloud Postman Collection").
		SetContentLength(true).
		SetCookie(cookie).
		SetBody(oauth2Client).
		Put(path)

	common.RaiseForStatus(err, resp.Error())

	zap.S().Infow("IG OAuth2 Client", "statusCode", resp.StatusCode())
}

// CreateIGPolicyAgent -
func CreateIGPolicyAgent(cookie *http.Cookie) {
	zap.L().Debug("Creating IG Policy agent")
	policyAgent := &PolicyAgent{
		Userpassword: "password",
		IgTokenIntrospection: IgTokenIntrospection{
			Value:     "Realm",
			Inherited: false,
		},
	}
	path := "https://" + viper.GetString("IAM_FQDN") + "/am/json/alpha/realm-config/agents/IdentityGatewayAgent/ig-agent"
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetHeader("Connection", "keep-alive").
		SetHeader("X-Requested-With", "ForgeRock Identity Cloud Postman Collection").
		SetContentLength(true).
		SetCookie(cookie).
		SetBody(policyAgent).
		Put(path)

	common.RaiseForStatus(err, resp.Error())

	zap.S().Infow("IG Policy Agent", "statusCode", resp.StatusCode())
}

func CreateIDMAdminClient(cookie *http.Cookie) {
	zap.L().Debug("Creating IDM admin oauth2 client")
	b, err := ioutil.ReadFile(viper.GetString("REQUEST_BODY_PATH") + "idm-admin-client.json")
	if err != nil {
		panic(err)
	}
	config := &OAuth2Client{}
	json.Unmarshal(b, config)
	var redirect string
	for _, uri := range config.CoreOAuth2ClientConfig.RedirectionUris.Value {
		redirect = strings.ReplaceAll(uri, "IAM_FQDN", viper.GetString("IAM_FQDN"))
	}
	config.CoreOAuth2ClientConfig.RedirectionUris.Value = []string{redirect}
	zap.S().Debugw("Admin client request", "body", config)
	path := "https://" + viper.GetString("IAM_FQDN") + "/am/json/realm-config/agents/OAuth2Client/idmAdminClient"
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetHeader("Connection", "keep-alive").
		SetHeader("X-Requested-With", "ForgeRock Identity Cloud Postman Collection").
		SetContentLength(true).
		SetCookie(cookie).
		SetBody(config).
		Put(path)

	common.RaiseForStatus(err, resp.Error())

	zap.S().Infow("IDM Admin Client", "statusCode", resp.StatusCode(), "redirect", config.CoreOAuth2ClientConfig.RedirectionUris.Value)
}
