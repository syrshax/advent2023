package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isDigit(word byte) bool {
	return (word >= '0' && word <= '9')
}
func isLetter(word byte) bool {
	return (word >= 'a' && word <= 'z')
}

func giveNumber(lines [][]byte) []string {
	var results []string
	var right, left string
	var tempWord string

	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
	m["four"] = 4
	m["five"] = 5
	m["six"] = 6
	m["seven"] = 7
	m["eight"] = 8
	m["nine"] = 9

	for _, l := range lines {

	normal:
		for i := 0; i < len(l); i++ {
			if isDigit(l[i]) {
				right = string(l[i])
				tempWord = ""
				break normal
			}

			if isLetter(l[i]) {
				tempWord += string(l[i])
				fmt.Println("Asi va la letra: - ", tempWord)
				for key, val := range m {
					if strings.Contains(tempWord, key) {
						fmt.Println("Habemus cosas encontradas, ", tempWord, val)
						right = strconv.Itoa(val)
						tempWord = ""
						break normal
					}
				}
			}
		}
	reverse:
		for i := len(l) - 1; i >= 0; i-- {
			if isDigit(l[i]) {
				left = string(l[i])
				tempWord = ""
				break reverse
			}
			if isLetter(l[i]) {
				tempWord = string(l[i]) + tempWord
				fmt.Println("Asi va la letra: - ", tempWord)
				for key, val := range m {
					if strings.Contains(tempWord, key) {
						fmt.Println("Habemus cosas encontradas, ", tempWord, val)
						left = strconv.Itoa(val)
						tempWord = ""
						break reverse
					}
				}
			}
		}
		combination := right + left
		right = ""
		left = ""
		results = append(results, combination)
		combination = ""
	}
	fmt.Println("This is the final results!", results)

	return results
}

func arraySum(str []string) int {
	var sum int
	for _, num := range str {
		numInt, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println("Error converting this number to things!", err)
			continue
		}
		sum = sum + numInt
	}
	return sum

}
func main() {
	data, err := os.ReadFile("input1.txt")
	fmt.Println(data)
	check(err)
	answer := arraySum(giveNumber(bytes.Split(data, []byte("\n"))))
	fmt.Println("La suma es: ", answer)

}
