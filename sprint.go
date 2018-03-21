package main

import (
	"sort"
)

//https://codefights.com/arcade/intro/level-1/egbueTZRRL5Mm4TXN
func centuryFromYear(year int) int {
	century := year / 100
	if year < (century*100 + 1) {
		return century
	}
	return century + 1

}

//https://codefights.com/arcade/intro/level-1/s5PbmwxfECC52PWyQ
func checkPalindrome(inputString string) bool {
	sizeString := len(inputString)
	isEqual := false
	if sizeString == 1 {
		return true
	}
	for i := range inputString {
		if inputString[i] == inputString[sizeString-(i+1)] {
			isEqual = true
		} else {
			isEqual = false
			return isEqual
		}
	}
	return isEqual
}

//sum of a array elements
func sum(args ...int) int {
	sumVar := 0
	for _, item := range args {
		sumVar += item
	}
	return sumVar
}

//https://codefights.com/arcade/intro/level-2/xzKiBHjhoinnpdh6m
func adjacentElementsProduct(inputArray []int) int {
	var maxValue int
	for i := range inputArray {
		if i == 0 {
			maxValue = inputArray[i] * inputArray[i+1]
		} else if i <= len(inputArray)-2 {
			if inputArray[i]*inputArray[i-1] > maxValue {
				maxValue = inputArray[i] * inputArray[i-1]
			} else if inputArray[i+1]*inputArray[i] > maxValue {
				maxValue = inputArray[i+1] * inputArray[i]
			}
		}
	}
	return maxValue
}

//https://codefights.com/arcade/intro/level-2/yuGuHvcCaFCKk56rJ
func shapeArea(n int) int {
	area := 1
	for i := range make([]int, n-1) {
		area += (0 + (i+1)*4)
	}
	return area

}

//https://codefights.com/arcade/intro/level-2/bq2XnSr5kbHqpHGJC
func rangeNew(start, end, step int) []int {
	array := make([]int, end)
	for i := range array {
		if i == 0 {
			array[i] = start
		} else {
			start += step
			array[i] = start
		}
	}
	return array
}
func makeArrayConsecutive2(statues []int) int {
	sort.Ints(statues)
	values := []int{}
	for i := range statues {
		if i <= len(statues)-2 {
			if statues[i] != statues[i+1]-1 {
				values = append(values,
					rangeNew(statues[i]+1, statues[i+1]-statues[i]-1, 1)...)
			}
		}
	}
	return len(values)
}

func inArray(key int, array []int) (bool, int) {
	for i, item := range array {
		if item == key {
			return true, i
		}
	}
	return false, 0
}

//https://codefights.com/arcade/intro/level-2/2mxbGwLzvkTCKAJMG
func myAreSorted(array []int) bool {
	for i := range array {
		if i <= len(array)-2 {
			if array[i] > array[i+1] || array[i] == array[i+1] {
				return false
			}
		}
	}
	return true
}

func myreverse(array []int) []int {
	values := make([]int, len(array))
	for index := len(array) - 1; index >= 0; index-- {
		values[(len(array)-1)-index] = array[index]
	}
	return values
}

func almostIncreasingSequence(sequence []int) bool {
	var item int
	var values []int

	for index := len(sequence) - 1; index >= 0; index-- {

		item, sequence = sequence[index], sequence[:index]
		values = append(values, item)
		sequence = append(sequence, myreverse(values[:len(values)-1])...)
		if myAreSorted(sequence) {
			return true
		}
	}
	return false
}
