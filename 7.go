/*What is the 10001st prime number?*/
/*go run 7.go*/

package euler

import "fmt"

const target = 10001

func Seven() {
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
