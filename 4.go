package main

import "fmt"
import "strconv"

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
	i := 100
	for i <= 999 {
		j := 100
		for j <= 999 {
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
