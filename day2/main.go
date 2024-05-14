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
func isComma(byte byte) bool {
	return byte == ','
}
func isOtherTurn(byte byte) bool {
	return byte == ';'
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

	var counter string
	var slicedLine []string //Game 1 [0] || 2 red, 3 blue, 4 green [1]

	m := make(map[string]int)
	m["red"] = 0
	m["blue"] = 0
	m["green"] = 0
	var gameIDpossible []string
	var sumgame int

	for _, lines := range fileLines {

		slicedLine = strings.Split(lines, ":")
		counter = slicedLine[0][5:]

		var digits string
		var color string
		var listDigits []string
		var listColor []string
		problemColors := []string{"red", "green", "blue"}
		var colorResult []string
		var digitsResult []string
		max_red := 12
		max_green := 13
		max_blue := 14
		var possibleGame bool
		possibleGame = false

	GameCounter:
		for i := 0; i < len(slicedLine[1]); i++ {

			if isDigit(slicedLine[1][i]) {
				digits += string(slicedLine[1][i])
			}
			if isLetter(slicedLine[1][i]) {
				color += string(slicedLine[1][i])
			}
			if isComma(slicedLine[1][i]) {
				listDigits = append(listDigits, digits)
				listColor = append(listColor, color)
				digits = ""
				color = ""
			}
			if isOtherTurn(slicedLine[1][i]) || i == len(slicedLine[1])-1 {
				listDigits = append(listDigits, digits)
				listColor = append(listColor, color)
				digits = ""
				color = ""
				for i := 0; i < len(listColor); i++ {
					for _, substring := range problemColors {
						if strings.Contains(listColor[i], substring) {
							colorResult = append(colorResult, listColor[i])
							digitsResult = append(digitsResult, listDigits[i])
						}
					}
				}
			}

			for i := 0; i < len(colorResult); i++ {
				if colorResult[i] == "red" {
					numb1, err := strconv.Atoi(digitsResult[i])
					fmt.Println("error sir", digitsResult[i], err)
					if numb1 > max_red {
						possibleGame = false
						break GameCounter
					}
					possibleGame = true
				}
				if colorResult[i] == "blue" {
					numb1, err := strconv.Atoi(digitsResult[i])
					fmt.Println("error sir", digitsResult[i], err)
					if numb1 > max_blue {
						possibleGame = false
						break GameCounter
					}
					possibleGame = true
				}
				if colorResult[i] == "green" {
					numb1, err := strconv.Atoi(digitsResult[i])
					fmt.Println("error sir", digitsResult[i], err)
					if numb1 > max_green {
						possibleGame = false
						break GameCounter
					}
					possibleGame = true
				}
			}
		}
		if possibleGame {
			gameIDpossible = append(gameIDpossible, counter)
		}

		fmt.Println("Lista de digitos: ", listDigits)
		fmt.Println("Lista de colores: ", listColor)
		fmt.Println("Possible games ID's: ", gameIDpossible)
	}
	for _, gameid := range gameIDpossible {
		numbergame, err := strconv.Atoi(gameid)
		fmt.Println("numbergame", numbergame)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("the sumgame?", sumgame)
		sumgame = sumgame + numbergame

	}

	fmt.Println("The sum of games possble ID's are: ", sumgame)
}
