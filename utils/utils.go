package utils

import (
	"regexp"
	"strings"
	"unicode"
)

func CleanProjectName(name string) string {
	name = strings.TrimSpace(name)

	name = strings.Map(func(r rune) rune {
		switch {
		case unicode.IsLetter(r):
			return unicode.ToLower(r)
		case unicode.IsNumber(r):
			return r
		case unicode.IsSpace(r) || r == '_' || r == '-':
			return '-'
		default:
			return -1
		}
	}, name)

	name = regexp.MustCompile(`-+`).ReplaceAllString(name, "-")

	return strings.Trim(name, "-")
}

func CleanRepoUrl(url string) string {
	url = strings.TrimSpace(url)

	url = strings.TrimSuffix(url, ".git")

	url = strings.TrimPrefix(url, "https://")
	url = strings.TrimPrefix(url, "http://")

	url = strings.ToLower(url)

	return url
}

func ValidateInputs(projectName, repoURL string) (bool, string) {
	if len(projectName) == 0 {
		return false, "Project name cannot be empty"
	}

	if len(projectName) > 50 {
		return false, "Project name too long (max 50 characters)"
	}

	validName := regexp.MustCompile(`^[a-z0-9][a-z0-9]*[a-z0-9]$`).MatchString(projectName)
	if !validName {
		return false, "Project name can only contain lowercase letters, numbers and hyphens"
	}

	if len(repoURL) == 0 {
		return false, "Repository URL cannot be empty"
	}

	validRepo := regexp.MustCompile(`^github\.com/[\w-]+/[\w-]+$`).MatchString(repoURL)
	if !validRepo {
		return false, "Invalid repository URL format (should be github.com/username/repo)"
	}

	return true, ""
}
