package main

import "log"

func main() {
	log.Println("isPrime(5) => ", isPrime(5)) // [1*5 or 5*1] => true
	log.Println("isPrime(4) => ", isPrime(4)) // [1*4 or 2*2 or 4*1] => false
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
