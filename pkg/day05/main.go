package day05

import (
	"strconv"
	"strings"

	"github.com/nandajavarma/aoc2021/pkg/filereader"
	log "github.com/sirupsen/logrus"
)

type point struct {
	x int
	y int
}

func makeRange(min, max int) []int {
	if min < max {
		a := make([]int, max-min+1)
		for i := range a {
			a[i] = min + i
		}
		return a
	} else {
		a := make([]int, min-max+1)
		for i := range a {
			a[i] = min - i
		}
		return a
	}
}

// max returns the larger of x or y.
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// min returns the smaller of x or y.
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func absDiff(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func zip(a, b []int) []point {

	r := make([]point, len(a), len(a))

	for i, e := range a {
		r[i] = point{x: e, y: b[i]}
	}

	return r
}

func getDiaganols(pointA point, pointB point) []point {
	points := getPointInBetween(pointA, pointB)
	if absDiff(pointA.x, pointB.x) == absDiff(pointA.y, pointB.y) {
		pointxs := makeRange(pointA.x, pointB.x)
		pointys := makeRange(pointA.y, pointB.y)
		points = append(points, zip(pointxs, pointys)...)
	}

	return points
}

func getPointInBetween(pointA point, pointB point) []point {
	points := []point{}
	if pointA.x == pointB.x {
		for _, y := range makeRange(min(pointA.y, pointB.y), max(pointA.y, pointB.y)) {
			points = append(points, point{
				x: pointA.x,
				y: y,
			})
		}
	} else if pointA.y == pointB.y {
		for _, x := range makeRange(min(pointA.x, pointB.x), max(pointA.x, pointB.x)) {
			points = append(points, point{
				x: x,
				y: pointA.y,
			})
		}
	}

	return points
}

func newPoint(points []string) point {
	x, _ := strconv.Atoi(points[0])
	y, _ := strconv.Atoi(points[1])
	return point{x: x,
		y: y}
}

func getGrid(input []string, diagnols bool) map[point]int {
	grids := map[point]int{}
	for i := 0; i < len(input); i++ {
		coords := strings.Split(input[i], " -> ")
		if len(coords) < 2 {
			continue
		}
		pointA := newPoint(strings.Split(coords[0], ","))
		pointB := newPoint(strings.Split(coords[1], ","))

		points := []point{}
		if diagnols {
			points = getDiaganols(pointA, pointB)
		} else {
			points = getPointInBetween(pointA, pointB)
		}
		for j := 0; j < len(points); j++ {
			grids[points[j]] += 1
		}
	}

	return grids
}

func getOverlapPoints(grid map[point]int, count int) int {
	overlapPoints := 0

	for _, v := range grid {
		if v >= count {
			overlapPoints++
		}
	}

	return overlapPoints
}

func Run(inputfile string) error {
	input, err := filereader.ReadFile(inputfile)
	if err != nil {
		return err
	}

	// part 1
	grid1 := getGrid(input, false)

	log.Info(getOverlapPoints(grid1, 2))

	// part 2
	grid2 := getGrid(input, true)

	log.Info(getOverlapPoints(grid2, 2))

	return nil

}
