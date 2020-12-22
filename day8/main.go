package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Instruction struct {
	Op    string
	Parm1 int
}

var errLoopDetected = errors.New("loop detected")

func main() {
	pgm := processInput(input)
	fmt.Println("Part1:")
	fmt.Println(execute(pgm))
	fmt.Println("Part2:")
	fmt.Println(part2(pgm))
}

func execute(pgm []Instruction) (int, error) {
	acc := 0
	ip := 0
	hitcounts := make([]int, len(pgm))
	for {
		if ip >= len(pgm) {
			return acc, nil
		}
		if hitcounts[ip] > 0 {
			return acc, errLoopDetected
		}
		hitcounts[ip]++
		switch pgm[ip].Op {
		case "nop":
			ip++
		case "acc":
			acc += pgm[ip].Parm1
			ip++
		case "jmp":
			ip += pgm[ip].Parm1
		}
	}
}

func part2(pgm []Instruction) int {
	pgmcopy := make([]Instruction, len(pgm))
	testip := -1
	for {
		copy(pgmcopy, pgm)
		// Flip next nop/jmp
		for testip < len(pgmcopy) {
			testip++
			if pgmcopy[testip].Op == "nop" {
				pgmcopy[testip].Op = "jmp"
				break
			}
			if pgmcopy[testip].Op == "jmp" {
				pgmcopy[testip].Op = "nop"
				break
			}
		}
		lastacc, err := execute(pgmcopy)
		if err == nil {
			return lastacc
		}
	}
}

func processInput(input []string) []Instruction {
	a := make([]Instruction, 0)
	for _, l := range input {
		if i, err := strconv.Atoi(l[4:]); err == nil {
			a = append(a, Instruction{Op: l[0:3], Parm1: i})
		}
	}
	return a
}
