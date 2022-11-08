package securebanking

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"secure-banking-uk-initializer/pkg/httprest"
	"secure-banking-uk-initializer/pkg/types"

	"secure-banking-uk-initializer/pkg/common"

	"go.uber.org/zap"
)

func CreateSecureBankingRemoteConsentService() {
	remoteConsentId := common.Config.Identity.RemoteConsentId
	if remoteConsentExists(remoteConsentId) {
		zap.L().Info("Remote consent exists. skipping")
		return
	}
	zap.L().Info("Creating remote consent service")
	rc := &types.RemoteConsent{
		RemoteConsentRequestEncryptionAlgorithm: types.InheritedValueString{
			Inherited: false,
			Value:     "RSA-OAEP-256",
		},
		PublicKeyLocation: types.InheritedValueString{
			Inherited: false,
			Value:     "jwks",
		},
		JwksCacheTimeout: types.InheritedValueInt{
			Inherited: false,
			Value:     0,
		},
		RemoteConsentResponseSigningAlg: types.InheritedValueString{
			Inherited: false,
			Value:     "PS256",
		},
		RemoteConsentRequestSigningAlgorithm: types.InheritedValueString{
			Inherited: false,
			Value:     "RS256",
		},
		JwkSet: types.JwkSet{
			Inherited: false,
			Value:     "{\"keys\":[{\"kty\":\"RSA\",\"kid\":\"jwt-signer\",\"use\":\"sig\",\"x5t\":\"Vk3cJPwNXkXybNmQ7Urf187Nq2E\",\"x5c\":[\"MIIC1TCCAb2gAwIBAgIIfyDBBVavuiswDQYJKoZIhvcNAQELBQAwGTEXMBUGA1UEAxMOVGVzdCBBdXRob3JpdHkwHhcNMjExMTI2MDgwMjI1WhcNMjIwMjI0MDgwMjI1WjAZMRcwFQYDVQQDEw5UZXN0IEF1dGhvcml0eTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAIMY1aJ0GKMCI+baO6BWdsuXYRVSANCOrqAC3yc+tNbth6j9ukgFJKHC685yUez3S0raT8vD5RIKc9T5Y+kyP9VZBe+m8z1zd6UJ+MydPIqMfgkGFY3pnd8vA9ed3t3gwMo4rxx/DZ+z6l206ngUOo5BebcMWTrzlJb2/HYim+E87sQ7ZztkL1fghET0qGvd76spySxcxJVfy0U8WVMygzePqPlfGjSeANuXXA06Cai5AtS6HGOMZ4ACDDuwY8EutFPZp6H9/X2imc9/OrSqMQxbQ3zLknlxziyiAIJG/cBDyp4TR5C+W1df8qdpIscNtdw9N4kJX7ns8X3ptNWFUw8CAwEAAaMhMB8wHQYDVR0OBBYEFPvx7emzZMIWq5Dgh5S3hVLa9Rj+MA0GCSqGSIb3DQEBCwUAA4IBAQBG36MtMsIGts5Sg3opKVVtQ28hayJEhszlo9GjvI3fZtX07cZdcjX0ItImfGAkEJ3wvXR1W8AqhLTWoUqCY8Wlem1eGVuvvANIuJRNCkE+paCYQ836+D+fb85xEe922I0dUyuDrfHZ/PEBAGwaL7Y2atweBSS0XZhuQIa/kWyMcfY+TDHU7hZgoJ8HIUGOq1eirz4L9qdUzkcKHHWRrYraFarGgjOhPJ2YUuyJVrl9gAK9mXF6BrDTpct46LHnq7imEfKTN87+j+jDyRJ3YdKIIQBJmYkUp5I7P7Nuu+ptWJBml+zbOoBt/ITU6WkBo6dkxSYpVTG0Gh3D+cJyHgyk\"],\"n\":\"gxjVonQYowIj5to7oFZ2y5dhFVIA0I6uoALfJz601u2HqP26SAUkocLrznJR7PdLStpPy8PlEgpz1Plj6TI_1VkF76bzPXN3pQn4zJ08iox-CQYVjemd3y8D153e3eDAyjivHH8Nn7PqXbTqeBQ6jkF5twxZOvOUlvb8diKb4TzuxDtnO2QvV-CERPSoa93vqynJLFzElV_LRTxZUzKDN4-o-V8aNJ4A25dcDToJqLkC1LocY4xngAIMO7BjwS60U9mnof39faKZz386tKoxDFtDfMuSeXHOLKIAgkb9wEPKnhNHkL5bV1_yp2kixw213D03iQlfuezxfem01YVTDw\",\"e\":\"AQAB\",\"d\":\"XuH1lVujjS96XpYqu7R4zIemy3CLiGcMemE5s8TNzBUkr6ncTk3yomVamBPjubeONgHl6RvCSploFofdySUGUFrbUgWqXRqaSMf729QdwkVG3y8ZIJoqGiOEC2WGrV4DCxmVm_FVIfZstR_A5-H0M4uuFU8JsgIj1FO0i6gm3BBTQxGzDv-iRt15vygOS3hRaNd0AZ0XEViOxwdoAGJvuztL2vfATMOQHGAPH0Q1hSYzWM3tn6I_tKfjhkVurp3u4ZRlRtwnFYlNOGV44rrGBgdg0ulqmhSWM8qsxtyvsr_VJWpi31bNojf9cXwCobIDWLg4sJxi5_4GDjWHBSZXYQ\",\"p\":\"8jn49N-0_oKssTrqVjtomNd10wSlSbhwckpEQcdqSIfYCC_WJ9OQRLwZxE7Vg8iJUN42CcDtVX3Cmh9fMG9l4-I51LD2ho1dEniUHfhnkDDIcunk9DyhoHeD4P5Dcx84nV1tRI0UHpwMe9UMwy8uKtvCkFTlAnCOkNvGguFM0d8\",\"q\":\"io0t-ZFM97qcwbHpkE6SH1rXY4r9fh7aAb_dWAaxakY4ZY2EHW3MAoMb7gZyh5AmJoNUbi7YAJvtOb4ZuVx6ZxWlSsFlxgo5Xs8c-XK6D9KnxFvRH8f4fOQ5XrOtq7zjx1Jixfgnnfnwjbaun9d9l02jYJmTEQ--jaTPsK9JhNE\",\"dp\":\"ApzzfZjIOBKq0EKlcoaziyqP39Xl_pSZyfHZKKiBEgU9JGF3uvhCTyuET6TWEtTf_lpXVOWa6dgweD8sZLZe8AVpwEykbDEsNt3MI0Khw2FzWCID0UqyJ6wCZTP5AE3u62utmRa4h9gBHnje6WAh7F5wi-QOkGcsco2cZ58MmBs\",\"dq\":\"M5gxV69xHwtiFos_M9redUipzsrSbSXl_yLItWAAr1eo5sBVQ3RAtWrHetLx0WOvoUXkqgdNrqRiKc-N2sYCWuLno7fzQ1VJWfH8kzPS79N9YBTlAlXARhni904nzT1RAUe_uoMXla-ekddGngVsImzp-y4VYxQe3LZUFTKhSRE\",\"qi\":\"os8rcOf_xa-epMFCfc7sUaDtwHj8pXkc5L-rf7MzdWYJzzk56aC5zPVealKUCPiRinD95LGhZ3R2ObYCkhxWnLHyipx0wsmp-ZBYbt4ZnFLXjhRAVhTWKtcZTk_nNHosLbovcWu0KF570S4s90UO_bhZhVwh6Hqqh29o3eejWWY\"}]}",
		},
		JwkStoreCacheMissCacheTime: types.InheritedValueInt{
			Inherited: false,
			Value:     0,
		},
		RemoteConsentResponseEncryptionMethod: types.InheritedValueString{
			Inherited: false,
			Value:     "A128GCM",
		},
		RemoteConsentRedirectURL: types.InheritedValueString{
			Inherited: false,
			Value:     fmt.Sprintf("https://%s", common.Config.Hosts.RcsUiFQDN),
		},
		RemoteConsentRequestEncryptionEnabled: types.InheritedValueBool{
			Inherited: false,
			Value:     false,
		},
		RemoteConsentRequestEncryptionMethod: types.InheritedValueString{
			Inherited: false,
			Value:     "A128GCM",
		},
		RemoteConsentResponseEncryptionAlgorithm: types.InheritedValueString{
			Inherited: false,
			Value:     "RSA-OAEP-256",
		},
		RequestTimeLimit: types.InheritedValueInt{
			Inherited: false,
			Value:     0,
		},
		JwksURI: types.InheritedValueString{
			Inherited: false,
			//Value:     "http://securebanking-openbanking-uk-rcs:8080/api/rcs/consent/jwk_pub",
		},
		Type: types.Type{
			ID:         "RemoteConsentAgent",
			Name:       "OAuth2 Remote Consent Service",
			Collection: true,
		},
		Userpassword: common.Config.Ig.IgRcsSecret,
	}
	path := "/am/json/realms/root/realms/" + common.Config.Identity.AmRealm + "/realm-config/agents/RemoteConsentAgent/" + remoteConsentId

	s := httprest.Client.Put(path, rc, map[string]string{
		"Accept":             "*/*",
		"Connection":         "keep-alive",
		"Accept-API-Version": "protocol=2.0,resource=1.0",
	})

	zap.S().Infow("Remote Consent Service", "statusCode", s)
}

