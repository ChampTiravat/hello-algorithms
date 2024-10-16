package main

import "fmt"

func main() {
	list := [10]int{2, 3, 6, 1, 3, 9, 28, 6, 0, -10}
	fmt.Printf("unsorted: %v\n", list)

	BubbleSort(&list)
	fmt.Printf("  sorted: %v\n", list)
}

func BubbleSort(list *[10]int) {
	isSorted := false
	for !isSorted {
		isSorted = true
		for i := 0; i < len(*list)-1; i++ {
			if (*list)[i] > (*list)[i+1] {
				(*list)[i], (*list)[i+1] = (*list)[i+1], (*list)[i]
				isSorted = false
			}
		}
	}
}

func InsertionSort() {}

func SelectionSort() {}

func MergeSort() {}

func QuickSort() {}

func HeapSort() {}

func CountingSort() {}

func RadixSort() {}

func BucketSort() {}
