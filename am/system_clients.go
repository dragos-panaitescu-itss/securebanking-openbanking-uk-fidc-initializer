package am

import (
	"encoding/json"
	"fmt"
	"github.com/secureBankingAccessToolkit/securebanking-openbanking-uk-fidc-initialiszer/common"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func ApplySystemClients(cookie *http.Cookie) {
	zap.L().Debug("Creating oauth2 system client")
	b, err := ioutil.ReadFile(common.IamDirectoryPath() + "am-oauth2-system-clients/" + "end-user-ui-oauth2-client.json")
	if err != nil {
		panic(err)
	}
	config := &OAuth2Client{}
	json.Unmarshal(b, config)
	var redirects []string
	for _, uri := range config.CoreOAuth2ClientConfig.RedirectionUris.Value {
		s := strings.ReplaceAll(uri, "{{IAM_FQDN}}", viper.GetString("IAM_FQDN"))
		redirects = append(redirects, s)
	}
	config.CoreOAuth2ClientConfig.RedirectionUris.Value = redirects
	zap.S().Debugw("oauth2 system client request", "body", config)
	path := fmt.Sprintf("https://%s/am/json/alpha/realm-config/agents/OAuth2Client/end-user-ui", viper.GetString("IAM_FQDN"))
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

	zap.S().Infow("oauth2 system Client", "statusCode", resp.StatusCode(), "redirect", config.CoreOAuth2ClientConfig.RedirectionUris.Value)
}
