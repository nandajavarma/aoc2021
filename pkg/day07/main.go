package day07

import (
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/nandajavarma/aoc2021/pkg/filereader"
	log "github.com/sirupsen/logrus"
)

func absDiff(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func mean(list []string) int {
	sum := 0

	for i := 0; i < len(list); i++ {

		num, _ := strconv.Atoi(list[i])
		sum += num
	}

	return int(math.Round((float64(sum)) / (float64(len(list)))))
}

func getMedianFuel(list []string) int {
	sort.Strings(list)
	num, _ := strconv.Atoi(list[len(list)/2])

	fuel := 0
	for _, v := range list {
		dat, _ := strconv.Atoi(v)
		steps := absDiff(dat, num)
		fuel += steps //* (steps + 1) / 2
	}

	return fuel
}

func getMeanFuel(list []string) int {
	sort.Strings(list)
	mean := mean(list)

	fuel := 0
	for _, v := range list {
		dat, _ := strconv.Atoi(v)
		steps := absDiff(dat, int(mean))
		fuel += steps * (steps + 1) / 2
	}

	return fuel
}

func getAlignPoint(pos []string, mean bool) int {

	if mean {
		return getMeanFuel(pos)
	}

	return getMedianFuel(pos)
}

func Run(inputfile string) error {
	input, err := filereader.ReadFile(inputfile)
	if err != nil {
		return err
	}

	pos := strings.Split(input[0], ",")

	log.Infof("Solution to part 1 is %d", getAlignPoint(pos, false))
	log.Infof("Solution to part 2 is %d", getAlignPoint(pos, true))

	return nil
}
