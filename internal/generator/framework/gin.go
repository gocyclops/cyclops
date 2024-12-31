package framework

type Gin struct {
	ProjectName string
}

func NewGin(projectName string) *Gin {
	return &Gin{ProjectName: projectName}
}

func (g *Gin) GenerateFiles() error {
	// Implementation
	return nil
}
