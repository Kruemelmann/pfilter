package main_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRadixSort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, (rand.Intn(100-1) + 1))
	}
	fmt.Println(arr)
	//fmt.Println(sort.RadixSort(arr))
}
