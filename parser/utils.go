package parser

import (
	"math"
	"os"
	"strings"
)

var (
	p          = &JSONRepetedParser{}
	parser     = NewParser(p)
	file, _    = os.ReadFile("../examples/repeated-json.txt")
	players, _ = parser.ParseFile(strings.TrimSpace(string(file)))
	minmax     = NewTuple(math.MaxInt, math.MinInt)
)

func High() string {
	var player1 string
	for _, player := range players {
		if player.High_score > minmax.max {
			minmax.max = player.High_score
			player1 = player.Name
		}
	}
	return player1
}

func Low() string {
	var player2 string
	for _, player := range players {
		if player.High_score < minmax.min {
			minmax.min = player.High_score
			player2 = player.Name
		}
	}
	return player2
}
