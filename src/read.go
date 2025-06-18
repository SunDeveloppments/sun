package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"gopkg.in/yaml.v3"
)

type Hosting struct {
	Platform string `yaml:"platform"`
	Repo     string `yaml:"repo"`
}

type Config struct {
	Name           string `yaml:"name"`
	Language       string `yaml:"language"`
	Author         string `yaml:"author"`
	AuthorEmail    string `yaml:"author-email"`
	Maintainer     string `yaml:"maintainer"`
	MaintainerEmail string `yaml:"maintainer-email"`
	Hosting        Hosting `yaml:"hosting"`
}

func ifNotSpecified(field *string, defaultValue string) {
	if *field == "" {
		*field = defaultValue
	}
}

func readConfig(filePath string) (*Config, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	ifNotSpecified(&config.Name, " not specified")
	ifNotSpecified(&config.Language, " not specified")
	ifNotSpecified(&config.Author, " not specified")
	ifNotSpecified(&config.AuthorEmail, " not specified")
	ifNotSpecified(&config.Maintainer, " not specified")
	ifNotSpecified(&config.MaintainerEmail, " not specified")
	ifNotSpecified(&config.Hosting.Platform, " not specified")
	ifNotSpecified(&config.Hosting.Repo, " not specified")

	return &config, nil
}

func printConfig(config *Config) {
	fmt.Printf("This app is : %s\n", config.Name)
	fmt.Printf("Written in : %s\n", config.Language)
	fmt.Printf("Created by : %s, email: %s\n", config.Author, config.AuthorEmail)
	fmt.Printf("Maintained by : %s, email: %s\n", config.Maintainer, config.MaintainerEmail)
	fmt.Printf("Hosting platform : %s\n", config.Hosting.Platform)
	fmt.Printf("Repository : %s\n", config.Hosting.Repo)
}

func Read(jsonOutput bool) {
	config, err := readConfig("./.sunenv.yaml")
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}
	if jsonOutput {
		output, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling to JSON: %v", err)
		}
		fmt.Println(string(output))
	} else {
		printConfig(config)
	}
}