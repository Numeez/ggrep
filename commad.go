package main

import "log"

type CommandData struct {
	pattern  string
	fileName string
	flag     string
}

func GetCommandData(arguments []string) CommandData {
	if len(arguments) == 2 {
		return CommandData{
			pattern:  arguments[0],
			fileName: arguments[1],
		}
	}
	return CommandData{
		flag:     arguments[0],
		pattern:  arguments[1],
		fileName: arguments[2],
	}
}



func executeCommand(command CommandData){
	if err:=findMatchFromFile(command.pattern,command.fileName);err!=nil{
		log.Fatal(err)
	}
}
