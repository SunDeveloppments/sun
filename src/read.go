package main

import (
	"fmt"
	"os"
	// "log"
	"gopkg.in/yaml.v3"
)

type Hosting struct {
	Platform string `yaml:"platform"`
	Repo     string `yaml:"repo"`
}

type Config struct {
	Name     string `yaml:"name"`
	Language string `yaml:"language"`
	Author   string `yaml:"author"`
	Maintainer string `yaml:"maintainer"`
	Hosting  Hosting `yaml:"hosting"`
}

func Read() {
	filePath := ".sunenv.yaml"

	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	var config Config

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("This app is : %s\n", config.Name)
	fmt.Printf("Written in : %s\n", config.Language)
	fmt.Printf("Created by : %s\n", config.Author)
	fmt.Printf("Maintained by : %s\n", config.Maintainer)
	fmt.Printf("Hosting platform : %s\n", config.Hosting.Platform)
	fmt.Printf("Repository : %s\n", config.Hosting.Repo)
}
