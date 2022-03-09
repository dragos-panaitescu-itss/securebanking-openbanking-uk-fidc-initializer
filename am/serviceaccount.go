package am

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/secureBankingAccessToolkit/securebanking-openbanking-uk-fidc-initialiszer/common"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"strings"
)

var client = resty.New().SetRedirectPolicy(resty.NoRedirectPolicy()).SetError(common.RestError{})

// CreateIGServiceUser -
func CreateIGServiceUser() {
	if ServiceIdentityExists(viper.GetString("IG_IDM_USER")) {
		zap.L().Info("Skipping creation of IG service user")
		return
	}

	zap.L().Info("Creating IG service user")

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
	_, s := common.Client.Post(path, user, map[string]string{
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

	zap.L().Info("Creating IG OAuth2 client")
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
	path := fmt.Sprintf("/am/json/alpha/realm-config/agents/OAuth2Client/%s", viper.GetString("IG_CLIENT_ID"))
	s := common.Client.Put(path, oauth2Client, map[string]string{
		"Accept":           "application/json",
		"Content-Type":     "application/json",
		"Connection":       "keep-alive",
		"X-Requested-With": "ForgeRock Identity Cloud Postman Collection",
	})

	zap.S().Infow("IG OAuth2 Client", "statusCode", s)
}

// CreateIGPolicyAgent -
func CreateIGPolicyAgent() {
	zap.L().Info("Creating IG Policy agent")
	policyAgent := &PolicyAgent{
		Userpassword: viper.GetString("IG_AGENT_PASSWORD"),
		IgTokenIntrospection: IgTokenIntrospection{
			Value:     "Realm",
			Inherited: false,
		},
	}
	path := fmt.Sprintf("/am/json/realms/root/realms/alpha/realm-config/agents/IdentityGatewayAgent/%s", viper.GetString("IG_AGENT_ID"))
	s := common.Client.Put(path, policyAgent, map[string]string{
		"Accept":           "application/json",
		"Content-Type":     "application/json",
		"Connection":       "keep-alive",
		"X-Requested-With": "ForgeRock Identity Cloud Postman Collection",
	})

	zap.S().Infow("IG Policy Agent", "statusCode", s)
}

func CreateIDMAdminClient(cookie *http.Cookie) {
	zap.L().Info("Creating IDM admin oauth2 client")
	b, err := ioutil.ReadFile(common.IamDirectoryPath() + "idm-admin-client.json")
	if err != nil {
		panic(err)
	}
	config := &OAuth2Client{}
	json.Unmarshal(b, config)
	var redirects []string
	for _, uri := range config.CoreOAuth2ClientConfig.RedirectionUris.Value {
		redirects = append(redirects, strings.ReplaceAll(uri, "{{IAM_FQDN}}", viper.GetString("IAM_FQDN")))
	}
	config.CoreOAuth2ClientConfig.RedirectionUris.Value = redirects
	zap.S().Infow("Admin client request", "body", config)
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

	common.RaiseForStatus(err, resp.Error(), resp.StatusCode())

	zap.S().Infow("IDM Admin Client", "statusCode", resp.StatusCode(), "redirect", config.CoreOAuth2ClientConfig.RedirectionUris.Value)
}

// ServiceIdentityExists will check for service identities in the alpha realm
//   When CDK is removed, these entities might still be persisted. this gives us
//   an indication that we do not need to initialize the environment
func ServiceIdentityExists(identity string) bool {
	filter := "?_queryFilter=uid+eq+%22" + identity + "%22&_fields=username"
	path := "/am/json/realms/root/realms/alpha/users" + filter
	//path := "/am/json/realms/root/realms/alpha/users/" + identity + "?_fields=username"
	serviceIdentityFilter := &common.ResultFilter{}
	b, _ := common.Client.Get(path, map[string]string{
		"Accept":             "application/json",
		"X-Requested-With":   "ForgeRock Identity Cloud Postman Collection",
		"Accept-Api-Version": "protocol=2.1, resource=4.0",
	})

	err := json.Unmarshal(b, serviceIdentityFilter)
	if err != nil {
		panic(err)
	}
	return serviceIdentityFilter.ResultCount > 0
}

// ServiceIdentityExists will check for service identities in the alpha realm
//   When CDK is removed, these entities might still be persisted. this gives us
//   an indication that we do not need to initialize the environment
func GetIdentityIdByUsername(identity string) string {
	filter := "?_queryFilter=uid+eq+%22" + identity + "%22&_fields=username"
	path := "/am/json/realms/root/realms/alpha/users" + filter
	//path := "/am/json/realms/root/realms/alpha/users/" + identity + "?_fields=username"
	result := &common.ResultFilter{}
	b, _ := common.Client.Get(path, map[string]string{
		"Accept":             "application/json",
		"X-Requested-With":   "ForgeRock Identity Cloud Postman Collection",
		"Accept-Api-Version": "protocol=2.1, resource=4.0",
	})

	err := json.Unmarshal(b, result)
	if err != nil {
		panic(err)
		return ""
	}

	zap.S().Info(result)
	userId := result.Result[0].ID
	if userId == "" {
		panic("The user with the username " + identity + " does not exist")
		return ""
	}
	zap.S().Info("The user with the usename ", identity, " has the following id ", userId)

	return userId
}
