package main

import (
	"fmt"
	"os"
	//	"strings"
)

func main() {

	arg := os.Args[1]
	fmt.Printf(arg)

	f, err := os.Open(arg)
	check(err)

	b1 := make([]byte, 100)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	//	lines := strings.Split(string(b1), "\n")
	fmt.Printf("--%s--\n", b1)
}

// from https://gobyexample.com/reading-files
func check(e error) {
	if e != nil {
		panic(e)
	}
}
