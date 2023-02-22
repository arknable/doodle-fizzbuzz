package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/arknable/doodle-fizzbuzz/numbers"
)

func main() {
	p, err := numbers.NewPrinter(1, 100)
	if err != nil {
		fmt.Println("ERROR: ", err)
		os.Exit(1)
	}

	p.WithStringer(func(n int) string {
		if n > 0 {
			if (n%5 == 0) && (n%3 == 0) {
				return "FizzBuzz"
			}

			if n%3 == 0 {
				return "Fizz"
			}

			if n%5 == 0 {
				return "Buzz"
			}
		}
		return strconv.Itoa(n)
	})

	p.Print()
}
