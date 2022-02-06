package main

import (
	"fmt"

	"github.com/mrtyunjaygr8/aoc_2020/common"
)

type counter map[string]int
type checked map[string]bool

func init_checked() checked {
	checked_map := checked{}
	for x := 97; x <= 122; x++ {
		checked_map[string(x)] = false
	}
	return checked_map
}

func main() {
	data := common.Read_file("./input.txt")
	// checked_map := init_checked()
	count_map := counter{}
	local_counter := counter{}
	party := 0
	for _, x := range data {
		if len(x) == 0 {
			fmt.Print("mapping: ")
			// checked_map = init_checked()
			for z := 97; z <= 122; z++ {
				if local_counter[string(z)] == party {
					count_map[string(z)]++
					fmt.Print(string(z))
				}
			}
			fmt.Println(" :asd")
			fmt.Println()
			local_counter = counter{}
			party = 0
			continue
		}
		fmt.Println(x)
		party++
		for _, y := range x {
			local_counter[string(y)]++
		}
	}
	COUNT := 0
	for _, count := range count_map {
		COUNT += count
	}
	fmt.Println(COUNT)
}
