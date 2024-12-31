package redis

import (
	"os"
	"path/filepath"
)

func Generate(projectPath string) error {
	redisDir := filepath.Join(projectPath, "myredis")
	if err := os.MkdirAll(redisDir, 0755); err != nil {
		return err
	}

	return nil
}
