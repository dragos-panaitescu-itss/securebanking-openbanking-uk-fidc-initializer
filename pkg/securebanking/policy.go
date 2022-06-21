package securebanking

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"secure-banking-uk-initializer/pkg/httprest"
	"secure-banking-uk-initializer/pkg/types"
	"strings"

	"secure-banking-uk-initializer/pkg/common"

	"go.uber.org/zap"
)

// CreatePolicyServiceUser -
func CreatePolicyServiceUser() {
	if httprest.ServiceIdentityExists(common.Config.Identity.ServiceAccountPolicy) {
		zap.L().Info("Skipping creation of Policy service user")
		return
	}

	b, err := ioutil.ReadFile(common.Config.Environment.Paths.ConfigSecureBanking + "create-policy-service-user.json")
	if err != nil {
		panic(err)
	}
	path := "/openidm/managed/user/?_action=create"
	_, s := httprest.Client.Post(path, b, map[string]string{
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

	b, err := ioutil.ReadFile(common.Config.Environment.Paths.ConfigSecureBanking + "policy-evaluation-script.js")
	if err != nil {
		panic(err)
	}
	rawPolicyString := string(b)
	rawPolicyString = strings.ReplaceAll(rawPolicyString, "{{IDM_CLIENT_ID}}", common.Config.Identity.IdmClientId)
	rawPolicyString = strings.ReplaceAll(rawPolicyString, "{{IDM_CLIENT_SECRET}}", common.Config.Identity.IdmClientSecret)
	rawPolicyString = strings.ReplaceAll(rawPolicyString, "{{IG_IDM_USER}}", common.Config.Ig.IgIdmUser)
	rawPolicyString = strings.ReplaceAll(rawPolicyString, "{{IG_IDM_PASSWORD}}", common.Config.Ig.IgIdmPassword)

	scriptB64 := b64.StdEncoding.EncodeToString([]byte(rawPolicyString))

	policyScript := &types.PolicyEvaluationScript{
		Name:        "Open Banking Dynamic Policy",
		Context:     "POLICY_CONDITION",
		Description: "Open Banking Dynamic Policy",
		Language:    "JAVASCRIPT",
		Script:      scriptB64,
	}

	zap.L().Info("Creating policy evaluation script")

	path := fmt.Sprintf("https://%s/am/json/alpha/scripts/?_action=create", common.Config.Hosts.IdentityPlatformFQDN)
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
	b, err := ioutil.ReadFile(common.Config.Environment.Paths.ConfigSecureBanking + "ob-policy-set.json")
	if err != nil {
		panic(err)
	}
	ps := &types.OpenBankingPolicySet{}
	err = json.Unmarshal(b, ps)
	if err != nil {
		zap.S().Fatalw("Error unmarshalling policy set", "error", err)
	}
	ps.Realm = "/alpha"
	zap.S().Infow("Open Banking Policy set unmarshaled", "policy-set", ps)
	path := "/am/json/alpha/applications/?_action=create"
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
	b, err := ioutil.ReadFile(common.Config.Environment.Paths.ConfigSecureBanking + "aisp-policy.json")
	if err != nil {
		panic(err)
	}
	aisp := &types.CreatePolicy{}
	err = json.Unmarshal(b, aisp)
	if err != nil {
		panic(err)
	}
	var resources []string
	for _, s := range aisp.Resources {
		resources = append(resources, strings.ReplaceAll(s, "{{HOSTS.BASE_DOMAIN}}", common.Config.Hosts.BaseDomain))
	}
	aisp.Resources = resources
	aisp.Condition.ScriptID = policyScriptId
	path := "/am/json/alpha/policies/?_action=create"
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
	zap.L().Info("Creating PISP policy")
	b, err := ioutil.ReadFile(common.Config.Environment.Paths.ConfigSecureBanking + "pisp-policy.json")
	if err != nil {
		panic(err)
	}
	pisp := &types.CreatePolicy{}
	err = json.Unmarshal(b, pisp)
	if err != nil {
		panic(err)
	}
	var resources []string
	for _, s := range pisp.Resources {
		resources = append(resources, strings.ReplaceAll(s, "{{HOSTS.BASE_DOMAIN}}", common.Config.Hosts.BaseDomain))
	}
	pisp.Resources = resources
	pisp.Condition.ScriptID = policyScriptId
	zap.S().Infow("PISP Policy", "policy", pisp)
	path := "/am/json/alpha/policies/?_action=create"
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
	if httprest.AlphaClientsExist("policy-client") {
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
	path := "/am/json/alpha/realm-config/agents/OAuth2Client/policy-client"
	s := httprest.Client.Put(path, engineClient, map[string]string{
		"Accept":           "application/json",
		"Content-Type":     "application/json",
		"Connection":       "keep-alive",
		"X-Requested-With": "ForgeRock Identity Cloud Postman Collection",
	})

	zap.S().Infow("Policy engine OAuth2 client", "statusCode", s)
}
