package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/mrtyunjaygr8/aoc_2020/common"
)

type password_policy struct {
	min  int
	max  int
	char string
}

type password struct {
	policy   password_policy
	password string
	valid    bool
}

func (p *password) validate() {
	count := strings.Count(p.password, p.policy.char)
	p.valid = p.policy.min <= count && count <= p.policy.max
}

func (p *password) validate_2() {
	v1 := string(p.password[p.policy.min-1]) == p.policy.char
	v2 := string(p.password[p.policy.max-1]) == p.policy.char
	p.valid = (v1 || v2) == true && (v1 && v2) == false
}

func parse_line(line string) password {
	re, _ := regexp.Compile(`(?P<min>[\d]+)-(?P<max>[\d]+) (?P<char>[\w]): (?P<pass>[\w]+)`)
	match := re.FindStringSubmatch(line)

	min, err := strconv.Atoi(match[1])
	if err != nil {
		log.Fatal("err parsing min")
	}
	max, err := strconv.Atoi(match[2])
	if err != nil {
		log.Fatal("err parsing max")
	}

	pass_pol := password_policy{min: min, max: max, char: match[3]}
	pass := password{policy: pass_pol, password: match[4]}

	return pass

}

func main() {
	data := common.Read_file("./input.txt")
	var pass_s []password

	for _, line := range data {
		// fmt.Println(line)
		pass_s = append(pass_s, parse_line(line))
	}

	count := 0
	for _, pass := range pass_s {
		pass.validate_2()
		if pass.valid {
			count++
		}
	}

	fmt.Println(count)
}
