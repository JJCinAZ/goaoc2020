package main

import (
	"fmt"
	"sort"
)

func main() {
	part1()
	part2()
}

func part1() {
	highest := 0
	for _, boardingpass := range input {
		row, col := decodepass(boardingpass)
		seatid := row*8 + col
		if seatid > highest {
			highest = seatid
		}
	}
	fmt.Printf("Highest ID: %d\n", highest)
}

func part2() {
	ids := make([]int, 0, 1024)
	for _, boardingpass := range input {
		row, col := decodepass(boardingpass)
		seatid := row*8 + col
		ids = append(ids, seatid)
	}
	sort.Ints(ids)
	for i := 1; i < len(ids); i++ {
		if ids[i] != (ids[i-1] + 1) {
			fmt.Printf("MISSING\n")
		}
		fmt.Println(ids[i])
	}
}

func decodepass(pass string) (row, col int) {
	rl, rh := 0, 127
	cl, ch := 0, 7
	for _, c := range pass {
		switch c {
		case 'F':
			rh -= (rh - rl + 1) / 2
		case 'B':
			rl += (rh - rl + 1) / 2
		case 'L':
			ch -= (ch - cl + 1) / 2
		case 'R':
			cl += (ch - cl + 1) / 2
		}
	}
	row = rl
	col = cl
	return
}
