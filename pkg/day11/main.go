package day11

import (
	"strconv"
	"strings"

	"github.com/nandajavarma/aoc2021/pkg/filereader"
	log "github.com/sirupsen/logrus"
)

// create a grid with x, y as indexes
// increase each by one, if > 9, set to zero and add to just exploded column
// iterate through just exploded again, increment the neighbors one, if they have not yet exploded. (iterate this function)
//
//
//
var explosions int = 0

type coord struct {
	x int
	y int
}

func getNeighborPos(i, j, iBound, jBound int) []coord {
	neighbours := []coord{
		{
			x: i,
			y: j + 1,
		},
		{
			x: i + 1,
			y: j,
		},
		{
			x: i - 1,
			y: j,
		},
		{
			x: i,
			y: j - 1,
		},
		{
			x: i + 1,
			y: j + 1,
		},
		{
			x: i - 1,
			y: j + 1,
		},
		{
			x: i + 1,
			y: j - 1,
		},
		{
			x: i - 1,
			y: j - 1,
		},
	}

	result := []coord{}

	for _, val := range neighbours {
		if val.x < 0 || val.x >= iBound {
			continue
		}

		if val.y < 0 || val.y >= jBound {
			continue
		}

		if val.x == i && val.y == j {
			continue
		}

		result = append(result, val)
	}

	return result
}

func incrementNeighboursExplosion(grid [][]int, lastExplosion []coord, exploded map[coord]bool, iter int) [][]int {
	newExploded := []coord{}
	for _, point := range lastExplosion {
		neighbours := getNeighborPos(point.x, point.y, len(grid), len((grid)[0]))
		if len(neighbours) < 1 {
			return grid
		}

		for _, x := range neighbours {
			if !exploded[x] {
				newVal := (grid)[x.x][x.y] + 1
				if newVal > 9 {
					explosions++
					newExploded = append(newExploded, x)
					exploded[x] = true
					newVal = 0
				}
				grid[x.x][x.y] = newVal
			}
		}
	}

	if len(exploded) == len(grid)*len(grid[0]) {
		log.Infof("yep all, synchronized on iter %d", iter+1)
	}

	if len(newExploded) < 1 {
		return grid
	}

	return incrementNeighboursExplosion(grid, newExploded, exploded, iter)
}

func lightCycleOctopus(grid [][]int, iter int) [][]int {
	exploded := map[coord]bool{}
	newExplosions := []coord{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			newVal := grid[i][j] + 1
			if newVal > 9 {
				explosions++
				newVal = 0
				point := coord{
					x: i,
					y: j,
				}
				exploded[point] = true
				newExplosions = append(newExplosions, point)
			}
			grid[i][j] = newVal
		}
	}

	return incrementNeighboursExplosion(grid, newExplosions, exploded, iter)
}

func createGrid(input []string) [][]int {
	grid := make([][]int, len(input)-1)
	for i := 0; i < len(input); i++ {
		line := input[i]
		if len(line) < 1 {
			continue
		}

		items := strings.Split(line, "")

		grid[i] = make([]int, len(items))
		for j := 0; j < len(items); j++ {
			val, err := strconv.Atoi(items[j])
			if err != nil {
				continue
			}

			grid[i][j] = val

		}
	}

	return grid
}

func getExplosions(grid [][]int, iters int) [][]int {
	for i := 0; i < iters; i++ {
		grid = lightCycleOctopus(grid, i)
	}

	return grid
}

func Run(inputfile string) error {
	input, _ := filereader.ReadFile(inputfile)
	grid := createGrid(input)

	grid = getExplosions(grid, 100)

	log.Infof("Solution to part 1 is %d", explosions)

	return nil
}
