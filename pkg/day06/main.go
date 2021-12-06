package day06

import (
	"strconv"
	"strings"

	"github.com/nandajavarma/aoc2021/pkg/filereader"
	log "github.com/sirupsen/logrus"
)

func simulateFishLife(fishCount map[int]int) map[int]int {
	for i := 0; i < 9; i++ {
		fishCount[i-1] = fishCount[i]
	}

	fishCount[8] = fishCount[-1]
	fishCount[6] += fishCount[-1]
	fishCount[-1] = 0

	return fishCount
}

func getTotalFish(initial string, day int) int {
	fish := strings.Split(initial, ",")
	fishCount := make(map[int]int, 9)
	for i := 0; i < len(fish); i++ {
		count, _ := strconv.Atoi(fish[i])
		fishCount[count]++
	}

	for i := 0; i < day; i++ {
		fishCount = simulateFishLife(fishCount)
	}

	count := 0
	for _, v := range fishCount {
		count += v
	}

	return count
}

func Run(inputfile string) error {
	data, err := filereader.ReadFile(inputfile)
	if err != nil {
		return err
	}

	log.Infof("Solution to part 1 is %d", getTotalFish(data[0], 80))
	log.Infof("Solution to part 1 is %d", getTotalFish(data[0], 256))
	return nil
}
