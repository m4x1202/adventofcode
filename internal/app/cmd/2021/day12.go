package cmd2021

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

type Path struct {
	joker uint8
	nodes []*Node
}

type Node struct {
	name      string
	singleUse bool
	next      []*Node
}

func (n *Node) GetAllPathsAsStrings(path Path) []string {
	if n.singleUse && nodePoolContains(path.nodes, n.name) > 0 {
		if n.name == "start" || path.joker == 0 {
			return nil
		} else {
			path.joker--
		}
	}
	path.nodes = append(path.nodes, n)
	day12logger.Trace().Msgf("Visiting %s", ToString(path.nodes))
	if n.name == "end" {
		return []string{ToString(path.nodes)}
	}
	res := []string{}
	for _, node := range n.next {
		res = append(res, node.GetAllPathsAsStrings(path)...)
	}
	return res
}

func init() {
	cmd2021.AddCommand(day12Cmd)

	day12Cmd.AddCommand(day12part1Cmd)
	day12Cmd.AddCommand(day12part2Cmd)
}

var (
	day12logger = log.With().
			Int("day", 12).
			Logger()
	day12Cmd = &cobra.Command{
		Use:   "day12",
		Short: "Day 12 Challenge",
	}
	day12part1Cmd = &cobra.Command{
		Use:   "part1",
		Short: "Day 12 Part 1 Challenge",
		Run: func(cmd *cobra.Command, args []string) {
			day12part1logger := day12logger.With().
				Int("part", 1).
				Logger()
			day12part1logger.Info().Msg("Start")
			converted := prepareday12Input()

			allPaths := converted.GetAllPathsAsStrings(Path{})

			fmt.Printf("%d\n", len(allPaths))
		},
	}
	day12part2Cmd = &cobra.Command{
		Use:   "part2",
		Short: "Day 12 Part 2 Challenge",
		Run: func(cmd *cobra.Command, args []string) {
			day12part2logger := day12logger.With().
				Int("part", 2).
				Logger()
			day12part2logger.Info().Msg("Start")
			converted := prepareday12Input()

			allPaths := converted.GetAllPathsAsStrings(Path{joker: 1})

			fmt.Printf("%d\n", len(allPaths))
		},
	}
)

func prepareday12Input() *Node {
	content, err := os.ReadFile("resources/day12.txt")
	if err != nil {
		day12logger.Fatal().Err(err).Send()
	}

	input := strings.Split(string(content), "\n")
	day12logger.Info().Msgf("length of input file: %d", len(input))
	day12logger.Debug().Msgf("plain input: %v", input)

	var startNode *Node

	nodePool := make([]string, 0, len(input))
	for _, connection := range input {
		nodes := strings.Split(connection, "-")
		nodePool = append(nodePool, nodes...)
	}

	converted := make([]*Node, 0, len(input))
	for _, nodeCandidate := range nodePool {
		if nodePoolContains(converted, nodeCandidate) == 0 {
			newNode := &Node{name: nodeCandidate, singleUse: IsLower(nodeCandidate)}
			if newNode.name == "start" {
				startNode = newNode
			}
			converted = append(converted, newNode)
		}
	}

	for _, connection := range input {
		nodes := strings.Split(connection, "-")
		a := nodePoolGet(converted, nodes[0])
		b := nodePoolGet(converted, nodes[1])
		a.next = append(a.next, b)
		b.next = append(b.next, a)
	}

	return startNode
}

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func nodePoolContains(pool []*Node, elem string) uint8 {
	var res uint8
	for _, node := range pool {
		if node.name == elem {
			res++
		}
	}
	return res
}

func nodePoolGet(pool []*Node, elem string) *Node {
	for _, node := range pool {
		if node.name == elem {
			return node
		}
	}
	return nil
}

func ToString(path []*Node) string {
	stringPath := make([]string, 0, len(path))
	for _, node := range path {
		stringPath = append(stringPath, node.name)
	}
	return strings.Join(stringPath, ",")
}
