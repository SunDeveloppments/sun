package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func WriteYaml(key string, value string) {

	f, err := os.OpenFile("./.sunenv.yaml", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	var file_content string = fmt.Sprintf("%s : %s\n", key, value)

	io.WriteString(f, file_content)
}

func Init() {

	f_r := os.Remove("./.sunenv.yaml")

	if f_r != nil {
		fmt.Println("")
	}

	if len(os.Args) >= 1 {

		name := flag.String("name", "Sun", "The name of your package")
		language := flag.String("language", "Go", "The langauge in which is written your software")
		author := flag.String("author", "Jellyfish", "You ( here me )")

		flag.Parse()

		WriteYaml("name", *name)

		WriteYaml("language", *language)

		WriteYaml("author", *author)

	} else {
		GreetInit()

	}

}
