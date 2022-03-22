package securebanking

import (
	"encoding/json"
	"io/ioutil"
	"secure-banking-uk-initializer/pkg/common"
	"secure-banking-uk-initializer/pkg/httprest"
	"strings"

	"go.uber.org/zap"
)

//objectNames - retrieve filenames from a path. the .json extension will be trimed and
//  a list of filenames will be returned
func objectNames(relativePath string) []string {
	files, err := ioutil.ReadDir(relativePath)
	if err != nil {
		zap.L().Fatal(err.Error())
	}

	var names []string
	for _, f := range files {
		name := strings.TrimSuffix(f.Name(), ".json")
		names = append(names, name)
	}
	return names
}

//MissingObjects - return a list of missing managed object names in idm.
//  supply an array of managed object names to query against.
func missingObjects(objectNames []string) []string {
	path := "/openidm/config/managed"
	result := &OBManagedObjects{}
	b, _ := httprest.Client.Get(path, map[string]string{
		"Accept":           "application/json",
		"X-Requested-With": "ForgeRock Identity Cloud Postman Collection",
	})

	err := json.Unmarshal(b, result)
	if err != nil {
		panic(err)
	}

	var missObjects = objectNames
	for _, o := range result.Objects {
		for i, objectName := range missObjects {
			zap.S().Infow("checking", "object", o)
			if strings.Contains(o.Name, objectName) {
				zap.S().Infow("ManagedObject found", "name", objectName)
				missObjects = append(missObjects[:i], missObjects[i+1:]...)
				break
			}
		}
	}
	return missObjects
}

// OBManagedObjects model
type OBManagedObjects struct {
	ID      string `json:"_id"`
	Objects []struct {
		Name string `json:"name"`
	} `json:"objects"`
}

// AddOBManagedObjects - Add managed objects to IDM. This will look for json in the managed objects OB config directory
//  and add them to IDM if they dont already exist.
func AddOBManagedObjects() {
	configPath := common.Config.Environment.Paths.ConfigSecureBanking + "managed-objects/"
	managedObjectFilenames := objectNames(configPath)
	mObjects := missingObjects(managedObjectFilenames)

	for _, o := range mObjects {
		addManagedObject(o, configPath)
	}
}

// AddManagedObject - Will add a managed object in IDM. retrieve a filename (minus the suffix) in a supplied directory
//  and apply patch to idm.
func addManagedObject(name string, objectFolderPath string) {
	b, err := ioutil.ReadFile(objectFolderPath + name + ".json")
	if err != nil {
		panic(err)
	}

	path := "/openidm/config/managed"
	s := httprest.Client.Patch(path, b, map[string]string{
		"Accept":       "*/*",
		"Content-Type": "application/json",
		"Connection":   "keep-alive",
	})

	zap.S().Infow("Managed object patched", "statusCode", s, "name", name)
}

func CreateApiJwksEndpoint() {
	zap.L().Info("Creating API JWKS Endpoint")
	b, err := ioutil.ReadFile(common.Config.Environment.Paths.ConfigSecureBanking + "create-jwks-endpoint.json")
	if err != nil {
		panic(err)
	}

	path := "/openidm/config/endpoint/apiclientjwks"
	s := httprest.Client.Put(path, b, map[string]string{
		"Accept":       "*/*",
		"Content-Type": "application/json",
		"Connection":   "keep-alive",
	})

	zap.S().Infow("JWKS endpoint", "statusCode", s)
}
