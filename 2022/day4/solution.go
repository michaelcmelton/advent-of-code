package main

import (
	"advent-of-code-go/utils"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var regex = regexp.MustCompile(`(?P<startOne>\d+)-(?P<endOne>\d+),(?P<startTwo>\d+)-(?P<endTwo>\d+)`)

func CreateRegion(start, end int) string {
	var region string
	for i := start; i <= end; i++ {
		region = region + strconv.Itoa(i)
	}
	return region
}

func Part1Solution() {
	fileContents, err := utils.GetInput(".\\input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(strings.NewReader(fileContents))

	var contiguousSum int
	for scanner.Scan() {
		input := scanner.Text()
		m := regex.FindStringSubmatch(input)
		var bounds map[string]int = make(map[string]int)
		for i, name := range regex.SubexpNames() {
			if i != 0 && name != "" {
				num, _ := strconv.Atoi(m[i])
				bounds[name] = num
			}
		}

		if bounds["startOne"] <= bounds["startTwo"] && bounds["endOne"] >= bounds["endTwo"] || bounds["startOne"] >= bounds["startTwo"] && bounds["endOne"] <= bounds["endTwo"] {
			contiguousSum++
		}
	}

	fmt.Println("Number of Contiguous Regions: ", contiguousSum)
}

func Part2Solution() {
	fileContents, err := utils.GetInput(".\\input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(strings.NewReader(fileContents))

	var contiguousSum int
	for scanner.Scan() {
		input := scanner.Text()
		m := regex.FindStringSubmatch(input)
		var bounds map[string]int = make(map[string]int)
		for i, name := range regex.SubexpNames() {
			if i != 0 && name != "" {
				num, _ := strconv.Atoi(m[i])
				bounds[name] = num
			}
		}

		if bounds["startOne"] <= bounds["endTwo"] && bounds["endOne"] >= bounds["startTwo"] {
			contiguousSum++
		}
	}

	fmt.Println("Number of Overlapped Regions: ", contiguousSum)
}

func main() {
	Part1Solution()
	Part2Solution()
}