func remoteConsentExists(name string) bool {
	path := "/am/json/realms/root/realms/" + common.Config.Identity.AmRealm + "/realm-config/agents/RemoteConsentAgent?_queryFilter=true&_pageSize=10&_fields=agentgroup"
	consent := &types.AmResult{}
	b, _ := httprest.Client.Get(path, map[string]string{
		"Accept":             "application/json",
		"X-Requested-With":   "ForgeRock Identity Cloud Postman Collection",
		"Accept-Api-Version": "protocol=2.0,resource=1.0",
	})

	err := json.Unmarshal(b, consent)
	if err != nil {
		panic(err)
	}

	return common.Find(name, consent, func(r *types.Result) string {
		return r.ID
	})
}

func CreateSoftwarePublisherAgentOBRI() {
	if softwarePublisherAgentExists(common.Config.Identity.ObriSoftwarePublisherAgent) {
		zap.L().Info("Skipping creation of Software publisher agent")
		return
	}

	zap.L().Info("Creating software publisher agent")
	pa := types.PublisherAgent{
		PublicKeyLocation: types.InheritedValueString{
			Inherited: false,
			Value:     "jwks_uri",
		},
		JwksCacheTimeout: types.InheritedValueInt{
			Inherited: false,
			Value:     3600000,
		},
		SoftwareStatementSigningAlgorithm: types.InheritedValueString{
			Inherited: false,
			Value:     "PS256",
		},
		JwkSet: types.JwkSet{
			Inherited: false,
		},
		Issuer: types.InheritedValueString{
			Inherited: false,
			Value:     "ForgeRock",
		},
		JwkStoreCacheMissCacheTime: types.InheritedValueInt{
			Inherited: false,
			Value:     60000,
		},
		JwksURI: types.InheritedValueString{
			Inherited: false,
			Value:     "https://service.directory.ob.forgerock.financial/api/directory/keys/jwk_uri",
		},
	}
	path := "/am/json/realms/root/realms/" + common.Config.Identity.AmRealm + "/realm-config/agents/SoftwarePublisher/" + common.Config.Identity.ObriSoftwarePublisherAgent
	s := httprest.Client.Put(path, pa, map[string]string{
		"Accept":             "*/*",
		"Connection":         "keep-alive",
		"Accept-API-Version": "protocol=2.0,resource=1.0",
	})

	zap.S().Infow("Software Publisher Agent", "statusCode", s)
}

