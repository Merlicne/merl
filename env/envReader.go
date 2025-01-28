package env

import (
	"github.com/spf13/viper"
)

var (
	envReader *viper.Viper
)

func NewEnvReader(ConfigPath string) error {

	viper.SetConfigFile(ConfigPath)

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	envReader = viper.GetViper()
	return nil
}

func GetStringValue(key string) string {
	result := envReader.GetString(key)
	return result
}

func SetStringValue(key string, value string) {
	envReader.Set(key, value)
}
