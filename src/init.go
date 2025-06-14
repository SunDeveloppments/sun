package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

func exists(path string) bool {

	_, err := os.Stat(path)

	return !errors.Is(err, os.ErrNotExist)

}

func CreateFile() {

	fc, e := os.Create("./.sunenv.yaml")

	if e != nil {

		fmt.Println(e)

	}

	fc.Close()

}

func WriteYaml(key string, value string) {

	path := "./.sunenv.yaml"

	fe := exists(path)

	if fe == true {

		CreateFile()

		return

	} else {

		f, err := os.OpenFile("./.sunenv.yaml", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

		if err != nil {

			fmt.Println(err)

		}

		defer f.Close()

		var file_content string = fmt.Sprintf("%s : %s\n", key, value)

		io.WriteString(f, file_content)
	}

}
func Init() {

	// f_r := os.Remove("./.sunenv.yaml")

	// if f_r != nil {
	// fmt.Println("Reinit sun.")
	// }

	name := flag.String("name", "Sun", "The name of your package")
	language := flag.String("language", "Go", "The langauge in which is written your software")
	author := flag.String("author", "Jellyfish", "You ( here me )")

	flag.Parse()

	WriteYaml("name", *name)

	WriteYaml("language", *language)

	WriteYaml("author", *author)

}
