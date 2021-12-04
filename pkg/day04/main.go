package day04

import (
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/nandajavarma/aoc2021/pkg/filereader"
	log "github.com/sirupsen/logrus"
)

func stripSpaces(s string) []string {
	r := regexp.MustCompile("[^\\s]+")
	return r.FindAllString(s, -1)
}

func BuildGrid(nums []string) map[[2]int]int {
	grid := make(map[[2]int]int)

	x := 0
	for i := 0; i < len(nums); i++ {
		elems := stripSpaces(nums[i])
		if len(elems) < 1 {
			continue
		}

		for j := 0; j < len(elems); j++ {
			num, _ := strconv.Atoi(elems[j])
			grid[[2]int{x, j}] = num
		}
		x++
	}

	return grid
}

func findCell(grid map[[2]int]int, num int) []int {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if grid[[2]int{i, j}] == num {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

func calculateScore(grid map[[2]int]int, markedGrid map[int]bool, num int) int {
	totalUnmarked := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			val := grid[[2]int{i, j}]
			if !markedGrid[val] {
				totalUnmarked += val

			}
		}
	}

	return totalUnmarked * num

}

func isLineCompleted(grid map[[2]int]int, x, y int, row bool, markCalled map[int]bool) bool {
	for iter := 0; iter < 5; iter++ {
		var val int

		if row {
			val = grid[[2]int{x, iter}]
		} else {
			val = grid[[2]int{iter, y}]
		}

		if !markCalled[val] {
			return false
		}

		if iter == 4 {
			return true
		}
	}

	return false
}

func isComplete(grid map[[2]int]int, mapCalled map[int]bool, num int) bool {
	cell := findCell(grid, num)
	if len(cell) == 0 {
		// this means the newly called number is not present in the grid; skip!
		return false
	}

	rowCompleted := isLineCompleted(grid, cell[0], cell[1], true, mapCalled)
	if rowCompleted {
		return true
	}

	colCompleted := isLineCompleted(grid, cell[0], cell[1], false, mapCalled)
	if colCompleted {
		return true
	}

	return false
}

func simulateWin(grids []map[[2]int]int, calledNums []string, squidwin bool) int {
	markCalled := make(map[int]bool)
	gridsCompleted := make(map[int]bool)
	score := 0

	for i := 0; i < len(calledNums); i++ {
		if len(gridsCompleted) == len(grids) {
			break
		}

		called_num, _ := strconv.Atoi(calledNums[i])
		markCalled[called_num] = true

		for k := 0; k < len(grids); k++ {
			if gridsCompleted[k] {
				continue
			}

			if isComplete(grids[k], markCalled, called_num) {
				gridsCompleted[k] = true

				log.Infof("Grid %d is complete, on call of number %d", k+1, called_num)

				score = calculateScore(grids[k], markCalled, called_num)

				if !squidwin {
					return score
				}
			}
		}
	}

	return score
}

func Run(inputfile string) error {
	input, err := filereader.ReadFile(inputfile)
	if err != nil {
		return err
	}

	calledNums := strings.Split(input[0], ",")

	grids := []map[[2]int]int{}
	for i := 1; i <= len(input)-5; i = i + 6 {
		grid_ret := BuildGrid(input[i : i+6])
		if grid_ret != nil && !reflect.DeepEqual(grid_ret, map[[2]int]int{}) {
			grids = append(grids, grid_ret)
		}
	}

	log.Infof("Solution for part 1 is: %d", simulateWin(grids, calledNums, false))
	log.Infof("Solution for part 2 is: %d", simulateWin(grids, calledNums, true))

	return nil
}
