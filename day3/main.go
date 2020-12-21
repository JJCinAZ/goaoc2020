package main

import "fmt"

func main() {
	fmt.Println("Part1:")
	fmt.Println(countTrees(3, 1))
	fmt.Println("\nPart2:")
	fmt.Println(
		countTrees(1, 1) *
			countTrees(3, 1) *
			countTrees(5, 1) *
			countTrees(7, 1) *
			countTrees(1, 2))
}

func countTrees(slopeRight, slopeDown int) int {
	var (
		r, c, trees int
	)
	c = slopeRight
	for r = slopeDown; r < len(input); r += slopeDown {
		l := input[r]
		cPrime := c % len(l)
		if l[cPrime] == '#' {
			trees++
		}
		c += slopeRight
	}
	return trees
}
