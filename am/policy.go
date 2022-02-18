package am

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/secureBankingAccessToolkit/securebanking-openbanking-uk-fidc-initialiszer/common"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// CreatePolicyServiceUser -
func CreatePolicyServiceUser() {
	if ServiceIdentityExists("service_account.policy") {
		zap.L().Info("Skipping creation of Policy service user")
		return
	}

	b, err := ioutil.ReadFile(common.IamDirectoryPath() + "create-policy-service-user.json")
	if err != nil {
		panic(err)
	}
	path := "/openidm/managed/user/?_action=create"
	s := Client.Post(path, b, map[string]string{
		"Accept":       "*/*",
		"Content-Type": "application/json",
		"Connection":   "keep-alive",
	})

	zap.S().Infow("Policy Service User", "statusCode", s)
}

type PolicyEvaluationScript struct {
	Name        string `json:"name"`
	Context     string `json:"context"`
	Description string `json:"description"`
	Language    string `json:"language"`
	Script      string `json:"script"`
}

// CreatePolicyEvaluationScript - and returns the created ID
func CreatePolicyEvaluationScript(cookie *http.Cookie) string {
	id := GetScriptIdByName("Open Banking Dynamic Policy")
	if id != "" {
		zap.L().Info("Script exists")
		return id
	}

	b, err := ioutil.ReadFile(common.IamDirectoryPath() + "policy-evaluation-script.groovy")
	if err != nil {
		panic(err)
	}
	rawPolicyString := string(b)
	rawPolicyString = strings.ReplaceAll(rawPolicyString, "{{IDM_CLIENT_ID}}", viper.GetString("IDM_CLIENT_ID"))
	rawPolicyString = strings.ReplaceAll(rawPolicyString, "{{IDM_CLIENT_SECRET}}", viper.GetString("IDM_CLIENT_SECRET"))
	rawPolicyString = strings.ReplaceAll(rawPolicyString, "{{IG_IDM_USER}}", viper.GetString("IG_IDM_USER"))
	rawPolicyString = strings.ReplaceAll(rawPolicyString, "{{IG_IDM_PASSWORD}}", viper.GetString("IG_IDM_PASSWORD"))

	scriptB64 := b64.StdEncoding.EncodeToString([]byte(rawPolicyString))

	policyScript := &PolicyEvaluationScript{
		Name:        "Open Banking Dynamic Policy",
		Context:     "POLICY_CONDITION",
		Description: "Open Banking Dynamic Policy",
		Language:    "JAVASCRIPT",
		Script:      scriptB64,
	}

	zap.L().Debug("Creating policy evaluation script")

	path := fmt.Sprintf("https://%s/am/json/alpha/scripts/?_action=create", viper.GetString("IAM_FQDN"))
	scriptBody := &RequestScript{}
	resp, err := client.R().
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
	if PolicySetExists("Open Banking") {
		zap.L().Info("Skipping policy set creation")
		return
	}

	zap.L().Debug("Creating Open Banking policy set")
	b, err := ioutil.ReadFile(common.IamDirectoryPath() + "ob-policy-set.json")
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
	path := "/am/json/alpha/applications/?_action=create"
	s := Client.Post(path, ps, map[string]string{
		"Accept":             "*/*",
		"Content-Type":       "application/json",
		"Connection":         "keep-alive",
		"Accept-API-Version": "protocol=1.0,resource=2.0",
	})

	zap.S().Infow("Open Banking Policy Set", "statusCode", s)
}

func PolicySetExists(name string) bool {
	path := "/am/json/alpha/applications?_pageSize=20&_sortKeys=name&_queryFilter=name+eq+%22%5E(%3F!sunAMDelegationService%24).*%22&_pagedResultsOffset=0"
	serviceIdentity := &AmResult{}
	b, _ := Client.Get(path, map[string]string{
		"Accept":             "application/json",
		"X-Requested-With":   "ForgeRock Identity Cloud Postman Collection",
		"Accept-Api-Version": "protocol=1.0,resource=2.0",
	})

	err := json.Unmarshal(b, serviceIdentity)
	if err != nil {
		panic(err)
	}

	return Find(name, serviceIdentity, func(r *Result) string {
		return r.Name
	})
}

func PolicyExists(name string) bool {
	path := "/am/json/alpha/policies?_pageSize=20&_sortKeys=name&_queryFilter=applicationName+eq+%22Open%20Banking%22&_pagedResultsOffset=0"
	serviceIdentity := &AmResult{}
	b, _ := Client.Get(path, map[string]string{
		"Accept":             "application/json",
		"X-Requested-With":   "ForgeRock Identity Cloud Postman Collection",
		"Accept-Api-Version": "protocol=1.0,resource=2.0",
	})

	err := json.Unmarshal(b, serviceIdentity)
	if err != nil {
		panic(err)
	}

	return Find(name, serviceIdentity, func(r *Result) string {
		return r.Name
	})
}

// CreateAISPPolicy -
func CreateAISPPolicy(policyScriptId string) {
	if PolicyExists("AISP Policy") {
		zap.L().Info("Skipping creation of AISP policy")
		return
	}
	zap.L().Debug("Creating AISP policy")
	b, err := ioutil.ReadFile(common.IamDirectoryPath() + "aisp-policy.json")
	if err != nil {
		panic(err)
	}
	aisp := &CreatePolicy{}
	err = json.Unmarshal(b, aisp)
	if err != nil {
		panic(err)
	}
	aisp.Condition.ScriptID = policyScriptId
	path := "/am/json/alpha/policies/?_action=create"
	s := Client.Post(path, aisp, map[string]string{
		"Accept":             "*/*",
		"Content-Type":       "application/json",
		"Connection":         "keep-alive",
		"Accept-API-Version": "protocol=1.0,resource=2.0",
	})

	zap.S().Infow("AISP policy", "statusCode", s)
}

// CreatePISPPolicy -
func CreatePISPPolicy(policyScriptId string) {
	if PolicyExists("PISP Policy") {
		zap.L().Info("Skipping creation of PISP policy")
		return
	}
	zap.L().Debug("Creating PISP policy")
	b, err := ioutil.ReadFile(common.IamDirectoryPath() + "pisp-policy.json")
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
	path := "/am/json/alpha/policies/?_action=create"
	s := Client.Post(path, pisp, map[string]string{
		"Accept":             "*/*",
		"Content-Type":       "application/json",
		"Connection":         "keep-alive",
		"Accept-API-Version": "protocol=1.0,resource=2.0",
	})

	zap.S().Infow("PISP policy", "statusCode", s)
}

// CreatePolicyEngineOAuth2Client -
func CreatePolicyEngineOAuth2Client() {
	if AlphaClientsExist("policy-client") {
		zap.L().Info("Skipping creation of policy engine oauth2 client")
		return
	}

	zap.L().Debug("Creating policy engine oauth2 client")
	b, err := ioutil.ReadFile(common.IamDirectoryPath() + "create-policy-engine-oauth2-client.json")
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
	path := "/am/json/alpha/realm-config/agents/OAuth2Client/policy-client"
	s := Client.Put(path, engineClient, map[string]string{
		"Accept":           "application/json",
		"Content-Type":     "application/json",
		"Connection":       "keep-alive",
		"X-Requested-With": "ForgeRock Identity Cloud Postman Collection",
	})

	zap.S().Infow("Policy engine OAuth2 client", "statusCode", s)
}
