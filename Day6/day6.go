package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	time, distance, timePt2, distancePt2 := leggiInput()
	pt1 := calcolaCombinazioni(time, distance)
	pt2 := calcolaCombinazioni([]int{timePt2}, []int{distancePt2})
	fmt.Printf("Pt1: %d\nPt2: %d\n", pt1, pt2)
}

func leggiInput() (time []int, distance []int, timePt2 int, distancePt2 int) {
	scanner := bufio.NewScanner(os.Stdin)
	var tp, dp string
	scanner.Scan()
	time, tp = stringToInt(strings.Fields(scanner.Text())[1:])
	scanner.Scan()
	distance, dp = stringToInt(strings.Fields(scanner.Text())[1:])
	timePt2, _ = strconv.Atoi(tp)
	distancePt2, _ = strconv.Atoi(dp)
	return time, distance, timePt2, distancePt2
}

func stringToInt(tmp []string) (ris []int, numPt2 string) {
	for _, v := range tmp {
		numPt2 += v
		n, _ := strconv.Atoi(v)
		ris = append(ris, n)
	}
	return
}

func calcolaCombinazioni(time []int, distance []int) int {
	tot := 1
	for i, t := range time {
		count := 0
		for j := 1; j < t; j++ {
			dist := j * (t - j)
			if dist > distance[i] {
				count++
			}
		}
		tot *= count
	}
	return tot
}
