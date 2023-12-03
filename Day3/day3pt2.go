package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	leggiFile()
}

func leggiFile() {
	scanner := bufio.NewScanner(os.Stdin)
	numeri := make(map[string][]int)
	//caratteri := make(map[int]string)

	re := regexp.MustCompile(`\d\d\d*`)
	for scanner.Scan() {
		line := scanner.Text()
		match := re.FindAllStringIndex(line, -1)
		for _, slc := range match {
			if len(slc) != 0 {
				numeri[line[slc[0]:slc[1]]] = slc
			}
		}
	}
	fmt.Println(numeri)
}
