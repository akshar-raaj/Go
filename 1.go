/*Find the sum of all the multiples of 3 or 5 below 1000.*/

package main 

import "fmt"

func is_divisible_by_num(dividend, divisor int) bool {
    if (dividend % divisor == 0) {
        return true
    }
    return false
}

func main() {
    var sum = 0
    for i := 0; i < 1000; i++ {
        if (is_divisible_by_num(i, 3) || is_divisible_by_num(i, 5)) {
            sum += i;
        }
    }
    fmt.Println(sum)
}