package am

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/secureBankingAccessToolkit/securebanking-openbanking-uk-fidc-initialiszer/common"
	"go.uber.org/zap"
)

//ObjectNames - retrieve filenames from a path. the .json extension will be trimed and
//  a list of filenames will be returned
func ObjectNames(relativePath string) []string {
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
func MissingObjects(objectNames []string) []string {
	path := "/openidm/config/managed"
	result := &OBManagedObjects{}
	b, _ := Client.Get(path, map[string]string{
		"Accept":           "application/json",
		"X-Requested-With": "ForgeRock Identity Cloud Postman Collection",
	})

	err := json.Unmarshal(b, result)
	if err != nil {
		panic(err)
	}

	var missingObjects []string = objectNames
	for _, o := range result.Objects {
		for i, objectName := range missingObjects {
			zap.S().Debugw("checking", "object", o)
			if strings.Contains(o.Name, objectName) {
				zap.S().Infow("ManagedObject found", "name", objectName)
				missingObjects = append(missingObjects[:i], missingObjects[i+1:]...)
				break
			}
		}
	}
	return missingObjects
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
	path := managedObjectsObDirectory()
	managedObjectFilenames := ObjectNames(path)
	missingObjects := MissingObjects(managedObjectFilenames)

	for _, o := range missingObjects {
		AddManagedObject(o, path)
	}
}

func managedObjectsObDirectory() string {
	return common.ManagedObjectsDirectoryPath() + "openbanking/"
}

// AddManagedObject - Will add a managed object in IDM. retrieve a filename (minus the suffix) in a supplied directory
//  and apply patch to idm.
func AddManagedObject(name string, objectFolderPath string) {
	b, err := ioutil.ReadFile(objectFolderPath + name + ".json")
	if err != nil {
		panic(err)
	}

	path := "/openidm/config/managed"
	s := Client.Patch(path, b, map[string]string{
		"Accept":       "*/*",
		"Content-Type": "application/json",
		"Connection":   "keep-alive",
	})

	zap.S().Infow("Managed object patched", "statusCode", s, "name", name)
}

func CreateApiJwksEndpoint() {
	zap.L().Debug("Creating API JWKS Endpoint")
	b, err := ioutil.ReadFile(common.IamDirectoryPath() + "create-jwks-endpoint.json")
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
	b, err := ioutil.ReadFile(common.IamDirectoryPath() + "create-user.json")
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
