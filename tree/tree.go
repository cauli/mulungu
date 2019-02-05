package tree

import (
	"fmt"
)

type Tree struct {
	Id   string `json:"id"`
	Root *Node  `json:"root"`
}

type SubordinatesResponse struct {
	Subordinates SubordinatesInfo `json:"subordinates"`
}

type SubordinatesInfo struct {
	Count     SubordinatesCount `json:"count"`
	Hierarchy []*Node           `json:"hierarchy,omitempty"`
}

type SubordinatesCount struct {
	Direct int `json:"direct"`
	Total  int `json:"total"`
}

func Create(treeId string) (*Tree, error) {
	rootNode := Node{
		ID: "root",
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

func (tree Tree) AttachNode(newNode *Node, parent *Node) error {
	if parent == nil || newNode == nil {
		return fmt.Errorf("Must provide a new node and parent to attach it to")
	}

	if parent == newNode {
		return fmt.Errorf("It is not possible to attach a Node to itself")
	}

	foundParentNode, err := tree.FindNode(parent.ID, nil)
	if err != nil {
		return fmt.Errorf("An error has occurred when searching for parent node on the tree.\nDetails: %s", err.Error())
	}

	if foundParentNode == nil {
		return fmt.Errorf("The parent node does not exist on the tree")
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
		return nil, fmt.Errorf("Could not find parent node to detach on this tree.\nDetails: '%v'", err)
	}

	parentNode.RemoveChildren(node)

	return node, nil
}

func (tree Tree) MoveNode(node *Node, newParent *Node) error {
	if node == nil || newParent == nil {
		return fmt.Errorf("Must provide a node to move and parent to attach it to")
	}

	detachedNode, err := tree.DetachNode(node)
	if err != nil {
		return fmt.Errorf("Unable to detach node from tree.\nDetails: %s", err.Error())
	}

	tree.AttachNode(detachedNode, newParent)
	if err != nil {
		return fmt.Errorf("Unable to attach node to new parent.\nDetails: %s", err.Error())
	}

	return nil
}
