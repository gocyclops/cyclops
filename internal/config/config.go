package config

type Config struct {
	ProjectName	string	`yaml:"project_name"`
	Framework	string `yaml:"framework"`
	Features	map[string]bool	`yaml:"features"`
}

func Load(path string) (*Config, error) {
	return nil, nil
}

func Save(config *Config, path string) error {
	return nil
}
