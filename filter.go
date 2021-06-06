package main

import "math"

type filter func(elem int) (bool, error)

//filterfunc is the function used to check if an element is wanted or not
func filterfunc(elem int) (bool, error) {
	return isPrime(elem), nil
}

//example 1
func isPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

//example 2
func isGreaderOne(value int) bool {
	if value > 1 {
		return true
	}
	return false
}
