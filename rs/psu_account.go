package rs

import (
	"encoding/json"
	"github.com/secureBankingAccessToolkit/securebanking-openbanking-uk-fidc-initialiszer/common"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

// CreatePSU - create the psu user if necessary and always return the userId if exist to populate de user data into RS
func CreatePSU() string {
	exist, userId := PSUIdentityExists(viper.GetString("PSU_USERNAME"))
	if exist {
		zap.S().Infow("Skipping creation of Payment Services User", "userID", userId)
		return userId
	}

	zap.L().Debug("Creating Payment Services User")

	user := &PSU{
		UserName:  viper.GetString("PSU_USERNAME"),
		SN:        "Payment Services User",
		GivenName: "PSU",
		Mail:      "psu@acme.com",
		Password:  viper.GetString("PSU_PASSWORD"),
	}

	path := "/openidm/managed/user/?_action=create"
	body, s := common.Client.Post(path, user, map[string]string{
		"Accept":       "*/*",
		"Content-Type": "application/json",
		"Connection":   "keep-alive",
	})
	userRes := &UserResponse{}
	err := json.Unmarshal(body, userRes)
	if err != nil {
		panic(err)
	}
	zap.S().Infow("PSU created", "Response", userRes, "UserId", userRes.UserId)

	zap.S().Infow("Payment Services User", "statusCode", s)
	return userRes.UserId
}

// PSUIdentityExists will check for psu identities in the alpha realm
func PSUIdentityExists(identity string) (bool, string) {
	filter := "?_queryFilter=uid+eq+%22" + identity + "%22&_fields=username"
	path := "/am/json/realms/root/realms/alpha/users" + filter
	//path := "/am/json/realms/root/realms/alpha/users/" + identity + "?_fields=username"
	serviceIdentityFilter := &common.ResultFilter{}
	b, _ := common.Client.Get(path, map[string]string{
		"Accept":             "application/json",
		"X-Requested-With":   "ForgeRock Identity Cloud Postman Collection",
		"Accept-Api-Version": "protocol=2.1, resource=4.0",
	})

	err := json.Unmarshal(b, serviceIdentityFilter)
	if err != nil {
		panic(err)
	}
	var psuID = ""
	if serviceIdentityFilter.ResultCount > 0 {
		psuID = serviceIdentityFilter.Result[0].ID
	}
	return serviceIdentityFilter.ResultCount > 0, psuID
}

// PopulateRSData -
func PopulateRSData(userId string) {
	if userId == "" {
		zap.L().Info("Skipping populate PSU Data to RS service")
		return
	}
	// need to be the same namespaces set in https://raw.githubusercontent.com/ForgeCloud/sbat-infra/master/NAMESPACES.md
	namespaces := []string{"dev", "nightly", "jorgesanchezperez", "bohocode", "mariantiris", "andra-racovita", "christian-brindley"}
	for index, namespace := range namespaces {
		zap.S().Infow("*", "index", index, "namespace", namespace)
		path := "https://rs." + namespace + ".forgerock.financial/admin/data/user/has-data?userId=" + userId
		if mustPopulateUserData(path, namespace) {
			zap.S().Infow("Populate with RS Data the Payment Services User with the userId: " + userId)
			params := "userId=" + userId + "&username=" + userId + "&profile=random"
			path := "https://rs." + namespace + ".forgerock.financial/admin/fake-data/generate?" + params
			s := common.Client.PostRS(path, map[string]string{
				"Accept":     "*/*",
				"Connection": "keep-alive",
			})
			zap.S().Infow("Populate RS Data response", "namespace", namespace, "statusCode", s)
		}
	}
}

// mustPopulateUserData check is the user has data and if the environment is initialised, return true/false
func mustPopulateUserData(path string, namespace string) bool {
	b, state := common.Client.GetRS(path, map[string]string{
		"Accept": "*/*",
	})
	if state != http.StatusOK {
		zap.S().Infow("No environment initialised", "namespace", namespace, "request status", state)
		return false
	}
	value := string(b)
	zap.S().Infow("User has data?", "namespace", namespace, "result", value)
	result, err := strconv.ParseBool(value)
	if err != nil {
		panic(err.Error())
	}
	return !result
}
