package main

import "fmt"

type HasDuplicate struct{}

func (t *HasDuplicate) Run(someArray []int) string {
	var has string
	if !t.hasDuplicate(someArray) {
		has = "not "
	}
	return fmt.Sprintf("array %v has %sduplicates", someArray, has)
}

func (t *HasDuplicate) hasDuplicate(someArray []int) bool {
	var tmpList []int
	for _, value := range someArray {
		if !t.contains(tmpList, value) {
			tmpList = append(tmpList, value)
		} else {
			return true
		}
	}
	return false
}

func (t *HasDuplicate) contains(someArray []int, searchValue int) bool {
	for _, value := range someArray {
		if value == searchValue {
			return true
		}
	}
	return false
}
