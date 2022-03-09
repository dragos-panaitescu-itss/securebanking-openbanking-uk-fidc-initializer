package am

import (
	"io/ioutil"

	"github.com/secureBankingAccessToolkit/securebanking-openbanking-uk-fidc-initialiszer/common"
	"go.uber.org/zap"
)

// ApplyAmAuthenticationConfig will attempt to create the username/password nodes
func ApplyAmAuthenticationConfig() {
	CreateCaUsernameNode()
	CreateCaPasswordNode()
	CreateCa()
	CreateScaUsernameNode()
	CreateScaPasswordNode()
	CreateSca()
}

func CreateCaUsernameNode() {
	zap.L().Info("Creating CA Username Node")
	b, err := ioutil.ReadFile(common.IamDirectoryPath() + "am-authentication-config/" + "ca-username-node.json")
	if err != nil {
		panic(err)
	}

	zap.S().Infow("CA username node", "body", string(b))
	path := "/am/json/realms/root/realms/alpha/realm-config/authentication/authenticationtrees/nodes/UsernameCollectorNode/ada9ef86-d550-4591-b9dc-5751e7adbb62"
	status := common.Client.Put(path, b, map[string]string{
		"Accept":             "*/*",
		"Content-Type":       "application/json",
		"Accept-Api-Version": "protocol=2.0, resource=1.0",
	})

	zap.S().Infow("CA Node Username", "statusCode", status)
}

func CreateCaPasswordNode() {
	zap.L().Info("Creating CA Password Node")
	b, err := ioutil.ReadFile(common.IamDirectoryPath() + "am-authentication-config/" + "ca-password-node.json")
	if err != nil {
		panic(err)
	}

	zap.S().Infow("CA Password node", "body", string(b))
	path := "/am/json/realms/root/realms/alpha/realm-config/authentication/authenticationtrees/nodes/PasswordCollectorNode/1db869b1-09de-4a8e-b340-e0563891c3bf"
	status := common.Client.Put(path, b, map[string]string{
		"Accept":             "*/*",
		"Content-Type":       "application/json",
		"Accept-Api-Version": "protocol=2.0, resource=1.0",
	})

	zap.S().Infow("CA Node Password", "statusCode", status)
}

func CreateCa() {
	zap.L().Info("Creating login tree CA")
	b, err := ioutil.ReadFile(common.IamDirectoryPath() + "am-authentication-config/" + "ca.json")
	if err != nil {
		panic(err)
	}

	zap.S().Infow("Login tree CA", "body", string(b))
	path := "/am/json/realms/root/realms/alpha/realm-config/authentication/authenticationtrees/trees/PSD2CustomerAuthentication"
	status := common.Client.Put(path, b, map[string]string{
		"Accept":             "*/*",
		"Content-Type":       "application/json",
		"Accept-Api-Version": "protocol=2.0, resource=1.0",
	})

	zap.S().Infow("CA", "statusCode", status)
}

func CreateScaUsernameNode() {
	zap.L().Info("Creating SCA Username Node")
	b, err := ioutil.ReadFile(common.IamDirectoryPath() + "am-authentication-config/" + "sca-username-node.json")
	if err != nil {
		panic(err)
	}

	zap.S().Infow("SCA username node", "body", string(b))
	path := "/am/json/realms/root/realms/alpha/realm-config/authentication/authenticationtrees/nodes/UsernameCollectorNode/ee0efdc1-9fba-4323-95ef-ec468f6ad30c"
	status := common.Client.Put(path, b, map[string]string{
		"Accept":             "*/*",
		"Content-Type":       "application/json",
		"Accept-Api-Version": "protocol=2.0, resource=1.0",
	})

	zap.S().Infow("SCA Node Username", "statusCode", status)
}

func CreateScaPasswordNode() {
	zap.L().Info("Creating SCA Password Node")
	b, err := ioutil.ReadFile(common.IamDirectoryPath() + "am-authentication-config/" + "sca-password-node.json")
	if err != nil {
		panic(err)
	}

	zap.S().Infow("SCA Password node", "body", string(b))
	path := "/am/json/realms/root/realms/alpha/realm-config/authentication/authenticationtrees/nodes/PasswordCollectorNode/4785b3c1-5dc9-4883-b01e-2f1b6bfda50e"
	status := common.Client.Put(path, b, map[string]string{
		"Accept":             "*/*",
		"Content-Type":       "application/json",
		"Accept-Api-Version": "protocol=2.0, resource=1.0",
	})

	zap.S().Infow("SCA Node Password", "statusCode", status)
}

func CreateSca() {
	zap.L().Info("Creating login tree SCA")
	b, err := ioutil.ReadFile(common.IamDirectoryPath() + "am-authentication-config/" + "sca.json")
	if err != nil {
		panic(err)
	}

	zap.S().Infow("Login tree SCA", "body", string(b))
	path := "/am/json/realms/root/realms/alpha/realm-config/authentication/authenticationtrees/trees/PSD2SecureCustomerAuthentication"
	status := common.Client.Put(path, b, map[string]string{
		"Accept":             "*/*",
		"Content-Type":       "application/json",
		"Accept-Api-Version": "protocol=2.0, resource=1.0",
	})

	zap.S().Infow("SCA", "statusCode", status)
}
