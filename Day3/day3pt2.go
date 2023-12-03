package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	numeri, caratteri, dimRiga := leggiFile()
	calcolaPunteggio(numeri, caratteri, dimRiga)
}

func calcolaPunteggio(numeri map[string][]int, caratteri map[int][]string, dimRiga int) {
	for num, slc := range numeri {
		start := slc[0]
		finish := slc[len(slc)-1]
		for index := start; index <= finish; index++ {
			//fmt.Println(num, index)
			if foundKey(caratteri, index+1) {
				caratteri[index+1] = append(caratteri[index+1], num)
				break
			}
			if foundKey(caratteri, index-1) {
				caratteri[index-1] = append(caratteri[index-1], num)
				break
			}
			if foundKey(caratteri, dimRiga+index+1) {
				caratteri[dimRiga+index+1] = append(caratteri[dimRiga+index+1], num)
				break
			}
			if foundKey(caratteri, dimRiga+index-1) {
				caratteri[dimRiga+index-1] = append(caratteri[dimRiga+index-1], num)
				break
			}
			if foundKey(caratteri, index-dimRiga+1) {
				caratteri[index-dimRiga+1] = append(caratteri[index-dimRiga+1], num)
				break
			}
			if foundKey(caratteri, index-dimRiga-1) {
				caratteri[index-dimRiga-1] = append(caratteri[index-dimRiga-1], num)
				break
			}
		}
	}
	fmt.Println(caratteri)

}

func foundKey(caratteri map[int][]string, pos int) bool {
	if _, ok := caratteri[pos]; ok {
		return true
	}
	return false
}

func leggiFile() (map[string][]int, map[int][]string, int) {
	scanner := bufio.NewScanner(os.Stdin)
	numriga := 0
	dimRiga := 0
	numeri := make(map[string][]int)
	caratteri := make(map[int][]string)
	re := regexp.MustCompile(`\d\d\d*`)
	re2 := regexp.MustCompile(`\*`)

	for scanner.Scan() {
		line := scanner.Text()
		dimRiga = len(line)
		match := re.FindAllStringIndex(line, -1)
		if len(match) != 0 {
			for _, slc := range match {
				numeri[line[slc[0]:slc[1]]] = []int{slc[0] + (numriga * 10) + 1, slc[1] + (numriga * 10)}
			}
		}
		match2 := re2.FindAllStringIndex(line, -1)
		for _, s := range match2 {
			caratteri[(s[0]+1)+(numriga*10)] = []string{}
		}
		numriga++
	}
	return numeri, caratteri, dimRiga
}
