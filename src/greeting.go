package main

import "github.com/charmbracelet/glamour"
import "fmt"

func GreetSun(){
	in := `# Sun Development Environment                                                                                                                                                                      
                                                                                                                                                                                                                       
     Sun is an opensource, free development projects manager, written in Go.                                                                                                                                
                                                                                                                                                                                                                                                                                                                                                                  
     `                                                                                                                                                                                                                 
                                                                                                                                                                                                                       
             out, err := glamour.Render(in, "dark")                                                                                                                                                                    
             if err != nil {                                                                                                                                                                                           
                 panic(err)                                                                                                                                                                                            
           }                                                                                                                                                                                                         
                                                                                                                                                                                                                      
	          fmt.Print(out)                                                                                                                                                                                            
              return        
}

func GreetInit(){
		in := `# Sun init

Usage :

- sun init --name : initialize the name of your app
- sun init --language : initialize the programming language of your app
- sun init --author : initialize the author of your app
`

out, err := glamour.Render(in, "dark")

if err != nil {
	fmt.Println(err)
}
 
fmt.Print(out)
}
