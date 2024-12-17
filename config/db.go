package config

import (
	"github.com/kiyuu10/2fa-sys/models"
	"github.com/kiyuu10/2fa-sys/utils"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	logger := utils.NewLogger()
	// read url connection database from config
	dbUrl := viper.GetString("db.url")
	if dbUrl == "" {
		logger.LogError("Database URL is not set in config",
			errors.New("Database URL"))
		return
	}

	var (
		database *gorm.DB
		err      error
	)

	database, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		logger.LogError("Failed to connect database", err)
		return
	}

	DB = database
	logger.LogInfo("Connected to database successfully")

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		logger.LogError("Failed to migrate database", err)
		return
	}
}
