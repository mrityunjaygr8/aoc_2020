package main

import (
	"fmt"
	"regexp"

	"github.com/mrtyunjaygr8/aoc_2020/common"
)

type person struct {
	byr   string
	iyr   string
	eyr   string
	hgt   string
	hcl   string
	ecl   string
	pid   string
	cid   string
	valid bool
}

func (p *person) validate() {
	byr, err := common.Atoi(p.byr)
	if err != nil {
		return
	}
	if !(1920 <= byr && byr <= 2002) {
		return
	}
	iyr, err := common.Atoi(p.iyr)
	if err != nil {
		return
	}
	if !(2010 <= iyr && iyr <= 2020) {
		return
	}
	eyr, err := common.Atoi(p.eyr)
	if err != nil {
		return
	}
	if !(2020 <= eyr && eyr <= 2030) {
		return
	}

	if len(p.hgt) == 0 {
		return
	}
	if !(p.hgt[len(p.hgt)-2:] == "cm" || p.hgt[len(p.hgt)-2:] == "in") {
		return
	}
	if p.hgt[len(p.hgt)-2:] == "cm" {
		h, err := common.Atoi(p.hgt[:len(p.hgt)-2])
		if err != nil {
			return
		}
		if !(150 <= h && h <= 193) {
			return
		}
	}
	if p.hgt[len(p.hgt)-2:] == "in" {
		h, err := common.Atoi(p.hgt[:len(p.hgt)-2])
		if err != nil {
			return
		}
		if !(59 <= h && h <= 76) {
			return
		}
	}

	if len(p.hcl) != 7 {
		return
	}

	for _, x := range p.hcl[1:len(p.hcl)] {
		if !((48 <= x && x <= 57) || (97 <= x && x <= 102)) {
			return
		}
	}
	if !common.Contains([]string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}, p.ecl) {
		return
	}
	if len(p.pid) != 9 {
		return
	}
	_, err = common.Atoi(p.pid)
	if err != nil {
		return
	}
	p.valid = true

}

func parse_lines(lines []string) []person {
	r, _ := regexp.Compile(`([\w]+):([#\w]+)`)
	var peoples []person
	var tmp_p person
	for _, line := range lines {
		matches := r.FindAllStringSubmatch(line, -1)
		if len(matches) == 0 {
			peoples = append(peoples, tmp_p)
			tmp_p = person{}
			continue
		}
		for _, x := range matches {
			switch x[1] {
			case "byr":
				tmp_p.byr = x[2]
			case "iyr":
				tmp_p.iyr = x[2]
			case "eyr":
				tmp_p.eyr = x[2]
			case "hgt":
				tmp_p.hgt = x[2]
			case "hcl":
				tmp_p.hcl = x[2]
			case "ecl":
				tmp_p.ecl = x[2]
			case "pid":
				tmp_p.pid = x[2]
			case "cid":
				tmp_p.cid = x[2]
			}
		}
	}
	return peoples
}

func main() {
	data := common.Read_file("./input.txt")
	peoples := parse_lines(data)
	VALID := 0
	for _, x := range peoples {
		x.validate()
		if x.valid {
			VALID++
		}
	}
	fmt.Println(VALID)
}
