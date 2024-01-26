package config

import (
	"encoding/json"
	"os"

	log "github.com/sirupsen/logrus"
)

type Config struct {
	SrcDirectory    string `json:"sourceDirectory"`
	TargetDirectory string `json:"ReportDirectory"`
	MaxGoRoutines   int    `json:"maxGoRoutines"`
}

type ConfigProcessor interface {
	LoadConfig(path string) (Config, error)
}

func LoadConfig(path string) (config *Config, err error) {
	if path == "" {
		path = "../config.json"
		log.Info("No config path provided, using default path: ", path)
	}
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	config = &Config{}
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
