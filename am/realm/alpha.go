package realm

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/secureBankingAcceleratorToolkit/securebanking-openbanking-uk-fidc-initialiszer/common"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var client = resty.New().SetRedirectPolicy(resty.NoRedirectPolicy()).SetError(common.RestError{})

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

// CheckAlphaRealm will check if alpha realm exists and if there are any clients within
//   exit with success code if 200 resonse and clients exist.
func CheckAlphaRealm(cookie *http.Cookie) bool {
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

type resultCount struct {
	Count int `json:"resultCount"`
}

//https://iam.idhub.cc/am/json/realms/root/realms/alpha/realm-config/agents/OAuth2Client?_queryFilter=true&_pageSize=10&_fields=coreOAuth2ClientConfig/status,coreOAuth2ClientConfig/agentgroup
func CheckAlphaClients(cookie *http.Cookie) bool {
	path := "https://" + viper.GetString("IAM_FQDN") + "/am/json/realms/root/realms/alpha/realm-config/agents/OAuth2Client?_queryFilter=true&_pageSize=10&_fields=coreOAuth2ClientConfig/status,coreOAuth2ClientConfig/agentgroup"
	result := &resultCount{}
	_, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("X-Requested-With", "ForgeRock Identity Cloud Postman Collection").
		SetHeader("Accept-Api-Version", "protocol=2.0,resource=1.0").
		SetCookie(cookie).
		SetResult(result).
		Get(path)
	if err != nil {
		panic(err)
	}
	if result.Count > 0 {
		zap.L().Info("Clients exist in the alpha realm, inilialization assummed")
		return true
	}
	return false
}

// CheckServiceIdentities will check for service identities in the alpha realm
//   When CDK is removed, these entities might still be persisted. this gives us
//   an indication that we do not need to initialize the environment
func CheckServiceIdentities(cookie *http.Cookie) bool {
	path := "https://" + viper.GetString("IAM_FQDN") + "/am/json/realms/root/realms/alpha/users?_queryFilter=true&_pageSize=10&_fields=cn,mail,username,inetUserStatus"
	result := &resultCount{}
	_, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("X-Requested-With", "ForgeRock Identity Cloud Postman Collection").
		SetHeader("Accept-Api-Version", "protocol=2.1, resource=4.0").
		SetCookie(cookie).
		SetResult(result).
		Get(path)
	if err != nil {
		panic(err)
	}
	if result.Count > 0 {
		zap.L().Info("Identities exist in the alpha realm, inilialization assummed")
		return true
	}
	return false
}

// https://iam.idhub.cc/openidm/config/managed
func CheckObjects(cookie *http.Cookie, accessToken string, objectName string) bool {
	path := "https://" + viper.GetString("IAM_FQDN") + "/openidm/config/managed"
	result := &OBManagedObjects{}
	_, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("X-Requested-With", "ForgeRock Identity Cloud Postman Collection").
		SetAuthToken(accessToken).
		SetCookie(cookie).
		SetResult(result).
		Get(path)
	if err != nil {
		panic(err)
	}
	for _, o := range result.Objects {
		zap.S().Infow("checking", "object", o)
		if strings.Contains(o.Name, "objectName") {
			zap.L().Info("obTpp found, skipping managed objects")
			return true
		}
	}
	return false
}

// OBManagedObjects model
type OBManagedObjects struct {
	ID      string `json:"_id"`
	Objects []struct {
		Name string `json:"name"`
	} `json:"objects"`
}
