package templates

import "embed"

//go:embed base/* features/* frameworks/*
var Files embed.FS
