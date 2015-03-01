/*What is the 10001st prime number?*/
/*go run 7.go*/

package main

import "fmt"

const target = 10001

func isPrime(num int) bool {
	/* Finds if the passed number is prime or not*/
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var currentNum = 2
	var foundPrimes = 0
	for {
		if isPrime(currentNum) {
			foundPrimes++
		}
		if foundPrimes == target {
			break
		}
		currentNum++
	}
	fmt.Println(currentNum)
}
