package main

import (
	"fmt"
)

func rotateKPositionsLeft(s string, k int) string {
	stringArray := []rune(s)
	n := len(stringArray)
	k = k % n
	rotated := make([]rune, n)
	for i := 0; i < n; i++ {
		rotated[i] = stringArray[(i+k)%n]
	}
	return string(rotated)
}

func main() {
	s1 := "abcdef"
	k1 := 2
	result1 := rotateKPositionsLeft(s1, k1)
	fmt.Println("Test case 1 - Rotated string:", result1)

	s2 := "abcdef"
	k2 := 9
	result2 := rotateKPositionsLeft(s2, k2)
	fmt.Println("Test case 2 - Rotated string:", result2)

	s3 := "abcdef"
	k3 := 6
	result3 := rotateKPositionsLeft(s3, k3)
	fmt.Println("Test case 3 - Rotated string:", result3)
}
