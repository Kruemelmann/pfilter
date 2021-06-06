package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

//size the size of the slice that is filled with random numbers
const size = 1000000

//chunksize of the parts that are processed in parallel
const chunksize = 1000

func init() {
	//seeding the random numbers
	rand.Seed(time.Now().UnixNano())

	//set the count of used cores
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)
}

func main() {
	primes := make([]int, size)
	for i := 0; i < size; i++ {
		primes[i] = rand.Intn(100-1) + 1
	}
	fmt.Println("Size of the slice is ", size)

	//parallel
	res := ParallelWrapper(primes, filterfunc)
	fmt.Println("result len:", len(res))

	//seriell
	res = SeriellWrapper(primes, filterfunc)
	fmt.Println("result len:", len(res))
}
