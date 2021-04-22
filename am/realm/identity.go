package realm

import (
	"encoding/json"

	"github.com/secureBankingAcceleratorToolkit/securebanking-openbanking-uk-fidc-initialiszer/am"
	"go.uber.org/zap"
)

type ServiceIdentity struct {
	Result                  []Result    `json:"result"`
	Resultcount             int         `json:"resultCount"`
	Pagedresultscookie      interface{} `json:"pagedResultsCookie"`
	Totalpagedresultspolicy string      `json:"totalPagedResultsPolicy"`
	Totalpagedresults       int         `json:"totalPagedResults"`
	Remainingpagedresults   int         `json:"remainingPagedResults"`
}

type Result struct {
	ID             string   `json:"_id"`
	Rev            string   `json:"_rev"`
	Cn             []string `json:"cn"`
	Mail           []string `json:"mail"`
	Username       string   `json:"username"`
	Inetuserstatus []string `json:"inetUserStatus"`
}

// ServiceIdentityExists will check for service identities in the alpha realm
//   When CDK is removed, these entities might still be persisted. this gives us
//   an indication that we do not need to initialize the environment
func ServiceIdentityExists(identity string) bool {
	path := "/am/json/realms/root/realms/alpha/users?_queryFilter=true&_pageSize=10&_fields=cn,mail,username,inetUserStatus"
	serviceIdentity := &ServiceIdentity{}
	b := am.Client.Get(path, map[string]string{
		"Accept":             "application/json",
		"X-Requested-With":   "ForgeRock Identity Cloud Postman Collection",
		"Accept-Api-Version": "protocol=2.1, resource=4.0",
	})

	err := json.Unmarshal(b, serviceIdentity)
	if err != nil {
		panic(err)
	}

	for _, r := range serviceIdentity.Result {
		if r.ID == identity {
			zap.L().Info("Identity " + identity + " exists")
			return true
		}
	}
	return false
}
