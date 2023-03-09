package config

import "github.com/spf13/viper"

func InitConfigs() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config") // Register config file name (no extension)
	viper.SetConfigType("json")   // Look for specific type

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}
