package main

import (
	"os"
	"fmt"
	"io"
)


var key string

func CreateFile(){
	
         f, err := os.Create("./.sunenv.yaml")
   
         if err != nil {
   
             fmt.Println(err)
   
         }
   
   defer f.Close()
}

func WriteYaml(key string){

//	CreateFile()
	 f, err := os.Create("./.sunenv.yaml")
   
            if err != nil {
   
                fmt.Println(err)
    
	         }

       defer f.Close()

	//arg := os.Args[1]
    
    var file_content string = fmt.Sprintf("name: %s", key)
    
    io.WriteString(f, file_content)
}

func Init(){

	var arg string = os.Args[2]


switch arg {

case "--author":

			WriteYaml("author")

case "--language":

			WriteYaml("language")

default:

	WriteYaml("name")

	}

}
