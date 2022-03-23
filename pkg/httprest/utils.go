package httprest

import (
	"encoding/json"
	"net/url"
	"secure-banking-uk-initializer/pkg/common"
	"secure-banking-uk-initializer/pkg/types"
)

// AlphaClientsExist - Will return true if clients exist in the alpha realm.
func AlphaClientsExist(clientName string) bool {
	path := "/am/json/realms/root/realms/alpha/realm-config/agents/OAuth2Client?_queryFilter=true&_pageSize=10&_fields=coreOAuth2ClientConfig/status,coreOAuth2ClientConfig/agentgroup"
	result := &types.AmResult{}
	b, _ := Client.Get(path, map[string]string{
		"Accept":             "application/json",
		"X-Requested-With":   "ForgeRock Identity Cloud Postman Collection",
		"Accept-Api-Version": "protocol=2.0,resource=1.0",
	})

	err := json.Unmarshal(b, result)
	if err != nil {
		panic(err)
	}

	return common.Find(clientName, result, func(r *types.Result) string {
		return r.ID
	})
}

func GetScriptIdByName(name string) string {
	path := "/am/json/realms/alpha/scripts?_prettyPrint=true&_queryFilter=name+eq+%22" + url.QueryEscape(name) + "%22&_sortKeys=name"
	//path := "/am/json/alpha/scripts?_pageSize=20&_sortKeys=name&_queryFilter=true&_pagedResultsOffset=0"
	consent := &types.AmResult{}
	b, _ := Client.Get(path, map[string]string{
		"Accept":             "application/json",
		"X-Requested-With":   "ForgeRock Identity Cloud Postman Collection",
		"Accept-Api-Version": "protocol=1.0,resource=1.0",
	})

	err := json.Unmarshal(b, consent)
	if err != nil {
		panic(err)
	}

	return common.FindIdByName(name, consent, func(r *types.Result) string {
		return r.Name
	})
}

func PolicySetExists(name string) bool {
	path := "/am/json/alpha/applications?_pageSize=20&_sortKeys=name&_queryFilter=name+eq+%22%5E(%3F!sunAMDelegationService%24).*%22&_pagedResultsOffset=0"
	serviceIdentity := &types.AmResult{}
	b, _ := Client.Get(path, map[string]string{
		"Accept":             "application/json",
		"X-Requested-With":   "ForgeRock Identity Cloud Postman Collection",
		"Accept-Api-Version": "protocol=1.0,resource=2.0",
	})

	err := json.Unmarshal(b, serviceIdentity)
	if err != nil {
		panic(err)
	}

	return common.Find(name, serviceIdentity, func(r *types.Result) string {
		return r.Name
	})
}

func PolicyExists(name string) bool {
	path := "/am/json/alpha/policies?_pageSize=20&_sortKeys=name&_queryFilter=applicationName+eq+%22Open%20Banking%22&_pagedResultsOffset=0"
	serviceIdentity := &types.AmResult{}
	b, _ := Client.Get(path, map[string]string{
		"Accept":             "application/json",
		"X-Requested-With":   "ForgeRock Identity Cloud Postman Collection",
		"Accept-Api-Version": "protocol=1.0,resource=2.0",
	})

	err := json.Unmarshal(b, serviceIdentity)
	if err != nil {
		panic(err)
	}

	return common.Find(name, serviceIdentity, func(r *types.Result) string {
		return r.Name
	})
}

// ServiceIdentityExists will check for service identities in the alpha realm
//   When CDK is removed, these entities might still be persisted. this gives us
//   an indication that we do not need to initialize the environment
func ServiceIdentityExists(identity string) bool {
	filter := "?_queryFilter=uid+eq+%22" + url.QueryEscape(identity) + "%22&_fields=username"
	path := "/am/json/realms/root/realms/alpha/users" + filter
	//path := "/am/json/realms/root/realms/alpha/users/" + identity + "?_fields=username"
	serviceIdentityFilter := &types.ResultFilter{}
	b, _ := Client.Get(path, map[string]string{
		"Accept":             "application/json",
		"X-Requested-With":   "ForgeRock Identity Cloud Postman Collection",
		"Accept-Api-Version": "protocol=2.1, resource=4.0",
	})

	err := json.Unmarshal(b, serviceIdentityFilter)
	if err != nil {
		panic(err)
	}
	return serviceIdentityFilter.ResultCount > 0
}
