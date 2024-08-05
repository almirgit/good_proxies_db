package config

import (
	"fmt"
	"good_proxies_go_ai/shared"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	//"log"
	"os"
	"strings"
)

// // Config holds the overall configuration for the application
// type Config struct {
// 	Database        DatabaseConfig `yaml:"database"`
// 	CheckURLEndPoint string        `yaml:"CheckURLEndPoint"`
// }

// // DatabaseConfig holds the database-related configuration
// type DatabaseConfig struct {
// 	Username string `yaml:"username"`
// 	Password string `yaml:"password"`
// 	Host     string `yaml:"host"`
// 	SSLMode  string `yaml:"sslmode"`
// 	DBName   string `yaml:"dbname"`
// }

// LoadConfig reads the YAML configuration file and returns a Config struct
func LoadConfig(filename string) (*shared.Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config shared.Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	setDefaults(&config)
	overrideConfigWithEnvVars(&config)

	if err := validateConfig(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

// setDefaults assigns default values to config fields if they are not set
func setDefaults(config *shared.Config) {
	if config.Database.SSLMode == "" {
		config.Database.SSLMode = "disable"
	}
}

// overrideConfigWithEnvVars allows environment variables to override config values
func overrideConfigWithEnvVars(config *shared.Config) {
	if val, exists := os.LookupEnv("DB_USERNAME"); exists {
		config.Database.Username = val
	}
	if val, exists := os.LookupEnv("DB_PASSWORD"); exists {
		config.Database.Password = val
	}
	if val, exists := os.LookupEnv("CHECK_URL_ENDPOINT"); exists {
		config.CheckURLEndPoint = val
	}
}

// validateConfig checks that all required configuration fields are set
func validateConfig(config *shared.Config) error {
	if strings.TrimSpace(config.Database.Username) == "" {
		return fmt.Errorf("database username is required")
	}
	if strings.TrimSpace(config.Database.Password) == "" {
		return fmt.Errorf("database password is required")
	}
	if strings.TrimSpace(config.CheckURLEndPoint) == "" {
		return fmt.Errorf("CheckURLEndPoint is required")
	}
	return nil
}
