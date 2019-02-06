package tree

import (
	"encoding/json"
)

func (tree Tree) ToJSON() (string, error) {
	json, err := json.Marshal(tree)
	if err != nil {
		return "", err
	}

	return string(json), nil
}

func FromJSON(value string) (*Tree, error) {
	tree := Tree{}

	err := json.Unmarshal([]byte(value), &tree)
	if err != nil {
		return nil, err
	}

	return &tree, nil
}
