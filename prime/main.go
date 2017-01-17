package main

import (
	"fmt"
	"os"
	"strconv"
)

func IsPrime(x int64) bool {
	var max, current, step, nextstep int64

	// get rid of the small numbers right away
	if x <= 3 {
		return true
	}
	if x%2 == 0 || x%3 == 0 {
		return false
	}

	max = x / 3
	current = 5
	step, nextstep = 2, 4

	for current < max {
		if x%current == 0 {
			return false
		}

		max = x / current
		current += step
		step, nextstep = nextstep, step
	}

	return true
}

func main() {
	for _, x := range os.Args {
		n, err := strconv.ParseInt(x, 10, 64)
		if err != nil {
			continue
		}
		fmt.Println("Is Prime? ", x, IsPrime(n))
	}
}
