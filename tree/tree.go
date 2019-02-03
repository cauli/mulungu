package tree

import (
	"encoding/json"
	"fmt"

	"github.com/TylerBrock/colorjson"
)

type Tree struct {
	Id   string `json:"id"`
	Root *Node  `json:"root"`
}

type Node struct {
	Id       string   `json:"id"`
	Data     MetaData `json:"metadata,omitempty"`
	Children []*Node  `json:"children,omitempty"`
	parent   *Node    `json:"parent,omitempty"`
}

type MetaData struct {
	Name  string `json:"name,omitempty"`
	Title string `json:"title,omitempty"`
}

func Create(treeId string) (*Tree, error) {
	rootNode := Node{
		Id: "1", //uuid.New().String(),
		Data: MetaData{
			Name:  "#1",
			Title: "Founder",
		},
	}

	tree := &Tree{
		Id:   treeId,
		Root: &rootNode,
	}

	return tree, nil
}

func (tree Tree) GetRoot() (*Node, error) {
	if tree.Root == nil {
		return nil, fmt.Errorf("Tree does not contain a root node")
	}

	return tree.Root, nil
}

func (tree Tree) FindNode(id string, currentNode *Node) (*Node, error) {
	if currentNode == nil {
		rootNode, err := tree.GetRoot()
		if err != nil {
			return nil, err
		}

		currentNode = rootNode
	}

	if currentNode.Id == id {
		return currentNode, nil
	}

	if currentNode.Children != nil {
		for _, child := range currentNode.Children {
			foundNode, err := tree.FindNode(id, child)
			if err != nil {
				return nil, err
			}
			if foundNode != nil {
				return foundNode, nil
			}
		}
	}

	return nil, nil
}

func (tree Tree) InsertNode(newNode *Node, parent *Node) error {
	if parent == nil || newNode == nil {
		return fmt.Errorf("Must provide a new node and parent to attach it to")
	}

	if parent == newNode {
		return fmt.Errorf("Cannot attach a node to itself")
	}
	newNode.parent = parent
	parent.Children = append(parent.Children, newNode)

	return nil
}

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
