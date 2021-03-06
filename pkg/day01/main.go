package day01

import (
	"strconv"

	"github.com/nandajavarma/aoc2021/pkg/filereader"
	log "github.com/sirupsen/logrus"
)

func getInt(data string) int {
	v, _ := strconv.Atoi(data)

	return v
}

func isIncreasing(a, b int) bool {
	return a > b
}

func getWindowSum(data []string, idx int, window int) int {
	sum := 0
	if len(data[idx:]) < window {
		return 0
	}

	for i := 0; i < window; i++ {
		sum += getInt(data[idx+i])
	}

	return sum
}

func countIncreasing(data []string, window int) int {
	increasing_count := 0

	prevSum := 0

	for i := -1; i < len(data)-1; i++ {
		currSum := getWindowSum(data, i+1, window)
		if currSum == 0 {
			// the windowSum returns 0 if there are not enough items left
			break
		}

		if prevSum > 0 && prevSum < currSum {
			increasing_count++
		}

		prevSum = currSum
	}

	return increasing_count
}

func Run(inputfile string) error {
	lines, err := filereader.ReadFile(inputfile)
	if err != nil {
		return err
	}

	log.Infof("Solution for part 1: %d", countIncreasing(lines, 1))

	log.Infof("Solution for part 2: %d", countIncreasing(lines, 3))

	return nil
}
