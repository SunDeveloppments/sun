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
	checkDirectory := func(dir string) error {
		files, err := os.ReadDir(dir)
		if err != nil {
			return err
		}

		for _, file := range files {
			if file.IsDir() {
				if err := checkDirectory(filepath.Join(dir, file.Name())); err != nil {
					return err
				}
			} else {
				for _, framework := range frameworks {
					if strings.EqualFold(file.Name(), framework.Filename) {
						detected = append(detected, framework.Name)
					}
				}
			}
		}
		return nil
	}
	if err := checkDirectory("."); err != nil {
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
