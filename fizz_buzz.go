package main

import "fmt"

const (
	fizzbuzz, fizz, buzz string = "FizzBuzz", "Fizz", "Buzz"
)

func main() {
	var input int

	fmt.Print("Enter integer: ")
	fmt.Scanf("%d", &input)

	for i := 1; i <= input; i++ {
		fmt.Println(fizzBuzz(i))
	}
}

func fizzBuzz(i int) interface{} {
	switch {
	case i%15 == 0:
		return fizzbuzz
	case i%3 == 0:
		return fizz
	case i%5 == 0:
		return buzz
	default:
		return i
	}
}
