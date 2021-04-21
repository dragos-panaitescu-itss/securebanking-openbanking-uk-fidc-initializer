package main

import (
	"time"

	"github.com/secureBankingAcceleratorToolkit/securebanking-openbanking-uk-fidc-initialiszer/am"
	"github.com/secureBankingAcceleratorToolkit/securebanking-openbanking-uk-fidc-initialiszer/am/oauth2"
	"github.com/secureBankingAcceleratorToolkit/securebanking-openbanking-uk-fidc-initialiszer/am/policy"
	"github.com/secureBankingAcceleratorToolkit/securebanking-openbanking-uk-fidc-initialiszer/am/realm"
	"github.com/secureBankingAcceleratorToolkit/securebanking-openbanking-uk-fidc-initialiszer/am/serviceaccount"
	"github.com/secureBankingAcceleratorToolkit/securebanking-openbanking-uk-fidc-initialiszer/idm"
	"github.com/secureBankingAcceleratorToolkit/securebanking-openbanking-uk-fidc-initialiszer/platform"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	configureVariables()
	logger, err := configureLogger()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	undo := zap.ReplaceGlobals(logger)
	defer undo()

	if !platform.IsValidX509() {
		zap.L().Fatal("No Valid SSL certificate present in the cdk")
	}

	s := platform.FromAmAdminSession()
	serviceaccount.CreateIDMAdminClient(s.Cookie)
	if !realm.AlphaRealmExists(s.Cookie) {
		realm.CreateAlphaRealm(s.Cookie)
	}

	s.Authenticate()

	am.InitRestReaderWriter(s.Cookie, s.AuthToken.AccessToken)

	if !realm.AlphaClientsExist("policy-client") {
		oauth2.CreateRemoteConsentService(s.Cookie)
		oauth2.CreateSoftwarePublisherAgent(s.Cookie)
		id := oauth2.CreateOIDCClaimsScript(s.Cookie)
		oauth2.UpdateOAuth2Provider(s.Cookie, id)

		time.Sleep(5 * time.Second)

		policy.CreatePolicyServiceUser(s.Cookie, s.AuthToken.AccessToken)
		scriptID := policy.CreatePolicyEvaluationScript(s.Cookie)
		policy.CreateOpenBankingPolicySet(s.Cookie)
		policy.CreateAISPPolicy(s.Cookie)
		policy.CreatePISPPolicy(s.Cookie, scriptID)
		policy.CreatePolicyEngineOAuth2Client(s.Cookie)
	}

	if !realm.AlphaClientsExist(viper.GetString("IG_CLIENT_ID")) {
		serviceaccount.CreateIGServiceUser(s.Cookie, s.AuthToken.AccessToken)
		serviceaccount.CreateIGOAuth2Client(s.Cookie)
		serviceaccount.CreateIGPolicyAgent(s.Cookie)
	}

	time.Sleep(5 * time.Second)
	if !realm.ManagedObjectExists("apiClient") {
		idm.AddOBManagedObjects(s.Cookie, s.AuthToken.AccessToken)
		idm.CreateApiJwksEndpoint(s.Cookie, s.AuthToken.AccessToken)
	}
	if viper.GetString("ENVIRONMENT_TYPE") == "CDK" &&
		!realm.ManagedObjectExists("alpha_user") {
		idm.AddAdditionalCDKObjects(s.Cookie, s.AuthToken.AccessToken)
	}
}

func configureLogger() (*zap.Logger, error) {
	verbose := viper.GetBool("VERBOSE")

	if verbose {
		config := zap.NewProductionConfig()
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		// disable sampling to ensure we get all log messages
		config.Sampling = nil
		return config.Build(zap.AddCaller())
	}
	return zap.NewProduction(zap.AddCaller())
}

func configureVariables() {
	viper.AutomaticEnv()
	viper.SetDefault("VERBOSE", false)
	viper.SetDefault("STRICT", true)
	viper.SetDefault("ENVIRONMENT_TYPE", "CDK")
	viper.SetDefault("FQDN", "obdemo-bank.idhub.cc")
	viper.SetDefault("IAM_FQDN", "iam.idhub.cc")
	viper.SetDefault("AM_REALM", "alpha")
	viper.SetDefault("IG_CLIENT_ID", "ig-client")
	viper.SetDefault("IG_CLIENT_SECRET", "password")
	viper.SetDefault("IG_IDM_USER", "service_account.ig")
	viper.SetDefault("IG_IDM_PASSWORD", "0penBanking!")
	viper.SetDefault("IG_AGENT_ID", "ig-agent")
	viper.SetDefault("OPEN_AM_PASSWORD", "password")
	viper.SetDefault("REQUEST_BODY_PATH", "config/")
}
