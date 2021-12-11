package day09

import (
	"strconv"
	"strings"

	"github.com/nandajavarma/aoc2021/pkg/filereader"
	log "github.com/sirupsen/logrus"
)

var basinlength []int

func getMatrix(input []string) [][]int {
	matrix := make([][]int, len(input)-1)
	for i := 0; i < len(input); i++ {
		items := strings.Split(input[i], "")
		if len(items) < 1 {
			continue
		}

		matrix[i] = make([]int, len(items))

		for j := 0; j < len(items); j++ {
			v, _ := strconv.Atoi(items[j])
			matrix[i][j] = v
		}
	}

	return matrix
}

type coords struct {
	x          int
	y          int
	riskFactor int
}

func getNeighborPos(i, j, iBound, jBound int) []coords {
	coord := []coords{
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
	}

	result := []coords{}

	for _, val := range coord {
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

// func findBasins(matrix [][]int, neighbours []coords, basinCount int, visited map[coords]bool) int {
// 	if len(neighbours) == 0 {
// 		return basinCount
// 	}

// 	for _, coord := range neighbours {
// 		if visited[coord] {
// 			return basinCount
// 		}

// 		if matrix[coord.x][coord.y] == 9 {
// 			return basinCount
// 		}
// 		log.Infof("Considering coord %v", coord)

// 		visited[coord] = true
// 		basinCount = findBasins(matrix, getNeighborPos(coord.x, coord.y, len(matrix), len(matrix[0])), basinCount+1, visited)
// 	}

// 	return basinCount
// }
//
func findBasins(matrix [][]int, neighbours []coords, basinCount int, visited map[coords]bool) int {
	if len(neighbours) == 0 {
		return basinCount
	}

	for i := 0; i < len(neighbours); i++ {
		item := neighbours[i]
		if visited[item] {
			return basinCount
		}

		visited[item] = true
		if matrix[item.x][item.y] == 9 {
			return basinCount
		}

		return findBasins(matrix, getNeighborPos(item.x, item.y, len(matrix), len(matrix[0])), basinCount+1, visited)
	}

	return basinCount
}

func isLowest(matrix [][]int, i, j int) bool {
	neighborCoords := getNeighborPos(i, j, len(matrix), len(matrix[0]))

	for k := 0; k < len(neighborCoords); k++ {
		coord := neighborCoords[k]
		if matrix[coord.x][coord.y] <= matrix[i][j] {
			return false
		}
	}

	visits := map[coords]bool{}
	visits[coords{
		x: i,
		y: j,
	}] = true
	basinlength = append(basinlength, findBasins(matrix, neighborCoords, 1, visits))

	return true
}

func findLowPoints(matrix [][]int) []coords {
	points := []coords{}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if isLowest(matrix, i, j) {
				points = append(points, coords{
					x:          i,
					y:          j,
					riskFactor: matrix[i][j] + 1,
				})
			}
		}
	}

	return points
}

func Run(inputfile string) error {
	input, err := filereader.ReadFile(inputfile)
	if err != nil {
		return err
	}

	matrix := getMatrix(input)
	lowPoints := findLowPoints(matrix)

	risk := 0
	for _, val := range lowPoints {
		risk += val.riskFactor
	}

	log.Infof("Solution to problem 1 is: %d", risk)
	log.Info(basinlength)

	return nil
}
