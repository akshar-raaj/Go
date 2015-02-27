/*Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.*/

package main

import "fmt"

const N = 100

func square(x int) int {
	return x * x
}

func sumOfSquares(n int) int {
	product := n * (n + 1) * ((2 * n) + 1)
	return product / 6
}

func sumN(n int) int {
	product := n * (n + 1)
	return product / 2
}

func main() {
	var sumOfN = sumN(N)
	var squareOfSumN = square(sumOfN)
	var sumOfSquaresN = sumOfSquares(N)
	fmt.Println(squareOfSumN - sumOfSquaresN)
}
