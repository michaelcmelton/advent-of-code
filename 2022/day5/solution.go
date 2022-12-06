package main

import (
	"advent-of-code-go/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getStartingState(state string) [][]string {
	var intermediate [][]string

	result := [][]string{{}, {}, {}, {}, {}, {}, {}, {}, {}}

	scanner := bufio.NewScanner(strings.NewReader(state))

	for scanner.Scan() {
		var line []string
		input := scanner.Text()
		for i, _ := range input {
			if i%4 == 0 {
				line = append(line, string(input[i+1]))
			}
		}
		intermediate = append(intermediate, line)
		line = []string{}
	}

	for i := 0; i < len(intermediate)-1; i++ {
		for j := 0; j < len(intermediate[i]); j++ {
			if intermediate[i][j] != " " {
				result[j] = append(result[j], intermediate[i][j])
			}
		}
	}

	return result
}

func makeSingularMoves(moves string, startState [][]string) [][]string {
	var qty, from, to int
	scanner := bufio.NewScanner(strings.NewReader(moves))

	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "move %d from %d to %d", &qty, &from, &to)

		for i := 0; i < qty; i++ {
			startState[to-1] = append([]string{startState[from-1][0]}, startState[to-1]...)
			startState[from-1] = startState[from-1][1:]
		}
	}

	return startState
}

func makeMultipleMoves(moves string, startState [][]string) [][]string {
	var qty, from, to int
	scanner := bufio.NewScanner(strings.NewReader(moves))

	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "move %d from %d to %d", &qty, &from, &to)

		intermediate := make([]string, qty)

		copy(intermediate, startState[from-1][:qty])
		startState[to-1] = append(intermediate, startState[to-1]...)
		startState[from-1] = startState[from-1][qty:]
	}

	return startState
}

func Part1Solution() {
	fileContents, err := utils.GetInput("./input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var result string
	fileParts := strings.Split(fileContents, "\n\n")
	startingState := getStartingState(fileParts[0])
	finalState := makeSingularMoves(fileParts[1], startingState)
	for l := 0; l < len(finalState); l++ {
		result = result + finalState[l][0]
	}

	fmt.Println(result)
}

func Part2Solution() {
	fileContents, err := utils.GetInput("./input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var result string
	fileParts := strings.Split(fileContents, "\n\n")
	startingState := getStartingState(fileParts[0])
	finalState := makeMultipleMoves(fileParts[1], startingState)
	for l := 0; l < len(finalState); l++ {
		result = result + finalState[l][0]
	}

	fmt.Println(result)
}

func main() {
	Part1Solution()
	Part2Solution()
}
