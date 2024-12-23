package main

import "fmt"

func main() {
	array := []int{18, 43, 69, 68, 54, 70}
	fmt.Println(binarySearch(array, 69))
}

func binarySearch(array []int, x int) int {
	l := 0
	r := len(array) - 1
	for l < r {
		mid := (l + r) / 2
		if l < r && array[mid] == x {
			return mid
		} else if l < r && x < array[mid] {
			r = r - 1
		} else {
			l = l + 1
		}
	}
	return -1
}
