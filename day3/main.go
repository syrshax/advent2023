package main

import (
	"context"
	"fmt"
	"os"
	"strings"
)

type Data struct {
	Lines []string
}

type Engine struct {
	Pieces int
	Number string
}

func isDigit(byte byte) bool {
	return (byte >= '0' && byte <= '9')
}

func symbolFinder(byte byte) bool {
	return (byte != '.')
}

func isInRange(index int, string string) bool {
	if index >= len(string) || index < 0 {
		return false
	}
	return true
}

func parseInput(filename string) (*Data, error) {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Read input file: %w", err)
	}

	d := &Data{}
	lines := strings.Split(string(buf), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		e := Engine{}
		d.Lines = append(d.Lines, line)
		for x := 0; x < len(line); x++ {
			actual := x
			pre := x - 1
			post := x + 1
			if isInRange(actual, line) && isInRange(pre, line) && isInRange(post, line) {
				if isDigit(line[actual]) {
					e.Number += string(line[actual])
				}
				e.Number = ""
			}
			fmt.Println("Number", string(e.Number))
		}
	}
	return d, nil
}

func run(_ context.Context) error {
	d, err := parseInput("input.txt")
	if err != nil {
		return fmt.Errorf("parse %w", err)
	}

	fmt.Println("Lines printed", len(d.Lines))
	fmt.Println("Data", d.Lines)

	return nil
}

func main() {

	if err := run(context.Background()); err != nil {
		fmt.Println("Fail", err.Error())
		return
	}

	fmt.Println("success!")
}
