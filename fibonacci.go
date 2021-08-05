package main

import (
	"fmt"
)

type Fibonacci struct{}

func (t *Fibonacci) Run(number int) string {

	result := t.magic(number)
	return fmt.Sprintf("Fibonacci for %d:\t%d", number, result)
}

func (t *Fibonacci) magic(number int) int {
	a := 1
	b := 1

	for i := 3; i <= number; i++ {
		var c = a + b
		a = b
		b = c
	}
	return b
}
