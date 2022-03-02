// Homework 0: Hello Go!
// Due January 24, 2017 at 11:59pm
package main

import (
	"fmt"
	"math"
)

func main() {
	// Feel free to use the main function for testing your functions
	fmt.Println("Hello World")
	// fmt.Println(Fizzbuzz(31));
	// fmt.Println(IsPrime(18));
	// fmt.Println(IsPalindrome("Li"));
}

// Fizzbuzz is a classic introductory programming problem.
// If n is divisible by 3, return "Fizz"
// If n is divisible by 5, return "Buzz"
// If n is divisible by 3 and 5, return "FizzBuzz"
// Otherwise, return the empty string
func Fizzbuzz(n int) string {
	if n % 3 == 0 && n % 5 == 0 {
		return "FizzBuzz"
	} else if n % 3 == 0 {
		return "Fizz"
	} else if n % 5 == 0 {
		return "Buzz"
	} else {
		return ""
	}
}

// IsPrime checks if the number is prime.
// You may use any prime algorithm, but you may NOT use the standard library.
func IsPrime(n int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(n)))); i++ {
		if n % i == 0 {
			return false;
		}
	}

	return n > 1;
}

// IsPalindrome checks if the string is a palindrome.
// A palindrome is a string that reads the same backward as forward.
func IsPalindrome(s string) bool {
	for i := 0; i < len(s); i++ {
		j := len(s) - 1 - i;
		if s[i] != s[j] {
			return false
		}
	}

	return true
}