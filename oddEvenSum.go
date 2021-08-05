package main

import (
	"fmt"
)

type OddEvenSum struct{}

func (t *OddEvenSum) Run(numberList []int) string {
	sumEven, sumOdd := t.magic(numberList)
	return fmt.Sprintf("For given slice : \t%v\n\tSum of evens: \t%d\n\tSum of odds: \t%d", numberList, sumEven, sumOdd)
}

func (t *OddEvenSum) magic(numberList []int) (sumEven int, sumOdd int) {
	for _, number := range numberList {
		if t.even(number) {
			sumEven += number
		} else {
			sumOdd += number
		}
	}
	return
}

func (t OddEvenSum) even(val int) bool {
	return val%2 == 0
}
