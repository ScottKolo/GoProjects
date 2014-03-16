// Implement two types of sorting algorithms: Merge sort and bubble sort.

package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {

	var n int
	fmt.Printf("How long of an array should we sort? (≤1000) ")
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if n > 1000 {
		fmt.Println("Error:", n, "> 1000 - too long!")
		return
	} else if n < 1 {
		fmt.Println("Error:", n, "< 1")
		return
	}

	array := make([]int, n)
	rand.Seed(time.Now().UTC().UnixNano())

	fmt.Println("Original Array:")
	for i := range array {
		array[i] = rand.Intn(10000)
		fmt.Printf("%d\t", array[i])
	}
	fmt.Println()

	bubbleSortedArray := BubbleSort(array)

	fmt.Println("Bubble Sorted Array: ")
	for i := range bubbleSortedArray {
		fmt.Printf("%d\t", bubbleSortedArray[i])
	}
	fmt.Println()

	mergeSortedArray := MergeSort(array)

	fmt.Println("Merge Sorted Array: ")
	for i := range mergeSortedArray {
		fmt.Printf("%d\t", mergeSortedArray[i])
	}
	fmt.Println()
}

// Implementing a fairly "dumb" bubble sort
func BubbleSort(inputArray []int) (sortedArray []int) {
	array := make([]int, len(inputArray))
	copy(array, inputArray)
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-1; j++ {
			if array[j] > array[j+1] {
				// Idiomatic Go swap
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

// Implementing a top-down merge sort
// Adapted from the Merge Sort C code from Wikipedia
// Used idiomatic Go where possible
func MergeSort(array []int) (sortedArray []int) {
	sortedArray = make([]int, len(array), cap(array))
	splitMerge(array, sortedArray)
	return array
}

func splitMerge(A []int, B []int) {
    if(len(A) < 2) {   // if run size == 1, consider it sorted
    	return
    }                       
    // recursively split runs into two halves until run size == 1,
    // then merge them and return back up the call chain
    iMiddle := len(A) / 2              // iMiddle = mid point
    splitMerge(A[:iMiddle], B[:iMiddle])  // split / merge left  half
    splitMerge(A[iMiddle:], B[iMiddle:])  // split / merge right half
    merge(A, B, iMiddle)  // merge the two half runs
    copy(A, B)             // copy the merged runs back to A
}
 
func merge(A []int, B []int, iMiddle int) {
    i0 := 0
    i1 := len(A) / 2
    iEnd := len(A)
 
    for j := range A {
        // If left run head exists and is <= existing right run head.
        if (i0 < iMiddle && (i1 >= iEnd || A[i0] <= A[i1])) {
            B[j] = A[i0]  
            i0++  // Increment i0 after using it as an index.
        } else {
            B[j] = A[i1]
            i1++  // Increment i1 after using it as an index.
        }
    }
 
}