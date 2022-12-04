package main

import (
	"advent-of-code-go/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findItem(compartment1 string, compartment2 string) string {
	items := strings.Split(compartment1, "")

	for _, item := range items {
		if strings.Contains(compartment2, item) {
			return item
		}
	}

	return ""
}

func getGroupItem(group []string) string {
	items := strings.Split(group[0], "")

	for _, item := range items {
		if strings.Contains(group[1], item) && strings.Contains(group[2], item) {
			return item
		}
	}

	fmt.Println(fmt.Errorf("nothing found in group below: %s", group))
	return ""
}

func assignPriority(item string) int {
	runes := []rune(item)

	if runes[0] > 91 {
		return int(runes[0] - 96)
	} else {
		return int(runes[0] - 38)
	}
}

func Part1Solution() {
	fileContents, err := utils.GetInput("./2022/day3/input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var prioritySum int
	for _, rucksack := range strings.Split(fileContents, "\n") {
		compartment1 := rucksack[0 : len(rucksack)/2]
		compartment2 := rucksack[len(rucksack)/2:]

		item := findItem(compartment1, compartment2)
		priority := assignPriority(item)
		prioritySum = priority + prioritySum
	}

	fmt.Println("Sum of priorities: ", prioritySum)
}

func Part2Solution() {
	fileContents, err := utils.GetInput("./2022/day3/input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(strings.NewReader(fileContents))

	var block []string
	var blocks [][]string

	for scanner.Scan() {
		block = append(block, scanner.Text())
		if len(block) == 3 {
			blocks = append(blocks, block)
			block = []string{}
		}
	}

	var sumPriority int

	for _, block := range blocks {
		item := getGroupItem(block)
		sumPriority += assignPriority(item)
	}

	fmt.Println("Sum of priority for grouped elves:", sumPriority)
}

func main() {
	Part1Solution()
	Part2Solution()
}