func CreateSoftwarePublisherAgentTestPublisher() {
	if softwarePublisherAgentExists(common.Config.Identity.TestSoftwarePublisherAgent) {
		zap.L().Info("Skipping creation of Software publisher agent")
		return
	}

	zap.L().Info("Creating software publisher agent")
	pa := types.PublisherAgent{
		Userpassword: common.Config.Ig.IgSsaSecret,
		PublicKeyLocation: types.InheritedValueString{
			Inherited: false,
			Value:     "jwks_uri",
		},
		JwksCacheTimeout: types.InheritedValueInt{
			Inherited: false,
			Value:     3600000,
		},
		SoftwareStatementSigningAlgorithm: types.InheritedValueString{
			Inherited: false,
			Value:     "HS256",
		},
		JwkSet: types.JwkSet{
			Inherited: false,
		},
		Issuer: types.InheritedValueString{
			Inherited: false,
			Value:     "test-publisher",
		},
		JwkStoreCacheMissCacheTime: types.InheritedValueInt{
			Inherited: false,
			Value:     60000,
		},
		JwksURI: types.InheritedValueString{
			Inherited: false,
		},
	}
	path := "/am/json/realms/root/realms/" + common.Config.Identity.AmRealm + "/realm-config/agents/SoftwarePublisher/" + common.Config.Identity.TestSoftwarePublisherAgent
	s := httprest.Client.Put(path, pa, map[string]string{
		"Accept":             "*/*",
		"Connection":         "keep-alive",
		"Accept-API-Version": "protocol=2.0,resource=1.0",
	})

	zap.S().Infow("Software Publisher Agent", "statusCode", s)
}

