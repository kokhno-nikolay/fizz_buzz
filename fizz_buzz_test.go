package main

import (
	"fmt"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	for i, tt := range []struct {
		number   int
		response string
	}{
		{
			18,
			fizz,
		},
		{
			5,
			buzz,
		},
		{
			1,
			"1",
		},
		{
			15,
			fizzbuzz,
		},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			res := fizzBuzz(tt.number)
			if res != tt.response {
				t.Errorf("want %v; got %v", tt.response, res)
			}
		})
	}
}
