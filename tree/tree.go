package tree

import (
	"encoding/json"
	"fmt"
)

type Tree struct {
	Id   string `json:"id"`
	Root *Node  `json:"root"`
}

type Node struct {
	Id       string   `json:"id"`
	Data     MetaData `json:"metadata"`
	Children *[]Node  `json:"children"`
	parent   *Node    `json:"parent"`
}

type MetaData struct {
	Name  string `json:"name"`
	Title string `json:"title"`
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

	fmt.Println("Checking node: ", currentNode.Id)

	if currentNode.Children != nil {
		for _, child := range *currentNode.Children {
			foundNode, err := tree.FindNode(id, &child)
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
	if parent == nil {
		return fmt.Errorf("Must provide a parent node")
	}

	if newNode == nil {
		return fmt.Errorf("Must provide a new node")
	}

	newNode.parent = parent

	if parent.Children == nil {
		parent.Children = &[]Node{}
	}

	*parent.Children = append(*parent.Children, *newNode)
	return nil
}

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
