// Homework 1: Finger Exercises
// Due January 31, 2017 at 11:59pm
package main

import (
	"fmt"
	"unicode"
)

func main() {
	// Feel free to use the main function for testing your functions
	// fmt.Println("Hello, World!")
	// fmt.Println(ParsePhone("1 2 3 4 5 6 7 8 9 0"))
	// fmt.Println(Anagram("Undertale", "Deltarune"))
	// fmt.Println(FindEvens([]int{1, 2, 3, 4, 5}))
	// fmt.Println(SliceProduct([]int{1, 2, 3, 4}))
	// fmt.Println(Unique([]int{1, 1, 2, 2, 4, 5, 6, 7, 7}))
	// fmt.Println(InvertMap(map[string]int{"apples": 1, "bananas": 2}))
	fmt.Println(TopCharacters("Hello", 2))
}

// ParsePhone parses a string of numbers into the format (123) 456-7890.
// This function should handle any number of extraneous spaces and dashes.
// All inputs will have 10 numbers and maybe extra spaces and dashes.
// For example, ParsePhone("123-456-7890") => "(123) 456-7890"
//              ParsePhone("1 2 3 4 5 6 7 8 9 0") => "(123) 456-7890"
func ParsePhone(phone string) string {
	var parsedPhone []byte
	count := 0

	for _, c := range phone {
		if c == ' ' || c == '-' {
			continue
		} else if len(parsedPhone) == 0 {
			parsedPhone = append(parsedPhone, '(')
			parsedPhone = append(parsedPhone, byte(c))
		} else {
			parsedPhone = append(parsedPhone, byte(c))
		}

		count += 1

		if len(parsedPhone) == 4 {
			parsedPhone = append(parsedPhone, ')')
			parsedPhone = append(parsedPhone, ' ')
		}

		if count == 6 {
			parsedPhone = append(parsedPhone, '-')
		}
	}

	return string(parsedPhone)
}

// Anagram tests whether the two strings are anagrams of each other.
// This function is NOT case sensitive and should handle UTF-8
func Anagram(s1, s2 string) bool {
	map1 := make(map[rune]int)
	map2 := make(map[rune]int)

	if len(s1) != len(s2) {
		return false
	}

	for _, r := range s1 {
		map1[unicode.ToLower(r)] += 1
	}

	for _, r := range s2 {
		map2[unicode.ToLower(r)] += 1
	}

	for k, v := range map1 {
		val, ok := map2[k]
		if !ok {
			return false
		} else if (val != v) {
			return false
		}
	}

	return true
}

// FindEvens filters out all odd numbers from input slice.
// Result should retain the same ordering as the input.
func FindEvens(e []int) []int {
	result := make([]int, 0)
	for _, item := range e {
		if item % 2 == 0 {
			result = append(result, item)
		}
	}

	return result
}

// SliceProduct returns the product of all elements in the slice.
// For example, SliceProduct([]int{1, 2, 3}) => 6
func SliceProduct(e []int) int {
	if e == nil {
		return 0
	}
	if len(e) == 0 {
		return 1
	}
	
	return e[0] * SliceProduct(e[1:])
}

// Unique finds all distinct elements in the input array.
// Result should retain the same ordering as the input.
func Unique(e []int) []int {
	if len(e) <= 1 {
		return e
	}
	result := make([]int, 0)
	record := make(map[int]int)
	for _, item := range e {
		record[item] += 1
	}
	for _, item := range e {
		if (record[item] == 1) {
			result = append(result, item)
		}
	}

	return result
}

// InvertMap inverts a mapping of strings to ints into a mapping of ints to strings.
// Each value should become a key, and the original key will become the corresponding value.
// For this function, you can assume each value is unique.
func InvertMap(kv map[string]int) map[int]string {
	result := make(map[int] string)
	for k, v := range kv {
		result[v] = k
	}

	return result
}

// TopCharacters finds characters that appear more than k times in the string.
// The result is the set of characters along with their occurrences.
// This function MUST handle UTF-8 characters.
func TopCharacters(s string, k int) map[rune]int {
	result := make(map[rune]int)
	for _, item := range s {
		result[item] += 1
	}
	for key, v := range result {
		if v < k {
			delete(result, key)
		}
	}

	return result
}