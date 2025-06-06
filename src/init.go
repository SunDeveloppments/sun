package main

import (
	"os"
	"fmt"
	"io"
)

func Init(){
	var arg string = os.Args[2]

	f, err := os.Create("./.sunenv.yaml")       

	if err != nil {

        fmt.Println(err)

	}

	defer f.Close()

	var file_content string = fmt.Sprintf("name: %s", arg)

	io.WriteString(f, file_content) 
}
