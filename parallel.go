package main

import (
	"fmt"
	"sync"
	"time"
)

type Workeroutput struct {
	Results []int
	Id      int
}

func worker(wg *sync.WaitGroup, primes []int, res_chan chan Workeroutput, workerid int, filterfunc filter) {
	defer wg.Done()
	//Phase 1: create the index
	index := [chunksize]int{}
	for i := 0; i < len(primes); i++ {
		b, _ := filterfunc(primes[i])
		if b == true {
			index[i] = 1
		}
	}

	//Phase 2: prefix sum -> here the offsets are calculated
	offsets := [chunksize]int{}
	j := 0
	for i := 0; i < len(primes); i++ {
		offsets[i] = j
		if index[i] == 1 {
			j++
		}
	}

	//Phase 3: Moving the valid elements to the correct spot in the result slice
	arr := [chunksize]int{}
	for i := 0; i < len(arr); i++ {
		if index[i] == 1 {
			arr[offsets[i]] = primes[i]
		}
	}

	//shrink the array to a slice with the matching size
	res_arr := arr[:j]
	res := Workeroutput{
		Id:      workerid,
		Results: res_arr,
	}
	res_chan <- res
}

func ParallelWrapper(primes []int, filterfunc filter) []int {
	res_chan := make(chan Workeroutput, (size / chunksize))
	var wg sync.WaitGroup

	parallelstart := time.Now().UnixNano()

	//Phase 1: Start workers
	for i := 0; i < len(primes); i += chunksize {
		wg.Add(1)
		batch := primes[i:min(i+chunksize, len(primes))]
		go worker(&wg, batch, res_chan, i, filterfunc)
	}

	wg.Wait()
	//close the response channel so i can iterate in the next step
	close(res_chan)
	//unsored list of array parts of all workers
	wres := []Workeroutput{}

	for elem := range res_chan {
		wres = append(wres, elem)
	}

	//Phase 2: Sort of the Responsechannel
	wres = RadixSort(wres)

	res_arr := []int{}
	for _, v := range wres {
		res_arr = append(res_arr, v.Results...)
	}

	parallelend := time.Now().UnixNano()
	t := (parallelend - parallelstart)
	fmt.Printf("parallel: \t%dns -> %dms\n", t, (t / 1000000))
	return res_arr
}
