package am

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/secureBankingAccessToolkit/securebanking-openbanking-uk-fidc-initialiszer/common"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// CreateRemoteConsentService -
func CreateRemoteConsentService() {
	if RemoteConsentExists("forgerock-rcs") {
		zap.L().Info("Remote consent exists. skipping")
		return
	}
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
			Value:     "HS256",
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
			Value:     fmt.Sprintf("https://%s", viper.GetString("RCS_UI_FQDN")),
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
			Value:     "http://securebanking-openbanking-uk-rcs:8080/api/rcs/consent/jwk_pub",
		},
		Type: Type{
			ID:         "RemoteConsentAgent",
			Name:       "OAuth2 Remote Consent Service",
			Collection: true,
		},
		Userpassword: viper.GetString("IG_RCS_SECRET"),
	}
	path := "/am/json/realms/root/realms/alpha/realm-config/agents/RemoteConsentAgent/forgerock-rcs"

	s := Client.Put(path, rc, map[string]string{
		"Accept":             "*/*",
		"Connection":         "keep-alive",
		"Accept-API-Version": "protocol=2.0,resource=1.0",
	})

	zap.S().Infow("Remote Consent Service", "statusCode", s)
}

func RemoteConsentExists(name string) bool {
	path := "/am/json/realms/root/realms/alpha/realm-config/agents/RemoteConsentAgent?_queryFilter=true&_pageSize=10&_fields=agentgroup"
	consent := &AmResult{}
	b := Client.Get(path, map[string]string{
		"Accept":             "application/json",
		"X-Requested-With":   "ForgeRock Identity Cloud Postman Collection",
		"Accept-Api-Version": "protocol=2.0,resource=1.0",
	})

	err := json.Unmarshal(b, consent)
	if err != nil {
		panic(err)
	}

	return Find(name, consent, func(r *Result) string {
		return r.ID
	})
}

// CreateSoftwarePublisherAgent OBRI
func CreateSoftwarePublisherAgentOBRI() {
	if SoftwarePublisherAgentExists("OBRI") {
		zap.L().Info("Skipping creation of Software publisher agent")
		return
	}

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
	s := Client.Put(path, pa, map[string]string{
		"Accept":             "*/*",
		"Connection":         "keep-alive",
		"Accept-API-Version": "protocol=2.0,resource=1.0",
	})

	zap.S().Infow("Software Publisher Agent", "statusCode", s)
}

// CreateSoftwarePublisherAgent test-publisher
func CreateSoftwarePublisherAgentTestPublisher() {
	if SoftwarePublisherAgentExists("test-publisher") {
		zap.L().Info("Skipping creation of Software publisher agent")
		return
	}

	zap.L().Debug("Creating software publisher agent")
	pa := PublisherAgent{
		Userpassword: viper.GetString("IG_SSA_SECRET"),
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
			Value:     "HS256",
		},
		JwkSet: JwkSet{
			Inherited: false,
		},
		Issuer: InheritedValueString{
			Inherited: false,
			Value:     "test-publisher",
		},
		JwkStoreCacheMissCacheTime: InheritedValueInt{
			Inherited: false,
			Value:     60000,
		},
		JwksURI: InheritedValueString{
			Inherited: false,
		},
	}
	path := "/am/json/realms/root/realms/alpha/realm-config/agents/SoftwarePublisher/test-publisher"
	s := Client.Put(path, pa, map[string]string{
		"Accept":             "*/*",
		"Connection":         "keep-alive",
		"Accept-API-Version": "protocol=2.0,resource=1.0",
	})

	zap.S().Infow("Software Publisher Agent", "statusCode", s)
}

func SoftwarePublisherAgentExists(name string) bool {
	path := "/am/json/realms/root/realms/alpha/realm-config/agents/SoftwarePublisher?_queryFilter=true&_pageSize=10&_fields=agentgroup"
	agent := &AmResult{}
	b := Client.Get(path, map[string]string{
		"Accept":             "application/json",
		"X-Requested-With":   "ForgeRock Identity Cloud Postman Collection",
		"Accept-Api-Version": "protocol=2.0,resource=1.0",
	})

	err := json.Unmarshal(b, agent)
	if err != nil {
		panic(err)
	}

	return Find(name, agent, func(r *Result) string {
		return r.ID
	})
}

// CreateOIDCClaimsScript -
func CreateOIDCClaimsScript(cookie *http.Cookie) string {
	id := GetScriptIdByName("Open Banking OIDC Claims Script")
	if id != "" {
		zap.L().Info("Script exists")
		return id
	}
	zap.L().Debug("Creating OIDC claims script")
	b, err := ioutil.ReadFile(common.IamDirectoryPath() + "oidc.json")
	if err != nil {
		panic(err)
	}

	path := fmt.Sprintf("https://%s/am/json/alpha/scripts/?_action=create", viper.GetString("IAM_FQDN"))
	claimsScript := &RequestScript{}
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

func GetScriptIdByName(name string) string {
	path := "/am/json/alpha/scripts?_pageSize=20&_sortKeys=name&_queryFilter=true&_pagedResultsOffset=0"
	consent := &AmResult{}
	b := Client.Get(path, map[string]string{
		"Accept":             "application/json",
		"X-Requested-With":   "ForgeRock Identity Cloud Postman Collection",
		"Accept-Api-Version": "protocol=1.0,resource=1.0",
	})

	err := json.Unmarshal(b, consent)
	if err != nil {
		panic(err)
	}

	return FindIdByName(name, consent, func(r *Result) string {
		return r.Name
	})
}

// UpdateOAuth2Provider - update the oauth 2 provider, must supply the claimScript ID
func UpdateOAuth2Provider(claimsScriptID string) {
	if Oauth2ProviderExists("oauth-oidc") {
		zap.L().Info("OAuth2 provider exists")
		return
	}
	b, err := ioutil.ReadFile(common.IamDirectoryPath() + "oauth2provider.json")
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
	s := Client.Put(path, oauth2Provider, map[string]string{
		"Accept":           "*/*",
		"Content-Type":     "application/json",
		"Connection":       "keep-alive",
		"X-Requested-With": "ForgeRock Identity Cloud Postman Collection",
	})

	zap.S().Infow("OAuth2 provider", "statusCode", s)
}

func Oauth2ProviderExists(id string) bool {
	path := "/am/json/realms/root/realms/alpha/realm-config/services?_queryFilter=true"
	r := &AmResult{}
	b := Client.Get(path, map[string]string{
		"Accept":             "application/json",
		"X-Requested-With":   "ForgeRock Identity Cloud Postman Collection",
		"Accept-Api-Version": "protocol=1.0,resource=1.0",
	})

	err := json.Unmarshal(b, r)
	if err != nil {
		panic(err)
	}

	return Find(id, r, func(r *Result) string {
		return r.ID
	})
}
