package main

import (
	"fmt"
	"os"
	"github.com/charmbracelet/glamour"
)

func greeting(){

in := `# Sun Development Environment                                                                                                                                                                      
                                                                                                                                                                                                                       
     Sun is a development projects manager, written in Go. Sun is free and open source.                                                                                                                                
                                                                                                                                                                                                                       
     Bye!                                                                                                                                                                                                              
     `                                                                                                                                                                                                                 
                                                                                                                                                                                                                       
             out, err := glamour.Render(in, "dark")                                                                                                                                                                    
             if err != nil {                                                                                                                                                                                           
                 panic(err)                                                                                                                                                                                            
           }                                                                                                                                                                                                         
                                                                                                                                                                                                                      
	          fmt.Print(out)                                                                                                                                                                                            
              return                                                                                                                                                                                          	
}


func main() {

	var arg string

	if len(os.Args) > 1 {

		arg = os.Args[1]


	switch arg {

	case "settings":
		fmt.Println("This feature is not implemented.")

	case "new":
		fmt.Println("This feature is not implemented.")

	default:

		greeting()

}

	} else {

	greeting()


	}



}
