package main

import (
    "fmt"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "os/exec"
    "strings"
    "github.com/charmbracelet/huh"
)

var (
    LocalInstall   bool
    SysInstall     bool
    PortableInstall bool
    Version        = "dev (version unknown)"
)

type Config struct {
    Lang string `yaml:"lang"`
    Frm  string `yaml:"frm"`
    Init struct {
        Script string `yaml:"script"`
    } `yaml:"init"`
    Rm struct {
        Script string `yaml:"script"`
    } `yaml:"rm"`
}

func readConfig(filename string) (Config, error) {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        return Config{}, err
    }
    var config Config
    if err := yaml.Unmarshal(data, &config); err != nil {
        return Config{}, err
    }
    return config, nil
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

        configFile := fmt.Sprintf("%s.yaml", strings.ToLower(selectedFramework))
        config, err := readConfig(configFile)

        if err != nil {
            fmt.Println("Error reading config file:", err)
            return
        }

        if wantsToInitialize {
            if err := executeScript(config.Init.Script); err != nil {
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

func executeScript(script string) error {
    cmd := exec.Command(script)
    output, err := cmd.CombinedOutput()
    if err != nil {
        return err
    }
    fmt.Printf("Output of init script:\n%s\n", output)
    return nil
}