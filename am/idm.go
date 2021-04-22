package am

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// ManagedObjectExists - checks if a managed object exists, must supply the object name
func ManagedObjectExists(objectName string) bool {
	path := "/openidm/config/managed"
	result := &OBManagedObjects{}
	b := Client.Get(path, map[string]string{
		"Accept":           "application/json",
		"X-Requested-With": "ForgeRock Identity Cloud Postman Collection",
	})

	err := json.Unmarshal(b, result)
	if err != nil {
		panic(err)
	}

	for _, o := range result.Objects {
		zap.S().Infow("checking", "object", o)
		if strings.Contains(o.Name, objectName) {
			zap.L().Debug("ManagedObject " + objectName + " found")
			return true
		}
	}
	return false
}

// OBManagedObjects model
type OBManagedObjects struct {
	ID      string `json:"_id"`
	Objects []struct {
		Name string `json:"name"`
	} `json:"objects"`
}

// AddOBManagedObjects -
func AddOBManagedObjects() {
	zap.L().Debug("Adding OB managed objects")
	b, err := ioutil.ReadFile(viper.GetString("REQUEST_BODY_PATH") + "ob-managed-objects.json")
	if err != nil {
		panic(err)
	}

	path := "/openidm/config/managed"
	s := Client.Patch(path, b, map[string]string{
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
	s := Client.Patch(path, b, map[string]string{
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
	s := Client.Put(path, b, map[string]string{
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
	s := Client.Patch(path, b, map[string]string{
		"Accept":       "*/*",
		"Content-Type": "application/json",
		"Connection":   "keep-alive",
	})

	zap.S().Infow("User created", "statusCode", s)
}
