/*What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?*/
/*go run 5.go*/

package euler

import "fmt"

func LCM(num1, num2 int) int {
	//This looks like a good implementation of LCM
	var product int
	//This function works with num1 being greater
	if num1 >= num2 {
		//pass
	}else {
		num1, num2 = num2, num1
	}
	i := 1
	for {
		product = num1 * i
		if product%num2 == 0 {
			break
		}
		i++
	}
	return product
}

func Five() {
    const MIN = 1
    const MAX = 20

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
