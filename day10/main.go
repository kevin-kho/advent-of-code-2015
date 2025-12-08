package main

import (
	"fmt"
)

func lookAndSay(slc []int) []int {

	var res []int
	p := 1

	digit := slc[0]
	count := 1

	for p < len(slc) {
		if slc[p] != digit {
			res = append(res, count)
			res = append(res, digit)

			digit = slc[p]
			count = 1

		} else {
			count++
		}
		p++
	}

	res = append(res, count)
	res = append(res, digit)

	return res

}

func main() {
	input := []int{1, 1, 1, 3, 2, 2, 2, 1, 1, 3}

	for range 40 {
		input = lookAndSay(input)
	}

	fmt.Println(len(input))

}
