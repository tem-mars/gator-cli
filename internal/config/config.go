package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

// Config is the structure that represents the data in the JSON file
type Config struct {
	DBURL          string `json:"db_url"`
	CurrentUserName string `json:"current_user_name,omitempty"`
}

// getConfigFilePath returns the full path of the config file
func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, configFileName), nil
}

// Read reads the config file and converts it to a Config struct
func Read() (Config, error) {
	var cfg Config

	configPath, err := getConfigFilePath()
	if err != nil {
		return cfg, err
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return cfg, err
	}

	err = json.Unmarshal(data, &cfg)
	return cfg, err
}

// write writes the Config struct to the config file
func write(cfg Config) error {
	configPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0644)
}

// SetUser sets the current user name and saves it to the config file
func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username
	return write(*c)
} 