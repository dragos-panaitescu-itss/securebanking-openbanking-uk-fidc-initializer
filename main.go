package main

import (
	"strings"
	"time"

	"github.com/secureBankingAccessToolkit/securebanking-openbanking-uk-fidc-initialiszer/am"
	"github.com/secureBankingAccessToolkit/securebanking-openbanking-uk-fidc-initialiszer/platform"
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

	if !strings.HasSuffix(viper.GetString("MANAGED_OBJECTS_DIRECTORY_PATH"), "/") {
		zap.S().Fatalw("MANAGED_OBJECTS_DIRECTORY_PATH must have a trailing slash /", "MANAGED_OBJECTS_DIRECTORY_PATH", viper.GetString("MANAGED_OBJECTS_DIRECTORY_PATH"))
	}

	if !platform.IsValidX509() {
		zap.L().Fatal("No Valid SSL certificate present in the cdk")
	}

	c := platform.GetCookieNameFromAm()
	s := platform.FromUserSession(c)

	am.CreateIDMAdminClient(s.Cookie)
	if !am.AlphaRealmExists(s.Cookie) {
		am.CreateAlphaRealm(s.Cookie)
	}

	s.Authenticate()
	am.InitRestReaderWriter(s.Cookie, s.AuthToken.AccessToken)

	am.ApplyAmAuthenticationConfig()

	am.CreateRemoteConsentService()
	am.CreateSoftwarePublisherAgentOBRI()
	am.CreateSoftwarePublisherAgentTestPublisher()

	id := am.CreateOIDCClaimsScript(s.Cookie)
	am.UpdateOAuth2Provider(id)

	time.Sleep(5 * time.Second)

	am.CreatePolicyServiceUser()
	scriptID := am.CreatePolicyEvaluationScript(s.Cookie)
	am.CreateOpenBankingPolicySet()
	am.CreateAISPPolicy(scriptID)
	am.CreatePISPPolicy(scriptID)
	am.CreatePolicyEngineOAuth2Client()

	am.CreateIGServiceUser()
	am.CreateIGOAuth2Client()
	am.CreateIGPolicyAgent()
	am.CreatePSU()
	am.PopulateRSData()

	am.ApplySystemClients(s.Cookie)

	time.Sleep(5 * time.Second)
	am.AddOBManagedObjects()

	am.CreateApiJwksEndpoint()
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
	viper.SetDefault("RCS_UI_FQDN", "rcs-ui.dev.forgerock.financial")
	viper.SetDefault("IAM_FQDN", "iam.idhub.cc")
	viper.SetDefault("IDM_FQDN", "idm")
	viper.SetDefault("AM_REALM", "alpha")
	viper.SetDefault("IDM_CLIENT_ID", "policy-client")
	viper.SetDefault("IDM_CLIENT_SECRET", "password")
	viper.SetDefault("IG_CLIENT_ID", "ig-client")
	viper.SetDefault("IG_CLIENT_SECRET", "password")
	viper.SetDefault("IG_RCS_SECRET", "password")
	viper.SetDefault("IG_SSA_SECRET", "password")
	viper.SetDefault("IG_IDM_USER", "service_account.ig")
	viper.SetDefault("IG_IDM_PASSWORD", "0penBanking!")
	viper.SetDefault("IG_AGENT_ID", "ig-agent")
	viper.SetDefault("IG_AGENT_PASSWORD", "password")
	viper.SetDefault("OPEN_AM_USERNAME", "amadmin")
	viper.SetDefault("OPEN_AM_PASSWORD", "password")
	viper.SetDefault("MANAGED_OBJECTS_DIRECTORY_PATH", "config/defaults/managed-objects/")
	viper.SetDefault("IAM_DIRECTORY_PATH", "config/defaults/")
	viper.SetDefault("SCHEME", "https")
	viper.SetDefault("PSU_USERNAME", "psu")
	viper.SetDefault("PSU_PASSWORD", "0penBanking!")
}
