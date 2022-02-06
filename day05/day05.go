package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/mrtyunjaygr8/aoc_2020/common"
)

type seat struct {
	str string
	row int64
	col int64
	sid int64
	bin int64
}

func parse_lines(s string) seat {
	m_s := seat{}
	r := ""
	for _, x := range s[:7] {
		if string(x) == "F" {
			r = r + "0"
		} else if string(x) == "B" {
			r = r + "1"
		}
	}
	R, err := strconv.ParseInt(r, 2, 0)
	if err != nil {
		panic(err)
	}
	m_s.str = s
	m_s.row = R

	c := ""
	for _, x := range s[7:] {
		if string(x) == "R" {
			c = c + "1"
		} else if string(x) == "L" {
			c = c + "0"
		}
	}
	C, err := strconv.ParseInt(c, 2, 0)
	m_s.col = C
	m_s.sid = R*8 + C
	if err != nil {
		panic(err)
	}
	return m_s
}

func main() {
	data := common.Read_file("./input.txt")
	seats := []seat{}
	for _, line := range data {
		if len(line) == 0 {
			continue
		}
		seats = append(seats, parse_lines(line))
	}

	sort.Slice(seats, func(i, j int) bool {
		return seats[i].sid > seats[j].sid
	})
	for s := 0; s < len(seats)-1; s++ {
		if (seats[s].sid - seats[s+1].sid) > 1 {
			fmt.Println(seats[s].sid, seats[s+1].sid)
		}
	}
	// fmt.Println(len(seats), seats[0].sid, seats[len(seats)-1].sid)
}
