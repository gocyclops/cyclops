package generator

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"github.com/gocyclops/cyclops/internal/config"
)

type Project struct {
	Name      string
	ModuleName	string
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
	".github/workflows",
}

var featureDirectories = map[string]string{
	"redis": "myredis",
	"s3":    "mys3",
	"mail":  "mail",
	"utils": "utils",
}

func (p *Project) InitGitRepo() error {
	projectPath := filepath.Join(p.RootDir, p.Name)

	commands := []struct {
		name string
		args []string
	}{
		{"git", []string{"init"}},
		{"git", []string{"add", "."}},
		{"git", []string{"commit", "-m", "Initial commit from Cyclops 🚀"}},
	}

	for _, cmd := range commands {
		command := exec.Command(cmd.name, cmd.args...)
		command.Dir = projectPath
		if err := command.Run(); err != nil {
			return fmt.Errorf("failed to execute git command '%s': %w", cmd.name, err)
		}
	}

	return nil
}

func (p *Project) InitGoModule() error {
	projectPath := filepath.Join(p.RootDir, p.Name)

	command := exec.Command("go", "mod", "init", p.ModuleName)
	command.Dir = projectPath
	if err := command.Run(); err != nil {
		return fmt.Errorf("failed to create go module")
	}
	return nil
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
		return fmt.Errorf("failed to parse template %s: %w", templatePath, err)
	}

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file %s: %w", outputPath, err)
	}
	defer outputFile.Close()

	if err := template.Execute(outputFile, data); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
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

func (p *Project) generateFiles(templates map[string]string) error {
	for template, output := range templates {
		if err := p.generateFile(template, output); err != nil {
			return err
		}
	}
	return nil
}

func (p *Project) generateBaseFiles() error {
	baseFileTemplates := map[string]string{
		"base/main.go.tmpl": "main.go",
		"base/README.md.tmpl": "README.md",
		"base/Dockerfile.tmpl": "Dockerfile",
		"base/.air.toml.tmpl": ".air.toml",
		"base/.env.tmpl": ".env",
		"base/.gitignore.tmpl": ".gitignore",
		"base/.dockerignore.tmpl": ".dockerignore",
		"base/database/config.go.tmpl": "database/config.go",
		"base/models/models.go.tmpl": "models/models.go",
		"base/repository/repository.go.tmpl": "repository/repository.go",
		"base/migrations/migrations.go.tmpl": "migrations/migrations.go",
		"base/test.yml.tmpl": ".github/workflows/test.yml",
	}
	return p.generateFiles(baseFileTemplates)
}

func (p *Project) generateFrameworkFiles() error {
	frameworkFilesTemplates := map[string]string{
		"frameworks/routes.tmpl": "routes/routes.go",
		"frameworks/controllers.tmpl": "controllers/example_controller.go",
	}
	return p.generateFiles(frameworkFilesTemplates)
}

func (p *Project) generateConfigYaml() error {
	cfg := &config.Config{
		ProjectName: p.Name,
		Framework: p.Framework,
		Features: p.Features,
	}

	if err := config.Save(cfg, filepath.Join(p.Name, "config.yaml")); err != nil {
		return fmt.Errorf("error saving config.yaml: %w", err)
	}

	return nil
}

func (p *Project) generateFeatureFiles() error {
	for feature, enabled := range p.Features {
		if enabled {
			var featureFilesTemplates map[string]string
			switch feature {
			case "s3":
				featureFilesTemplates = map[string]string{
					"features/s3/s3.go.tmpl": "mys3/s3.go",
				}
			case "mail":
				featureFilesTemplates = map[string]string{
					"features/mail/mail.go.tmpl": "mail/mail.go",
				}
			case "redis":
				featureFilesTemplates = map[string]string{
					"features/redis/redis.go.tmpl": "myredis/redis.go",
				}
			case "auth":
				featureFilesTemplates = map[string]string{
					"features/utils/jwt.go.tmpl": "utils/jwt.go",
				}
			}
			if err := p.generateFiles(featureFilesTemplates); err != nil {
				return err
			}
		}
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

	if err := p.generateConfigYaml(); err != nil {
		return fmt.Errorf("failed to generate config.yaml: %w", err)
	}

	if err := p.InitGitRepo(); err != nil {
		return fmt.Errorf("failed to initialize git repository: %w", err)
	}

	if err := p.InitGoModule(); err != nil {
		return fmt.Errorf("failed to initialize GO module: %w", err)
	}

	return nil
}
