package main

import (
//	"fmt"
	"os"
// 	"github.com/charmbracelet/glamour"
)


func main() {

	var arg string

	if len(os.Args) > 1 {

		arg = os.Args[1]


	switch arg {

	case "read":
		Read()

	case "init":
		Init()

	default:

		GreetSun()

}

	} else {

	GreetSun()


	}



}
