package model

import "github.com/cauli/mulungu/tree"

// Employee is the public representation of a tree.Node
type Employee struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Title  string `json:"title"`
	Leader string `json:"leader"`
}

// CreateNode generates a *tree.Node given input from the API
func (employee *Employee) CreateNode() *tree.Node {
	return &tree.Node{
		ID: employee.ID,
		Data: tree.MetaData{
			Name:  employee.Name,
			Title: employee.Title,
		},
		ParentID: employee.Leader,
	}
}
