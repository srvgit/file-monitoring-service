package config

import (
	"encoding/json"
	"os"
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
