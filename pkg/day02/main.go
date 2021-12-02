package day02

import (
	"strconv"
	"strings"

	"github.com/nandajavarma/aoc2021/pkg/filereader"
	log "github.com/sirupsen/logrus"
)

func get_position(data []string, aimed bool) int {
	x, y, displacement := 0, 0, 0
	for i := 0; i < len(data); i++ {
		path := strings.Split(data[i], " ")
		if len(path) < 2 {
			break
		}
		distance, _ := strconv.Atoi(path[1])
		switch path[0] {
		case "forward":
			x = x + distance
			displacement = displacement + distance*y
		case "up":
			y = y - distance
		case "down":
			y = y + distance
		}
	}

	if aimed {
		return x * displacement
	}

	return x * y
}

func Run(inputfile string) error {
	lines, err := filereader.ReadFile(inputfile)
	if err != nil {
		return err
	}

	log.Infof("Solution for part 1: %d", get_position(lines, false))
	log.Infof("Solution for part 2: %d", get_position(lines, true))

	return nil
}
