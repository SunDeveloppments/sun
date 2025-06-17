package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Language struct {
	Name       string
	Extension  string
	LineCount  int
}

func CountLines(filePath string) (int, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return 0, err
	}
	lines := strings.Split(string(content), "\n")
	return len(lines), nil
}

func DetectLanguages() (map[string]*Language, error) {
	languages := []Language{
		{"Go", ".go"},
		{"Python", ".py"},
		{"JavaScript", ".js"},
		{"Java", ".java"},
		{"Ruby", ".rb"},
		{"PHP", ".php"},
		{"C", ".c"},
		{"C++", ".cpp"},
		{"HTML", ".html"},
		{"CSS", ".css"},
	}

	languageMap := make(map[string]*Language)
	checkDirectory := func(dir string) error {
		files, err := os.ReadDir(dir)
		if err != nil {
			return err
		}

		for _, file := range files {
			if file.IsDir() {
				if err := checkDirectory(filepath.Join(dir, file.Name())); err != nil {
					return err
				}
			} else {
				for _, lang := range languages {
					if strings.HasSuffix(file.Name(), lang.Extension) {
						lineCount, err := CountLines(filepath.Join(dir, file.Name()))
						if err != nil {
							return err
						}
						if _, exists := languageMap[lang.Name]; !exists {
							languageMap[lang.Name] = &Language{Name: lang.Name, Extension: lang.Extension}
						}
						languageMap[lang.Name].LineCount += lineCount
					}
				}
			}
		}
		return nil
	}

	if err := checkDirectory("."); err != nil {
		return nil, fmt.Errorf("error reading directory: %w", err)
	}

	return languageMap, nil
}
func CalculatePercentages(languageMap map[string]*Language) map[string]float64 {
	totalLines := 0
	for _, lang := range languageMap {
		totalLines += lang.LineCount
	}

	percentages := make(map[string]float64)
	for _, lang := range languageMap {
		if totalLines > 0 {
			percentages[lang.Name] = (float64(lang.LineCount) / float64(totalLines)) * 100
		} else {
			percentages[lang.Name] = 0
		}
	}

	return percentages
}

func main() {
	languageMap, err := DetectLanguages()
	if err != nil {
		fmt.Println("Error detecting languages:", err)
		return
	}

	percentages := CalculatePercentages(languageMap)

	fmt.Println("Language usage percentages:")
	for lang, percent := range percentages {
		fmt.Printf("%s: %.2f%%\n", lang, percent)
	}
}
