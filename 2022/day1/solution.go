package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var blankLine = regexp.MustCompile(`(?m)^\D`)

func Part1Solution() {
	fileContents, err := utils.GetInput("/usr/local/projects/advent-of-code-go/day1/input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var sumSlice = make([]int, 0)

	splitString := blankLine.Split(fileContents, -1)

	for _, set := range splitString {

		numberSlice := strings.Split(set, "\n")

		sum := 0
		for _, number := range numberSlice {
			convInt, _ := strconv.Atoi(number)
			sum += convInt
		}

		sumSlice = append(sumSlice, sum)
	}

	var max int = 0
	for _, sum := range sumSlice {
		if sum > max {
			max = sum
		}
	}

	fmt.Println("Maximum Calories:", max)
}

func Part2Solution() {
	fileContents, err := utils.GetInput("/usr/local/projects/advent-of-code-go/day1/input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var sumSlice = make([]int, 0)

	splitString := blankLine.Split(fileContents, -1)

	for _, set := range splitString {

		numberSlice := strings.Split(set, "\n")

		sum := 0
		for _, number := range numberSlice {
			convInt, _ := strconv.Atoi(number)
			sum += convInt
		}

		sumSlice = append(sumSlice, sum)
	}

	sort.Ints(sumSlice)

	var sumOfThree int = 0
	for _, sum := range sumSlice[len(sumSlice)-3:] {
		sumOfThree += sum
	}

	fmt.Println("Sum of Top Three:", sumOfThree)
}

func main() {
	Part1Solution()
	Part2Solution()
}
