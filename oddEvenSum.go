package main

import (
	"fmt"
)

func PrintOddEvenSum(numberList []int) {
	sumEven, sumOdd := magic(numberList)
	fmt.Printf("For given slice : \t%v\n\tSum of evens: \t%d\n\tSum of odds: \t%d\n", numberList, sumEven, sumOdd)
}

func magic(numberList []int) (sumEven int, sumOdd int) {
	for _, number := range numberList {
		if even(number) {
			sumEven += number
		} else {
			sumOdd += number
		}
	}
	return
}

func even(val int) bool {
	return val%2 == 0
}
