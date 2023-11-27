package parser

import (
	"encoding/json"
	"fmt"
)

type JSONParser struct{}

func (csv *JSONParser) ParseFile(file string) ([]Players, error) {
	players := []Players{}
	err := json.Unmarshal([]byte(file), &players)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return players, nil
}
