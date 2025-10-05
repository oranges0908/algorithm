package main

import "fmt"

func TestMonotoneIncreasingDigits() {
	fmt.Println(monotoneIncreasingDigits(10) == 9)
	fmt.Println(monotoneIncreasingDigits(1234) == 1234)
	fmt.Println(monotoneIncreasingDigits(332) == 299)
}

func monotoneIncreasingDigits1(n int) int {
	if n < 10 {
		return n
	}

	//Break an integer into digits
	digits := make([]int, 0)
	for n > 0 {
		digits = append(digits, n%10)
		n = n / 10
	}
	//compare i and i-1
	for i := len(digits) - 1; i > 0; i-- {

		if digits[i] > digits[i-1] {
			// Subtract one from the high digit
			digits[i] = digits[i] - 1
			// set the remains to 9
			for j := i - 1; j >= 0; j-- {
				digits[j] = 9
			}
			// recheck the whole number
			i = len(digits)
		}
	}
	rc := 0
	for i := len(digits) - 1; i >= 0; i-- {
		rc = rc*10 + digits[i]
	}
	return rc
}

func monotoneIncreasingDigits(n int) int {
	if n < 10 {
		return n
	}

	//Break an integer into digits
	digits := make([]int, 0)
	for n > 0 {
		digits = append(digits, n%10)
		n = n / 10
	}
	//compare i and i-1
	for i := 0; i < len(digits)-1; i++ {
		if digits[i] < digits[i+1] {
			// Subtract one from the high digit
			digits[i+1] = digits[i+1] - 1
			// set the remains to 9
			for j := 0; j <= i; j++ {
				digits[j] = 9
			}
		}
	}
	rc := 0
	for i := len(digits) - 1; i >= 0; i-- {
		rc = rc*10 + digits[i]
	}
	return rc
}
