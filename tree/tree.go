package tree

import (
	"fmt"
)

type Tree struct {
	Id   string `json:"id"`
	Root *Node  `json:"root"`
}

type Node struct {
	ID       string   `json:"id"`
	Data     MetaData `json:"metadata,omitempty"`
	Children []*Node  `json:"children,omitempty"`
	ParentID string   `json:"parentId,omitEmpty"`
}

type MetaData struct {
	Name  string `json:"name,omitempty"`
	Title string `json:"title,omitempty"`
}

type SubordinatesResponse struct {
	subordinates SubordinatesInfo `json:"subordinates"`
}

type SubordinatesInfo struct {
	count     SubordinatesCount `json:"count"`
	hierarchy []*Node           `json:"hierarchy"`
}

type SubordinatesCount struct {
	direct int `json:"direct,omitempty"`
	total  int `json:"total,omitempty"`
}

func Create(treeId string) (*Tree, error) {
	rootNode := Node{
		ID: "1", //uuid.New().String(),
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

	if currentNode.ID == id {
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

func (tree Tree) GetDescendants(id string) ([]*Node, error) {
	node, err := tree.FindNode(id, nil)
	if err != nil {
		return nil, err
	}

	if node == nil {
		return nil, fmt.Errorf("ID `%s` was not found", id)
	}

	return node.Children, nil
}

func (node *Node) GetDescendants() (SubordinatesResponse, error) {
	directCount := len(node.Children)

	totalCount := 0
	for _, descendant := range node.Children {
		totalCount += descendant.countAllDescendants() + 1
	}

	response := SubordinatesResponse{
		subordinates: SubordinatesInfo{
			count: SubordinatesCount{
				direct: directCount,
				total:  totalCount,
			},
			hierarchy: node.Children,
		},
	}

	return response, nil
}

func (tree Tree) InsertNode(newNode *Node, parent *Node) error {
	if parent == nil || newNode == nil {
		return fmt.Errorf("Must provide a new node and parent to attach it to")
	}

	if parent == newNode {
		return fmt.Errorf("It is not possible to attach a Node to itself")
	}

	newNode.ParentID = parent.ID
	parent.Children = append(parent.Children, newNode)

	return nil
}

func (tree Tree) DetachNode(node *Node) (*Node, error) {
	if node == nil {
		return nil, fmt.Errorf("Must provide a new node to detach")
	}

	if node == tree.Root {
		return nil, fmt.Errorf("Cannot detach root node from tree. Where would you attach it?")
	}

	parentNode, err := tree.FindNode(node.ParentID, nil)
	if err != nil {
		return nil, fmt.Errorf("Could no find find parent node to detach.\nDetails: '%v'", err)
	}

	parentNode.RemoveChildren(node)

	return node, nil
}

func (node *Node) RemoveChildren(childToRemove *Node) {
	indexToRemove := -1

	for index, child := range node.Children {
		if child.ID == childToRemove.ID {
			indexToRemove = index
			childToRemove.ParentID = ""
			break
		}
	}

	if indexToRemove != -1 {
		node.Children = append(node.Children[:indexToRemove], node.Children[indexToRemove+1:]...)

		if len(node.Children) == 0 {
			node.Children = nil
		}
	}
}

func (node *Node) countAllDescendants() int {
	var totalCount int

	for _, descendant := range node.Children {
		totalCount += descendant.countAllDescendants() + 1
	}

	return totalCount
}
