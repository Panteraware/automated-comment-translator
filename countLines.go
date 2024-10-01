package main

import (
	"bufio"
	"log"
	"os"
)

func CountLines(filePath string) int {
	f, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return 0
	}
	defer f.Close()

	index := 0

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		if CheckingString(sc.Text()).Matches {
			index++
		}
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return 0
	}

	return index
}
