package config

import (
	"fmt"
	"os"

	"go.yaml.in/yaml/v3"
)

type Config struct {
	Hall    Hall    `yaml:"hall"`
	Kitchen Kitchen `yaml:"kitchen"`
	App     App     `yaml:"app"`
}

type Hall struct {
	MaxCapacity int `yaml:"max_capacity"`
}

type Kitchen struct {
	DishCountToSetProduct int `yaml:"dish_count_to_set_product"`
}

type App struct {
	PathToJSON string `yaml:"path_to_json"`
}

func LoadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML: %w", err)
	}

	return &config, nil
}
