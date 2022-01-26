package main

import (
	"fmt"
	"strconv"
)

const (
	fizzbuzz, fizz, buzz string = "FizzBuzz", "Fizz", "Buzz"
)

func main() {
	var input int

	_, err := fmt.Scanf("%d", &input)
	if err != nil {
		panic("invalid input value")
	}

	for i := 1; i <= input; i++ {
		fmt.Println(fizzBuzz(i))
	}
}

func fizzBuzz(i int) string {
	switch {
	case i%15 == 0:
		return fizzbuzz
	case i%3 == 0:
		return fizz
	case i%5 == 0:
		return buzz
	default:
		return strconv.Itoa(i)
	}
}
