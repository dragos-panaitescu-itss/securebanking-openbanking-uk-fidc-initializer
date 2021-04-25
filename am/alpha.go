package am

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/secureBankingAccessToolkit/securebanking-openbanking-uk-fidc-initialiszer/common"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// CreateAlphaRealm creates the alpha realm for a new deployment of CDK
func CreateAlphaRealm(cookie *http.Cookie) {
	zap.L().Debug("Creating Alpha Realm")
	b, err := ioutil.ReadFile(viper.GetString("REQUEST_BODY_PATH") + "alpha-realm.json")
	if err != nil {
		panic(err)
	}
	path := "https://" + viper.GetString("IAM_FQDN") + "/am/json/global-config/realms?_action=create"
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Requested-With", "ForgeRock Identity Cloud Postman Collection").
		SetContentLength(true).
		SetCookie(cookie).
		SetBody(b).
		Post(path)
	common.RaiseForStatus(err, resp.Error())

	zap.S().Infow("Alpha Realm Created", "statusCode", resp.StatusCode())
}

// AlphaRealmExists will check if alpha realm exists
func AlphaRealmExists(cookie *http.Cookie) bool {
	path := "https://" + viper.GetString("IAM_FQDN") + "/am/json/global-config/realms/L2FscGhh"
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("X-Requested-With", "ForgeRock Identity Cloud Postman Collection").
		SetHeader("Accept-Api-Version", "protocol=2.0,resource=1.0").
		SetCookie(cookie).
		Get(path)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode() == 200 {
		zap.L().Info("Alpha realm exists.")
		return true
	}
	return false
}

// AlphaClientsExist - Will return true if clients exist in the alpha realm.
func AlphaClientsExist(clientName string) bool {
	path := "/am/json/realms/root/realms/alpha/realm-config/agents/OAuth2Client?_queryFilter=true&_pageSize=10&_fields=coreOAuth2ClientConfig/status,coreOAuth2ClientConfig/agentgroup"
	result := &AmResult{}
	b := Client.Get(path, map[string]string{
		"Accept":             "application/json",
		"X-Requested-With":   "ForgeRock Identity Cloud Postman Collection",
		"Accept-Api-Version": "protocol=2.0,resource=1.0",
	})

	err := json.Unmarshal(b, result)
	if err != nil {
		panic(err)
	}

	return Find(clientName, result, func(r *Result) string {
		return r.ID
	})
}
