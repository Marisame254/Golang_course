// create a func with the identifier foo that
// 			○ takes in a variadic parameter of type int
// 			○ pass in a value of type []int into your func (unfurl the []int)
// 			○ returns the sum of all values of type int passed in
// create a func with the identifier bar that
// 			○ takes in a parameter of type []int
// 			○ returns the sum of all values of type int passed in

package main

import "fmt"

func main() {
	fmt.Println("Hi")

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("the numbers are:", a)

	f := foo(a...)
	fmt.Println("the foo sum is:", f)

	f1 := bar(a)
	fmt.Println("the bar sum is:", f1)
}

func foo(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

func bar(x []int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
