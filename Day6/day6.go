package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	time, distance := leggiInput()
	fmt.Println(time, distance)
	calcolaCombinazioni(time, distance)
}

func leggiInput() (time []int, distance []int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	time = stringToInt(strings.Fields(scanner.Text())[1:])
	scanner.Scan()
	distance = stringToInt(strings.Fields(scanner.Text())[1:])
	return time, distance
}

func calcolaCombinazioni(time []int, distance []int) {
	tot := 1
	for i, t := range time {
		count := 0
		for j := 1; j < t; j++ {
			dist := j * (t - j)
			if dist > distance[i] {
				count++
			}
		}
		fmt.Println("count", count)
		tot *= count
	}
	fmt.Println("tot: ", tot)

}

func stringToInt(tmp []string) (ris []int) {
	for _, v := range tmp {
		n, _ := strconv.Atoi(v)
		ris = append(ris, n)
	}
	return
}
