package main

import "fmt"

/*
  Pseudocodigo

  function linearSearch(array, target):
    for index from 0 to length(array) -1:
      if array[index] == target:
	return index
      return -1: // not found value in the list
*/

func linearSearch(list []int, target int) int {
	for i, num := range list {
		if num == target {
			return i
		}
	}
	return -1
}

func main() {
	searchList := []int{12, 13, 24, 99}
	target := 14

	if index := linearSearch(searchList, target); index != -1 {
		fmt.Printf("Element %d found at index %d\n", target, index)
	} else {
		fmt.Printf("Element %d not found in the list\n", target)
	}
}
