package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//fmt.Println(partOne())
	fmt.Println(partTwo())
}

func partOne() int {
	instructions := parseInput()
	dialPosition := 50
	numZeroes := 0

	for _, instruction := range instructions {
		dialPosition += instruction
		dialPosition %= 100

		if dialPosition == 0 {
			numZeroes++
		}
	}

	return numZeroes
}

func partTwo() int {
	instructions := parseInput()
	dialPosition := 50
	numZeroes := 0

	for _, instruction := range instructions {
		if dialPosition == 0 && instruction < 0 {
			dialPosition = 100
		}

		dialPosition += instruction

		if dialPosition <= 0 {
			dialPosition = 100*((dialPosition/-100)+1) + modDialPosition(dialPosition)
		}

		numZeroes += dialPosition / 100
		dialPosition -= (dialPosition / 100) * 100
	}

	return numZeroes
}

func modDialPosition(n int) int {
	const m = 100
	return ((n % m) + m) % m
}

func parseInput() []int {
	ret := []int{}
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ret = append(ret, parseLine(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("Encountered error while scanning: %s", err))
	}

	return ret
}

func parseLine(line string) int {
	line = strings.ReplaceAll(line, "L", "-")
	line = strings.ReplaceAll(line, "R", "")
	ret, err := strconv.Atoi(line)
	if err != nil {
		panic("could not parse line as int")
	}

	return ret
}
