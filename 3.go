package main

import "fmt"

var num = 600851475143
var next_prime = 2

func is_prime(num int) bool {
    /* Finds if the passed number is prime or not*/
    for i:=2; i<=num/2; i++ {
        if (num % i == 0) {
            return false
        }
    } 
    return true;
}

func set_next_prime() {
    /* While factorizing if we find number is not divisible by 2
    no need to check with 2 during further factorization.
    We need to check with next prime numbers thereafter*/
    var i = next_prime+1
    for is_prime(i) != true {
        i += 1
    }
    next_prime = i
}

func is_divisible(dividend, divisor int) bool {
    if (dividend % divisor == 0) {
        return true
    }
    return false
}

func main() {
    var largest = 0
    for num != 1 {
        if(is_divisible(num, next_prime)) {
            num = num/next_prime
            largest = next_prime
        } else {
            set_next_prime()
        }
    }
    fmt.Println(largest)
}
