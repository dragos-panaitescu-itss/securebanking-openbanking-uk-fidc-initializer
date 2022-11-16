package securebanking

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
	"net/http"
	"secure-banking-uk-initializer/pkg/common"
)

func CreateCrestResourceFromConfigFile(url string, configFileName string, cookie *http.Cookie) {
	CreateOrUpdateCrestResourceFromConfigFile(resty.MethodPost, url, configFileName, cookie)
}

func UpdateCrestResourceFromConfigFile(url string, configFileName string, cookie *http.Cookie) {
	CreateOrUpdateCrestResourceFromConfigFile(resty.MethodPut, url, configFileName, cookie)
}

// CreateOrUpdateCrestResourceFromConfigFile
// Generic method for creating or updating resources using the FR Crest API
// Accepts a target CREST url to create or update the resource, and the name of a json config file to unmarshall.
//
// This method is suitable when the config is complete with only template value substitution, it will not work when the
// config needs to be edited post unmarshalling e.g. to set an id value to that of a resource created in a previous step.
func CreateOrUpdateCrestResourceFromConfigFile(httpMethod string, url string, configFileName string, cookie *http.Cookie) {
	zap.L().Info("Attempting to create resource using CREST, url: " + url + ", configFileName: " + configFileName)

	var jsonConfig map[string]interface{}
	err := common.Unmarshal(common.Config.Environment.Paths.ConfigSecureBanking+configFileName, &common.Config, &jsonConfig)
	if err != nil {
		zap.S().Fatalw(fmt.Sprintf("Failed to log jsonConfig: %s , error: %v", configFileName, err))
	}

	var responsePayload map[string]interface{}
	resp, err := restClient.R().
		SetHeader("Accept", "*/*").
		SetHeader("Content-Type", "application/json").
		SetHeader("Connection", "keep-alive").
		SetHeader("X-Requested-With", "XMLHttpRequest").
		SetHeader("Accept-API-Version", "protocol=1.0,resource=1.0").
		SetContentLength(true).
		SetCookie(cookie).
		SetBody(jsonConfig).
		SetResult(&responsePayload).
		Execute(httpMethod, url)

	zap.S().Info("resp is " + resp.String())
	if resp != nil && resp.StatusCode() == 409 {
		zap.S().Info("Nothing created, resource already exists for url: " + url + " , configFileName: " + configFileName)
	} else {
		common.RaiseForStatus(err, resp.Error(), resp.StatusCode())
		zap.S().Info("Created resource, _id: ", responsePayload["_id"])
	}
}
