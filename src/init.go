package main

import (
	"errors"
	"fmt"
	"os"
	"log"
	"io/ioutil"
	"strings"
	"path/filepath"
	"encoding/json"
	"os/exec"
)

type ConfigType struct {
	Name            string
	Language        string
	Author          string
	AuthorEmail     string
	Maintener       string
	MaintenerEmail  string
	Platform        string
	Repo            string
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

func CreateFile() {
	fc, err := os.Create("./.sunenv.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fc.Close()
}

func WriteYaml(configtype ConfigType) {
	path := "./.sunenv.yaml"
	_, err := ioutil.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			CreateFile()
		} else {
			fmt.Println("Error reading file:", err)
			return
		}
	}

	var newContent strings.Builder
	newContent.WriteString(fmt.Sprintf("name: %s\n", configtype.Name))
	newContent.WriteString(fmt.Sprintf("maintener-email: %s\n", configtype.MaintenerEmail))
	newContent.WriteString(fmt.Sprintf("language: %s\n", configtype.Language))
	newContent.WriteString(fmt.Sprintf("author: %s\n", configtype.Author))
	newContent.WriteString(fmt.Sprintf("author-email: %s\n", configtype.AuthorEmail))
	newContent.WriteString(fmt.Sprintf("maintener: %s\n", configtype.Maintener))
	newContent.WriteString("hosting:\n")
	newContent.WriteString(fmt.Sprintf("  platform: \"%s\"\n", configtype.Platform))
	newContent.WriteString(fmt.Sprintf("  repo: %s\n", configtype.Repo))

	err = ioutil.WriteFile(path, []byte(newContent.String()), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

func Init(configtype ConfigType, helpflag bool, y, nohosting bool) {
	path := "./.sunenv.yaml"

    if exists(path) {
        log.Println("Warning: File .sunenv.yaml already exists.")
		if !y {
			response := Input("Overwrite existing file [y/N]? ")
        	if response != "yes" && response != "y" {
            	log.Println("Cancelled.")
    	        os.Exit(0)
        	}
		}
	}
	
	if !y {
		if configtype.Name == "default" {
			configtype.Name = Input("package name: ")
		}
		if configtype.Language == "default" {
			configtype.Language = Input("Programming language: ")
		}
		if configtype.Author == "default" {
			configtype.Author = Input("Author name: ")
		}
		if configtype.AuthorEmail == "default" {
			configtype.AuthorEmail = Input("Author email: ")
		}
		if configtype.Maintener == "default" {
			configtype.Maintener = Input("Maintener name: ")
		}
		if configtype.MaintenerEmail == "default" {
			configtype.MaintenerEmail = Input("Maintener email: ")
		}
		if nohosting {
			if configtype.Platform == "default" {
				configtype.Platform = Input("Hosting platform: ")
			}
			if configtype.Repo == "default" {
				configtype.Repo = Input("Repository URL: ")
			}
		}
	} else {
		configtype.Name = filepath.Dir(".")
		cmd := exec.Command("./sun", "detect", "--json")
		output, err := cmd.Output()
		if err != nil {
			fmt.Println("Error executing command:", err)
			return
		}
		var percentages map[string]float64
		if err := json.Unmarshal(output, &percentages); err != nil {
			fmt.Println("Error unmarshaling JSON:", err)
			return
		}
		var maxLang string
		var maxPercent float64
		for lang, percent := range percentages {
			if percent > maxPercent {
				maxPercent = percent
				maxLang = lang
			}
		}
		configtype.Language = maxLang
		configtype.Author = "John Doe"
		configtype.AuthorEmail = ""
		configtype.Maintener = "John Doe"
		configtype.MaintenerEmail = ""
		if nohosting {
			configtype.Platform = "Github"
			configtype.Repo = ""
		}
	}
	
	WriteYaml(configtype)
}
