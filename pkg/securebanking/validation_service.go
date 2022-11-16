package securebanking

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"secure-banking-uk-initializer/pkg/common"
)

func CreateAmValidationService(cookie *http.Cookie) {
	zap.L().Info("Creating Validation Service")

	createValidationServiceUrl := fmt.Sprintf("https://%s/am/json/realms/root/realms/%s/realm-config/services/validation?_action=create",
		common.Config.Hosts.IdentityPlatformFQDN, common.Config.Identity.AmRealm)
	CreateCrestResourceFromConfigFile(createValidationServiceUrl, "create-validation-service.json", cookie)
}
