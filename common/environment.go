package common

import "github.com/spf13/viper"

func ConfigDirectoryPath() string {
	return viper.GetString("CONFIG_DIRECTORY_PATH")
}
