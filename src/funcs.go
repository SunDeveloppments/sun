package main

import (
	"os"
	"bufio"
	"fmt"
)

func Input(question string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(question)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(1)
	}
	return input[:len(input)-1]
}