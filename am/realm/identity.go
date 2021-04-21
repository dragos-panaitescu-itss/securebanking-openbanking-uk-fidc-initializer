package realm

import (
	"github.com/secureBankingAcceleratorToolkit/securebanking-openbanking-uk-fidc-initialiszer/am"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type ServiceIdentity struct {
	Result []struct {
		ID             string   `json:"_id"`
		Rev            string   `json:"_rev"`
		Cn             []string `json:"cn"`
		Mail           []string `json:"mail"`
		Username       string   `json:"username"`
		Inetuserstatus []string `json:"inetUserStatus"`
	} `json:"result"`
	Resultcount             int         `json:"resultCount"`
	Pagedresultscookie      interface{} `json:"pagedResultsCookie"`
	Totalpagedresultspolicy string      `json:"totalPagedResultsPolicy"`
	Totalpagedresults       int         `json:"totalPagedResults"`
	Remainingpagedresults   int         `json:"remainingPagedResults"`
}

// ServiceIdentityExists will check for service identities in the alpha realm
//   When CDK is removed, these entities might still be persisted. this gives us
//   an indication that we do not need to initialize the environment
func ServiceIdentityExists(identity string) bool {
	path := "https://" + viper.GetString("IAM_FQDN") + "/am/json/realms/root/realms/alpha/users?_queryFilter=true&_pageSize=10&_fields=cn,mail,username,inetUserStatus"
	serviceIdentity := &ServiceIdentity{}
	am.Client.Get(path, map[string]string{
		"Accept":             "application/json",
		"X-Requested-With":   "ForgeRock Identity Cloud Postman Collection",
		"Accept-Api-Version": "protocol=2.1, resource=4.0",
	}, serviceIdentity)

	for _, r := range serviceIdentity.Result {
		if r.Username == identity {
			zap.L().Info("Identity " + identity + " exists")
			return true
		}
	}
	return false
}
