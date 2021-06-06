package main

const digit = 4
const maxbit = -1 << 31

// Radix Sort
func RadixSort(array []Workeroutput) []Workeroutput {
	maxvalue := maxValue(array)
	size := len(array)
	significantDigit := 1
	semiSorted := make([]Workeroutput, size, size)

	for maxvalue/significantDigit > 0 {
		bucket := [10]int{0}
		for i := 0; i < size; i++ {
			bucket[(array[i].Id/significantDigit)%10]++
		}
		for i := 1; i < 10; i++ {
			bucket[i] += bucket[i-1]
		}
		for i := size - 1; i >= 0; i-- {
			bucket[(array[i].Id/significantDigit)%10]--
			semiSorted[bucket[(array[i].Id/significantDigit)%10]].Id = array[i].Id
			semiSorted[bucket[(array[i].Id/significantDigit)%10]].Results = array[i].Results
		}
		for i := 0; i < size; i++ {
			array[i] = semiSorted[i]
		}
		significantDigit *= 10
	}
	return array
}

// Finds the largest number in an array
func maxValue(array []Workeroutput) int {
	largestNum := 0

	for i := 0; i < len(array); i++ {
		if array[i].Id > largestNum {
			largestNum = array[i].Id
		}
	}
	return largestNum
}
