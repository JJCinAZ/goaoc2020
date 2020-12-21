package main

import (
	"fmt"
)

type Answers string

type Group struct {
	People []Answers
}

func main() {
	groups := processData(input)
	part1(groups)
	part2(groups)
}

func part1(groups []Group) {
	count := 0
	for _, g := range groups {
		count += countUniqueAnswers(g)
	}
	fmt.Println(count)
}

func part2(groups []Group) {
	count := 0
	for _, g := range groups {
		count += countUnanomAnswers(g)
	}
	fmt.Println(count)
}

func countUniqueAnswers(grp Group) int {
	x := make(map[rune]struct{})
	for _, p := range grp.People {
		for _, a := range p {
			if _, found := x[a]; !found {
				x[a] = struct{}{}
			}
		}
	}
	return len(x)
}

func countUnanomAnswers(grp Group) int {
	x := make(map[rune]int)
	for _, p := range grp.People {
		for _, a := range p {
			if _, found := x[a]; !found {
				x[a] = 1
			} else {
				x[a] = x[a] + 1
			}
		}
	}
	c := 0
	for _, i := range x {
		if i == len(grp.People) {
			c += 1
		}
	}
	return c
}

func processData(input []string) []Group {
	var (
		grp Group
	)
	groups := make([]Group, 0)
	grp.People = make([]Answers, 0)
	for _, l := range input {
		if len(l) == 0 {
			if len(grp.People) > 0 {
				groups = append(groups, grp)
			}
			grp.People = make([]Answers, 0)
			continue
		}
		grp.People = append(grp.People, Answers(l))
	}
	if len(grp.People) > 0 {
		groups = append(groups, grp)
	}
	fmt.Printf("Read %d groups\n", len(groups))
	return groups
}
