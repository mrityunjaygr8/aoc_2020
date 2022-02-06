package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/mrtyunjaygr8/aoc_2020/common"
)

type tree map[string][]string

func main() {
	data := common.Read_file("./input.txt")
	tree := tree{}
	for _, x := range data {
		if len(x) == 0 {
			continue
		}
		r_holder, _ := regexp.Compile(`([\w]+ [\w]+) bags`)
		r_holdee, _ := regexp.Compile(`([\d]) ([\w]+ [\w]+) bags?`)

		m_holder := r_holder.FindStringSubmatch(x)
		m_holdee := r_holdee.FindAllStringSubmatch(x, -1)
		// fmt.Println(m_holder, m_holdee)
		holder := strings.ReplaceAll(m_holder[1], " ", "-")
		if len(tree[holder]) == 0 {
			tree[holder] = make([]string, 0)
		}
		for _, y := range m_holdee {
			// tree[m_holder[1]] = append(tree[m_holder[1]], y[2])
			holdee := strings.ReplaceAll(y[2], " ", "-")
			tree[holdee] = append(tree[holdee], holder)
		}
	}
	search := "shiny-gold"

	options := search_tree(search, tree)
	fmt.Println(options)

}

func search_tree(term string, tree tree) int {
	options := 1
	// fmt.Println(term)
	for _, x := range tree[term] {
		// fmt.Println("\t", x)
		// fmt.Println("\t\t", tree[x])
		options += search_tree(x, tree)
	}
	return options
}
