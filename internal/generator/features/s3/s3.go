package s3

import (
	"os"
	"path/filepath"
)

func Generate(projectPath string) error {
	s3Dir := filepath.Join(projectPath, "mys3")
	if err := os.MkdirAll(s3Dir, 0755); err != nil {
			return err
	}

	// Write s3.go file
	return nil
}
