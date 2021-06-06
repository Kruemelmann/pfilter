package main

import (
	"fmt"
	"time"
)

func SeriellWrapper(primes []int, filterfunc filter) []int {
	s_arr := []int{}
	seriellstart := time.Now().UnixNano()
	for i := 0; i < len(primes); i++ {
		b, _ := filterfunc(primes[i])
		if b == true {
			s_arr = append(s_arr, primes[i])
		}
	}
	seriellend := time.Now().UnixNano()
	t := (seriellend - seriellstart)
	fmt.Printf("seriell: \t%dns -> %dms\n", t, (t / 1000000))
	return s_arr
}
