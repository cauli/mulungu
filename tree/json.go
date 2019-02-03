package tree

import (
	"encoding/json"
	"fmt"

	"github.com/TylerBrock/colorjson"
)

func (tree Tree) ToJSON() (string, error) {
	json, err := json.Marshal(tree)
	if err != nil {
		return "", err
	}

	return string(json), nil
}

func (tree Tree) Print() error {
	str, err := tree.ToJSON()
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	json.Unmarshal([]byte(str), &obj)

	formatter := colorjson.NewFormatter()
	formatter.Indent = 2

	s, _ := formatter.Marshal(obj)
	fmt.Println("\n", string(s))

	return nil
}

func FromJSON(value string) (*Tree, error) {
	tree := Tree{}

	err := json.Unmarshal([]byte(value), &tree)
	if err != nil {
		return nil, err
	}

	return &tree, nil
}
