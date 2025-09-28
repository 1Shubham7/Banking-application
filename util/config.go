package util

import (
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver             string        `mapstructure:"DB_DRIVER"`
	DBSource             string        `mapstructure:"DB_SOURCE"`
	ServerAddress        string        `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	EmailSenderName      string        `mapstructure:"EMAIL_SENDER_NAME"`
	EmailSenderAddress   string        `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword  string        `mapstructure:"EMAIL_SENDER_PASSWORD"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	// Enable automatic environment variable reading
	viper.AutomaticEnv()

	// Try to read config file, but don't fail if it doesn't exist
	err = viper.ReadInConfig()
	if err != nil {
		// Check if it's a "file not found" error
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found, that's okay - we'll use environment variables
			// Reset the error since we can continue without the file
			err = nil
		} else {
			// Some other error occurred while reading the file
			return
		}
	}

	// Unmarshal the configuration (from file + environment variables)
	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}

	// Fallback: If TOKEN_SYMMETRIC_KEY is empty, try to get it directly from env
	if config.TokenSymmetricKey == "" {
		config.TokenSymmetricKey = os.Getenv("TOKEN_SYMMETRIC_KEY")
	}

	// Fallback: If DB_SOURCE is empty, try to get it directly from env
	if config.DBSource == "" {
		config.DBSource = os.Getenv("DB_SOURCE")
	}

	// Fallback: If SERVER_ADDRESS is empty, try to get it directly from env
	if config.ServerAddress == "" {
		config.ServerAddress = os.Getenv("SERVER_ADDRESS")
	}

	return config, err
}
