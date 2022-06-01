package array

import "fmt"

// ArrayOf is a variadic function to create array
func ArrayOf(elem ...int) []int {
	return elem
}

// First returns first element of an array
func First(arr []int) int {
	return arr[0]
}

// Last returns last element of an array
func Last(arr []int) int {
	return arr[len(arr)-1]
}

// Reverse reverses an array
func Reverse(arr []int) {
	fmt.Println(arr)
	i := 0
	j := len(arr) - 1
	for i != len(arr)/2 {
		t := arr[i]
		arr[i] = arr[j]
		arr[j] = t
		i++
		j--
	}
	fmt.Println(arr)
}

// Rotate rotates an array by n elements
// Time O(n), Auxiliary space O(n)
func Rotate(arr []int, n int) []int {
	fmt.Println(arr)
	t := append(arr[n:], arr[0:n]...)
	fmt.Println(t)
	return t
}
