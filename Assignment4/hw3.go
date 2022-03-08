// Homework 3: Interfaces
// Due February 14, 2017 at 11:59pm
package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// Feel free to use the main function for testing your functions
	// hello := map[string]string{
	// 	"hello":   "world",
	// 	"hola":    "mundo",
	// 	"bonjour": "monde",
	// }
	// for k, v := range hello {
	// 	fmt.Printf("%s, %s\n", strings.Title(k), v)
	// }

	fmt.Println(Fold([]int{1, 2, 3, 4}, 1, Multiply))

	person1 := NewPerson("Rakhat", "Zhumekin")
	person2 := NewPerson("First", "Last")

	var personSlice = make(PersonSlice, 0)

	personSlice = append(personSlice, person1)
	personSlice = append(personSlice, person2)
	personSlice = append(personSlice, person1)

	fmt.Println("The length of the slice is", personSlice.Len())
	fmt.Println("Is this a palindrome?", IsPalindrome(personSlice))

}

// Problem 1: Sorting Names
// Sorting in Go is done through interfaces!
// To sort a collection (such as a slice), the type must satisfy sort.Interface,
// which requires 3 methods: Len() int, Less(i, j int) bool, and Swap(i, j int).
// To actually sort a slice, you need to first implement all 3 methods on a
// custom type, and then call sort.Sort on your the PersonSlice type.
// See the Go documentation: https://golang.org/pkg/sort/ for full details.

// Person stores a simple profile. These should be sorted by alphabetical order
// by last name, followed by the first name, followed by the ID. You can assume
// the ID will be unique, but the names need not be unique.
// Sorting should be case-sensitive and UTF-8 aware.
type Person struct {
	ID        int
	FirstName string
	LastName  string
}

type PersonSlice []*Person

var counter int = 1

// NewPerson is a constructor for Person. ID should be assigned automatically in
// sequential order, starting at 1 for the first Person created.
func NewPerson(first, last string) *Person {
	person := &Person {
		ID: counter,
		FirstName: first,
		LastName: last,
	}
	
	counter++

	return person
}

func (personSlice PersonSlice) Len() int {
	return len(personSlice)
}

func (personSlice PersonSlice) Less(i, j int) bool {
	last := strings.Compare(personSlice[i].LastName, personSlice[j].LastName)
	first := strings.Compare(personSlice[i].FirstName, personSlice[j].FirstName)

	if last != 0 {
		return last < 0
	}

	if first != 0 {
		return first < 0
	}

	return personSlice[i].ID < personSlice[j].ID
}

func (personSlice PersonSlice) Swap(i, j int) {
	personSlice[i], personSlice[j] = personSlice[j], personSlice[i]
}

// Problem 2: IsPalindrome Redux
// Using a function that simply requires sort.Interface, you should be able to
// check if a sequence is a palindrome. You may use, adapt, or modify your code
// from HW0. Note that the input does not need to be a string: any type which
// satisfies sort.Interface can (and will) be used to test. This means that the
// only functionality you should use should come from the sort.Interface methods
// Ex: [1, 2, 1] => true

// IsPalindrome checks if the string is a palindrome.
// A palindrome is a string that reads the same backward as forward.
func IsPalindrome(s sort.Interface) bool {
	for i := 0; i < s.Len(); i++ {
		j := s.Len() - 1 - i;
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}

	return true
}

// Problem 3: Functional Programming
// Write a function Fold which applies a function repeatedly on a slice,
// producing a single value via repeated application of an input function.
// The behavior of Fold should be as follows:
//   - When s is empty, return v (default value)
//   - When s has 1 value (x0), apply f once: f(v, x0)
//   - When s has 2 values (x0, x1), apply f twice, from left to right: f(f(v, x0), x1)
//   - Continue this pattern recursively to obtain the final result.

// Fold applies a left to right application of f on s starting with v.
// Note the argument signature of f - func(int, int) int.
// This means f is a function which has 2 int arguments and returns an int.
func Fold(s []int, v int, f func(int, int) int) int {
	if len(s) == 0 {
		return v
	}

	return Fold(s[1:], f(v, s[0]), f)
}

func Add(a, b int) int {
	return a + b
}

func Multiply(a, b int) int {
	return a * b
}