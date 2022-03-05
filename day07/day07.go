package main

import (
	"fmt"
	"regexp"

	"github.com/mrtyunjaygr8/aoc_2020/common"
)

type tree map[string]map[string]int
type visited map[string]bool

func main() {
	data := common.Read_file("./input.txt")
	tree := make(tree)
	for _, x := range data {
		if len(x) == 0 {
			continue
		}
		r_holder, _ := regexp.Compile(`([\w]+ [\w]+) bags`)
		r_holdee, _ := regexp.Compile(`([\d]) ([\w]+ [\w]+) bags?`)

		m_holder := r_holder.FindStringSubmatch(x)
		m_holdee := r_holdee.FindAllStringSubmatch(x, -1)
		// fmt.Println(m_holder, m_holdee)
		temp := make(map[string]int)
		for _, x := range m_holdee {
			str, err := common.Atoi(x[1])
			if err != nil {
				panic(err)
			}
			temp[x[2]] = str
		}

		tree[m_holder[1]] = temp

	}

	outer := make([]string, 0)
	// for part 1
	for x := range tree {
		if tree.search_for_bag(x) {
			outer = append(outer, x)
		}
	}
	fmt.Println(len(outer))

	// for part 2
	fmt.Println(tree.held_count("shiny gold") - 1)

}

func (tree tree) search_for_bag(term string) bool {
	if _, ok := tree[term]["shiny gold"]; ok {
		return true
	}

	for subBags := range tree[term] {
		if tree.search_for_bag(subBags) {
			return true
		}
	}

	return false
}

func (tree tree) held_count(entry string) int {
	bags := 1

	for subBag, count := range tree[entry] {
		bags += count * tree.held_count(subBag)

	}

	return bags
}
