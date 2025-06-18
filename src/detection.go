package main

import (
	"encoding/json"
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
			if file.Name() == "Makefile" {
				lineCount, err := CountLines(filePath)
				if err != nil {
					return err
				}
				detected["Makefile"] += lineCount
			}

			for _, lang := range languages {
				if strings.HasSuffix(file.Name(), lang.Extension) || strings.EqualFold(file.Name(), lang.Extension) {
					if _, exists := detected[lang.Name]; !exists {
						detected[lang.Name] = 0
					}
					lineCount, err := CountLines(filePath)
					if err != nil {
						return err
					}
					detected[lang.Name] += lineCount
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
		{"Swift", ".swift", 0},
		{"Kotlin", ".kt", 0},
		{"TypeScript", ".ts", 0},
		{"Rust", ".rs", 0},
		{"Dart", ".dart", 0},
		{"Scala", ".scala", 0},
		{"Shell", ".sh", 0},
		{"Perl", ".pl", 0},
		{"Haskell", ".hs", 0},
		{"Elixir", ".ex", 0},
		{"C#", ".cs", 0},
		{"Objective-C", ".m", 0},
		{"Groovy", ".groovy", 0},
		{"R", ".R", 0},
		{"Lua", ".lua", 0},
		{"COBOL", ".cob", 0},
		{"Fortran", ".f90", 0},
		{"Assembly", ".asm", 0},
		{"F#", ".fs", 0},
		{"Julia", ".jl", 0},
		{"Crystal", ".cr", 0},
		{"Nim", ".nim", 0},
		{"Pascal", ".pas", 0},
		{"Tcl", ".tcl", 0},
		{"VB.NET", ".vb", 0},
		{"Scratch", ".sb", 0},
		{"ActionScript", ".as", 0},
		{"Smalltalk", ".st", 0},
		{"OCaml", ".ml", 0},
		{"SAS", ".sas", 0},
		{"Solidity", ".sol", 0},
		{"APL", ".apl", 0},
		{"Logo", ".logo", 0},
		{"Prolog", ".pl", 0},
		{"XSLT", ".xsl", 0},
		{"VHDL", ".vhd", 0},
		{"Verilog", ".v", 0},
		{"Hare", ".ha", 0},
		{"Bash", ".bash", 0},
		{"Fish", ".fish", 0},
		{"Cmake", ".cmake", 0},
		{"LaTEX", ".tex", 0},
		{"Typst", ".typ", 0},
		{"Batch file", ".bat", 0},
		{"Haskell", ".hs", 0},
		{"OCaml", ".ml", 0}
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

func Detect(jsonOutput bool) {
	detected, err := DetectLanguages()
	if err != nil {
		fmt.Println("Error detecting languages:", err)
		return
	}

	percentages := CalculatePercentages(detected)

	if jsonOutput {
		output, err := json.MarshalIndent(percentages, "", "  ")
		if err != nil {
			fmt.Println("Error marshaling to JSON:", err)
			return
		}
		fmt.Println(string(output))
	} else {
		Cprint("Language usage percentages:", green)
		for lang, percent := range percentages {
			if percent > 0 {
				fmt.Printf("%s: %.2f%%\n", lang, percent)
			}
		}
	}
}
