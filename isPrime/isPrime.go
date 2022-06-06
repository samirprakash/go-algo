package main

import (
	"log"
	"math"
)

func main() {
	log.Println("isPrime(5) => ", isPrime(5))   // [1*5 or 5*1] => true
	log.Println("isPrime(4) => ", isPrime(4))   // [1*4 or 2*2 or 4*1] => false
	log.Println("isPrime(13) => ", isPrime(13)) // [11*12 or 12*11] => true
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	nn := math.Sqrt(float64(n))
	for i := 2; i <= int(nn); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
