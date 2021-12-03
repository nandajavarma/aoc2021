package day03

import (
	"strconv"
	"strings"

	"github.com/nandajavarma/aoc2021/pkg/filereader"
	log "github.com/sirupsen/logrus"
)

type bitCount struct {
	one  int
	zero int
}

func getPositionalCount(data []string) map[int]*bitCount {
	countMap := make(map[int]*bitCount, len(data))

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			if countMap[j] == nil {
				countMap[j] = &bitCount{
					one:  0,
					zero: 0,
				}
			}

			if int(data[i][j]) == 48 {
				countMap[j].zero = countMap[j].zero + 1
			} else {
				countMap[j].one = countMap[j].one + 1
			}
		}
	}
	return countMap
}

func countGammaEpsilon(data []string) int64 {
	countMap := getPositionalCount(data)
	gamma := make([]string, len(countMap))
	epsilon := make([]string, len(countMap))

	for i := 0; i < len(countMap); i++ {
		if countMap[i].one > countMap[i].zero {
			gamma[i] = "1"
			epsilon[i] = "0"
		} else {
			gamma[i] = "0"
			epsilon[i] = "1"
		}
	}

	gammaString := strings.Join(gamma, "")
	epsilonString := strings.Join(epsilon, "")
	gammaNum, _ := strconv.ParseInt(gammaString, 2, 64)
	epsilonNum, _ := strconv.ParseInt(epsilonString, 2, 64)

	return gammaNum * epsilonNum
}

func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func countOxygen(data []string) int64 {
	countMap := getPositionalCount(data)

	for i := 0; i < len(countMap); i++ {
		isCommonZero, isCommonOne, isEqual := false, false, false
		if countMap[i].one > countMap[i].zero {
			isCommonOne = true
		} else if countMap[i].one == countMap[i].zero {
			isEqual = true
		} else {
			isCommonZero = true
		}

		for j := 0; j < len(data); j++ {
			if data[j] == "" {
				continue
			}

			if int(data[j][i]) == 48 {
				if isCommonOne {
					data[j] = ""
				} else if isEqual {
					data[j] = ""
				}
			} else {
				if isCommonZero {
					data[j] = ""
				}
			}
		}
		data = deleteEmpty(data)
		if len(data) <= 1 {
			break
		}

		countMap = getPositionalCount(data)
	}

	oxygen, _ := strconv.ParseInt(data[0], 2, 64)

	return oxygen
}

func countCo2(data []string) int64 {
	countMap := getPositionalCount(data)

	for i := 0; i < len(countMap); i++ {
		isCommonZero, isCommonOne, isEqual := false, false, false
		if countMap[i].one > countMap[i].zero {
			isCommonOne = true
		} else if countMap[i].one == countMap[i].zero {
			isEqual = true
		} else {
			isCommonZero = true
		}

		for j := 0; j < len(data); j++ {
			if data[j] == "" {
				continue
			}

			if int(data[j][i]) == 48 {
				if isCommonZero {
					data[j] = ""
				}
			} else {
				if isCommonOne {
					data[j] = ""
				} else if isEqual {
					data[j] = ""
				}
			}
		}
		data = deleteEmpty(data)
		if len(data) <= 1 {
			break
		}

		countMap = getPositionalCount(data)
	}

	co2, _ := strconv.ParseInt(data[0], 2, 64)

	return co2
}

func countOxygenCo2(input []string) int64 {

	input2 := make([]string, len(input))
	copy(input2, input)

	return countOxygen(input) * countCo2(input2)
}

func Run(inputfile string) error {
	lines, err := filereader.ReadFile(inputfile)
	if err != nil {
		return err
	}

	log.Infof("Solution for part 1: %d", countGammaEpsilon(lines))
	log.Infof("Solution for part 2: %d", countOxygenCo2(lines))

	return nil
}
