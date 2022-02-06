package day03

import (
	"fmt"

	"github.com/mrtyunjaygr8/aoc_2020/common"
)

func main() {
	data := common.Read_file("./input.txt")
	length_width := len(data[0])
	length_height := len(data)
	INCREMENT_RIGHT := 1
	INCREMENT_DOWN := 2
	COUNT_TREE := 0
	start := 0
	for x := INCREMENT_DOWN; x < length_height-1; x += INCREMENT_DOWN {
		start = (start + INCREMENT_RIGHT) % length_width
		if string(data[x][start]) == "#" {
			COUNT_TREE++
		}
	}
	fmt.Println(COUNT_TREE)
}
