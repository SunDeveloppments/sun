package main

import (
	"fmt"
	"os"
	"github.com/charmbracelet/glamour"
)

func greeting(){

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


func main() {

	var arg string

	if len(os.Args) > 1 {

		arg = os.Args[1]


	switch arg {

	case "settings":
		fmt.Println("This feature is not implemented.")

	case "init":
		fmt.Println("This feature is not implemented.")

	default:

		greeting()

}

	} else {

	greeting()


	}



}
