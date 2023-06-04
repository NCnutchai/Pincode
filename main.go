package main

import (
	"fmt"
	"strconv"
)

type SerialNumber struct {
	first  int
	second int
	third  int
}

func convertString(b byte) (*int, error) {
	number, err := strconv.Atoi(string(b))
	if err != nil {
		return nil, err
	}
	return &number, nil
}

func validate_pincode(s string) (*bool, error) {
	var duplicated []string
	var serial []string
	var validate bool
	if len(s) < 6 {
		return nil, fmt.Errorf("invalid length")
	}
	for i := 0; i < len(s); i++ {
		var numbers SerialNumber
		if i+2 >= len(s) {
			validate = true
			break
		}
		first, err := convertString(s[i])
		if err != nil {
			return nil, err
		}
		second, err := convertString(s[i+1])
		if err != nil {
			return nil, err
		}
		third, err := convertString(s[i+2])
		if err != nil {
			return nil, err
		}

		numbers.first = *first
		numbers.second = *second
		numbers.third = *third

		if numbers.first == numbers.second {
			if i == 0 {
				duplicated = append(duplicated, strconv.Itoa(numbers.first)+strconv.Itoa(numbers.second))
			}

		}
		if numbers.second == numbers.third {
			duplicated = append(duplicated, strconv.Itoa(numbers.second)+strconv.Itoa(numbers.third))
		}
		if (numbers.first+1 == numbers.second) && (numbers.second+1 == numbers.third) {
			serial = append(serial, strconv.Itoa(numbers.third)+strconv.Itoa(numbers.first))
		}
		if (numbers.first-1 == numbers.second) && (numbers.second-1 == numbers.third) {
			serial = append(serial, strconv.Itoa(numbers.third)+strconv.Itoa(numbers.first))
		}

	}
	if len(duplicated) > 2 {
		fmt.Println(duplicated)
		validate = false
	} else if len(duplicated) == 2 {
		if duplicated[0] == duplicated[1] {
			fmt.Println(duplicated)
			validate = false
		}
	}
	if len(serial) > 0 {
		fmt.Println(serial)
		validate = false
	}
	return &validate, nil
}

func main() {
	var input string
	fmt.Scan(&input)
	valid, err := validate_pincode(input)
	if err != nil {
		fmt.Println(err)
	} else {
		if *valid {
			fmt.Println("valid pincode")
		} else {
			fmt.Println("invalid pincode")
		}
	}
}
