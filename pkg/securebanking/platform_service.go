package securebanking

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"secure-banking-uk-initializer/pkg/common"
)

func ConfigureAmPlatformService(cookie *http.Cookie) {
	zap.L().Info("Configuring AM Global Services Platform")
	UpdateCrestResourceFromConfigFile(fmt.Sprintf("https://%s/am/json/global-config/services/platform",
		common.Config.Hosts.IdentityPlatformFQDN), "global-services-platform.json", cookie)
}
