package main

import (
	"fmt"
	"slices"
)

func convertToRuneSlc(input string) []rune {
	var res []rune
	for _, c := range input {
		res = append(res, c)
	}
	return res
}

func incrementRuneSlc(slc []rune) {
	for i := len(slc) - 1; i >= 0; i-- {
		slc[i]++
		if slc[i]%123 == 0 {
			slc[i] = 97
			continue
		} else {
			break
		}
	}

}

func containsForbiddenRune(slc []rune) bool {
	return slices.Contains(slc, 105) || slices.Contains(slc, 111) || slices.Contains(slc, 108)
}

func containsStraightRunes(slc []rune) bool {
	for i := 2; i < len(slc); i++ {
		if (slc[i]-1 == slc[i-1]) && (slc[i-1]-1 == slc[i-2]) {
			return true
		}
	}
	return false
}

func containsTwoUniquePairs(slc []rune) bool {
	seen := make(map[rune]bool)
	for i := 1; i < len(slc); i++ {
		if slc[i-1] == slc[i] {
			seen[slc[i]] = true
		}
	}

	return len(seen) >= 2

}

func main() {

	oldPassword := "cqjxjnds"
	runeSlc := convertToRuneSlc(oldPassword)

	// Part 1
	for true {
		if containsStraightRunes(runeSlc) && containsTwoUniquePairs(runeSlc) && !containsForbiddenRune(runeSlc) {
			break
		}
		incrementRuneSlc(runeSlc)
	}
	fmt.Println(string(runeSlc))

	// Part 2
	incrementRuneSlc(runeSlc)
	for true {
		if containsStraightRunes(runeSlc) && containsTwoUniquePairs(runeSlc) && !containsForbiddenRune(runeSlc) {
			break
		}
		incrementRuneSlc(runeSlc)
	}
	fmt.Println(string(runeSlc))

}