func softwarePublisherAgentExists(name string) bool {
	path := "/am/json/realms/root/realms/" + common.Config.Identity.AmRealm + "/realm-config/agents/SoftwarePublisher?_queryFilter=true&_pageSize=10&_fields=agentgroup"
	agent := &types.AmResult{}
	b, _ := httprest.Client.Get(path, map[string]string{
		"Accept":             "application/json",
		"X-Requested-With":   "ForgeRock Identity Cloud Postman Collection",
		"Accept-Api-Version": "protocol=2.0,resource=1.0",
	})

	err := json.Unmarshal(b, agent)
	if err != nil {
		panic(err)
	}

	return common.Find(name, agent, func(r *types.Result) string {
		return r.ID
	})
}

// CreateOIDCClaimsScript -
func CreateOIDCClaimsScript(cookie *http.Cookie) string {

	zap.L().Info("Creating OIDC claims script")
	b, err := ioutil.ReadFile(common.Config.Environment.Paths.ConfigSecureBanking + "oidc-script.json")
	if err != nil {
		panic(err)
	}

	path := fmt.Sprintf("https://%s/am/json/"+common.Config.Identity.AmRealm+"/scripts/?_action=create", common.Config.Hosts.IdentityPlatformFQDN)

	claimsScript := &types.RequestScript{}

	err = json.Unmarshal(b, claimsScript)
	if err != nil {
		panic(err)
	}

	id := httprest.GetScriptIdByName(claimsScript.Name)
	if id != "" {
		zap.L().Info("Script exists")
		return id
	}

	resp, err := restClient.R().
		SetHeader("Accept", "*/*").
		SetHeader("Content-Type", "application/json").
		SetHeader("Connection", "keep-alive").
		SetHeader("Accept-API-Version", "protocol=2.0,resource=1.0").
		SetContentLength(true).
		SetCookie(cookie).
		SetResult(claimsScript).
		SetBody(b).
		Post(path)

	common.RaiseForStatus(err, resp.Error(), resp.StatusCode())

	zap.S().Infow("OIDC claims script", "statusCode", resp.StatusCode(), "claimsScriptID", claimsScript.ID, "createdBy", claimsScript.CreatedBy)
	return claimsScript.ID
}

