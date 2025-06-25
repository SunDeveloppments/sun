package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/charmbracelet/huh"
)

type YamlFileConfig struct {
	Lang string `yaml:"lang"`
	Frm  string `yaml:"frm"`
	Init struct {
		Script string `yaml:"script"`
	} `yaml:"init"`
	Rm struct {
		Script string `yaml:"script"`
	} `yaml:"rm"`
}

func readYamlFileConfig(filename string) (YamlFileConfig, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return YamlFileConfig{}, err
	}
	filename = strings.Replace(filename, "~", homeDir, 1)

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return YamlFileConfig{}, err
	}
	var config YamlFileConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return YamlFileConfig{}, err
	}
	return config, nil
}

func hasExecutePermission(filePath string) bool {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return false
	}
	return fileInfo.Mode().Perm()&0111 != 0
}

func setExecutePermission(filePath string) error {
	return os.Chmod(filePath, 0755)
}

func Ask() {
	var selectedLanguage, selectedFramework string
	var wantsToInitialize bool

	frameworks := map[string][]string{
		"go":         {"Gin", "Echo", "Beego"},
		"python":     {"Django", "Flask", "FastAPI"},
		"javascript": {"Express", "Koa", "NestJS"},
		"java":       {"Spring", "Java EE", "Grails"},
		"cpp":        {"Qt", "Boost", "Poco"},
		"ruby":       {"Rails", "Sinatra", "Hanami"},
		"rust":       {"Rocket", "Actix", "Warp"},
		"c":          {},
		"hare":       {},
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose a programming language").
				Options(
					huh.NewOption("Go", "go"),
					huh.NewOption("Python", "python"),
					huh.NewOption("JavaScript", "javascript"),
					huh.NewOption("Java", "java"),
					huh.NewOption("C++", "cpp"),
					huh.NewOption("Ruby", "ruby"),
					huh.NewOption("Rust", "rust"),
					huh.NewOption("C", "c"),
					huh.NewOption("Hare", "hare"),
				).
				Value(&selectedLanguage),
		),
	)

	if err := form.Run(); err != nil {
		fmt.Println("Error running form:", err)
		return
	}

	if selectedLanguage != "" {
		availableFrameworks := frameworks[selectedLanguage]
		frameworkSelect := huh.NewSelect[string]().
			Title("Choose a framework").
			Options(createFrameworkOptions(availableFrameworks)...).
			Value(&selectedFramework)

		if err := frameworkSelect.Run(); err != nil {
			fmt.Println("Error running framework select:", err)
			return
		}

		confirm := huh.NewConfirm().
			Title(fmt.Sprintf("Do you want to initialize %s?", selectedFramework)).
			Value(&wantsToInitialize)

		if err := confirm.Run(); err != nil {
			fmt.Println("Error running confirmation:", err)
			return
		}

		configFile := fmt.Sprintf("~/.config/sun/assets/init/yaml/%s.yaml", strings.ToLower(selectedFramework))
		config, err := readYamlFileConfig(configFile)

		if err != nil {
			fmt.Println("Error reading config file:", err)
			return
		}

		if wantsToInitialize {
			scriptPath := config.Init.Script
			if !hasExecutePermission(scriptPath) {
				fmt.Printf("Script %s does not have exec perms. Adding..\n", scriptPath)
				err := setExecutePermission(scriptPath)
				if err != nil {
					fmt.Printf("Error: %v\n", err)
					return
				}
			}

			cmd := exec.Command("bash", scriptPath)
			cmd.Dir = "."
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			err = cmd.Run()
			if err != nil {
				fmt.Printf("Error executing init script: %s\n", err)
			}
		}
	}
}

func createFrameworkOptions(frameworks []string) []huh.Option[string] {
	options := make([]huh.Option[string], len(frameworks))
	for i, framework := range frameworks {
		options[i] = huh.NewOption(framework, framework)
	}
	return options
}
