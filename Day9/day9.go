package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("part 1: %d\n", leggiFile())
}

func leggiFile() (tot int64) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ranges := parsStringToInts(strings.Split(scanner.Text(), " "))
		tot += calcRange(ranges)
	}
	return
}

func calcRange(ranges []int64) (sum int64) {
	var lastValue []int64
	for !checkAllZero(ranges) {
		var newRanges []int64
		for i := 0; i < len(ranges)-1; i++ {
			diff := ranges[i+1] - ranges[i]
			newRanges = append(newRanges, diff)
		}
		lastValue = append(lastValue, ranges[len(ranges)-1])
		ranges = newRanges
	}
	for _, val := range lastValue {
		sum += val
	}
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
