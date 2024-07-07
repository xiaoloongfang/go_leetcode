package main

import (
	"fmt"
	"sort"
)

type Element struct {
	value int
	index int
}

func nextGreaterElements(nums []int) []int {
	elements := make([]Element, len(nums))
	res := make([]int, len(nums))
	for i := 0; i < len(elements); i++ {
		res[i] = -1
	}
	for i := 0; i < len(nums); i++ {
		elements[i] = Element{
			value: nums[i],
			index: i,
		}
	}
	seenElements := make([]Element, 0)
	for i := 0; i < len(elements)*2; i++ {
		pop := 0
		index := i % len(elements)
		for _, seenElement := range seenElements {
			if elements[index].value > seenElement.value {
				res[seenElement.index] = elements[index].value
				pop++
			} else {
				break
			}
		}
		seenElements = seenElements[pop:]
		if i < len(elements) {
			seenElements = append(seenElements, elements[index])
		}
		sort.Slice(seenElements, func(i, j int) bool {
			return seenElements[i].value < seenElements[j].value
		})
	}
	return res
}

func main() {
	fmt.Printf("%v\n", nextGreaterElements([]int{1, 2, 3, 4, 3}))
	fmt.Printf("%v\n", nextGreaterElements([]int{1}))
	fmt.Printf("%v\n", nextGreaterElements([]int{1, 2, 1}))
	fmt.Printf("%v\n", nextGreaterElements([]int{5, 4, 3, 2, 1}))
	fmt.Printf("%v\n", nextGreaterElements([]int{1, 1, 1, 1, 1}))
}
