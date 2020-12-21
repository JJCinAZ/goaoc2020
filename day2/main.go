package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	part2()
}

func part1() {
	var (
		r1         = regexp.MustCompile(`^(\d+)-(\d+)\s(\S):\s(\S+)$`)
		validCount int
	)

	for _, s := range input {
		a := r1.FindStringSubmatch(s)
		if a == nil {
			fmt.Printf("unable to match on input '%s'\n", s)
			os.Exit(1)
		}
		min, _ := strconv.Atoi(a[1])
		max, _ := strconv.Atoi(a[2])
		if isValidPart1(a[4], min, max, []rune(a[3])[0]) {
			validCount++
		}
	}
	fmt.Printf("Found %d valid\n", validCount)
}

func isValidPart1(password string, min, max int, mustContain rune) bool {
	found := 0
	for _, c := range password {
		if c == mustContain {
			found++
		}
	}
	if found >= min && found <= max {
		return true
	}
	return false
}

func part2() {
	var (
		r1         = regexp.MustCompile(`^(\d+)-(\d+)\s(\S):\s(\S+)$`)
		validCount int
	)

	for _, s := range input {
		a := r1.FindStringSubmatch(s)
		if a == nil {
			fmt.Printf("unable to match on input '%s'\n", s)
			os.Exit(1)
		}
		p1, _ := strconv.Atoi(a[1])
		p2, _ := strconv.Atoi(a[2])
		if isValidPart2(a[4], p1-1, p2-1, []rune(a[3])[0]) {
			validCount++
		}
	}
	fmt.Printf("Found %d valid\n", validCount)
}

func isValidPart2(password string, p1, p2 int, reqChar rune) bool {
	found := 0
	passwordR := []rune(password)
	if passwordR[p1] == reqChar {
		found++
	}
	if passwordR[p2] == reqChar {
		found++
	}
	return found == 1
}
