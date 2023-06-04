package main

import (
	"fmt"
	"strconv"
)

func checkDoubleDigit(s string) bool {
	var double_digits []string
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			double_digits = append(double_digits, string(s[i])+string(s[i+1]))
		}
	}

	if len(double_digits) > 2 {
		return false
	}
	return true
}

func checkOrderNumber(s string) bool {
	var numbers []string
	for i := 0; i < len(s)-2; i++ {
		n1, _ := strconv.Atoi(string(s[i]))
		n2, _ := strconv.Atoi(string(s[i+1]))
		n3, _ := strconv.Atoi(string(s[i+2]))

		if n1+1 == n2 && n2+1 == n3 {
			numbers = append(numbers, string(s[i])+string(s[i+1])+string(s[i+2]))
			return false
		}

		if n1 == n2+1 && n2 == n3+1 {
			return false
		}
	}
	return true
}

func checkTripleDigit(s string) bool {
	var triple_digits []string
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+1] && s[i+1] == s[i+2] {
			triple_digits = append(triple_digits, string(s[i])+string(s[i+1]))
		}
	}

	if len(triple_digits) > 2 {
		return false
	}
	return true
}

func validate_pincode(s string) bool {
	checkLength := len(s) >= 6
	checkdouble := checkDoubleDigit(s)
	checkorder := checkOrderNumber(s)
	checktriple := checkTripleDigit(s)
	return checkLength && checkdouble && checkorder && checktriple
}

func main() {
	var input string
	fmt.Scan(&input)
	valid := validate_pincode(input)
	if valid {
		fmt.Println("valid pincode")
	} else {
		fmt.Println("invalid pincode")
	}
}
