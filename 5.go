package main

import "fmt"

const MIN = 1
const MAX = 20

func LCM(num1, num2 int) int {
	//This looks like a good implementation of LCM
	var greater int
	var smaller int
	var product int
	if num1 > num2 {
		greater = num1
		smaller = num2
	} else {
		greater = num2
		smaller = num1
	}
	i := 1
	for {
		product = greater * i
		if product%smaller == 0 {
			break
		}
		i++
	}
	return product
}

func main() {
	product := MAX
	for i := MIN; i <= MAX; {
		if product%i == 0 {
			i++
			continue
		} else {
			product = LCM(i, product)
		}
		i++
	}
	fmt.Println(product)
}
