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
	for scanner.Scan() {
		line := scanner.Text()
		match := regexp.MustCompile(`\d?=red`).FindAllString(line, -1)
		fmt.Println(match)
	}
}
