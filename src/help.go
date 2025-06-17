package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/charmbracelet/glamour"
)

func Help(section string) {
	helpDir := ".config/sun/assets/help/"
	homeDir := os.Getenv("HOME")
	var DirPath string
	var filePath string
	DirPath = filepath.Join(homeDir, helpDir)
	if section == "main" {
		filePath = filepath.Join(DirPath, "main.md")
	} else if section == "init" {
		filePath = filepath.Join(DirPath, "init.md")
	} else {
		fmt.Println("Error: unrecognized section.")
		os.Exit(127)
	}
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", filePath, err)
		return
	}
	out, err := glamour.Render(string(content), "dark")
	if err != nil {
		panic(err)
	}
	fmt.Print(out)
}
