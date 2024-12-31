package framework

type Chi struct {
    ProjectName string
}

func NewChi(projectName string) *Chi {
    return &Chi{ProjectName: projectName}
}

func (c *Chi) GenerateFiles() error {
    // Implementation
    return nil
}
