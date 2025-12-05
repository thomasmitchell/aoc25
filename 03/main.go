package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(partTwo())
}

func partOne() int {
	lines := parseInput()
	ret := 0

	for _, line := range lines {
		ret += solveLine(line, 2)
	}

	return ret
}

func partTwo() int {
	lines := parseInput()
	ret := 0

	for _, line := range lines {
		ret += solveLine(line, 12)
	}

	return ret
}

func solveLine(line []int, numBatteries int) int {
	acc := new(int)
	//Use helper func with accumulator to structure for tail call recursion
	solveLineWorker(line, numBatteries, acc)
	return *acc
}

func solveLineWorker(line []int, numBatteries int, acc *int) {
	if numBatteries <= 0 {
		return
	}

	val, idx := findFirstLargestBattery(line[0 : len(line)+1-numBatteries])
	*acc += intPowBase10(numBatteries-1) * val
	solveLineWorker(line[idx+1:], numBatteries-1, acc)
}

func findFirstLargestBattery(line []int) (int, int) {
	if len(line) == 1 {
		return line[0], 0
	}

	val, idx := findFirstLargestBattery(line[1:])
	if val > line[0] {
		return val, idx + 1
	}

	return line[0], 0
}

func intPowBase10(pow int) int {
	ret := 1
	for j := 0; j < pow; j++ {
		ret *= 10
	}

	return ret
}

func parseInput() [][]int {
	ret := [][]int{}
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ret = append(ret, parseLine(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("Encountered error while scanning: %s", err))
	}

	return ret
}

func parseLine(line string) []int {
	ret := []int{}
	for _, n := range line {
		nInt, err := strconv.Atoi(string(n))
		if err != nil {
			panic("could not parse line as int")
		}

		ret = append(ret, nInt)
	}
	return ret
}
