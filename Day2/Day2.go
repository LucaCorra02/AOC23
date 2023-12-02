package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	leggiFile()

}

func leggiFile() {
	scanner := bufio.NewScanner(os.Stdin)
	var pt1 int
	var pt2 int
	var i = 0
	for scanner.Scan() {
		line := strings.Split(strings.TrimSpace(scanner.Text()), ";")
		regex := regexp.MustCompile(`\d* r|\d* g|\d* b`)
		dict := map[string]int{"r": 0, "b": 0, "g": 0}
		for _, str := range line {

			match := regex.FindAllString(str, -1)

			for _, play := range match { //singoli round
				color := len(play) - 1
				n, _ := strconv.Atoi(play[:color-1])
				dict[string(play[color])] = slices.Max([]int{dict[string(play[color])], n})
			}
		}
		if dict["b"] <= 14 && dict["r"] <= 12 && dict["g"] <= 13 {
			pt1 += i + 1
		}
		pt2 += dict["r"] * dict["g"] * dict["b"]
		i++
	}

	fmt.Println(pt1, pt2)
}
