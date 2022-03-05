package main

import (
	"fmt"
	"regexp"

	"github.com/mrtyunjaygr8/aoc_2020/common"
)

func main() {
	data := common.Read_file("./input.txt")
	ins := len(data)
	fmt.Println(ins)
	idx := 0
	count := 0
	re := regexp.MustCompile(`([\w]+) (\+|\-)([\d]+)`)
	visited := make(map[int]bool)

	for idx < ins-1 {
		if visited[idx] {
			fmt.Println(count)
			break
		} else {
			visited[idx] = true
			items := re.FindStringSubmatch(data[idx])
			fmt.Println(idx, items[1], items[2], items[3])
			var update int
			c, err := common.Atoi(items[3])
			if err != nil {
				panic(err)
			}
			if items[2] == "+" {
				update = c
			} else if items[2] == "-" {
				update = -1 * c
			}
			if items[1] == "acc" {
				count += update
				idx += 1
				continue
			} else if items[1] == "nop" {
				idx += 1
				continue
			} else if items[1] == "jmp" {
				idx = idx + update
				continue
			}
		}
	}
	fmt.Println(count)
}
