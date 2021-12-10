package day10

import (
	"sort"
	"strings"

	"github.com/nandajavarma/aoc2021/pkg/filereader"
	log "github.com/sirupsen/logrus"
)

var incomplete [][]string = [][]string{}

func pop(alist *[]string) string {
	f := len(*alist)
	rv := (*alist)[f-1]
	*alist = (*alist)[:f-1]
	return rv
}

func point(paran string) int {
	points := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}

	return points[paran]
}

func findCorruptedLines(expr []string, unmatched map[string]int) map[string]int {

	stack := []string{}
	for i := 0; i < len(expr); i++ {
		char := expr[i]

		if char == "{" || char == "[" || char == "(" || char == "<" {
			stack = append(stack, char)
			continue
		}

		if len(stack) < 1 {
			continue
		}

		x := pop(&stack)

		switch x {
		case "{":
			if char != "}" {
				unmatched[char] = unmatched[char] + 1
				return unmatched
			}
		case "[":
			if char != "]" {
				unmatched[char] = unmatched[char] + 1
				return unmatched
			}
		case "(":
			if char != ")" {
				unmatched[char] = unmatched[char] + 1
				return unmatched
			}
		case "<":
			if char != ">" {
				unmatched[char] = unmatched[char] + 1
				return unmatched
			}
		}
	}

	if len(stack) >= 1 {
		incomplete = append(incomplete, stack)
	}

	return unmatched
}

func getUnmatchedParanthesesScore(input []string) int {

	unmatched := map[string]int{
		")": 0,
		"}": 0,
		"]": 0,
	}

	for i := 0; i < len(input); i++ {
		line := strings.Split(input[i], "")
		if len(line) < 1 {
			continue
		}

		unmatched = findCorruptedLines(line, unmatched)

	}

	finalPoints := 0
	for k, v := range unmatched {
		finalPoints += v * point(k)

	}

	return finalPoints
}

func getMatch(char string) string {
	switch char {
	case "{":
		return "}"
	case "[":
		return "]"
	case "(":
		return ")"
	case "<":
		return ">"
	}

	return ""
}

func getMatchPoints(char string) int {
	switch char {
	case "}":
		return 3
	case "]":
		return 2
	case ")":
		return 1
	case ">":
		return 4
	}

	return 0
}

func matchParantheses(expr []string) []string {
	complete := make([]string, len(expr))

	for i, v := range expr {
		complete[len(expr)-1-i] = getMatch(v)
	}

	return complete
}

func matchParanthesesScore(incomplete [][]string) int {
	points := make([]int, len(incomplete))
	for i := 0; i < len(incomplete); i++ {
		total := 0
		completed := matchParantheses(incomplete[i])
		for j := 0; j < len(completed); j++ {
			if len(completed[j]) > 0 {
				total = total*5 + getMatchPoints(completed[j])
			}
		}
		points[i] = total
	}

	sort.Ints(points)

	return points[(len(points)-1)/2]
}

func Run(inputfile string) error {
	input, _ := filereader.ReadFile(inputfile)

	log.Infof("solution to part 1 is: %d", getUnmatchedParanthesesScore(input))
	log.Infof("solution to part 2 is: %d", matchParanthesesScore(incomplete))

	return nil
}
