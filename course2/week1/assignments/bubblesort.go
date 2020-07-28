package main

import "fmt"

// this function will swap the ith index of slice with its i+1 index element

func swap(sli []int, i int) {
	var tempI int = sli[i]
	var tempIP1 int = sli[i+1]

	sli[i] = tempIP1
	sli[i+1] = tempI

}

// bubble sort pseudo code:https://www.chegg.com/homework-help/questions-and-answers/q1-bubblesort-another-simple-sorting-algorithm-whose-pseudocode-follows-1-write-c-bubbleso-q8984069

// BubbleSort func
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

func main() {

	var ex = make([]int, 0)
	var UserInput int
	// user input the integers
	for i := 0; i < 10; i++ {
		fmt.Println("Please input an integer...")
		_, _ = fmt.Scan(&UserInput)
		ex = append(ex, UserInput)

	}

	// sort
	bubbleSort(ex)
	fmt.Println(ex)

}
