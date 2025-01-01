package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ProjectName	string	`yaml:"project_name"`
	Framework	string `yaml:"framework"`
	Features	map[string]bool	`yaml:"features"`
}

func Load(path string) (*Config, error) {
	return nil, nil
}

func Save(config *Config, path string) error {
	data, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %W", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}
