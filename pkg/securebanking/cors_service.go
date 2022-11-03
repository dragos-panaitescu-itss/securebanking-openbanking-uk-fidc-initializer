package securebanking

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"secure-banking-uk-initializer/pkg/common"
)

func ConfigureAmCorsService(cookie *http.Cookie) {
	zap.L().Info("Configuring CORS")
<<<<<<< HEAD
	CreateCrestResourceFromConfigFile(fmt.Sprintf("https://%s/am/json/global-config/services/CorsService/configuration?_action=create",
		common.Config.Hosts.IdentityPlatformFQDN),
		"cors-login-ui.json", cookie)
=======

	var corsConfig map[string]interface{}
	common.Unmarshal(common.Config.Environment.Paths.ConfigSecureBanking+"cors-login-ui.json", &common.Config, &corsConfig)
	path := fmt.Sprintf("https://%s/am/json/global-config/services/CorsService/configuration?_action=create", common.Config.Hosts.IdentityPlatformFQDN)

	var responsePayload map[string]interface{}
	resp, err := restClient.R().
		SetHeader("Accept", "*/*").
		SetHeader("Content-Type", "application/json").
		SetHeader("Connection", "keep-alive").
		SetHeader("X-Requested-With", "XMLHttpRequest").
		SetHeader("Accept-API-Version", "protocol=1.0,resource=1.0").
		SetContentLength(true).
		SetCookie(cookie).
		SetBody(corsConfig).
		SetResult(&responsePayload).
		Post(path)

	zap.S().Info("resp is " + resp.String())
	if resp != nil && resp.StatusCode() == 409 {
		zap.S().Info("CORS Service configuration already exists")
	} else {
		common.RaiseForStatus(err, resp.Error(), resp.StatusCode())
		zap.S().Info("Created CORS Service configuration, _id: ", responsePayload["_id"])
	}
>>>>>>> 629 configure cors for login UI (#127)
}
