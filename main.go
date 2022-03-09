package main

import (
	"github.com/secureBankingAccessToolkit/securebanking-openbanking-uk-fidc-initialiszer/common"
	"github.com/secureBankingAccessToolkit/securebanking-openbanking-uk-fidc-initialiszer/rs"
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
	zap.L().Info("4")
	if !strings.HasSuffix(viper.GetString("MANAGED_OBJECTS_DIRECTORY_PATH"), "/") {
		zap.S().Fatalw("MANAGED_OBJECTS_DIRECTORY_PATH must have a trailing slash /", "MANAGED_OBJECTS_DIRECTORY_PATH", viper.GetString("MANAGED_OBJECTS_DIRECTORY_PATH"))
	}
	zap.L().Info("5")
	if !platform.IsValidX509() {
		zap.L().Fatal("No Valid SSL certificate present in the cdk")
	}
	zap.L().Info("6")
	c := platform.GetCookieNameFromAm()
	zap.L().Info("7")
	s := platform.FromUserSession(c)
	zap.L().Info("8")
	isCloud := viper.GetBool("IS_CLOUD")
	if !isCloud {
		am.CreateIDMAdminClient(s.Cookie)
	}
	if !am.RealmExist(s.Cookie, "alpha") {
		am.CreateAlphaRealm(s.Cookie)
	}
	zap.L().Info("9")
	s.Authenticate()
	common.InitRestReaderWriter(s.Cookie, s.AuthToken.AccessToken)
	zap.L().Info("10")
	am.ApplyAmAuthenticationConfig()
	zap.L().Info("11")
	am.CreateRemoteConsentService()
	zap.L().Info("12")
	am.CreateSoftwarePublisherAgentOBRI()
	zap.L().Info("13")
	am.CreateSoftwarePublisherAgentTestPublisher()
	zap.L().Info("14")

	id := am.CreateOIDCClaimsScript(s.Cookie)
	am.UpdateOAuth2Provider(id)
	zap.L().Info("15")
	time.Sleep(5 * time.Second)

	am.CreatePolicyServiceUser()
	zap.L().Info("16")
	scriptID := am.CreatePolicyEvaluationScript(s.Cookie)
	zap.L().Info("17")
	am.CreateOpenBankingPolicySet()
	zap.L().Info("18")
	am.CreateAISPPolicy(scriptID)
	zap.L().Info("19")
	am.CreatePISPPolicy(scriptID)
	zap.L().Info("20")
	am.CreatePolicyEngineOAuth2Client()
	zap.L().Info("21")
	am.CreateIGServiceUser()
	zap.L().Info("22")
	am.CreateIGOAuth2Client()
	am.CreateIGPolicyAgent()
	// Create and populate data for PSU user
	userId := rs.CreatePSU()
	rs.PopulateRSData(userId)

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
	viper.SetDefault("IS_CLOUD", false)
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
