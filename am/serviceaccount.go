package am

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/secureBankingAccessToolkit/securebanking-openbanking-uk-fidc-initialiszer/common"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var client = resty.New().SetRedirectPolicy(resty.NoRedirectPolicy()).SetError(common.RestError{})

// CreateIGServiceUser -
func CreateIGServiceUser() {
	if ServiceIdentityExists(viper.GetString("IG_IDM_USER")) {
		zap.L().Info("Skipping creation of IG service user")
		return
	}

	zap.L().Debug("Creating IG service user")

	user := &common.ServiceUser{
		UserName:  viper.GetString("IG_IDM_USER"),
		SN:        "Service Account",
		GivenName: "IG",
		Mail:      "ig@acme.com",
		Password:  viper.GetString("IG_IDM_PASSWORD"),
		AuthzRole: []common.AuthzRole{
			{
				Ref: "internal/role/openidm-admin",
			},
		},
	}
	path := "/openidm/managed/user/?_action=create"
	s := Client.Post(path, user, map[string]string{
		"Accept":       "*/*",
		"Content-Type": "application/json",
		"Connection":   "keep-alive",
	})

	zap.S().Infow("IG Service User", "statusCode", s)
}

// CreateIGOAuth2Client -
func CreateIGOAuth2Client() {
	if AlphaClientsExist(viper.GetString("IG_CLIENT_ID")) {
		zap.L().Info("Skipping creation of IG oauth2 client")
		return
	}

	zap.L().Debug("Creating IG OAuth2 client")
	b, err := ioutil.ReadFile(common.IamDirectoryPath() + "ig-oauth2-client.json")
	if err != nil {
		panic(err)
	}

	oauth2Client := &OAuth2Client{}
	err = json.Unmarshal(b, oauth2Client)
	if err != nil {
		panic(err)
	}
	oauth2Client.CoreOAuth2ClientConfig.Userpassword = "password"
	path := "/am/json/alpha/realm-config/agents/OAuth2Client/" + viper.GetString("IG_CLIENT_ID")
	s := Client.Put(path, oauth2Client, map[string]string{
		"Accept":           "application/json",
		"Content-Type":     "application/json",
		"Connection":       "keep-alive",
		"X-Requested-With": "ForgeRock Identity Cloud Postman Collection",
	})

	zap.S().Infow("IG OAuth2 Client", "statusCode", s)
}

// CreateIGPolicyAgent -
func CreateIGPolicyAgent() {
	if ServiceIdentityExists("service_account.policy") {
		zap.L().Info("Skipping creation of IG policy agent")
		return
	}
	zap.L().Debug("Creating IG Policy agent")
	policyAgent := &PolicyAgent{
		Userpassword: "password",
		IgTokenIntrospection: IgTokenIntrospection{
			Value:     "Realm",
			Inherited: false,
		},
	}
	path := "/am/json/alpha/realm-config/agents/IdentityGatewayAgent/ig-agent"
	s := Client.Put(path, policyAgent, map[string]string{
		"Accept":           "application/json",
		"Content-Type":     "application/json",
		"Connection":       "keep-alive",
		"X-Requested-With": "ForgeRock Identity Cloud Postman Collection",
	})

	zap.S().Infow("IG Policy Agent", "statusCode", s)
}

func CreateIDMAdminClient(cookie *http.Cookie) {
	zap.L().Debug("Creating IDM admin oauth2 client")
	b, err := ioutil.ReadFile(common.IamDirectoryPath() + "idm-admin-client.json")
	if err != nil {
		panic(err)
	}
	config := &OAuth2Client{}
	json.Unmarshal(b, config)
	var redirect string
	for _, uri := range config.CoreOAuth2ClientConfig.RedirectionUris.Value {
		redirect = strings.ReplaceAll(uri, "{{IAM_FQDN}}", viper.GetString("IAM_FQDN"))
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

// ServiceIdentityExists will check for service identities in the alpha realm
//   When CDK is removed, these entities might still be persisted. this gives us
//   an indication that we do not need to initialize the environment
func ServiceIdentityExists(identity string) bool {
	path := "/am/json/realms/root/realms/alpha/users?_queryFilter=true&_pageSize=10&_fields=cn,mail,username,inetUserStatus"
	serviceIdentity := &AmResult{}
	b := Client.Get(path, map[string]string{
		"Accept":             "application/json",
		"X-Requested-With":   "ForgeRock Identity Cloud Postman Collection",
		"Accept-Api-Version": "protocol=2.1, resource=4.0",
	})

	err := json.Unmarshal(b, serviceIdentity)
	if err != nil {
		panic(err)
	}

	return Find(identity, serviceIdentity, func(r *Result) string {
		return r.Username
	})
}
