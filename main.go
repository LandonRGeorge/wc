package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func count(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var count int
	for scanner.Scan() {
		count++
	}
	return count
}

func main() {
	fmt.Println(count(os.Stdin))
}
