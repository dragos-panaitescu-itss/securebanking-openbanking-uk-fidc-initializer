package platform

import (
	"fmt"
	"net/http"
	"secure-banking-uk-initializer/pkg/common"
	"secure-banking-uk-initializer/pkg/types"

	"go.uber.org/zap"
)

func ApplySystemClients(cookie *http.Cookie) {
	zap.L().Info("Creating oauth2 system client")
	oauth2Client := &types.OAuth2Client{}
	e := common.Unmarshal(common.Config.Environment.Paths.ConfigIdentityPlatform+"end-user-ui-oauth2-client.json", &common.Config, oauth2Client)
	if e != nil {
		panic(e)
	}

	zap.S().Debugw("oauth2 system client request", "body", oauth2Client)
	path := fmt.Sprintf("https://%s/am/json/alpha/realm-config/agents/OAuth2Client/end-user-ui", common.Config.Hosts.IdentityPlatformFQDN)
	resp, err := restClient.R().
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetHeader("Connection", "keep-alive").
		SetHeader("X-Requested-With", "ForgeRock Identity Cloud Postman Collection").
		SetContentLength(true).
		SetCookie(cookie).
		SetBody(oauth2Client).
		Put(path)

	common.RaiseForStatus(err, resp.Error(), resp.StatusCode())

	zap.S().Infow("oauth2 system Client", "statusCode", resp.StatusCode(), "redirect", oauth2Client.CoreOAuth2ClientConfig.RedirectionUris.Value)
}
