/*Find the largest palindrome made from the product of two 3-digit numbers.*/
/*go run 4.go*/

package main

import (
	"fmt"
	"strconv"
)

const MIN = 100
const MAX = 999

func checkPalindrome(s string) bool {
	length := len(s)
	half_length := length / 2
	rev_s := make([]rune, length)
	n := 0
	for _, r := range s {
		rev_s[n] = r
		n++
	}
	n = 0
	/*http://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go#answer-1754209*/
	for n < half_length {
		if rev_s[length-1-n] != rev_s[n] {
			return false
		}
		n++
	}
	return true
}

func main() {
	max := 1
	i := MIN
	for i <= MAX {
		j := MIN
		for j <= MAX {
			if i > j {
				j++
				continue
			}
			num := i * j
			if checkPalindrome(strconv.Itoa(num)) {
				if num > max {
					max = num
				}
			}
			j++
		}
		i++
	}
	fmt.Println(max)
}
