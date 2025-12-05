package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(partTwo())
}

func partOne() int {
	ret := 0

	ranges := getRanges()
	invalidIds := map[int]bool{}

	for _, r := range ranges {
		for _, id := range repeated(r, 2) {
			invalidIds[id] = true
		}
	}

	for id := range invalidIds {
		ret += id
	}

	return ret
}

func partTwo() int {
	ret := 0

	ranges := getRanges()

	invalidIds := map[int]bool{}

	for _, r := range ranges {
		for repeats := 2; repeats <= intLog10(r.High); repeats++ {
			for _, id := range repeated(r, repeats) {
				invalidIds[id] = true
			}
		}
	}

	for id := range invalidIds {
		ret += id
	}

	return ret
}

func repeated(r Range, numRepeats int) []int {
	digits := intLog10(r.Low)
	seqLen := (digits + numRepeats - 1) / numRepeats

	repeatedNum := mostSignificantDigits(r.Low, seqLen)
	if digits%numRepeats != 0 {
		repeatedNum = intPowBase10(seqLen - 1)
	}

	testNum := repeatNum(repeatedNum, numRepeats)

	ret := []int{}

	for testNum <= r.High {
		if testNum >= r.Low {
			ret = append(ret, testNum)
		}

		repeatedNum++
		testNum = repeatNum(repeatedNum, numRepeats)
	}

	return ret
}

type Range struct {
	Low, High int
}

func leastSignificantDigits(n, numDigits int) int {
	return n % intPowBase10(numDigits)
}

func mostSignificantDigits(n, numDigits int) int {
	totalDigits := intLog10(n)
	return n / intPowBase10(totalDigits-numDigits)
}

func repeatNum(n, reps int) int {
	ret := 0
	digits := intLog10(n)
	mult := 1
	multMult := intPowBase10(digits)

	for i := 0; i < reps; i++ {
		ret += n * mult
		mult *= multMult
	}

	return ret
}

func intLog10(n int) int {
	ret := 0
	for i := n; i > 0; i /= 10 {
		ret++
	}

	return ret
}

func intPowBase10(pow int) int {
	ret := 1
	for j := 0; j < pow; j++ {
		ret *= 10
	}

	return ret
}

func getRanges() []Range {
	ret := []Range{}

	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic("Couldn't read file")
	}

	atoiOrDie := func(n string) int {
		n = strings.TrimSpace(n)
		convertedNum, err := strconv.Atoi(n)
		if err != nil {
			panic(fmt.Sprintf("Couldn't convert string to int: %s", n))
		}

		return convertedNum
	}

	rangeStrings := strings.Split(string(b), ",")
	for _, rangeString := range rangeStrings {
		rangeVals := strings.Split(rangeString, "-")

		ret = append(ret,
			Range{
				Low:  atoiOrDie(rangeVals[0]),
				High: atoiOrDie(rangeVals[1]),
			},
		)
	}

	return ret
}
