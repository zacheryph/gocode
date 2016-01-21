package main

import (
	"fmt"
	"strconv"
)

// FizzBuzz is fizz buzz magic
type FizzBuzz int

func (fb FizzBuzz) String() (res string) {
	if fb%3 == 0 {
		res += "Fizz"
	}
	if fb%5 == 0 {
		res += "Buzz"
	}
	if res == "" {
		res = strconv.Itoa(int(fb))
	}
	return
}

func main() {
	var i FizzBuzz

	for i = 1; i <= 100; i++ {
		fmt.Println(i)
	}
}
