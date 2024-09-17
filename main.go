package main

import (
	"fmt"
	"log"
	"os"
	"asciiArtFs/myFunctions"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println()
		fmt.Println("EX: go run . --output=<fileName.txt> something standard")
		return
	}
	banner := os.Args[2]+".txt"
	standard, err := myfunctions.Read(banner)
	if err != nil {
		return
	}
	asciiChars := myfunctions.BytesToAsciiMap([]byte(standard))
	result, err := myfunctions.WriteResult(asciiChars)
	if err != nil {
		log.Println(err)
		return
	}
	myfunctions.PrintResult(result)
}
