package main

import (
	"fmt"
	//"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Name     string `yaml:"name"`
	Language string `yaml:"language"`
	Author   string `yaml:"author"`
}

func Read() {
	filePath := ".sunenv.yaml"

	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}

	var config Config

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Println(err)
	} else {

		fmt.Printf("This app is : %s\n", config.Name)
		fmt.Printf("Written in : %s\n", config.Language)
		fmt.Printf("Created by : %s\n", config.Author)
	}

}
