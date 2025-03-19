package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	DBHost       string `mapstructure:"DB_HOST"`
	DBPort       string `mapstructure:"DB_PORT"`
	DBUser       string `mapstructure:"DB_USER"`
	DBPassword   string `mapstructure:"DB_PASSWORD"`
	DBName       string `mapstructure:"DB_NAME"`
	JWTSecret    string `mapstructure:"JWT_SECRET"`
	KafkaBrokers string `mapstructure:"KAFKA_BROKERS"`
	KafkaTopic   string `mapstructure:"KAFKA_TOPIC"`
	ServerPort   string `mapstructure:"SERVER_PORT"`
}

func Load(path string) (config Config, err error) {
	viper.SetDefault("SERVER_PORT", ":8080")

	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	// Automatically override values with environment variables
	viper.AutomaticEnv()

	// Read the .env file
	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError

		if errors.As(err, &configFileNotFoundError) {
			log.Printf(".env file not found, using environment variables")
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return config, nil
}
