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
			Value:     "jwks",
		},
		JwksCacheTimeout: InheritedValueInt{
			Inherited: false,
			Value:     0,
		},
		RemoteConsentResponseSigningAlg: InheritedValueString{
			Inherited: false,
			Value:     "PS256",
		},
		RemoteConsentRequestSigningAlgorithm: InheritedValueString{
			Inherited: false,
			Value:     "RS256",
		},
		JwkSet: JwkSet{
			Inherited: false,
			Value: "{\"keys\":[{\"kty\":\"RSA\",\"kid\":\"jwt-signer\",\"use\":\"sig\",\"x5t\":\"Vk3cJPwNXkXybNmQ7Urf187Nq2E\",\"x5c\":[\"MIIC1TCCAb2gAwIBAgIIfyDBBVavuiswDQYJKoZIhvcNAQELBQAwGTEXMBUGA1UEAxMOVGVzdCBBdXRob3JpdHkwHhcNMjExMTI2MDgwMjI1WhcNMjIwMjI0MDgwMjI1WjAZMRcwFQYDVQQDEw5UZXN0IEF1dGhvcml0eTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAIMY1aJ0GKMCI+baO6BWdsuXYRVSANCOrqAC3yc+tNbth6j9ukgFJKHC685yUez3S0raT8vD5RIKc9T5Y+kyP9VZBe+m8z1zd6UJ+MydPIqMfgkGFY3pnd8vA9ed3t3gwMo4rxx/DZ+z6l206ngUOo5BebcMWTrzlJb2/HYim+E87sQ7ZztkL1fghET0qGvd76spySxcxJVfy0U8WVMygzePqPlfGjSeANuXXA06Cai5AtS6HGOMZ4ACDDuwY8EutFPZp6H9/X2imc9/OrSqMQxbQ3zLknlxziyiAIJG/cBDyp4TR5C+W1df8qdpIscNtdw9N4kJX7ns8X3ptNWFUw8CAwEAAaMhMB8wHQYDVR0OBBYEFPvx7emzZMIWq5Dgh5S3hVLa9Rj+MA0GCSqGSIb3DQEBCwUAA4IBAQBG36MtMsIGts5Sg3opKVVtQ28hayJEhszlo9GjvI3fZtX07cZdcjX0ItImfGAkEJ3wvXR1W8AqhLTWoUqCY8Wlem1eGVuvvANIuJRNCkE+paCYQ836+D+fb85xEe922I0dUyuDrfHZ/PEBAGwaL7Y2atweBSS0XZhuQIa/kWyMcfY+TDHU7hZgoJ8HIUGOq1eirz4L9qdUzkcKHHWRrYraFarGgjOhPJ2YUuyJVrl9gAK9mXF6BrDTpct46LHnq7imEfKTN87+j+jDyRJ3YdKIIQBJmYkUp5I7P7Nuu+ptWJBml+zbOoBt/ITU6WkBo6dkxSYpVTG0Gh3D+cJyHgyk\"],\"n\":\"gxjVonQYowIj5to7oFZ2y5dhFVIA0I6uoALfJz601u2HqP26SAUkocLrznJR7PdLStpPy8PlEgpz1Plj6TI_1VkF76bzPXN3pQn4zJ08iox-CQYVjemd3y8D153e3eDAyjivHH8Nn7PqXbTqeBQ6jkF5twxZOvOUlvb8diKb4TzuxDtnO2QvV-CERPSoa93vqynJLFzElV_LRTxZUzKDN4-o-V8aNJ4A25dcDToJqLkC1LocY4xngAIMO7BjwS60U9mnof39faKZz386tKoxDFtDfMuSeXHOLKIAgkb9wEPKnhNHkL5bV1_yp2kixw213D03iQlfuezxfem01YVTDw\",\"e\":\"AQAB\",\"d\":\"XuH1lVujjS96XpYqu7R4zIemy3CLiGcMemE5s8TNzBUkr6ncTk3yomVamBPjubeONgHl6RvCSploFofdySUGUFrbUgWqXRqaSMf729QdwkVG3y8ZIJoqGiOEC2WGrV4DCxmVm_FVIfZstR_A5-H0M4uuFU8JsgIj1FO0i6gm3BBTQxGzDv-iRt15vygOS3hRaNd0AZ0XEViOxwdoAGJvuztL2vfATMOQHGAPH0Q1hSYzWM3tn6I_tKfjhkVurp3u4ZRlRtwnFYlNOGV44rrGBgdg0ulqmhSWM8qsxtyvsr_VJWpi31bNojf9cXwCobIDWLg4sJxi5_4GDjWHBSZXYQ\",\"p\":\"8jn49N-0_oKssTrqVjtomNd10wSlSbhwckpEQcdqSIfYCC_WJ9OQRLwZxE7Vg8iJUN42CcDtVX3Cmh9fMG9l4-I51LD2ho1dEniUHfhnkDDIcunk9DyhoHeD4P5Dcx84nV1tRI0UHpwMe9UMwy8uKtvCkFTlAnCOkNvGguFM0d8\",\"q\":\"io0t-ZFM97qcwbHpkE6SH1rXY4r9fh7aAb_dWAaxakY4ZY2EHW3MAoMb7gZyh5AmJoNUbi7YAJvtOb4ZuVx6ZxWlSsFlxgo5Xs8c-XK6D9KnxFvRH8f4fOQ5XrOtq7zjx1Jixfgnnfnwjbaun9d9l02jYJmTEQ--jaTPsK9JhNE\",\"dp\":\"ApzzfZjIOBKq0EKlcoaziyqP39Xl_pSZyfHZKKiBEgU9JGF3uvhCTyuET6TWEtTf_lpXVOWa6dgweD8sZLZe8AVpwEykbDEsNt3MI0Khw2FzWCID0UqyJ6wCZTP5AE3u62utmRa4h9gBHnje6WAh7F5wi-QOkGcsco2cZ58MmBs\",\"dq\":\"M5gxV69xHwtiFos_M9redUipzsrSbSXl_yLItWAAr1eo5sBVQ3RAtWrHetLx0WOvoUXkqgdNrqRiKc-N2sYCWuLno7fzQ1VJWfH8kzPS79N9YBTlAlXARhni904nzT1RAUe_uoMXla-ekddGngVsImzp-y4VYxQe3LZUFTKhSRE\",\"qi\":\"os8rcOf_xa-epMFCfc7sUaDtwHj8pXkc5L-rf7MzdWYJzzk56aC5zPVealKUCPiRinD95LGhZ3R2ObYCkhxWnLHyipx0wsmp-ZBYbt4ZnFLXjhRAVhTWKtcZTk_nNHosLbovcWu0KF570S4s90UO_bhZhVwh6Hqqh29o3eejWWY\"}]}",
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
			//Value:     "http://securebanking-openbanking-uk-rcs:8080/api/rcs/consent/jwk_pub",
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
	b, status := Client.Get(path, map[string]string{
		"Accept":             "application/json",
		"X-Requested-With":   "ForgeRock Identity Cloud Postman Collection",
		"Accept-Api-Version": "protocol=2.0,resource=1.0",
	})
	if status == http.StatusNotFound {
		return false
	}
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
	b, status := Client.Get(path, map[string]string{
		"Accept":             "application/json",
		"X-Requested-With":   "ForgeRock Identity Cloud Postman Collection",
		"Accept-Api-Version": "protocol=2.0,resource=1.0",
	})
	if status == http.StatusNotFound {
		return false
	}
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

	common.RaiseForStatus(err, resp.Error(), resp.Status())

	zap.S().Infow("OIDC claims script", "statusCode", resp.StatusCode(), "claimsScriptID", claimsScript.ID, "createdBy", claimsScript.CreatedBy)
	return claimsScript.ID
}

func GetScriptIdByName(name string) string {
	path := "/am/json/alpha/scripts?_pageSize=20&_sortKeys=name&_queryFilter=true&_pagedResultsOffset=0"
	consent := &AmResult{}
	b, _ := Client.Get(path, map[string]string{
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
	b, status := Client.Get(path, map[string]string{
		"Accept":             "application/json",
		"X-Requested-With":   "ForgeRock Identity Cloud Postman Collection",
		"Accept-Api-Version": "protocol=1.0,resource=1.0",
	})
	if status == http.StatusNotFound {
		return false
	}
	err := json.Unmarshal(b, r)
	if err != nil {
		panic(err)
	}

	return Find(id, r, func(r *Result) string {
		return r.ID
	})
}
