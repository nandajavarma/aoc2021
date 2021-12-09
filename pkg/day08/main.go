package day08

import (
	"strings"

	"github.com/nandajavarma/aoc2021/pkg/filereader"
	log "github.com/sirupsen/logrus"
)

func isAcceptedLength(a string, lens []int) bool {
	for i := 0; i < len(lens); i++ {
		if lens[i] == len(a) {
			return true
		}
	}

	return false
}

func unique(arr string) bool {
	m := make(map[rune]bool)
	for _, i := range arr {
		_, ok := m[i]
		if ok {
			return false
		}

		m[i] = true
	}

	return true
}

func getUniqueDigits(data []string, lens []int) []string {
	digits := []string{}
	for j := 0; j < len(data); j++ {
		if unique(data[j]) && isAcceptedLength(data[j], lens) {
			digits = append(digits, data[j])
		}
	}
	return digits
}

func countUniqueOutput(input []string) int {
	count := 0
	for i := 0; i < len(input); i++ {
		inputAndOutput := strings.Split(input[i], "|")
		if len(inputAndOutput) < 2 {
			continue
		}
		output := strings.Split(inputAndOutput[1], " ")
		count += len(getUniqueDigits(output, []int{2, 3, 4, 7}))
	}

	return count

}
func getAlphabetMap(str string) map[string]int {
	alphMap := make(map[string]int, 7)
	for _, i := range str {
		char := string(i)
		if char != " " {
			alphMap[string(i)]++
		}
	}

	return alphMap
}

func calculateScore(digit string, alphaMap map[string]int) int {
	count := 0
	for _, i := range digit {
		count += alphaMap[string(i)]
	}

	return count
}

func buildDigitScore(input []string, alphaMap map[string]int) map[int]int {
	digitMap := make(map[int]int)
	digit7 := getUniqueDigits(input, []int{3})[0]
	digit4 := getUniqueDigits(input, []int{4})[0]
	digit1 := getUniqueDigits(input, []int{2})[0]
	digit8 := getUniqueDigits(input, []int{7})[0]

	digitMap[7] = calculateScore(digit7, alphaMap)
	digitMap[4] = calculateScore(digit4, alphaMap)
	digitMap[1] = calculateScore(digit1, alphaMap)
	digitMap[8] = calculateScore(digit8, alphaMap)

	e, g, b := 0, 0, 0

	for _, v := range alphaMap {
		switch v {
		case 6:
			e = v
		case 4:
			g = v
		case 9:
			b = v
		}
	}

	a := digitMap[1] - b

	digitMap[6] = digitMap[8] - a
	digitMap[3] = digitMap[8] - e - g
	digitMap[2] = digitMap[8] - e - b
	digitMap[9] = digitMap[8] - g
	digitMap[5] = digitMap[8] - a - g

	return digitMap
}

func getKey(hash map[int]int, val int) int {
	for k, v := range hash {
		if v == val {
			return k
		}
	}

	return 0
}

func decodeOutput(output []string, alphaMap map[string]int, digitScore map[int]int) int {
	outDigit := 0
	for i := 0; i < len(output); i++ {
		if len(output[i]) < 1 {
			continue
		}

		score := calculateScore(output[i], alphaMap)
		digit := getKey(digitScore, score)
		outDigit = outDigit*10 + digit
	}

	return outDigit
}

func countCharacterScore(input []string) int {
	count := 0

	for i := 0; i < len(input); i++ {
		inputAndOutput := strings.Split(input[i], "|")

		if len(inputAndOutput) < 2 {
			continue
		}

		alphaMap := getAlphabetMap(inputAndOutput[0])

		digitScore := buildDigitScore(strings.Split(inputAndOutput[0], " "), alphaMap)

		val := decodeOutput(strings.Split(inputAndOutput[1], " "), alphaMap, digitScore)
		count += val

	}

	return count
}

func Run(inputfile string) error {
	input, err := filereader.ReadFile(inputfile)
	if err != nil {
		return err
	}

	log.Infof("Solution to part 1 is %d", countUniqueOutput(input))

	log.Infof("Solution to part 2 is %d", countCharacterScore(input))

	return nil
}
