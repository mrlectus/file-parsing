package parser

import (
	"encoding/json"
	"fmt"
	"strings"
)

type JSONRepetedParser struct{}

func (csv *JSONRepetedParser) ParseFile(file string) ([]Players, error) {
	players := []Players{}
	text := strings.Split(file, "\n")
	for _, y := range text {
		player := Players{}
		if strings.HasPrefix(y, "#") {
		} else {
			err := json.Unmarshal([]byte(y), &player)
			if err != nil {
				return nil, fmt.Errorf("%w", err)
			}
		}
		players = append(players, player)
	}
	return players, nil
}
