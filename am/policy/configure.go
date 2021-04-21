package policy

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/secureBankingAcceleratorToolkit/securebanking-openbanking-uk-fidc-initialiszer/am"
	"github.com/secureBankingAcceleratorToolkit/securebanking-openbanking-uk-fidc-initialiszer/common"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var client = resty.New().SetRedirectPolicy(resty.NoRedirectPolicy()).SetError(common.RestError{})

// CreatePolicyServiceUser -
func CreatePolicyServiceUser(cookie *http.Cookie, accessToken string) {
	zap.S().Debugw("Creating policy service user", "accessToken", accessToken)
	b, err := ioutil.ReadFile(viper.GetString("REQUEST_BODY_PATH") + "create-policy-service-user.json")
	if err != nil {
		panic(err)
	}
	path := "https://" + viper.GetString("IAM_FQDN") + "/openidm/managed/user/?_action=create"
	resp, err := client.R().
		SetHeader("Accept", "*/*").
		SetHeader("Content-Type", "application/json").
		SetHeader("Connection", "keep-alive").
		SetAuthToken(accessToken).
		SetContentLength(true).
		SetCookie(cookie).
		SetBody(b).
		Post(path)

	common.RaiseForStatus(err, resp.Error())

	zap.S().Infow("Policy Service User", "statusCode", resp.StatusCode())
}

// CreatePolicyEvaluationScript - and returns the created ID
func CreatePolicyEvaluationScript(cookie *http.Cookie) string {
	zap.L().Debug("Creating policy evaluation script")
	b, err := ioutil.ReadFile(viper.GetString("REQUEST_BODY_PATH") + "policy-evaluation-script.json")
	if err != nil {
		panic(err)
	}
	path := "https://" + viper.GetString("IAM_FQDN") + "/am/json/alpha/scripts/?_action=create"
	scriptBody := &am.RequestScript{}
	resp, err := client.R().
		SetHeader("Accept", "*/*").
		SetHeader("Content-Type", "application/json").
		SetHeader("Connection", "keep-alive").
		SetHeader("Accept-API-Version", "protocol=2.0,resource=1.0").
		SetContentLength(true).
		SetResult(scriptBody).
		SetCookie(cookie).
		SetBody(b).
		Post(path)

	common.RaiseForStatus(err, resp.Error())

	zap.S().Infow("Policy Evaluation Script", "statusCode", resp.StatusCode(), "scriptId", scriptBody.ID)
	return scriptBody.ID
}

// CreateOpenBankingPolicySet -
func CreateOpenBankingPolicySet(cookie *http.Cookie) {
	zap.L().Debug("Creating Open Banking policy set")
	b, err := ioutil.ReadFile(viper.GetString("REQUEST_BODY_PATH") + "ob-policy-set.json")
	if err != nil {
		panic(err)
	}
	ps := &OpenBankingPolicySet{}
	err = json.Unmarshal(b, ps)
	if err != nil {
		zap.S().Fatalw("Error unmarshalling policy set", "error", err)
	}
	ps.Realm = "/alpha"
	zap.S().Debugw("Open Banking Policy set unmarshaled", "policy-set", ps)
	path := "https://" + viper.GetString("IAM_FQDN") + "/am/json/alpha/applications/?_action=create"
	resp, err := client.R().
		SetHeader("Accept", "*/*").
		SetHeader("Content-Type", "application/json").
		SetHeader("Connection", "keep-alive").
		SetHeader("Accept-API-Version", "protocol=1.0,resource=2.0").
		SetContentLength(true).
		SetCookie(cookie).
		SetBody(ps).
		Post(path)

	common.RaiseForStatus(err, resp.Error())

	zap.S().Infow("Open Banking Policy Set", "statusCode", resp.StatusCode())
}

// CreateAISPPolicy -
func CreateAISPPolicy(cookie *http.Cookie) {
	zap.L().Debug("Creating AISP policy")
	b, err := ioutil.ReadFile(viper.GetString("REQUEST_BODY_PATH") + "aisp-policy.json")
	if err != nil {
		panic(err)
	}
	path := "https://" + viper.GetString("IAM_FQDN") + "/am/json/alpha/policies/?_action=create"
	resp, err := client.R().
		SetHeader("Accept", "*/*").
		SetHeader("Content-Type", "application/json").
		SetHeader("Connection", "keep-alive").
		SetHeader("Accept-API-Version", "protocol=1.0,resource=2.0").
		SetContentLength(true).
		SetCookie(cookie).
		SetBody(b).
		Post(path)

	common.RaiseForStatus(err, resp.Error())

	zap.S().Infow("AISP policy", "statusCode", resp.StatusCode())
}

// CreatePISPPolicy -
func CreatePISPPolicy(cookie *http.Cookie, policyScriptId string) {
	zap.L().Debug("Creating PISP policy")
	b, err := ioutil.ReadFile(viper.GetString("REQUEST_BODY_PATH") + "pisp-policy.json")
	if err != nil {
		panic(err)
	}
	pisp := &CreatePolicy{}
	err = json.Unmarshal(b, pisp)
	if err != nil {
		panic(err)
	}
	pisp.Condition.ScriptID = policyScriptId
	zap.S().Debugw("PISP Policy", "policy", pisp)
	path := "https://" + viper.GetString("IAM_FQDN") + "/am/json/alpha/policies/?_action=create"
	resp, err := client.R().
		SetHeader("Accept", "*/*").
		SetHeader("Content-Type", "application/json").
		SetHeader("Connection", "keep-alive").
		SetHeader("Accept-API-Version", "protocol=1.0,resource=2.0").
		SetContentLength(true).
		SetCookie(cookie).
		SetBody(pisp).
		Post(path)

	common.RaiseForStatus(err, resp.Error())

	zap.S().Infow("PISP policy", "statusCode", resp.StatusCode())
}

// CreatePolicyEngineOAuth2Client -
func CreatePolicyEngineOAuth2Client(cookie *http.Cookie) {
	zap.L().Debug("Creating policy engine oauth2 client")
	b, err := ioutil.ReadFile(viper.GetString("REQUEST_BODY_PATH") + "create-policy-engine-oauth2-client.json")
	if err != nil {
		panic(err)
	}
	engineClient := &EngineOAuth2Client{}
	err = json.Unmarshal(b, engineClient)
	if err != nil {
		panic(err)
	}
	engineClient.CoreOAuth2ClientConfig.Userpassword = "password"
	zap.S().Debugw("Engine client body", "engine", engineClient)
	path := "https://" + viper.GetString("IAM_FQDN") + "/am/json/alpha/realm-config/agents/OAuth2Client/policy-client"
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetHeader("Connection", "keep-alive").
		SetHeader("X-Requested-With", "ForgeRock Identity Cloud Postman Collection").
		SetContentLength(true).
		SetCookie(cookie).
		SetBody(engineClient).
		Put(path)

	common.RaiseForStatus(err, resp.Error())

	zap.S().Infow("Policy engine OAuth2 client", "statusCode", resp.StatusCode())
}
