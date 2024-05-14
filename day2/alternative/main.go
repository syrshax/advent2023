package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Data struct {
	Lines []string
}

type Game struct {
	id    int
	Blue  int
	Red   int
	Green int
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

		d.Lines = append(d.Lines, line)
		parts := strings.Split(line, ":")
		gameNumberStr := strings.TrimPrefix(parts[0], "Game ")
		gameNumber, err := strconv.Atoi(gameNumberStr)
		if err != nil {
			return nil, fmt.Errorf("Atoi %q: w", gameNumberStr, err)
		}

		g := Game{}

	}

	return d, nil
}

func run(_ context.Context) error {
	d, err := parseInput("input.txt")
	if err != nil {
		return fmt.Errorf("parse %w", err)
	}
	fmt.Printf("Number of lines read: %d\n", len(d.Lines))
	return nil
}

func main() {

	if err := run(context.Background()); err != nil {
		fmt.Println("Fail", err.Error())
		return
	}

	fmt.Println("success")

}
