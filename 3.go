package main

import "fmt"

var (
	num = 600851475143
	nextPrime = 2
)

func isPrime(num int) bool {
	/* Finds if the passed number is prime or not*/
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func setNextPrime() {
	/* While factorizing if we find number is not divisible by 2
	   no need to check with 2 during further factorization.
	   We need to check with next prime numbers thereafter*/
	i := nextPrime + 1
	for isPrime(i) != true {
		i += 1
	}
	nextPrime = i
}

func isDivisible(dividend, divisor int) bool {
	if dividend%divisor == 0 {
		return true
	}
	return false
}

func main() {
	largest := 0
	for num != 1 {
		if isDivisible(num, nextPrime) {
			num = num / nextPrime
			largest = nextPrime
		} else {
			setNextPrime()
		}
	}
	fmt.Println(largest)
}
