package day12

import (
	"strings"
	"unicode"

	"github.com/nandajavarma/aoc2021/pkg/filereader"
	log "github.com/sirupsen/logrus"
)

type Node struct {
	small bool
	name  string
	next  []string
}

var start Node = Node{
	small: false,
	name:  "start",
}

var end Node = Node{
	small: false,
	name:  "end",
}

func isLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func buildGraph(input []string) map[string]Node {
	graph := map[string]Node{
		"start": start,
		"end":   end,
	}

	for i := 0; i < len(input); i++ {
		line := strings.Split(input[i], "-")
		var prevName string = ""
		for j := 0; j < len(line); j++ {
			item := graph[line[j]]
			if item.name == "" {
				item = Node{
					small: isLower(line[j]),
					name:  line[j],
				}
			}

			if prevName != "" {
				prev := graph[prevName]
				item.next = append(item.next, prev.name)
				prev.next = append(prev.next, item.name)

				graph[prevName] = prev
			}

			graph[line[j]] = item

			prevName = item.name

		}
	}

	return graph

}

func isInPath(path []string, name string) bool {
	for _, x := range path {
		if x == name {
			return true
		}
	}

	return false
}

func visitNodes(graph map[string]Node, node Node, visited []string, oneExtraSmall bool) int {
	if node.name == "end" {
		return 1
	}

	if isInPath(visited, node.name) {
		if node.name == "start" {
			return 0
		}

		if node.small {
			if !oneExtraSmall {
				return 0
			}

			oneExtraSmall = false
		}
	}

	visited = append(visited, node.name)

	sum := 0
	for _, child := range node.next {
		sum += visitNodes(graph, graph[child], visited, oneExtraSmall)
	}

	return sum
}

func Run(inputfile string) error {
	input, _ := filereader.ReadFile(inputfile)

	graph := buildGraph(input)

	log.Infof("Solution to part 1 is %d", visitNodes(graph, graph["start"], []string{}, false))
	log.Infof("Solution to part 2 is %d", visitNodes(graph, graph["start"], []string{}, true))

	return nil
}
