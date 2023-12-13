package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	pt1, pt2 := leggiFile()
	fmt.Printf("part1: %d, part2: %d\n", pt1, pt2)
}

func leggiFile() (tot1, tot2 int64) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ranges := parsStringToInts(strings.Split(scanner.Text(), " "))
		n1, n2 := calcRange(ranges)
		tot1 += n1
		tot2 += n2
	}
	return
}

func calcRange(ranges []int64) (sum1, sum2 int64) {
	var lastValue []int64
	var firstValue []int64
	for !checkAllZero(ranges) {
		var newRanges []int64
		for i := 0; i < len(ranges)-1; i++ {
			diff := ranges[i+1] - ranges[i]
			newRanges = append(newRanges, diff)
		}
		firstValue = append(firstValue, ranges[0])
		lastValue = append(lastValue, ranges[len(ranges)-1])
		ranges = newRanges
	}

	for _, val := range lastValue {
		sum1 += val
	}

	prev := firstValue[len(firstValue)-1]
	for i := len(firstValue) - 1; i > 0; i-- {
		prev = firstValue[i-1] - prev
	}
	sum2 = prev
	return
}

func checkAllZero(ranges []int64) bool {
	for _, val := range ranges {
		if val != 0 {
			return false
		}
	}
	return true
}

func parsStringToInts(toParse []string) (ris []int64) {
	for _, val := range toParse {
		n, _ := strconv.ParseInt(val, 10, 64)
		ris = append(ris, n)
	}
	return
}
