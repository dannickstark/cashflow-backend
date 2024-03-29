package config

import (
	"cashflow/backend/utils"
	"fmt"

	"github.com/spf13/viper"
)

type Configuration struct {
	Port int `mapstructure:"PORT"`
}

func LoadConfig(configPath, configName, configType string) (*Configuration, error) {
	var config *Configuration

	absPath := utils.GetAbsolutePath(configPath)

	viper.AddConfigPath(absPath)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("could not read the config file: %v", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal: %v", err)
	}

	return config, nil
}
