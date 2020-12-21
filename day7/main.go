package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Rule struct {
	Contains map[string]int // map[<color>]<quantity>
}

var (
	Rules map[string]Rule
)

func main() {
	Rules = processData(input)
	fmt.Printf("read %d rules\n", len(Rules))
	part1("shiny gold")
	part2("shiny gold")
}

func part1(targetColor string) {
	count := 0
	for _, rule := range Rules {
		if canContain1(rule, targetColor) {
			count++
		}
	}
	fmt.Printf("Part 1: %d\n", count)
}

func canContain1(rule Rule, targetColor string) bool {
	for color, _ := range rule.Contains {
		if color == targetColor {
			return true
		} else if r2, found := Rules[color]; found {
			if canContain1(r2, targetColor) {
				return true
			}
		}
	}
	return false
}

func part2(targetColor string) {
	if rule, found := Rules[targetColor]; found {
		fmt.Printf("Part 2: %d\n", countContainedBags(rule)-1)
	} else {
		panic(targetColor)
	}
}

func countContainedBags(rule Rule) int {
	totalQuan := 1
	for color, quan := range rule.Contains {
		if r2, found := Rules[color]; found {
			totalQuan += quan * countContainedBags(r2)
		}
	}
	return totalQuan
}

func processData(input []string) map[string]Rule {
	var (
		regx1 = regexp.MustCompile(`^(.+) bags contain (.+)$`)
		regx2 = regexp.MustCompile(`(\d+) (.+) bags?`)
	)
	rules := make(map[string]Rule)
	for _, l := range input {
		if a := regx1.FindStringSubmatch(l); a != nil {
			newRule := Rule{Contains: make(map[string]int)}
			if a[2] != "no other bags." {
				if strings.HasSuffix(a[2], ".") {
					a[2] = a[2][:len(a[2])-1]
				}
				for _, c := range strings.Split(a[2], ", ") {
					if a2 := regx2.FindStringSubmatch(c); a2 != nil {
						if count, err := strconv.Atoi(a2[1]); err != nil {
							panic(fmt.Sprintf("Invalid count in rule: %#v", a2))
						} else {
							newRule.Contains[a2[2]] = count
						}
					}
				}
			}
			rules[a[1]] = newRule
		}
	}
	return rules
}
