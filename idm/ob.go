package idm

import (
	"io/ioutil"

	"github.com/secureBankingAcceleratorToolkit/securebanking-openbanking-uk-fidc-initialiszer/am"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// AddOBManagedObjects -
func AddOBManagedObjects() {
	zap.L().Debug("Adding OB managed objects")
	b, err := ioutil.ReadFile(viper.GetString("REQUEST_BODY_PATH") + "ob-managed-objects.json")
	if err != nil {
		panic(err)
	}

	path := "/openidm/config/managed"
	s := am.Client.Patch(path, b, map[string]string{
		"Accept":       "*/*",
		"Content-Type": "application/json",
		"Connection":   "keep-alive",
	})

	zap.S().Infow("OpenBanking Managed Objects", "statusCode", s)
}

func AddAdditionalCDKObjects() {
	zap.L().Debug("Adding OB managed objects")
	b, err := ioutil.ReadFile(viper.GetString("REQUEST_BODY_PATH") + "cdk-additional-objects.json")
	if err != nil {
		panic(err)
	}

	path := "/openidm/config/managed"
	s := am.Client.Patch(path, b, map[string]string{
		"Accept":       "*/*",
		"Content-Type": "application/json",
		"Connection":   "keep-alive",
	})

	zap.S().Infow("OpenBanking Managed Objects", "statusCode", s)
}

func CreateApiJwksEndpoint() {
	zap.L().Debug("Adding OB managed objects")
	b, err := ioutil.ReadFile(viper.GetString("REQUEST_BODY_PATH") + "create-jwks-endpoint.json")
	if err != nil {
		panic(err)
	}

	path := "/openidm/config/endpoint/apiclientjwks"
	s := am.Client.Put(path, b, map[string]string{
		"Accept":       "*/*",
		"Content-Type": "application/json",
		"Connection":   "keep-alive",
	})

	zap.S().Infow("JWKS endpoint", "statusCode", s)
}

// CreateUser will create a user that will allow us to create new identities
//    in the alpha realm
func CreateUser() {
	zap.L().Debug("Creating new user")
	b, err := ioutil.ReadFile(viper.GetString("REQUEST_BODY_PATH") + "create-user.json")
	if err != nil {
		panic(err)
	}

	path := "/openidm/config/managed"
	s := am.Client.Patch(path, b, map[string]string{
		"Accept":       "*/*",
		"Content-Type": "application/json",
		"Connection":   "keep-alive",
	})

	zap.S().Infow("User created", "statusCode", s)
}
