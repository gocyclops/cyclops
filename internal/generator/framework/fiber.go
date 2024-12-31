package framework

type Fiber struct {
	ProjectName string
}

func NewFibre(projectName string) *Fiber {
	return &Fiber{ProjectName: projectName}
}

func (f *Fiber) GenerateFiles() error {
	return nil
}
