package day01

import (
	"fmt"
	"strconv"

	"github.com/nandajavarma/aoc2021/pkg/filereader"
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

func count_increasing(data []string, window int) int {
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

	// solving part 1
	fmt.Println(count_increasing(lines, 1))

	// solving part 2
	fmt.Println(count_increasing(lines, 3))

	return nil
}
