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
	zap.L().Debug("Creating CA Username Node")
	b, err := ioutil.ReadFile(common.IamDirectoryPath() + "am-authentication-config/" + "ca-username-node.json")
	if err != nil {
		panic(err)
	}

	zap.S().Debugw("CA username node", "body", string(b))
	path := "/am/json/realms/root/realms/alpha/realm-config/authentication/authenticationtrees/nodes/UsernameCollectorNode/ada9ef86-d550-4591-b9dc-5751e7adbb62"
	status := Client.Put(path, b, map[string]string{
		"Accept":             "*/*",
		"Content-Type":       "application/json",
		"Accept-Api-Version": "protocol=2.0, resource=1.0",
	})

	zap.S().Infow("CA Node Username", "statusCode", status)
}

func CreateCaPasswordNode() {
	zap.L().Debug("Creating CA Password Node")
	b, err := ioutil.ReadFile(common.IamDirectoryPath() + "am-authentication-config/" + "ca-password-node.json")
	if err != nil {
		panic(err)
	}

	zap.S().Debugw("CA Password node", "body", string(b))
	path := "/am/json/realms/root/realms/alpha/realm-config/authentication/authenticationtrees/nodes/PasswordCollectorNode/1db869b1-09de-4a8e-b340-e0563891c3bf"
	status := Client.Put(path, b, map[string]string{
		"Accept":             "*/*",
		"Content-Type":       "application/json",
		"Accept-Api-Version": "protocol=2.0, resource=1.0",
	})

	zap.S().Infow("CA Node Password", "statusCode", status)
}

func CreateCa() {
	zap.L().Debug("Creating login tree CA")
	b, err := ioutil.ReadFile(common.IamDirectoryPath() + "am-authentication-config/" + "ca.json")
	if err != nil {
		panic(err)
	}

	zap.S().Debugw("Login tree CA", "body", string(b))
	path := "/am/json/realms/root/realms/alpha/realm-config/authentication/authenticationtrees/trees/PSD2CustomerAuthentication"
	status := Client.Put(path, b, map[string]string{
		"Accept":             "*/*",
		"Content-Type":       "application/json",
		"Accept-Api-Version": "protocol=2.0, resource=1.0",
	})

	zap.S().Infow("CA", "statusCode", status)
}

func CreateScaUsernameNode() {
	zap.L().Debug("Creating SCA Username Node")
	b, err := ioutil.ReadFile(common.IamDirectoryPath() + "am-authentication-config/" + "sca-username-node.json")
	if err != nil {
		panic(err)
	}

	zap.S().Debugw("SCA username node", "body", string(b))
	path := "/am/json/realms/root/realms/alpha/realm-config/authentication/authenticationtrees/nodes/UsernameCollectorNode/ee0efdc1-9fba-4323-95ef-ec468f6ad30c"
	status := Client.Put(path, b, map[string]string{
		"Accept":             "*/*",
		"Content-Type":       "application/json",
		"Accept-Api-Version": "protocol=2.0, resource=1.0",
	})

	zap.S().Infow("SCA Node Username", "statusCode", status)
}

func CreateScaPasswordNode() {
	zap.L().Debug("Creating SCA Password Node")
	b, err := ioutil.ReadFile(common.IamDirectoryPath() + "am-authentication-config/" + "sca-password-node.json")
	if err != nil {
		panic(err)
	}

	zap.S().Debugw("SCA Password node", "body", string(b))
	path := "/am/json/realms/root/realms/alpha/realm-config/authentication/authenticationtrees/nodes/PasswordCollectorNode/4785b3c1-5dc9-4883-b01e-2f1b6bfda50e"
	status := Client.Put(path, b, map[string]string{
		"Accept":             "*/*",
		"Content-Type":       "application/json",
		"Accept-Api-Version": "protocol=2.0, resource=1.0",
	})

	zap.S().Infow("SCA Node Password", "statusCode", status)
}

func CreateSca() {
	zap.L().Debug("Creating login tree SCA")
	b, err := ioutil.ReadFile(common.IamDirectoryPath() + "am-authentication-config/" + "sca.json")
	if err != nil {
		panic(err)
	}

	zap.S().Debugw("Login tree SCA", "body", string(b))
	path := "/am/json/realms/root/realms/alpha/realm-config/authentication/authenticationtrees/trees/PSD2SecureCustomerAuthentication"
	status := Client.Put(path, b, map[string]string{
		"Accept":             "*/*",
		"Content-Type":       "application/json",
		"Accept-Api-Version": "protocol=2.0, resource=1.0",
	})

	zap.S().Infow("SCA", "statusCode", status)
}
