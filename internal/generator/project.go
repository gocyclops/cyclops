package generator

import (
	"fmt"
	"os"
	"path/filepath"
)

type Project struct {
	Name      string
	Framework string
	Features  map[string]bool
	RootDir   string
}

var baseDirectories = []string{
	"controllers",
	"database",
	"models",
	"repository",
	"routes",
	"migrations",
	"utils",
}

var featureDirectories = map[string]string{
	"redis": "myredis",
	"s3":    "mys3",
	"mail":  "mail",
}

func (p *Project) createDirectories() error {
	p.RootDir = filepath.Clean(p.Name)
	if err := os.MkdirAll(p.RootDir, 0755); err != nil {
		return fmt.Errorf("failed to create director: %w", err)
	}

	for _, dir := range baseDirectories {
		dirPath := filepath.Join(p.RootDir, dir)
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	for feature, enabled := range p.Features {
		if enabled {
			if dirName, exists := featureDirectories[feature]; exists {
				dirPath := filepath.Join(p.RootDir, dirName)
				if err := os.MkdirAll(dirPath, 0755); err != nil {
					return fmt.Errorf("failed to create feature directory %s: %w", dirName, err)
				}
			}
		}
	}

	return nil
}

func (p *Project) CreateDirectory(path string) error {
	fullPath := filepath.Join(p.RootDir, path)
	if err := os.MkdirAll(fullPath, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", path, err)
	}
	return nil
}

func (p *Project) generateBaseFiles() error {
	return nil
}

func (p *Project) generateFrameworkFiles() error {
	return nil
}

func (p *Project) generateFeatureFiles() error {
	return nil
}

func (p *Project) Generate() error {
	if err := p.createDirectories(); err != nil {
		return fmt.Errorf("failed to create project directories: %w", err)
	}

	if err := p.generateBaseFiles(); err != nil {
		return fmt.Errorf("failed to generate base files: %w", err)
	}

	if err := p.generateFrameworkFiles(); err != nil {
		return fmt.Errorf("failed to generate framework files: %w", err)
	}

	if err := p.generateFeatureFiles(); err != nil {
		return fmt.Errorf("failed to generate feature files: %w", err)
	}

	return nil
}
