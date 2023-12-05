package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type scale struct {
	start  uint64
	base   uint64
	offset uint64
}

func main() {
	fmt.Printf("Pt1 min: %d\n", leggiFile())
}

func leggiFile() int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var seed []uint64
	for _, s := range strings.Split(scanner.Text(), " ")[1:] {
		n, _ := strconv.ParseUint(s, 10, 64)
		seed = append(seed, n)
	}
	scanner.Scan()
	scanner.Text()
	var scales []scale

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			seed = doConvertion(&seed, &scales)
			scales = []scale{}
		} else if !strings.Contains(line, "-") {
			nums := strings.Split(line, " ")
			base, _ := strconv.ParseUint(nums[0], 10, 64)
			start, _ := strconv.ParseUint(nums[1], 10, 64)
			offset, _ := strconv.ParseUint(nums[2], 10, 64)
			scales = append(scales, scale{start: start, base: base, offset: offset})
		}
	}
	seed = doConvertion(&seed, &scales)
	return int(findMin(seed))
}

func doConvertion(seed *[]uint64, scales *[]scale) (convertiti []uint64) {
	for _, seme := range *seed {
		ris := seme
		for _, scala := range *scales {
			if seme >= (scala.start) && seme <= (scala.start+scala.offset) {
				ris = (scala.base + (seme - scala.start))
				break
			}
		}
		convertiti = append(convertiti, ris)
	}
	return
}

func findMin(seed []uint64) uint64 {
	min := seed[0]
	for i := 1; i < len(seed); i++ {
		if seed[i] < min {
			min = seed[i]
		}
	}
	return min
}
