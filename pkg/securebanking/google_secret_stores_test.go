package securebanking

import (
	"github.com/stretchr/testify/assert"
	"secure-banking-uk-initializer/pkg/common"
	"secure-banking-uk-initializer/pkg/types"
	"testing"
)

func configureGlobalConfig() {
	common.Config.Hosts.IdentityPlatformFQDN = "testhost"
	common.Config.Identity.AmRealm = "testrealm"
}

func Test_buildCreateStoreRequest(t *testing.T) {
	configureGlobalConfig()

	store := types.GoogleSecretStore{
		Name:                  "Test Store",
		ServiceAccount:        "test-acc",
		Project:               "test-proj",
		SecretFormat:          "PEM",
		ExpiryDurationSeconds: 600,
	}

	actualRequestUrl, actualRequestBody := buildCreateStoreRequest(store)

	expectedRequestUrl := "https://testhost/am/json/realms/root/realms/testrealm/realm-config/secrets/stores/GoogleSecretManagerSecretStoreProvider/Test%20Store"
	expectedRequestBody := map[string]interface{}{
		"_id":                   "Test Store",
		"serviceAccount":        "test-acc",
		"project":               "test-proj",
		"expiryDurationSeconds": 600,
		"secretFormat":          "PEM",
	}

	assert.Equalf(t, expectedRequestUrl, actualRequestUrl, "buildCreateStoreRequest(%v)", store)
	assert.Equalf(t, expectedRequestBody, actualRequestBody, "buildCreateStoreRequest(%v)", store)
}

func Test_buildSecretMappingRequest(t *testing.T) {
	configureGlobalConfig()

	storeName := "Test Store"
	secretMapping := types.SecretMapping{
		SecretId: "am.services.oauth2.tls.client.cert.authentication",
		Alias:    "a-secret-in-gsm",
	}
	actualRequestUrl, actualRequestBody := buildSecretMappingRequest(storeName, secretMapping)

	expectedRequestUrl := "https://testhost/am/json/realms/root/realms/testrealm/realm-config/secrets/stores/GoogleSecretManagerSecretStoreProvider/Test%20Store/mappings/am.services.oauth2.tls.client.cert.authentication"
	expectedRequestBody := map[string]interface{}{
		"secretId": "am.services.oauth2.tls.client.cert.authentication",
		"aliases":  []string{"a-secret-in-gsm"},
	}

	assert.Equalf(t, expectedRequestUrl, actualRequestUrl, "buildSecretMappingRequest(%v, %v)", storeName, secretMapping)
	assert.Equalf(t, expectedRequestBody, actualRequestBody, "buildCreateStoreRequest(%v, %v)", storeName, secretMapping)
}
