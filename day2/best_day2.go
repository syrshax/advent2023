package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Data struct {
	Lines       []string
	Games       []Game
	SumGame     int
	SumGameCube int
}

type Game struct {
	ID    int
	Blue  int
	Red   int
	Green int
}

type PossibleGames struct {
	Red      int
	Green    int
	Blue     int
	Possible bool
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
		if len(parts) != 2 {
			return nil, fmt.Errorf("Wrong parsing lines, bad split")
		}
		gameNumberStr := strings.TrimPrefix(parts[0], "Game ")
		gameNumber, err := strconv.Atoi(gameNumberStr)
		if err != nil {
			return nil, fmt.Errorf("Atoi %w", err)
		}

		g := Game{}
		g.ID = gameNumber

		sets := strings.Split(parts[1], ";")
		for _, set := range sets {
			for _, elem := range strings.Split(set, ",") {
				tmp := strings.Split(strings.TrimSpace(elem), " ")

				n, err := strconv.Atoi(tmp[0])
				if err != nil {
					return nil, fmt.Errorf("atoi2 %q: %w", elem, err)
				}
				switch tmp[1] {
				case "red":
					if g.Red < n {
						g.Red = n
					}
				case "green":
					if g.Green < n {
						g.Green = n
					}
				case "blue":
					if g.Blue < n {
						g.Blue = n
					}
				default:
					fmt.Printf("< %q\n", tmp[0])
				}
			}
		}
		d.Games = append(d.Games, g)
	}

	return d, nil
}

func run(_ context.Context) error {
	d, err := parseInput("input.txt")
	if err != nil {
		return fmt.Errorf("parse %w", err)
	}
	possibleGames := PossibleGames{}
	possibleGames.Red = 12
	possibleGames.Green = 13
	possibleGames.Blue = 14
	for _, g := range d.Games {
		possible := g.Red <= possibleGames.Red && g.Blue <= possibleGames.Blue && g.Green <= possibleGames.Green
		if possible {
			d.SumGame += g.ID
		}
		d.SumGameCube += g.Red * g.Blue * g.Green
	}
	fmt.Printf("Sum of game's ID which are possible: %d\n\n", d.SumGame)
	fmt.Printf("The cubes of game ids which are possible: %d\n\n", d.SumGameCube)
	return nil
}

func main() {

	if err := run(context.Background()); err != nil {
		fmt.Println("Fail", err.Error())
		return
	}

	fmt.Println("success")

}
