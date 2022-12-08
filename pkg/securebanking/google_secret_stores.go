package securebanking

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"net/url"
	"secure-banking-uk-initializer/pkg/common"
	"secure-banking-uk-initializer/pkg/types"
)

// ConfigureGoogleSecretStores Configures Google Secret Stores in AM if defined in the config
//
// This function processes the "GOOGLE_SECRET_STORES" config block, this is treated as an array and can
// be used to create multiple stores.
//
// Once the store has been created then one or more secrets may be mapped to it using the "SECRET_MAPPINGS" config block
// which is also treated as an array
func ConfigureGoogleSecretStores(cookie *http.Cookie) {
	stores := common.Config.Identity.GoogleSecretStores
	if stores == nil || len(stores) == 0 {
		zap.S().Infow("No Google Secret Stores found in config, nothing to do.")
		return
	}
	for _, store := range stores {
		configureGoogleSecretStore(store, cookie)
		configSecretMappings(store, cookie)
	}
}

func configureGoogleSecretStore(store types.GoogleSecretStore, cookie *http.Cookie) {
	createStoreUrl, storeRequest := buildCreateStoreRequest(store)
	zap.S().Infow("Attempting to configure Google Secret Store", "store", store,
		"requestUrl", createStoreUrl, "requestJson", storeRequest)
	CreateOrUpdateCrestResource("PUT", createStoreUrl, storeRequest, cookie)
}

func buildCreateStoreRequest(store types.GoogleSecretStore) (string, map[string]interface{}) {
	requestBody := make(map[string]interface{})
	requestBody["_id"] = store.Name
	requestBody["serviceAccount"] = store.ServiceAccount
	requestBody["project"] = store.Project
	requestBody["expiryDurationSeconds"] = store.ExpiryDurationSeconds
	requestBody["secretFormat"] = store.SecretFormat

	createStoreUrl := fmt.Sprintf("https://%s/am/json/realms/root/realms/%s/realm-config/secrets/stores/GoogleSecretManagerSecretStoreProvider/%s",
		common.Config.Hosts.IdentityPlatformFQDN, common.Config.Identity.AmRealm, url.PathEscape(store.Name))
	return createStoreUrl, requestBody
}

func configSecretMappings(store types.GoogleSecretStore, cookie *http.Cookie) {
	zap.S().Infow("Attempting to map secrets to store", "store", store)
	for _, mapping := range store.SecretMappings {
		createMappingUrl, mappingRequest := buildSecretMappingRequest(store.Name, mapping)
		CreateOrUpdateCrestResource("PUT", createMappingUrl, mappingRequest, cookie)
	}
}

func buildSecretMappingRequest(storeName string, secretMapping types.SecretMapping) (string, map[string]interface{}) {
	createMappingUrl := fmt.Sprintf("https://%s/am/json/realms/root/realms/%s/realm-config/secrets/stores/GoogleSecretManagerSecretStoreProvider/%s/mappings/%s",
		common.Config.Hosts.IdentityPlatformFQDN, common.Config.Identity.AmRealm, url.PathEscape(storeName), secretMapping.SecretId)

	requestBody := make(map[string]interface{})
	requestBody["secretId"] = secretMapping.SecretId
	requestBody["aliases"] = []string{secretMapping.Alias}

	return createMappingUrl, requestBody
}
