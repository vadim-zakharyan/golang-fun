package main

import (
	"strconv"
	"strings"
)

type FizBuzz struct{}

const (
	fiz     = "Fizz"
	buzz    = "Buzz"
	fizBuzz = "FizzBuzz"
)

func (t *FizBuzz) Run(numberList []int) (out string) {

	result := t.magic(numberList)

	return strings.Join(result[:], ", ")
}

func (t *FizBuzz) magic(numberList []int) (result []string) {

	for _, number := range numberList {
		var out = strconv.Itoa(number)
		if number%3 == 0 && number%5 == 0 {
			result = append(result, fizBuzz)
		} else if number%3 == 0 {
			result = append(result, fiz)
		} else if number%5 == 0 {
			result = append(result, buzz)
		} else {
			result = append(result, out)
		}
	}
	return
}
