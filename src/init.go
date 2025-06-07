package main

import (
	"os"
	"fmt"
	"io"
//    "github.com/charmbracelet/glamour"
)

var key string

func CreateFile() {
	f, err := os.Create("./.sunenv.yaml")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
}

func WriteYaml(key string) {
	f, err := os.OpenFile("./.sunenv.yaml", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	var arg string = os.Args[3]
	var file_content string = fmt.Sprintf("%s : %s\n", key, arg)
	io.WriteString(f, file_content)
}

func Init() {

	if len(os.Args) > 1 {

var args string = os.Args[2]

	switch args {
	case "--author":

		WriteYaml("author")

	case "--language":

		WriteYaml("language")

	case "--name":

		WriteYaml("name")

	default:

	GreetInit()

}
} else if len(os.Args) == 1 {

	GreetInit()

}

}

