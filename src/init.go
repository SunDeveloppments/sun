package main

import (
	"os"
	"fmt"
	"io"
)

func Init(){
	var arg string = os.Args[1]

	f, err := os.Create(".sunenv.yaml")       

	if err != nil {

        fmt.Println(err)

	}

	defer f.Close()

	var name string = arg

	var file_content string = fmt.Sprintf("name: %s", name)

	io.WriteString(f, file_content) 
}
