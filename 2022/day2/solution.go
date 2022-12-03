package main

// Rock, Paper, Scissors
// A, X = Rock,
// B, Y = Paper,
// C, Z = Scissors

import (
	"advent-of-code-go/utils"
	"fmt"
	"os"
	"strings"
)

func calcScore(theirChoice int, myChoice int) int {
	outcome := (((myChoice - theirChoice) % 3) + 3) % 3

	switch outcome {
	case 0: // outcome is a tie.
		return myChoice + 3
	case 1: // outcome is a win.
		return myChoice + 6
	case 2: // outcome is a loss.
		return myChoice + 0
	default:
		return myChoice
	}
}

func format(letter string) int {
	switch letter {
	// Rock
	case "A", "X":
		return 1
	// Paper
	case "B", "Y":
		return 2
	//Scissors
	case "C", "Z":
		return 3
	default:
		return 0
	}
}

func calcChoice(playerSelect string, outcome string) int {
	switch playerSelect {
	// Rock
	case "A":
		switch outcome {
		// Lose
		case "X":
			return 3
		// Draw
		case "Y":
			return 1
		// Win
		case "Z":
			return 2
		default:
			return 0
		}
	// Paper
	case "B":
		switch outcome {
		case "X":
			return 1
		case "Y":
			return 2
		case "Z":
			return 3
		default:
			return 0
		}
	case "C":
		switch outcome {
		case "X":
			return 2
		case "Y":
			return 3
		case "Z":
			return 1
		default:
			return 0
		}
	default:
		return 0
	}
}

func Part1Solution() {
	fileContents, err := utils.GetInput("/usr/local/projects/advent-of-code-go/2022/day2/input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	matches := strings.Split(fileContents, "\n")

	var sum int = 0

	for i, match := range matches {
		choices := strings.Split(match, " ")
		if i == 0 {
			fmt.Println(calcScore(format(choices[0]), format(choices[1])))
		}
		sum += calcScore(format(choices[0]), format(choices[1]))
	}

	fmt.Println("Total Score according to strategy guide: ", sum)
}

func Part2Solution() {
	fileContents, err := utils.GetInput("/usr/local/projects/advent-of-code-go/2022/day2/input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	matches := strings.Split(fileContents, "\n")

	var sum int = 0

	for i, match := range matches {
		choices := strings.Split(match, " ")
		if i == 0 {
			fmt.Println(calcScore(format(choices[0]), calcChoice(choices[0], choices[1])))
		}
		sum += calcScore(format(choices[0]), calcChoice(choices[0], choices[1]))
	}

	fmt.Println("Total Score according to strategy guide with Outcomes: ", sum)
}

func main() {
	Part1Solution()
	Part2Solution()
}
