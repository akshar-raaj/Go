/*By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.*/

package main

import "fmt"

const MAX = 4000000

func main() {
	first, second := 1, 2
	third := first + second
	sum := 2
	for third < MAX {
		if third%2 == 0 {
			sum += third
		}
		first, second = second, third
		third = first + second
	}
	fmt.Println(sum)
}
