package main

import (
	"bufio"
	"strings"
  "advent-of-code-go/utils"
  "fmt"
"os"
)

type Coordinate struct {
  XPos int
  YPos int
}

func (c *Coordinate) MoveHead(direction string) {
  switch direction {
  case "R":
    c.XPos++
    break
  case "L":
    c.XPos--
    break
  case "U":
    c.YPos++
    break
  case "D":
    c.YPos--
    break
}
}

func (c *Coordinate) MoveTail(head *Coordinate) {
  if c.XPos - head.XPos > 1 && c.YPos == head.YPos {
    c.XPos++
  }

  if c.XPos - head.XPos < -1 && c.YPos == head.YPos {
    c.XPos--
  }

  if c.YPos - head.YPos > 1 && c.XPos == head.XPos {
   c.YPos++
  }

  if c.YPos - head.YPos < -1 && c.XPos == head.XPos {
   c.YPos--
  }

}

func Part1Solution() {
 fileContents, err := utils.GetInput("./test.txt")

  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  scanner := bufio.NewScanner(strings.NewReader(fileContents))  

  head := &Coordinate{
    XPos: 0,
    YPos: 0,
  }

  tail := &Coordinate{
    XPos: 0,
    YPos: 0,
  }

  for scanner.Scan() {
    var qty int
    var direction string

    fmt.Sscanf(scanner.Text(), "%s %d", &direction, &qty)

    for i := 1; i <=qty; i++ {
      head.MoveHead(direction)
      tail.MoveTail(head)
    }
  }
}

func main() {
  Part1Solution()
}
