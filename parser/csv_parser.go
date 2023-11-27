package parser

import (
	"strconv"
	"strings"
)

type CSVParser struct{}

func (csv *CSVParser) ParseFile(file string) ([]Players, error) {
	players := []Players{}
	tokens := strings.Split(file, "\n")
	for index := 1; index < len(tokens); index++ {
		token := strings.Split(tokens[index], ",")
		val, err := strconv.ParseInt(token[1], 10, 64)
		if err != nil {
			return nil, err
		}
		player := Players{
			Name:       token[0],
			High_score: int(val),
		}
		players = append(players, player)
	}
	return players, nil
}
