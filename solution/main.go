package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode"
)

type LineInfo struct {
	line  string
	count int
}

type LineInfoList []*LineInfo

func main() {
	fmt.Println("")
	args := os.Args

	if len(args) <= 1 {
        fmt.Println("Write input, enter EOF (Ctrl-D) to end input") 
		reader := bufio.NewReader(os.Stdin)
		input := []byte{}
		for  {
			text, err := reader.ReadSlice('\n')
			if err == io.EOF {
				fmt.Println("")
				break
			}
			input = append(input, text...)
		}

		lineList := InputToLineList(input)
		PrettyPrint(ProcessLineList(lineList))
		return
	}


	lineList := ReadFileToLineList(args[1])
	PrettyPrint(ProcessLineList(lineList))
}

func InputToLineList(input []byte) []string {
	if !isValidASCII(input) {
		log.Println("Not valid ASCII")
		os.Exit(2)
	}

	// Handle case of \r\n line endings
	fileAsString := strings.ReplaceAll(string(input), "\r\n", "\n")

	return strings.Split(fileAsString, "\n")
}

func ReadFileToLineList(filename string) []string {
	fileAsByteSlice, err := os.ReadFile(filename)
	if err != nil {
		log.Println("Could not read file:", err.Error())
		os.Exit(1)
	}

	return InputToLineList(fileAsByteSlice)
}

func isValidASCII(s []byte) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}

func ProcessLineList(lineList []string) []string {
	seen := make(LineInfoList, 0, len(lineList))
	for _, line := range lineList {
		if seen.contains(line) {
			seen.increment(line)
		} else {
			seen = append(seen, &LineInfo{
				line:  line,
				count: 1,
			})
		}

		rev := reverseLine(line)
        if rev == line {
            continue
        }
		if seen.contains(rev) {
			seen.increment(rev)
			seen.increment(line)
		}
	}

	output := make([]string, 0)
	for _, lineInfo := range seen {
		if lineInfo.count > 1 {
			output = append(output, lineInfo.line)
		}
	}
	return output
}

func (list LineInfoList) contains(str string) bool {
	for _, info := range list {
		if info.line == str {
			return true
		}
	}
	return false
}

func (list LineInfoList) increment(str string) {
	for _, lineInfo := range list {
		if lineInfo.line == str {
			lineInfo.count++
		}
	}
}

// Because the string is only ASCII characters, we can use a byte array to reverse it
// Nevertheless, we'll use a rune array just in case
func reverseLine(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func PrettyPrint(s []string) {
	for _, l := range s {
		fmt.Println(l)
	}
}
