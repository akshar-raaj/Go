/*By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms. Answer: 4613732*/

package main

import "fmt"

const MAX = 4000000
var first = 1
var second = 2

func is_even(num int) bool {
	return num % 2 == 0
}

func main() {
	sum := 0
	if is_even(first) {
		sum += first
	}
	if is_even(second) {
		sum += second
	}
	for (first + second) < MAX {
		third := first+second
		if is_even(third) {
			sum += third
		}
		first, second = second, third
	}
	fmt.Println(sum)
}
