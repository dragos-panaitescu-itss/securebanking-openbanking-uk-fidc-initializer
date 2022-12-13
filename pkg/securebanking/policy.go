package securebanking

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"secure-banking-uk-initializer/pkg/httprest"
	"secure-banking-uk-initializer/pkg/types"

	"secure-banking-uk-initializer/pkg/common"

	"go.uber.org/zap"
)

// CreatePolicyServiceUser -
func CreatePolicyServiceUser() {
	if httprest.ServiceIdentityExists(common.Config.Identity.ServiceAccountPolicyUser) {
		zap.L().Info("Skipping creation of Policy service user")
		return
	}
	serviceUser := &types.ServiceUser{}
	common.Unmarshal(common.Config.Environment.Paths.ConfigSecureBanking+"create-policy-service-user.json", &common.Config, serviceUser)
	path := "/openidm/managed/user/?_action=create"
	//FIDC IDM default user managed objects use a different naming pattern <realm>_user Eg:alpha_user
	if common.Config.Environment.Type == "FIDC" {
		path = "/openidm/managed/" + common.Config.Identity.AmRealm + "_user/?_action=create"
	}
	_, s := httprest.Client.Post(path, serviceUser, map[string]string{
		"Accept":       "*/*",
		"Content-Type": "application/json",
		"Connection":   "keep-alive",
	})

	zap.S().Infow("Policy Service User", "statusCode", s)
}

// CreatePolicyEvaluationScript - and returns the created ID
func CreatePolicyEvaluationScript(cookie *http.Cookie) string {
	id := httprest.GetScriptIdByName("Open Banking Dynamic Policy")
	if id != "" {
		zap.L().Info("Script exists")
		return id
	}

	b, err := common.Template(common.Config.Environment.Paths.ConfigSecureBanking+"policy-evaluation-script.js", &common.Config)
	if err != nil {
		panic(err)
	}

	scriptB64 := b64.StdEncoding.EncodeToString(b)

	policyScript := &types.PolicyEvaluationScript{
		Name:        "Open Banking Dynamic Policy",
		Context:     "POLICY_CONDITION",
		Description: "Open Banking Dynamic Policy",
		Language:    "JAVASCRIPT",
		Script:      scriptB64,
	}

	zap.L().Info("Creating policy evaluation script")

	path := fmt.Sprintf("https://%s/am/json/"+common.Config.Identity.AmRealm+"/scripts/?_action=create", common.Config.Hosts.IdentityPlatformFQDN)
	scriptBody := &types.RequestScript{}
	resp, err := restClient.R().
		SetHeader("Accept", "*/*").
		SetHeader("Content-Type", "application/json").
		SetHeader("Connection", "keep-alive").
		SetHeader("Accept-API-Version", "protocol=2.0,resource=1.0").
		SetContentLength(true).
		SetResult(scriptBody).
		SetCookie(cookie).
		SetBody(policyScript).
		Post(path)

	common.RaiseForStatus(err, resp.Error(), resp.StatusCode())

	zap.S().Infow("Policy Evaluation Script", "statusCode", resp.StatusCode(), "scriptId", scriptBody.ID)
	return scriptBody.ID
}

// CreateOpenBankingPolicySet -
func CreateOpenBankingPolicySet() {
	if httprest.PolicySetExists("Open Banking") {
		zap.L().Info("Skipping policy set creation")
		return
	}

	zap.L().Info("Creating Open Banking policy set")
	ps := &types.OpenBankingPolicySet{}
	err := common.Unmarshal(common.Config.Environment.Paths.ConfigSecureBanking+"ob-policy-set.json", &common.Config, ps)
	if err != nil {
		panic(err)
	}
	zap.S().Infow("Open Banking Policy set unmarshaled", "policy-set", ps)
	path := "/am/json/" + common.Config.Identity.AmRealm + "/applications/?_action=create"
	_, s := httprest.Client.Post(path, ps, map[string]string{
		"Accept":             "*/*",
		"Content-Type":       "application/json",
		"Connection":         "keep-alive",
		"Accept-API-Version": "protocol=1.0,resource=2.0",
	})

	zap.S().Infow("Open Banking Policy Set", "statusCode", s)
}

// CreateAISPPolicy -
func CreateAISPPolicy(policyScriptId string) {
	if httprest.PolicyExists("AISP Policy") {
		zap.L().Info("Skipping creation of AISP policy")
		return
	}
	zap.L().Info("Creating AISP policy")

	aisp := &types.CreatePolicy{}
	err := common.Unmarshal(common.Config.Environment.Paths.ConfigSecureBanking+"aisp-policy.json", &common.Config, aisp)
	if err != nil {
		panic(err)
	}
	aisp.Condition.ScriptID = policyScriptId
	path := "/am/json/" + common.Config.Identity.AmRealm + "/policies/?_action=create"
	_, s := httprest.Client.Post(path, aisp, map[string]string{
		"Accept":             "*/*",
		"Content-Type":       "application/json",
		"Connection":         "keep-alive",
		"Accept-API-Version": "protocol=1.0,resource=2.0",
	})

	zap.S().Infow("AISP policy", "statusCode", s)
}

// CreatePISPPolicy -
func CreatePISPPolicy(policyScriptId string) {
	if httprest.PolicyExists("PISP Policy") {
		zap.L().Info("Skipping creation of PISP policy")
		return
	}
	pisp := &types.CreatePolicy{}
	err := common.Unmarshal(common.Config.Environment.Paths.ConfigSecureBanking+"pisp-policy.json", &common.Config, pisp)
	if err != nil {
		panic(err)
	}
	pisp.Condition.ScriptID = policyScriptId
	zap.S().Infow("PISP Policy", "policy", pisp)
	path := "/am/json/" + common.Config.Identity.AmRealm + "/policies/?_action=create"
	_, s := httprest.Client.Post(path, pisp, map[string]string{
		"Accept":             "*/*",
		"Content-Type":       "application/json",
		"Connection":         "keep-alive",
		"Accept-API-Version": "protocol=1.0,resource=2.0",
	})

	zap.S().Infow("PISP policy", "statusCode", s)
}

// CreatePolicyEngineOAuth2Client -
func CreatePolicyEngineOAuth2Client() {
	if httprest.OAuth2AgentClientsExist("policy-client") {
		zap.L().Info("Skipping creation of policy engine oauth2 client")
		return
	}

	zap.L().Info("Creating policy engine oauth2 client")
	b, err := ioutil.ReadFile(common.Config.Environment.Paths.ConfigSecureBanking + "create-policy-engine-oauth2-client.json")
	if err != nil {
		panic(err)
	}
	engineClient := &types.EngineOAuth2Client{}
	err = json.Unmarshal(b, engineClient)
	if err != nil {
		panic(err)
	}
	engineClient.CoreOAuth2ClientConfig.Userpassword = "password"
	zap.S().Infow("Engine client body", "engine", engineClient)
	path := "/am/json/" + common.Config.Identity.AmRealm + "/realm-config/agents/OAuth2Client/policy-client"
	s := httprest.Client.Put(path, engineClient, map[string]string{
		"Accept":           "application/json",
		"Content-Type":     "application/json",
		"Connection":       "keep-alive",
		"X-Requested-With": "ForgeRock Identity Cloud Postman Collection",
	})

	zap.S().Infow("Policy engine OAuth2 client", "statusCode", s)
}
