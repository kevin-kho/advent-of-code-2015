package main

import (
	"aoc-2015/common"
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Reindeer struct {
	Name         string
	Speed        int
	MoveDuration int
	RestDuration int
}

func createReindeers(data []byte) ([]Reindeer, error) {

	var res []Reindeer

	for entry := range bytes.SplitSeq(data, []byte{10}) {
		entryStr := strings.Split(string(entry), " ")

		name := entryStr[0]
		speed, err := strconv.Atoi(entryStr[3])
		if err != nil {
			return res, err
		}
		moveDuration, err := strconv.Atoi(entryStr[6])
		if err != nil {
			return res, err
		}
		restDuration, err := strconv.Atoi(entryStr[len(entryStr)-2])
		if err != nil {
			return res, err
		}

		res = append(res, Reindeer{
			Name:         name,
			Speed:        speed,
			MoveDuration: moveDuration,
			RestDuration: restDuration,
		})

	}

	return res, nil

}

func moveReindeer(r Reindeer, seconds int) int {
	position := 0
	move := r.MoveDuration
	rest := 0
	for i := 1; i < seconds+1; i++ {

		// case: movable
		if move > 0 && rest == 0 {
			position += r.Speed
			move--
			if move == 0 { // finish moving
				rest = r.RestDuration
				continue
			}
		}

		// case resting
		if rest > 0 {
			rest--
			if rest == 0 {
				move = r.MoveDuration
			}
		}

	}

	return position

}

func solvePartOne(reindeers []Reindeer, seconds int) int {
	var maxDistance int
	for _, r := range reindeers {
		maxDistance = max(maxDistance, moveReindeer(r, seconds))
	}

	return maxDistance

}

func main() {
	filePath := "inputExample.txt"
	filePath = "input.txt"
	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}

	data = common.TrimNewLineSuffix(data)

	reindeers, err := createReindeers(data)
	if err != nil {
		log.Fatal(err)
	}

	res := solvePartOne(reindeers, 2503)
	fmt.Println(res)

}