// UpdateOAuth2Provider - update the oauth 2 provider, must supply the claimScript ID
func UpdateOAuth2Provider(claimsScriptID string) {
	zap.S().Info("UpdateOAuth2Provider() Creating OAuth2Provider service in the " + common.Config.Identity.AmRealm + " realm")

	oauth2Provider := &types.OAuth2Provider{}
	err := common.Unmarshal(common.Config.Environment.Paths.ConfigSecureBanking+"oauth2provider-update.json", &common.Config, oauth2Provider)
	if err != nil {
		panic(err)
	}

	if oauth2ProviderExists(oauth2Provider.Type.ID) {
		zap.L().Info("UpdateOAuth2Provider() OAuth2 provider exists")
		return
	}

	oauth2Provider.PluginsConfig.OidcClaimsScript = claimsScriptID
	zap.S().Infow("UpdateOAuth2Provider() Updating OAuth2 provider", "claimScriptId", claimsScriptID)
	path := "/am/json/" + common.Config.Identity.AmRealm + "/realm-config/services/oauth-oidc"
	zap.S().Info("UpdateOAuth2Provider() Updating OAuth2Provider via the following path {}", path)
	s := httprest.Client.Put(path, oauth2Provider, map[string]string{
		"Accept":           "*/*",
		"Content-Type":     "application/json",
		"Connection":       "keep-alive",
		"X-Requested-With": "ForgeRock Identity Cloud Postman Collection",
	})

	zap.S().Infow("UpdateOAuth2Provider() OAuth2 provider", "statusCode", s)
}

func oauth2ProviderExists(id string) bool {
	path := "/am/json/realms/root/realms/" + common.Config.Identity.AmRealm + "/realm-config/services?_queryFilter=true"
	r := &types.AmResult{}
	b, _ := httprest.Client.Get(path, map[string]string{
		"Accept":             "application/json",
		"X-Requested-With":   "ForgeRock Identity Cloud Postman Collection",
		"Accept-Api-Version": "protocol=1.0,resource=1.0",
	})

	err := json.Unmarshal(b, r)
	if err != nil {
		panic(err)
	}

	return common.Find(id, r, func(r *types.Result) string {
		return r.ID
	})
}

func CreateBaseURLSourceService(cookie *http.Cookie) {
	zap.S().Info("Creating BaseURLSource service in the " + common.Config.Identity.AmRealm + " realm")

	s := &types.Source{}
	err := common.Unmarshal(common.Config.Environment.Paths.ConfigSecureBanking+"create-base-url-source.json", &common.Config, s)
	if err != nil {
		panic(err)
	}
	path := fmt.Sprintf("https://%s/am/json/realms/root/realms/"+common.Config.Identity.AmRealm+"/realm-config/services/baseurl?_action=create",
		common.Config.Hosts.IdentityPlatformFQDN)
	resp, err := restClient.R().
		SetHeader("Accept", "application/json").
		SetHeader("Accept-API-Version", "protocol=1.0,resource=1.0").
		SetHeader("Content-Type", "application/json").
		SetContentLength(true).
		SetCookie(cookie).
		SetBody(s).
		Post(path)

	zap.S().Info("resp is " + resp.String())
	if resp != nil && resp.StatusCode() == 409 {
		zap.S().Info("Did not create BaseURLSource service in " + common.Config.Identity.AmRealm + " realm. It already exists.")
	} else {
		common.RaiseForStatus(err, resp.Error(), resp.StatusCode())
		zap.S().Info("Created Base URL Service in AM's " + common.Config.Identity.AmRealm + " realm")
	}
}
