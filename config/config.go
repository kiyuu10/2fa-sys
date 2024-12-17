package config

import (
	"github.com/kiyuu10/2fa-sys/utils"
	"github.com/spf13/viper"
)

func LoadConfig() {
	// configure Viper
	viper.SetConfigName("config") // file config
	viper.SetConfigType("yaml")   // file extension
	viper.AddConfigPath(".")      // path of file

	// read file config
	logger := utils.NewLogger()
	if err := viper.ReadInConfig(); err != nil {
		logger.LogError("Failed to load config", err)
	}
	logger.LogInfo("Config loaded successfully")
}
