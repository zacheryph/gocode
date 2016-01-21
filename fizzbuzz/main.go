package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		str := ""

		if i%3 == 0 {
			str += "Fizz"
		}
		if i%5 == 0 {
			str += "Buzz"
		}

		if str == "" {
			fmt.Println(i)
		} else {
			fmt.Println(str)
		}
	}
}
