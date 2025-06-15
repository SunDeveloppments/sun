package main

import (
	"fmt"
	"github.com/charmbracelet/glamour"
)

func Help(section string) {
	if section == "main" {
		in := "# Main help \nSun version 0.1 \n### Commands: \n**init** \n- Init helps to initialize or re-initialize .sunenv.yaml config file. \n- Syntax: \n    - ```bash \n    sun init [options] \n    ``` \n - See \n    ```bash \n    sun init --help \n    ```  \n**read** \n- read reads content of an existing .sunenv.yaml file, and shows it."
		out, err := glamour.Render(in, "dark")
		if err != nil {
			panic(err)
		}

		fmt.Print(out)
		return
	}
}
