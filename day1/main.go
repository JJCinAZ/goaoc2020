package main

import "fmt"

func main() {
	part1()
	part2()
}

func part1() {
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i]+input[j] == 2020 {
				fmt.Printf("%d + %d = 2020, result=%d\n", input[i], input[j], input[i]*input[j])
			}
		}
	}
}

func part2() {
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			for k := j + 1; k < len(input); k++ {
				if input[i]+input[j]+input[k] == 2020 {
					fmt.Printf("%d + %d + %d = 2020, result=%d\n",
						input[i], input[j], input[k], input[i]*input[j]*input[k])
				}
			}
		}
	}
}
