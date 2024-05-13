package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isDigit(byte byte) bool {
	return (byte >= '0' && byte <= '9')
}
func isLetter(byte byte) bool {
	return ((byte >= 'a' && byte <= 'z') || (byte >= 'A' && byte <= 'Z'))
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Openning file error: ", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var fileLines []string
	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}
	file.Close()

	var counter int
	var slicedLine []string //Game 1 [0] || 2 red, 3 blue, 4 green [1]

	m := make(map[string]int)
	m["red"] = 0
	m["blue"] = 0
	m["green"] = 0

	for _, lines := range fileLines {

		slicedLine = strings.Split(lines, ":")
		fmt.Println(slicedLine[0][4:]) //ignore Game and get only the integer part
		counter, err = strconv.Atoi(slicedLine[0][5:])
		if err != nil {
			fmt.Println("Give me error code: ", err)
		}
		fmt.Println("Counter measures!: -", counter)
		fmt.Println("Array", slicedLine[1])

		var digits string
		var color string
		var listDigits []string
		var listColor []string


	slicedLineForloop:
		for i := 0; i < len(slicedLine[1]); i++ {
			if isDigit(slicedLine[1][i]) {
				digits += string(slicedLine[1][i])
			}
			if isLetter(slicedLine[1][i]) {
				color += string(slicedLine[1][i])
				fmt.Println("Good!", color)
			}
			if slicedLine[1][i]) == ',' {
				listDigits = strings.append(listDigits, digits)
			}
			if slicedLine[1][i] == ';' {
				fmt.Println("This is the digits of array: ", digits)
				fmt.Println("This si the color of the array: ", color)
				digits = ""
				color = ""
				break slicedLineForloop
			}
		}
	}
}
