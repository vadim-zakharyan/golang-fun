package main

import "fmt"

func main() {
	oddEvenSum := OddEvenSum{}
	display(oddEvenSum.Run([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))

	fibonacci := Fibonacci{}
	display(fibonacci.Run(35))

}

func display(value string) {
	fmt.Printf("%s\n\n", value)
}
