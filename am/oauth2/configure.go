package oauth2

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/secureBankingAcceleratorToolkit/securebanking-openbanking-uk-fidc-initialiszer/am"
	"github.com/secureBankingAcceleratorToolkit/securebanking-openbanking-uk-fidc-initialiszer/common"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var client = resty.New().SetRedirectPolicy(resty.NoRedirectPolicy()).SetError(common.RestError{})

// CreateRemoteConsentService -
func CreateRemoteConsentService() {
	zap.L().Debug("Creating remote consent service")
	rc := &RemoteConsent{
		RemoteConsentRequestEncryptionAlgorithm: InheritedValueString{
			Inherited: false,
			Value:     "RSA-OAEP-256",
		},
		PublicKeyLocation: InheritedValueString{
			Inherited: false,
			Value:     "jwks_uri",
		},
		JwksCacheTimeout: InheritedValueInt{
			Inherited: false,
			Value:     0,
		},
		RemoteConsentResponseSigningAlg: InheritedValueString{
			Inherited: false,
			Value:     "RS256",
		},
		RemoteConsentRequestSigningAlgorithm: InheritedValueString{
			Inherited: false,
			Value:     "RS256",
		},
		JwkSet: JwkSet{
			Inherited: false,
		},
		JwkStoreCacheMissCacheTime: InheritedValueInt{
			Inherited: false,
			Value:     0,
		},
		RemoteConsentResponseEncryptionMethod: InheritedValueString{
			Inherited: false,
			Value:     "A128GCM",
		},
		RemoteConsentRedirectURL: InheritedValueString{
			Inherited: false,
			Value:     "https://" + viper.GetString("FQDN") + "/rcs",
		},
		RemoteConsentRequestEncryptionEnabled: InheritedValueBool{
			Inherited: false,
			Value:     false,
		},
		RemoteConsentRequestEncryptionMethod: InheritedValueString{
			Inherited: false,
			Value:     "A128GCM",
		},
		RemoteConsentResponseEncryptionAlgorithm: InheritedValueString{
			Inherited: false,
			Value:     "RSA-OAEP-256",
		},
		RequestTimeLimit: InheritedValueInt{
			Inherited: false,
			Value:     0,
		},
		JwksURI: InheritedValueString{
			Inherited: false,
			Value:     "http://obdemo-rcs-api:8083/api/rcs/consent/jwk_pub",
		},
		Type: Type{
			ID:         "RemoteConsentAgent",
			Name:       "OAuth2 Remote Consent Service",
			Collection: true,
		},
	}
	path := "/am/json/realms/root/realms/alpha/realm-config/agents/RemoteConsentAgent/forgerock-rcs"

	s := am.Client.Put(path, rc, map[string]string{
		"Accept":             "*/*",
		"Connection":         "keep-alive",
		"Accept-API-Version": "protocol=2.0,resource=1.0",
	})

	zap.S().Infow("Remote Consent Service", "statusCode", s)
}

// CreateSoftwarePublisherAgent -
func CreateSoftwarePublisherAgent() {
	zap.L().Debug("Creating software publisher agent")
	pa := PublisherAgent{
		PublicKeyLocation: InheritedValueString{
			Inherited: false,
			Value:     "jwks_uri",
		},
		JwksCacheTimeout: InheritedValueInt{
			Inherited: false,
			Value:     3600000,
		},
		SoftwareStatementSigningAlgorithm: InheritedValueString{
			Inherited: false,
			Value:     "PS256",
		},
		JwkSet: JwkSet{
			Inherited: false,
		},
		Issuer: InheritedValueString{
			Inherited: false,
			Value:     "ForgeRock",
		},
		JwkStoreCacheMissCacheTime: InheritedValueInt{
			Inherited: false,
			Value:     60000,
		},
		JwksURI: InheritedValueString{
			Inherited: false,
			Value:     "https://service.directory.ob.forgerock.financial/api/directory/keys/jwk_uri",
		},
	}
	path := "/am/json/realms/root/realms/alpha/realm-config/agents/SoftwarePublisher/OBRI"
	s := am.Client.Put(path, pa, map[string]string{
		"Accept":             "*/*",
		"Connection":         "keep-alive",
		"Accept-API-Version": "protocol=2.0,resource=1.0",
	})

	zap.S().Infow("Software Publisher Agent", "statusCode", s)
}

// CreateOIDCClaimsScript -
func CreateOIDCClaimsScript(cookie *http.Cookie) string {
	zap.L().Debug("Creating OIDC claims script")
	b, err := ioutil.ReadFile(viper.GetString("REQUEST_BODY_PATH") + "oidc.json")
	if err != nil {
		panic(err)
	}

	path := "https://" + viper.GetString("IAM_FQDN") + "/am/json/alpha/scripts/?_action=create"
	claimsScript := &am.RequestScript{}
	resp, err := client.R().
		SetHeader("Accept", "*/*").
		SetHeader("Content-Type", "application/json").
		SetHeader("Connection", "keep-alive").
		SetHeader("Accept-API-Version", "protocol=2.0,resource=1.0").
		SetContentLength(true).
		SetCookie(cookie).
		SetResult(claimsScript).
		SetBody(b).
		Post(path)

	common.RaiseForStatus(err, resp.Error())

	zap.S().Infow("OIDC claims script", "statusCode", resp.StatusCode(), "claimsScriptID", claimsScript.ID, "createdBy", claimsScript.CreatedBy)
	return claimsScript.ID
}

// UpdateOAuth2Provider - update the oauth 2 provider, must supply the claimScript ID
func UpdateOAuth2Provider(claimsScriptID string) {
	b, err := ioutil.ReadFile(viper.GetString("REQUEST_BODY_PATH") + "oauth2provider.json")
	if err != nil {
		panic(err)
	}

	oauth2Provider := &OAuth2Provider{}
	err = json.Unmarshal(b, oauth2Provider)
	if err != nil {
		panic(err)
	}
	oauth2Provider.CoreOIDCConfig.OidcClaimsScript = claimsScriptID
	zap.S().Infow("Updating OAuth2 provider", "claimScriptId", oauth2Provider.CoreOIDCConfig.OidcClaimsScript)
	path := "/am/json/alpha/realm-config/services/oauth-oidc"
	s := am.Client.Put(path, oauth2Provider, map[string]string{
		"Accept":           "*/*",
		"Content-Type":     "application/json",
		"Connection":       "keep-alive",
		"X-Requested-With": "ForgeRock Identity Cloud Postman Collection",
	})

	zap.S().Infow("OAuth2 provider", "statusCode", s)
}
