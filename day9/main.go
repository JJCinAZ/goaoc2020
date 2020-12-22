package main

import (
	"fmt"
	"math"
)

func main() {
	n := findMissingPair(testinput, 5)
	min, max := findSumRange(testinput, n)
	fmt.Printf("%d, %d, %d\n", n, min, max)
	fmt.Println(findMissingPair(input, 25))
}

func findMissingPair(list []int, preamble int) int {
	idx := preamble
	for idx < len(list) {
		found := findSumPair(list[idx-preamble:idx], list[idx])
		if !found {
			return list[idx]
		}
		idx++
	}
	return -1
}

func findSumPair(list []int, sum int) bool {
	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i]+list[j] == sum {
				return true
			}
		}
	}
	return false
}

func findSumRange(list []int, sum int) (smallest, largest int) {
	for i := 0; i < len(list)-1; i++ {
		t := list[i]
		for j := i + 1; j < len(list); j++ {
			t += list[j]
			if t == sum {
				return minmax(list[i:j])
			} else if t > sum {
				break
			}
		}
	}
	return 0, 0
}

func minmax(list []int) (int, int) {
	min, max := int(math.MaxInt64), int(math.MinInt64)
	for _, n := range list {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max
}
