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

func checkDirectoryframes(dir string, frameworks []Framework, detected map[string]int) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			if err := checkDirectoryframes(filepath.Join(dir, file.Name()), frameworks, detected); err != nil {
				return err
			}
		} else {
			for _, framework := range frameworks {
				if strings.EqualFold(file.Name(), framework.Filename) {
					detected[framework.Name]++
				}
			}
		}
	}
	return nil
}

func DetectFrameworks() (map[string]int, error) {
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
		{"Make", "Makefile"},
		{"Just", "justfile"},
		{"Django", "manage.py"},
		{"Spring Boot", "application.properties"},
		{"Laravel", ".env"},
		{"React", "package.json"},
		{"Express", "app.js"},
		{"Symfony", "composer.json"},
		{"Svelte", "rollup.config.js"},
		{"Backbone.js", "backbone.js"},
		{"Ember.js", "ember-cli-build.js"},
		{"Gatsby", "gatsby-config.js"},
		{"Next.js", "next.config.js"},
		{"ASP.NET Core", "Startup.cs"},
		{"Flask", "app.py"},
		{"Kotlin with Ktor", "build.gradle.kts"},
		{"Phoenix", "mix.exs"},
		{"Crystal with Kemal", "shard.yml"},
		{"Haskell with Stack", "stack.yaml"},
		{"Clojure with Leiningen", "project.clj"},
		{"OpenCV", "CMakeLists.txt"},
		{"Apache Cordova", "config.xml"},
		{"Electron", "main.js"},
		{"Quasar", "quasar.config.js"},
		{"Nuxt.js", "nuxt.config.js"},
		{"Jekyll", "_config.yml"},
		{"Hugo", "config.toml"},
		{"Zig with Build.zig", "build.zig"},
		{"Ninja", "build.ninja"},
	}

	detected := make(map[string]int)

	if err := checkDirectoryframes(".", frameworks, detected); err != nil {
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
		Cprint("Detected frameworks:", green)
		for framework, count := range frameworks {
			if count > 1 {
				fmt.Printf("- %s (%d)\n", framework, count)
			} else {
				fmt.Println("-", framework)
			}
		}
	} else {
		fmt.Println("No frameworks detected.")
	}
}
