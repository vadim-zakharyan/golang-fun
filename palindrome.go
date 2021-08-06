package main

import "fmt"

type Palindrome struct{}

func (t *Palindrome) Run(leftString string) (out string) {
	symmetric := t.magic(leftString)
	return fmt.Sprintf("Palindrome for '%s':\t'%s'", leftString, symmetric)
}

func (t *Palindrome) magic(leftString string) (symmetric string) {
	symmetric = leftString
	for i := len(leftString) - 2; i >= 0; i-- {
		symmetric += fmt.Sprint(string([]rune(leftString)[i]))
	}
	return
}
