package main

import "fmt"

func main() {
	oddEvenSum := OddEvenSum{}
	display(oddEvenSum.Run([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))

	fibonacci := Fibonacci{}
	display(fibonacci.Run(35))

	fizBuzz := FizBuzz{}
	display(fizBuzz.Run([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}))

	palindrome := Palindrome{}
	display(palindrome.Run("rota"))

	hasDuplicate := HasDuplicate{}
	display(hasDuplicate.Run([]int{1, 2, 2}))
	display(hasDuplicate.Run([]int{1, 2, 3}))

	DisplayList()

	DisplayTasks()
}
// dumb comment
func display(value string) {
	fmt.Printf("%s\n\n", value)
}
