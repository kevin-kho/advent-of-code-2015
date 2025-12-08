package main

import (
	"fmt"
	"slices"
)

func lookAndSay(slc []int) []int {

	var res []int

	digit := slc[0]
	count := 1

	for p := 1; p < len(slc); p++ {
		if slc[p] != digit {
			res = append(res, count)
			res = append(res, digit)

			digit = slc[p]
			count = 1

		} else {
			count++
		}
	}

	res = append(res, count)
	res = append(res, digit)

	return res

}

func main() {
	input := []int{1, 1, 1, 3, 2, 2, 2, 1, 1, 3}
	part1 := slices.Clone(input)
	part2 := slices.Clone(input)

	for range 40 {
		part1 = lookAndSay(part1)
	}
	fmt.Println(len(part1))

	for range 50 {
		part2 = lookAndSay(part2)
	}
	fmt.Println(len(part2))

}
