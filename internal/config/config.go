package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const (
	CFG_FILE_NAME = ".gatorconfig.json"
)

type Config struct {
	DBURL string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (*Config, error) {
	cfgPath, err := getConfigFilePath()
	if err != nil {
		return &Config{}, err
	}

	file, err := os.Open(cfgPath)
	if err != nil {
		return &Config{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var cfg Config
	if err := decoder.Decode(&cfg); err != nil {
		return &Config{}, err
	}

	return &cfg, nil
}

func SetUser(userName string, cfg *Config)  error {
	cfg.CurrentUserName = userName
	if err := write(cfg); err != nil {
		return err
	}
	return nil
}

func write(cfg *Config) error {

	filepath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	if err := encoder.Encode(&cfg); err != nil {
		return err
	}
	return nil
}

func getConfigFilePath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	cfgPath := filepath.Join(homePath, CFG_FILE_NAME)
	return cfgPath, nil
}
