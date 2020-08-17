package main

import (
	"fmt"
	"strconv"
	"sync"
)

// this function will swap the ith index of slice with its i+1 index element

func swap(sli []int, i int) {
	var tempI int = sli[i]
	var tempIP1 int = sli[i+1]

	sli[i] = tempIP1
	sli[i+1] = tempI

}

// bubble sort pseudo code:https://www.chegg.com/homework-help/questions-and-answers/q1-bubblesort-another-simple-sorting-algorithm-whose-pseudocode-follows-1-write-c-bubbleso-q8984069

// BubbleSort func
// slice by default does passed by reference
func bubbleSort(sli []int) {
	var n int = len(sli)

	for i := n - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if sli[j] > sli[j+1] {
				swap(sli, j)
			}

		}
	}

}

// Concurrency sort
func concurrencySort(sli []int, ch chan []int, wg *sync.WaitGroup) {
	// print the unsorted array
	fmt.Println(sli)


	bubbleSort(sli)

	// push sli down the channel to store
	ch <- sli
	// say that we are done
	wg.Done()


}


/*
 this function takes in slice 1 and slice 2 and will merge them
 such that the resulting slice12 contains them in sorted order

 note that slice12 is an empty slice of len slice1 + slice2
*/
func MergeTwoSorted(slice1 []int, slice2 []int, slice12 []int) {
	fmt.Println()
	// fmt.Println("MergeTwoSorted():", len(slice1), len(slice2), len(slice12))
	// fmt.Println("MergeTwoSorted():", slice1, slice2, slice12)
	i, j, k := 0, 0, 0
	for {
		// fmt.Println("slice1[", i, "]=", slice1[i])
		// fmt.Println("slice2[", j, "] =", slice2[j])
		if slice1[i] < slice2[j] {
			slice12[k] = slice1[i]
			k++
			i++
		} else if slice1[i] == slice2[j] {
			slice12[k] = slice1[i]
			k++
			i++

			slice12[k] = slice2[j]
			k++
			j++
		} else if slice1[i] > slice2[j] {
			slice12[k] = slice2[j]
			k++
			j++
		}

		if i == len(slice1) {
			for j < len(slice2) {
				slice12[k] = slice2[j]
				k++
				j++
			}
		}

		if j == len(slice2) {
			for i < len(slice1) {
				slice12[k] = slice1[i]
				k++
				i++
			}
		}

		if k == len(slice12) {
			break
		}
	}

	// fmt.Println("MergeTwoSorted():", slice12)
}

func main() {
	var ex = make([]int, 0)
	var UserInput string
	var exit bool
	// user input the integers
	for exit == false {

		fmt.Println("Please input an integer. Input X to stop")
		_, _ = fmt.Scan(&UserInput)

		if UserInput == "X" {
			exit = true
		} else {
			convert, _ := strconv.Atoi(UserInput)
			ex = append(ex, convert)
		}


	}

	// create channel and wait group
	ch := make(chan []int,4)
	wg := sync.WaitGroup{}
	// other var for partition
	arrayLength := len(ex)
	chunkSize := int(arrayLength / 4) + 1

	// iterate and store the partitions into storage
	// split to chunks
	for i := 0; i < len(ex); i += chunkSize {
		end := i + chunkSize

		if end > arrayLength {
			end = arrayLength
		}

		// start to concurrent task

		// add waitgroup counter - each worker must decrease to unblock
		wg.Add(1)
		// start task concurrently
		go concurrencySort(ex[i:end], ch, &wg)
	}

	// force main routine to wait for all workers
	wg.Wait()

	// close channel
	close(ch)



	// iterate through channel and get the output
	var storageAllSlice []int
	for sli := range ch {
		storageAllSlice = append(storageAllSlice,sli...) // concat
		bubbleSort(storageAllSlice)
	}
	fmt.Println(storageAllSlice)



}
