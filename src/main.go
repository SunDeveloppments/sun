package main

import (
	"os"
	// "github.com/charmbracelet/glamour"
)

func main() {

	args := os.Args

	if len(args) >= 2 {

		var arg1 string = os.Args[1]

		switch arg1 {

		case "read":
			Read()

		case "init":
			Init()

		default:

			GreetSun()

		}

	} else if len(os.Args) == 2 {

		var arg1 string = os.Args[1]

		switch arg1 {

		case "init":

			GreetInit()

		case "read":

			Read()

		default:

			GreetSun()
		}

	} else {

		GreetSun()

	}

}
