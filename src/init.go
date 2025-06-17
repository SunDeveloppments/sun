package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"log"
	"io/ioutil"
	"strings"
)

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

func WriteYaml(key string, value string) {
	path := "./.sunenv.yaml"
	content, err := ioutil.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			content = []byte{}
		} else {
			fmt.Println("Error reading file:", err)
			return
		}
	}
	lines := strings.Split(string(content), "\n")
	data := make(map[string]string)
	for _, line := range lines {
		if line != "" {
			parts := strings.SplitN(line, ": ", 2)
			if len(parts) == 2 {
				data[parts[0]] = parts[1]
			}
		}
	}
	data[key] = value
	var newContent strings.Builder
	for k, v := range data {
		if k == "platform" || k == "repo" {
			newContent.WriteString(fmt.Sprintf("  %s: %s\n", k, v))
		} else {
			newContent.WriteString(fmt.Sprintf("%s: %s\n", k, v))
		}
	}
	err = ioutil.WriteFile(path, []byte(newContent.String()), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

func Init() {
	path := "./.sunenv.yaml"

	if exists(path) {
		log.Println("Warning: File .sunenv.yaml already exists.")
	}

	helpflag := flag.Bool("help", false, "Show help")
	yesFlag := flag.Bool("y", false, "Confirm action without ask questions")
	nohostingFlag := flag.Bool("no-hosting", false, "If the project has no hosting platform.")
	name := flag.String("name", "default", "The name of your package")
	language := flag.String("language", "default", "The language in which your software is written")
	author := flag.String("author", "default", "Your name")
	author_email := flag.String("author-email", "default", "Email of author")
	maintener := flag.String("maintener", "default", "Maintener of the repo")
	maintener_email := flag.String("maintener-email", "default", "Email of maintener")
	platform := flag.String("platform", "default", "Hosting platform")
	repo := flag.String("repo", "default", "Repository URL")

	flag.Parse()

	if *helpflag {
		Help("init")
	}
	
	if !*yesFlag {
		if *name == "default" {
			*name = Input("package name: ")
		}
		if *language == "default" {
			*language = Input("Programming language: ")
		}
		if *author == "default" {
			*author = Input("Author name: ")
		}
		if *author_email == "default" {
			*author_email = Input("Author email: ")
		}
		if *maintener == "default" {
			*maintener = Input("Maintener name: ")
		}
		if *maintener_email == "default" {
			*maintener_email = Input("Maintener email: ")
		}
		if *no_hosting {
			if *platform == "default" {
				*platform = Input("Hosting platform: ")
			}
			if *repo == "default" {
				*repo = Input("Repository URL: ")
			}
		}
	}

	WriteYaml("name", *name)
	WriteYaml("language", *language)
	WriteYaml("author", *author)
	WriteYaml("author-email", *author_email)
	WriteYaml("maintener", *maintener)
	WriteYaml("maintener-email", *maintener_email)
	WriteYaml("platform", *platform)
	WriteYaml("repo", *repo)
}
