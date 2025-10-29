package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func findMatchFromFile(pattern, fileName string) error {
	file, err := os.Open(fileName)
	fmt.Printf("\x1b[32m%s\x1b[0m\n", fileName)

	defer func() {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	if err != nil {
		return err
	}
	count := 1
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		found, line := findMatch(line, pattern)
		if found {
			fmt.Printf("\x1b[35m%d\x1b[0m : %s\n", count, line)
		}
		count++
	}

	return nil
}

func findMatch(line, pattern string) (bool, string) {
	matchIndex := strings.Index(line, pattern)
	if matchIndex != -1 {
		return true, fmt.Sprintf("%s\x1b[31m%s\x1b[0m%s", line[:matchIndex], line[matchIndex:matchIndex+len(pattern)],line[matchIndex+len(pattern):])
	}
	return false, ""
}
