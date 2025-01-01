package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
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

// Create a single directory
func (p *Project) CreateDirectory(path string) error {
	fullPath := filepath.Join(p.RootDir, path)
	if err := os.MkdirAll(fullPath, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", path, err)
	}
	return nil
}

// Create Base and Feature Directories.
func (p *Project) createDirectories() error {
	p.RootDir = filepath.Clean(p.Name)
	if err := os.MkdirAll(p.RootDir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	for _, dir := range baseDirectories {
		p.CreateDirectory(dir)
	}

	for feature, enabled := range p.Features {
		if enabled {
			if dirName, exists := featureDirectories[feature]; exists {
				p.CreateDirectory(dirName)
			}
		}
	}

	return nil
}

// Generate files from a template
func (p *Project) generateFromTemplate(templatePath, outputPath string, data interface{}) error {
	template, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	if err := template.Execute(outputFile, data); err != nil {
		return err
	}

	return nil
}

func (p *Project) generateFile(template, output string) error {
	templatePath := filepath.Join("templates", template)
	outputPath := filepath.Join(p.RootDir, output)

	if err := p.generateFromTemplate(templatePath, outputPath, p); err != nil {
		return fmt.Errorf("failed to generate file %s: %w", output, err)
	}

	return nil
}

func (p *Project) generateBaseFiles() error {
	baseFileTemplates := map[string]string{
		"main.tmpl": "main.go",
		"routes.tmpl": "routes/routes.go",
	}

	for template, output := range baseFileTemplates {
		p.generateFile(template, output)
	}

	return nil
}

func (p *Project) generateFrameworkFiles() error {
	frameworkFilesTemplates := map[string]string{
		"frameworks/routes.tmpl": "routes/routes.go",
		"frameworks/controllers.tmpl": "controllers/example_controller.go",
	}

	for template, output := range frameworkFilesTemplates {
		p.generateFile(template, output)
	}

	return nil
}

func (p *Project) generateFeatureFiles() error {
	featureFilesTemplates := map[string]string{
		"features/s3.tmpl": "mys3/s3.go",
		"features/mail.tmpl": "mail/mail.go",
	}

	for template, output := range featureFilesTemplates {
		p.generateFile(template, output)
	}

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
