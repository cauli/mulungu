package tree

import (
	"encoding/json"
)

// ToJSON returna a JSON representation of a given Tree
func (tree Tree) ToJSON() (string, error) {
	json, err := json.Marshal(tree)
	if err != nil {
		return "", err
	}

	return string(json), nil
}

// FromJSON receives JSON data for a tree and
// returns the unmarshalled struct of a Tree
func FromJSON(value string) (*Tree, error) {
	tree := Tree{}

	err := json.Unmarshal([]byte(value), &tree)
	if err != nil {
		return nil, err
	}

	return &tree, nil
}
