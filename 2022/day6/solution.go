package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"os"
	"strings"
)

func checkCharacters(chars string) bool {
	runeSlice := strings.Split(chars, "")
	result := true

	for i := 0; i < len(runeSlice); i++ {
		for j := 0; j < len(runeSlice); j++ {
			if runeSlice[i] == runeSlice[j] && i != j {
				result = false
			}
		}
	}
	return result
}

func Part1Solution() {
	fileContents, err := utils.GetInput("./input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	stringLength := 3

	for i, _ := range fileContents {
		if i >= stringLength {
			result := checkCharacters(fileContents[i-stringLength : i+1])

			if result {
				fmt.Println(i + 1)
				break
			}
		}
	}
}

func Part2Solution() {
	fileContents, err := utils.GetInput("./input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	stringLength := 13

	for i, _ := range fileContents {
		if i >= stringLength {
			result := checkCharacters(fileContents[i-stringLength : i+1])

			if result {
				fmt.Println(i + 1)
				break
			}
		}
	}
}

func main() {
	Part1Solution()
	Part2Solution()
}
