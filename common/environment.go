package common

import "github.com/spf13/viper"

func IamDirectoryPath() string {
	return viper.GetString("IAM_DIRECTORY_PATH")
}

func ManagedObjectsDirectoryPath() string {
	return viper.GetString("MANAGED_OBJECTS_DIRECTORY_PATH")
}
