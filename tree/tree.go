package tree

import (
	"fmt"
)

// Tree is the basic struct for any tree hierarchy
type Tree struct {
	Id   string `json:"id"`
	Root *Node  `json:"root"`
}

// SubordinatesResponse is the base
// response for the GET subordinates route
type SubordinatesResponse struct {
	Subordinates SubordinatesInfo `json:"subordinates"`
}

// SubordinatesInfo contains a `count` of subordinates
// and a full `hierarchy` of nodes
type SubordinatesInfo struct {
	Count     SubordinatesCount `json:"count"`
	Hierarchy []*Node           `json:"hierarchy,omitempty"`
}

// SubordinatesCount contains a `direct` and a `total`
// count of subordinates
type SubordinatesCount struct {
	Direct int `json:"direct"`
	Total  int `json:"total"`
}

// New will create a new tree given an id
// by default will also create a node with id `root`
func New(id string) *Tree {
	rootNode := Node{
		ID: "root",
		Data: MetaData{
			Name:  "#1",
			Title: "Founder",
		},
		Height: 0,
	}

	tree := &Tree{
		Id:   id,
		Root: &rootNode,
	}

	return tree
}

// GetRoot will retrieve the Root Node of a tree
func (tree Tree) GetRoot() (*Node, error) {
	if tree.Root == nil {
		return nil, fmt.Errorf("Tree does not contain a root node")
	}

	return tree.Root, nil
}

// FindNode will find a node with `id` in the tree in O(n) operations
// by default, `start` node will be the Root node
// it is possible to define the `start` node to speed up operations
func (tree Tree) FindNode(id string, start *Node) (*Node, error) {
	if start == nil {
		rootNode, err := tree.GetRoot()
		if err != nil {
			return nil, err
		}

		start = rootNode
	}

	if start.ID == id {
		return start, nil
	}

	if start.Children != nil {
		for _, child := range start.Children {
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

// AttachNode will receive two node pointers in a tree,
// and will attach the first one as a subordinate of the foremost
// this will update the entire subtree of the node being attached
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

	root, err := tree.GetRoot()
	if err != nil {
		return err
	}

	newNode.RootID = &root.ID
	newNode.ParentID = parent.ID
	newNode.updateHeight(parent.Height + 1)

	parent.Children = append(parent.Children, newNode)

	return nil
}

// DetachNode will remove a `node` from its parent
// and will return a pointer to the detached node
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

	parentNode.removeChildren(node)
	node.RootID = nil

	return node, nil
}

// MoveNode will detach a the subtree of a given `node`
// and then attach it as a descendant of `newParent`
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
