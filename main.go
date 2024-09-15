package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println()
		fmt.Println("EX: go run . --output=<fileName.txt> something standard")
		return
	}
	banner := os.Args[2]+".txt"
	standard, err := Read(banner)
	if err != nil {
		return
	}
	asciiChars := BytesToAsciiMap([]byte(standard))
	result, err := WriteResult(asciiChars)
	if err != nil {
		log.Println(err)
		return
	}
	PrintResult(result)
}

func PrintResult(result []string) {
	for _, v := range result {
		fmt.Println(v)
	}
}

func WriteResult(asciiChars map[int]string) ([]string, error) {
	result := []string{"","","","","","","",""}
	lineToWrite := 0
	str := os.Args[1]
	for i := 0; i < len(str); i++ {
		//Handle non ascii char.
		if str[i] < 32 || str[i] > 126 {
			return result, fmt.Errorf("A none ascii char has been found !!")
		}

		//Handle \n.
		if str[i] == '\\' && str[i+1] == 'n' {
			if len(str) == i+2 || (len(str) > i + 4 && str[i+2] == '\\' && str[i+3] == 'n' ) {
				result = append(result,	"")
				lineToWrite++
				i++
				continue
			} else {
				newLineAscii := []string{"","","","","","","",""}
				result = append(result,	newLineAscii...)
				lineToWrite += 8
				i++
				continue
			}
		}
		asciiChar := strings.Split(asciiChars[int(str[i])], "\n")
		for i, line := range asciiChar {
			result[i + lineToWrite] += line[:len(line)-1]

			//Fix spacing
			if i > 0 {
				spaces := strings.Repeat(" ", len(asciiChar[0])- len(line))
				result[i + lineToWrite] += spaces
			}
		}
	}
	return result, nil
}

func BytesToAsciiMap(style []byte) map[int]string {
	chars := make(map[int]string)
	line := 1
	next := 9
	char := ""
	nbrChar := 32
	for i := 1; i < len(style); i++ {
		if i < len(style) - 1 {
		 	if style[i] == '\n' {
				line++
			} else if line == next+1 {
				next += 9
				chars[nbrChar] = char[1:len(char)-2]
				nbrChar++
				char = ""
				continue
			}
			char += string(style[i])
		} else {
			char += string(style[i])
			chars[nbrChar] = char
		}
	}
	return chars
}

func Read(fileName string) (string, error) {
	
	//Open File.
    file, err := os.Open(fileName)
	if err != nil {
		log.Println("error opening file :", fileName)
		return "", err
	}

	defer file.Close()
	
	//Get file info.
    fileInfo, err := file.Stat()
    if err != nil {
        log.Println("Error getting file stats:", err)
		return "", err
    }

    //Get file size.
    fileSize := fileInfo.Size()
    data := make([]byte, fileSize)

	//Reading data.
    _, err = file.Read(data)
    if err != nil {
		log.Println("Error reading the file:", err)
		return "", err
    }
	return string(data), nil
}
