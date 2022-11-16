package securebanking

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"secure-banking-uk-initializer/pkg/common"
)

func ConfigureAmCorsService(cookie *http.Cookie) {
	zap.L().Info("Configuring CORS")
	CreateCrestResourceFromConfigFile(fmt.Sprintf("https://%s/am/json/global-config/services/CorsService/configuration?_action=create",
		common.Config.Hosts.IdentityPlatformFQDN),
		"cors-login-ui.json", cookie)
}
