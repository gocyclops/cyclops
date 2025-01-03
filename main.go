package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gocyclops/cyclops/cmd/create"
	"github.com/gocyclops/cyclops/internal/generator"
)

const Logo = `
     _________   ______  __      ____  ____  _____
    / ____/  /  / ___/ / /     / __ \/ __ \/ ___/
   / /   / /   / /    / /     / / / / /_/ /\__ \ 
  / /___/ /___/ /____/ /____ / /_/ / ____/___/ / 
  \____/_____/\____/_____/_/\____/_/    /____/  
`

func printLogo() string {
	return Logo
}

func showHelp() {
	help := printLogo()
	help += `


	Cyclops - Go Backend Boilerplate

		Usage:
			cyclops [command] [flags]
		
		Commands:
			new         Create a new project interactively
			add         Add features to existing project
		
		Flags:
			--help      Show help message
		
		Features available:
			- redis     Add Redis support
			- s3        Add S3 storage
			- mail      Add email functionality
			- auth      Add authentication
		`
	fmt.Println(help)
}

func checkConfig(path string) bool {
	configPath := filepath.Join(path, "config.yaml")
	_, err := os.Stat(configPath)
	return err == nil
}

func cyclopsSummary() string {
	sum := printLogo()
	sum += `


	Cyclops - A modern Go backend boilerplate
		Use 'cyclops new' to create a new project
		Use 'cyclops add' to add features to existing project
		Use 'cyclops --help' for more information`
	return sum
}

func featureAddition(feature string) error {
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	// Create feature adder
	featureAdder, err := generator.NewFeatureAdder(currentDir)
	if err != nil {
		return fmt.Errorf("failed to initialize feature adder: %w", err)
	}

	// Add the feature
	if err := featureAdder.AddFeature(feature); err != nil {
		return fmt.Errorf("failed to add feature: %w", err)
	}

	fmt.Printf("Successfully added %s feature to project!\n", feature)
	return nil
}

func main() {
	helpFlag := flag.Bool("help", false, "Show help message")
	flag.Parse()

	if *helpFlag {
		showHelp()
		return
	}

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println(cyclopsSummary())
		os.Exit(0)
	}

	command := args[0]
	switch command {
	case "new":
		if checkConfig(".") {
			fmt.Println("Error: Project found in current directory.")
			os.Exit(1)
		}

		if err := create.Execute(); err != nil {
			log.Fatal(err)
		}
	case "add":
		if len(args) < 2 {
			fmt.Println("Error: Feature name is required")
			fmt.Println("Available features: redis, s3, mail, auth")
			os.Exit(1)
		}

		if !checkConfig(".") {
			fmt.Println("Error: No project found in current directory (config.yaml missing)")
			os.Exit(1)
		}

		feature := args[1]
		validFeatures := map[string]bool{
			"redis": true,
			"s3":    true,
			"mail":  true,
			"auth":  true,
		}

		if !validFeatures[feature] {
			fmt.Printf("Error: Invalid feature '%s'\n", feature)
			fmt.Println("Available features: redis, s3, mail, auth")
			os.Exit(1)
		}

		fmt.Printf("Adding %s feature to project...\n", feature)
		if err := featureAddition(feature); err != nil {
			log.Fatal(err)
		}
	default:
		fmt.Printf("Error: Unknown command '%s'\n", command)
		fmt.Println("Run 'cyclops --help' for usage information")
		os.Exit(1)
	}
}
