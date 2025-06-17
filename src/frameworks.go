package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Framework struct {
	Name     string
	Filename string
}

func checkDirectoryframes(dir string, filenames []string, detected *[]string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			if err := checkDirectoryframes(filepath.Join(dir, file.Name()), filenames, detected); err != nil {
				return err
			}
		} else {
			for _, filename := range filenames {
				if strings.EqualFold(file.Name(), filename) {
					*detected = append(*detected, filename)
				}
			}
		}
	}
	return nil
}

func DetectFrameworks() ([]string, error) {
	frameworks := []Framework{
		{"Node.js", "package.json"},
		{"Go", "go.mod"},
		{"Python", "requirements.txt"},
		{"Ruby on Rails", "Gemfile"},
		{"PHP", "composer.json"},
		{"Java", "pom.xml"},
		{"Rust with Cargo", "Cargo.toml"},
		{"Vue.js", "app.vue"},
		{"Bootstrap", "bootstrap.min.css"},
		{"Flutter", "pubspec.yaml"},
		{"ASP.NET", ".csproj"},
		{"Angular", "angular.json"},
		{"jQuery", "jquery.min.js"},
	}

	var detected []string
	filenames := make([]string, len(frameworks))
	for i, framework := range frameworks {
		filenames[i] = framework.Filename
	}

	if err := checkDirectoryframes(".", filenames, &detected); err != nil {
		return nil, fmt.Errorf("error reading directory: %w", err)
	}

	return detected, nil
}

func Frameworks() {
	frameworks, err := DetectFrameworks()
	if err != nil {
		fmt.Println("Error detecting frameworks:", err)
		return
	}

	if len(frameworks) > 0 {
		fmt.Println("Detected frameworks:")
		for _, framework := range frameworks {
			fmt.Println("-", framework)
		}
	} else {
		fmt.Println("No frameworks detected.")
	}
}
