package main

import (
	"log"
	"os"
)

func main() {
	arguments := os.Args[1:]
	if len(arguments) == 0 {
		log.Fatal("Provide flags")
	}
	command := GetCommandData(arguments)
	executeCommand(command)

}
