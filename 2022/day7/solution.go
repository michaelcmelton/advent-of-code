package main

import (
	"advent-of-code-go/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type FileType uint

const (
	DIR FileType = iota
	FILE
)

type File struct {
	Name     string
	Type     FileType
	Size     int
	Children []*File
	Parent   *File
}

func (f *File) CalculateFileSize(childrenFiles []*File) {
	sum := 0
	for _, child := range childrenFiles {
		if child.Type == DIR && child.Size == 0 {
			child.CalculateFileSize(child.Children)
			sum += child.Size
		} else if child.Type == DIR && child.Size != 0 {
			sum += child.Size
		} else if child.Type == FILE {
			sum += child.Size
		}
	}

	f.Size = sum
}

func (f *File) AddChild(isDir bool, name string, size int) {
	if isDir {
		f.Children = append(f.Children, &File{
			Name:     name,
			Type:     DIR,
			Size:     0,
			Children: make([]*File, 0),
			Parent:   f,
		})
	} else {
		f.Children = append(f.Children, &File{
			Name:     name,
			Type:     FILE,
			Size:     size,
			Children: nil,
			Parent:   f,
		})
	}
}

type FileHistory []*File

func (fh *FileHistory) ChangeDirectory(command string) {
	if command == "$ cd .." {
		(*fh) = (*fh)[:len(*fh)-1]
	} else {
		dirName := strings.Split(command, " ")[2]
		for _, child := range (*fh)[len(*fh)-1].Children {
			if child.Name == dirName {
				(*fh) = append((*fh), child)
				return
			}
		}
	}
}

func SumLessThan10K(children []*File) int {
	sum := 0

	for _, child := range children {
		if child.Type == DIR && child.Size <= 100000 {
			sum = sum + child.Size
			sum = sum + SumLessThan10K(child.Children)
		} else {
			sum = sum + SumLessThan10K(child.Children)
		}
	}

	return sum
}

func TotalFreeSpace(file *File, totalUnused int, fileSizes *[]int) {
	neededSpace := 30000000 - totalUnused

	if file.Size >= neededSpace && file.Type == DIR {
		*fileSizes = append(*fileSizes, file.Size)
	}

	for _, child := range file.Children {
		TotalFreeSpace(child, totalUnused, fileSizes)
	}
}

func Part1Solution() {
	fileContents, err := utils.GetInput("./input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	root := &File{
		Name:     "/",
		Type:     DIR,
		Size:     0,
		Children: []*File{},
		Parent:   nil,
	}

	history := FileHistory{root}

	scanner := bufio.NewScanner(strings.NewReader(fileContents))

	for scanner.Scan() {
		if scanner.Text() != "$ cd /" && scanner.Text() != "$ ls" {
			if strings.Contains(scanner.Text(), "$ cd") {
				history.ChangeDirectory(scanner.Text())
			} else if strings.Contains(scanner.Text(), "dir ") {
				history[len(history)-1].AddChild(true, strings.Split(scanner.Text(), " ")[1], 0)
			} else {
				fileSize, err := strconv.Atoi(strings.Split(scanner.Text(), " ")[0])
				if err != nil {
					fmt.Println(fmt.Errorf("error occurred: %w", err))
					os.Exit(2)
				}
				history[len(history)-1].AddChild(false, strings.Split(scanner.Text(), " ")[1], fileSize)
			}
		}
	}

	root.CalculateFileSize(root.Children)
	fmt.Println(SumLessThan10K(root.Children))
}

func Part2Solution() {
	fileContents, err := utils.GetInput("./input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	root := &File{
		Name:     "/",
		Type:     DIR,
		Size:     0,
		Children: []*File{},
		Parent:   nil,
	}

	history := FileHistory{root}

	scanner := bufio.NewScanner(strings.NewReader(fileContents))

	for scanner.Scan() {
		if scanner.Text() != "$ cd /" && scanner.Text() != "$ ls" {
			if strings.Contains(scanner.Text(), "$ cd") {
				history.ChangeDirectory(scanner.Text())
			} else if strings.Contains(scanner.Text(), "dir ") {
				history[len(history)-1].AddChild(true, strings.Split(scanner.Text(), " ")[1], 0)
			} else {
				fileSize, err := strconv.Atoi(strings.Split(scanner.Text(), " ")[0])
				if err != nil {
					fmt.Println(fmt.Errorf("error occurred: %w", err))
					os.Exit(2)
				}
				history[len(history)-1].AddChild(false, strings.Split(scanner.Text(), " ")[1], fileSize)
			}
		}
	}

	root.CalculateFileSize(root.Children)
	totalUnused := 70000000 - root.Size
	var fileSizes []int = make([]int, 0)
	TotalFreeSpace(root, totalUnused, &fileSizes)

	min := 0

	for _, size := range fileSizes {
		if min == 0 {
			min = size
		}

		if size < min {
			min = size
		}
	}

	fmt.Println(min)
}

func main() {
	Part1Solution()
	Part2Solution()
}
