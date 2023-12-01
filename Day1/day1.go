package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	leggiFile()

}

func leggiFile() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
