package mail

import (
	"os"
	"path/filepath"
)

func Generate(projectPath string) error {
	mailDir := filepath.Join(projectPath, "mail")
	if err := os.MkdirAll(mailDir, 0755); err != nil {
			return err
	}

	// Write mail.go file
	return nil
}
