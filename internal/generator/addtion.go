package generator

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type Config struct {
	ProjectName string          `yaml:"projectName"`
	Framework   string          `yaml:"framework"`
	Features    map[string]bool `yaml:"features"`
}

type FeatureAdder struct {
	RootDir string
	Config  Config
}

func NewFeatureAdder(rootDir string) (*FeatureAdder, error) {
	configPath := filepath.Join(rootDir, "config.yaml")
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config.yaml: %w", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config.yaml: %w", err)
	}

	return &FeatureAdder{
		RootDir: rootDir,
		Config:  config,
	}, nil
}

func (f *FeatureAdder) AddFeature(feature string) error {
	// Check if feature is already enabled
	if f.Config.Features[feature] {
		return fmt.Errorf("feature '%s' is already enabled", feature)
	}

	// Create Project instance for using existing methods
	p := &Project{
		Name:     f.Config.ProjectName,
		RootDir:  f.RootDir,
		Features: map[string]bool{feature: true},
	}

	// Create feature directory
	switch feature {
	case "s3":
		if err := p.CreateDirectory("mys3"); err != nil {
			return err
		}
	case "mail":
		if err := p.CreateDirectory("mail"); err != nil {
			return err
		}
	case "redis":
		if err := p.CreateDirectory("myredis"); err != nil {
			return err
		}
	case "auth":
		if err := p.CreateDirectory("utils"); err != nil {
			return err
		}
	}

	// Generate feature files
	featureFiles := map[string]string{}
	switch feature {
	case "s3":
		featureFiles["features/s3/s3.go.tmpl"] = "mys3/s3.go"
	case "mail":
		featureFiles["features/mail/mail.go.tmpl"] = "mail/mail.go"
	case "redis":
		featureFiles["features/redis/redis.go.tmpl"] = "myredis/redis.go"
	case "auth":
		featureFiles["features/utils/jwt.go.tmpl"] = "utils/jwt.go"
	}

	if err := p.generateFiles(featureFiles); err != nil {
		return fmt.Errorf("failed to generate feature files: %w", err)
	}

	// Update config.yaml
	f.Config.Features[feature] = true
	configData, err := yaml.Marshal(f.Config)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	configPath := filepath.Join(f.RootDir, "config.yaml")
	if err := os.WriteFile(configPath, configData, 0644); err != nil {
		return fmt.Errorf("failed to update config.yaml: %w", err)
	}

	return nil
}
