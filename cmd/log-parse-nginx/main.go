package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	arg := os.Args[1]

	f, err := os.Open(arg)
	check(err)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading arginput:", err)
	}
}

// from https://gobyexample.com/reading-files
func check(e error) {
	if e != nil {
		panic(e)
	}
}
