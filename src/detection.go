package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Language struct {
	Name      string
	Extension string
	LineCount int
}

func CountLines(filePath string) (int, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return 0, err
	}
	lines := strings.Split(string(content), "\n")
	return len(lines), nil
}

func checkDirectory(dir string, extensions []string, detected map[string]int, languages []Language) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			if err := checkDirectory(filepath.Join(dir, file.Name()), extensions, detected, languages); err != nil {
				return err
			}
		} else {
			filePath := filepath.Join(dir, file.Name())
			// Check for files without extension (like Makefile)
			if file.Name() == "Makefile" {
				lineCount, err := CountLines(filePath)
				if err != nil {
					return err
				}
				detected["Makefile"] += lineCount
			}

			// Check for extensions
			for _, lang := range languages {
				if strings.HasSuffix(file.Name(), lang.Extension) || strings.EqualFold(file.Name(), lang.Extension) {
					if _, exists := detected[lang.Name]; !exists {
						detected[lang.Name] = 0
					}
					lineCount, err := CountLines(filePath)
					if err != nil {
						return err
					}
					detected[lang.Name] += lineCount // Utiliser le nom du langage ici
				}
			}
		}
	}
	return nil
}

func DetectLanguages() (map[string]int, error) {
	languages := []Language{
		{"Go", ".go", 0},
		{"Python", ".py", 0},
		{"JavaScript", ".js", 0},
		{"Java", ".java", 0},
		{"Ruby", ".rb", 0},
		{"PHP", ".php", 0},
		{"C", ".c", 0},
		{"C++", ".cpp", 0},
		{"HTML", ".html", 0},
		{"CSS", ".css", 0},
		{"Vue.js", ".vue", 0},
		{"JSP", ".jsp", 0},
	}

	detected := make(map[string]int)
	for _, lang := range languages {
		detected[lang.Name] = 0
	}

	extensions := make([]string, len(languages))
	for i, lang := range languages {
		extensions[i] = lang.Extension
	}

	detected["Makefile"] = 0
	if err := checkDirectory(".", extensions, detected, languages); err != nil {
		return nil, fmt.Errorf("error reading directory: %w", err)
	}
	return detected, nil
}

func CalculatePercentages(detected map[string]int) map[string]float64 {
	totalLines := 0
	for _, count := range detected {
		totalLines += count
	}

	percentages := make(map[string]float64)
	for lang, count := range detected {
		if totalLines > 0 {
			percentages[lang] = (float64(count) / float64(totalLines)) * 100
		} else {
			percentages[lang] = 0
		}
	}

	return percentages
}

func Detect() {
	detected, err := DetectLanguages()
	if err != nil {
		fmt.Println("Error detecting languages:", err)
		return
	}

	percentages := CalculatePercentages(detected)
	fmt.Println("Language usage percentages:")
	for lang, percent := range percentages {
		if percent > 0 {
			fmt.Printf("%s: %.2f%%\n", lang, percent)
		}
	}
}
